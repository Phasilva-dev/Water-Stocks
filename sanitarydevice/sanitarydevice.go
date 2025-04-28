package sanitarydevice

import (
	"math/rand/v2"
)

type SanitaryDevice interface {
	GenerateFlowLeak(rng *rand.Rand) float64
	GenerateDuration(rng *rand.Rand) int32
}

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


