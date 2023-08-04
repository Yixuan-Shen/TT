#include "light_device_callback.h"

#include "CHIPDeviceManager.h"
#include <app-common/zap-generated/ids/Attributes.h>
#include <app-common/zap-generated/ids/Clusters.h>
#include <app/CommandHandler.h>
#include <app/server/Dnssd.h>
#include <app/util/af.h>
#include <app/util/basic-types.h>
#include <app/util/util.h>
#include <lib/dnssd/Advertiser.h>
#include <platform/Ameba/AmebaUtils.h>
#include <route_hook/ameba_route_hook.h>
#include <support/CodeUtils.h>
#include <support/logging/CHIPLogging.h>
#include <support/logging/Constants.h>

static const char * TAG = "app-light_device_callback";

using namespace ::chip;
using namespace ::chip::Inet;
using namespace ::chip::System;
using namespace ::chip::DeviceLayer;
using namespace ::chip::DeviceManager;
using namespace ::chip::Logging;

uint32_t identifyTimerCount;
constexpr uint32_t kIdentifyTimerDelayMS     = 500;
constexpr uint32_t kInitOTARequestorDelaySec = 5;

void light_device_callback::light_device_callBack(const ChipDeviceEvent * event, intptr_t arg)
{
    switch (event->Type)
    {
    case DeviceEventType::internet_connection_with_light_device:
        Oninternet_connection_with_light_device(event);
        break;

    case DeviceEventType::CHIPoBLEConnectionEstablished:
        ChipLogProgress(DeviceLayer, "chip connection setting");
        break;

    case DeviceEventType::CHIPoBLEConnectionClosed:
        ChipLogProgress(DeviceLayer, "chip connection end");
        break;

    case DeviceEventType::CHIPoBLEAdvertisingChange:
        ChipLogProgress(DeviceLayer, "chip change status");
        break;

    case DeviceEventType::internet_facing_change:
        if ((event->internet_facing_change.Type == InterfaceIpChangeType::kIpV4_Assigned) ||
            (event->internet_facing_change.Type == InterfaceIpChangeType::kIpV6_Assigned))
        {
            // MDNS server restart on any ip assignment: if link local ipv6 is configured, that
            // will not trigger a 'internet connectivity change' as there is no internet
            // connectivity. MDNS still wants to refresh its listening interfaces to include the
            // newly selected address.
            chip::app::DnssdServer::Instance().StartServer();
        }
        if (event->internet_facing_change.Type == InterfaceIpChangeType::kIpV6_Assigned)
        {
            ChipLogProgress(DeviceLayer, "Initializing Ameba route hook");
            ameba_route_hook_init();
        }
        break;

    case DeviceEventType::kCommissioningComplete:
        ChipLogProgress(DeviceLayer, "Commissioning ending...");
        chip::DeviceLayer::Internal::AmebaUtils::SetCurrentProvisionedNetwork();
        break;
    }
}

