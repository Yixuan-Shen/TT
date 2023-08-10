#include <lib/core/CHIPCore.h>
#include <lib/core/CHIPError.h>
#include <platform/CHIPDeviceLayer.h>
#include <lib/support/DLLUtil.h>
#include <app/util/af-types.h>
#include <stdarg.h>
#include <stdlib.h>

namespace chip {
namespace DeviceManager {

class light_manager_callback
{
public:
    virtual void light_manager(const chip::DeviceLayer::ChipDeviceEvent * event, intptr_t arg) {}

    virtual void PostAttribute(chip::EndpointId endpoint, chip::ClusterId clusterId, chip::AttributeId attributeId,
                                             uint8_t type, uint16_t size, uint8_t * value)
    {}
    virtual ~light() {}
};

/**
 * A common class that drives other components of the CHIP stack
 */
class DLL_EXPORT light_managers
{
public:
    light_managers(const light_managers &)  = delete;
    light_managers(const light_managers &&) = delete;
    light_managers & operator=(const light_managers &) = delete;

    static light_managers & GetInstance()
    {
        static light_managers instance;
        return instance;
    }

    /**
     * Initialise light manager
     *
     * cb Application's instance of the light for consuming events
     */
    CHIP_ERROR Init(light * cb);

    /**
     * Fetch a pointer to the registered a light
     *
     */
    light * Getligh() { return mCB; }

    /**
     * Use internally for registration of the 
     */
    static void light_managerhandle(const chip::DeviceLayer::ChipDeviceEvent * event, intptr_t arg);

private:
    light * mCB = nullptr;
    CHIPDeviceManager() {}
};

} // namespace DeviceManager
} // namespace chip