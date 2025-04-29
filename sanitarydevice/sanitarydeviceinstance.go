package sanitarydevice



type SanitaryDeviceInstance struct {
	device *SanitaryDevice
	amount uint8
}

func NewSanitaryDeviceInstance(deviceType *SanitaryDevice, amount uint8) *SanitaryDeviceInstance {
	return &SanitaryDeviceInstance{
		device: deviceType,
		amount: amount,
	}
}

func (sdi *SanitaryDeviceInstance) Device() *SanitaryDevice {
    return sdi.device
}

func (sdi *SanitaryDeviceInstance) Amount() uint8 {
    return sdi.amount
}

