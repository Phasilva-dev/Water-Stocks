package sanitarydevice

import (
	"interfaces"
)

type SanitaryDeviceInstance struct {
	device interfaces.SanitaryDevice
	amount uint8
}

func NewSanitaryDeviceInstance(deviceType interfaces.SanitaryDevice, amount uint8) interfaces.SanitaryDeviceInstance {
	return &SanitaryDeviceInstance{
		device: deviceType,
		amount: amount,
	}
}

func (sdi *SanitaryDeviceInstance) Device() interfaces.SanitaryDevice {
    return sdi.device
}

func (sdi *SanitaryDeviceInstance) Amount() uint8 {
    return sdi.amount
}

