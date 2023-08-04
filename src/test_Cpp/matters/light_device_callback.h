#include "light_manager.h"
#include <app/util/af-types.h>
#include <app/util/basic-types.h>
#include <platform/CHIPDeviceLayer.h> // for chip::DeviceLayer::ChipDeviceEvent

class light_device_callback : public chip::DeviceManager::CHIPDeviceManagerCallbacks
{
public:
    virtual void light_device_callBack(const chip::DeviceLayer::ChipDeviceEvent * event, intptr_t arg);
    void PostAttributeChangeCallback(chip::EndpointId endpointId, chip::ClusterId clusterId, chip::AttributeId attributeId, uint8_t type, uint16_t size, uint8_t * value) override;

private:
    void OnInternetConnectivityChange(const chip::DeviceLayer::ChipDeviceEvent * event);
    void Oninternet_connection_with_light_device(chip::EndpointId endpointId, chip::AttributeId attributeId, uint8_t * value);
};