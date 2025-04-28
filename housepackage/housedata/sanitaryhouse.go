package housedata

import (
	"sanitarydevice"
)

type SanitaryHouse struct {
	toilet *sanitarydevice.SanitaryDeviceInstance
	shower *sanitarydevice.SanitaryDeviceInstance
	washbassin *sanitarydevice.SanitaryDeviceInstance

	washmachine *sanitarydevice.SanitaryDeviceInstance
	dishwasher *sanitarydevice.SanitaryDeviceInstance
	tanque *sanitarydevice.SanitaryDeviceInstance
}

func NewSanitaryHouse(devices map[string]*sanitarydevice.SanitaryDevice, amount uint8) *SanitaryHouse {
	return &SanitaryHouse{
		toilet:      sanitarydevice.NewSanitaryDeviceInstance(devices["toilet"], amount),
		shower:      sanitarydevice.NewSanitaryDeviceInstance(devices["shower"], amount),
		washbassin:  sanitarydevice.NewSanitaryDeviceInstance(devices["washbassin"], amount),
		washmachine: sanitarydevice.NewSanitaryDeviceInstance(devices["washmachine"], 1),
		dishwasher:  sanitarydevice.NewSanitaryDeviceInstance(devices["dishwasher"], 1),
		tanque:      sanitarydevice.NewSanitaryDeviceInstance(devices["tanque"], 1),
	}
}

func (h *SanitaryHouse) Toilet() *sanitarydevice.SanitaryDeviceInstance {
	return h.toilet
}

func (h *SanitaryHouse) Shower() *sanitarydevice.SanitaryDeviceInstance {
	return h.shower
}

func (h *SanitaryHouse) WashBassin() *sanitarydevice.SanitaryDeviceInstance {
	return h.washbassin
}

func (h *SanitaryHouse) WashMachine() *sanitarydevice.SanitaryDeviceInstance {
	return h.washmachine
}

func (h *SanitaryHouse) DishWasher() *sanitarydevice.SanitaryDeviceInstance {
	return h.dishwasher
}

func (h *SanitaryHouse) Tanque() *sanitarydevice.SanitaryDeviceInstance {
	return h.tanque
}