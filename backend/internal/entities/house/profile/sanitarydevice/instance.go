package sanitarydevice

import (
	"fmt"
)

type SanitaryDeviceInstance struct {
	device *SanitaryDevice
	amount uint8
}

// NewSanitaryDeviceInstancee creates a new instance of SanitaryDeviceInstance.
// Validations:
// 1. amount must be greater than 0 (cannot create an instance with zero devices).
// 2. device must not be nil (a valid device is required).
func NewSanitaryDeviceInstance(device *SanitaryDevice, amount uint8) (*SanitaryDeviceInstance, error) {
	if device == nil {
		return nil, fmt.Errorf("invalid SanitaryDeviceInstance: device cannot be nil")
	}

	if amount == 0 {
		return nil, fmt.Errorf("invalid SanitaryDeviceInstance: amount must be greater than 0, got %d", amount)
	}

	return &SanitaryDeviceInstance{
		device: device,
		amount: amount,
	}, nil
}

func (sdi *SanitaryDeviceInstance) Amount() uint8 {
	return sdi.amount
}

func (sdi *SanitaryDeviceInstance) Device() *SanitaryDevice {
	return sdi.device
}

