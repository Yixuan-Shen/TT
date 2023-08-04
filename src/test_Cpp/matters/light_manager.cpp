#include <stdlib.h>

#include "CHIPDeviceManager.h"
#include <app/util/basic-types.h>
#include <credentials/DeviceAttestationCredsProvider.h>
#include <credentials/examples/DeviceAttestationCredsExample.h>
#include <platform/Ameba/FactoryDataProvider.h>
#include <support/CHIPMem.h>
#include <support/CodeUtils.h>
#include <support/ErrorStr.h>

#include "Globals.h"
#include "LEDWidget.h"

#include <app-common/zap-generated/attributes/Accessors.h>
#include <app-common/zap-generated/ids/Attributes.h>
#include <app-common/zap-generated/ids/Clusters.h>
#include <app/util/af-types.h>
#include <app/util/attribute-storage.h>
#include <app/util/util.h>

using namespace ::chip;
using namespace ::chip::Credentials;

namespace chip {

namespace DeviceManager {

using namespace ::chip::DeviceLayer;

chip::DeviceLayer::FactoryDataProvider mFactoryDataProvider;

void CHIPDeviceManager::CommonDeviceEventHandler(const ChipDeviceEvent * event, intptr_t arg)
{
    CHIPDeviceManagerCallbacks * cb = reinterpret_cast<CHIPDeviceManagerCallbacks *>(arg);
    if (cb != nullptr)
    {
        cb->DeviceEventCallback(event, reinterpret_cast<intptr_t>(cb));
    }
}
