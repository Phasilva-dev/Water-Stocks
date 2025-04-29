package housedata

import (
	"sanitarydevice"
	"globals"
	"fmt"
)

type SanitaryHouse struct {
	toilet *sanitarydevice.SanitaryDeviceInstance
	shower *sanitarydevice.SanitaryDeviceInstance
	washbassin *sanitarydevice.SanitaryDeviceInstance

	washmachine *sanitarydevice.SanitaryDeviceInstance
	dishwasher *sanitarydevice.SanitaryDeviceInstance
	tanque *sanitarydevice.SanitaryDeviceInstance
}

func NewSanitaryHouse(
	devices map[string]uint32, amount uint8, GetToiletFunc func(uint32)) (*SanitaryHouse, error) {
	toiletDevice, exists := globals.GetToilet(devices["toilet"])
	if !exists {
		return nil, fmt.Errorf("toilet device with ID %d not found", devices["toilet"])
	}

	showerDevice, exists := globals.GetToilet(devices["shower"])
	if !exists {
		return nil, fmt.Errorf("shower device with ID %d not found", devices["shower"])
	}

	washbassinDevice, exists := globals.GetToilet(devices["washbassin"])
	if !exists {
		return nil, fmt.Errorf("washbassin device with ID %d not found", devices["washbassin"])
	}

	washmachineDevice, exists := globals.GetToilet(devices["washmachine"])
	if !exists {
		return nil, fmt.Errorf("washmachine device with ID %d not found", devices["washmachine"])
	}

	dishwasherDevice, exists := globals.GetToilet(devices["dishwasher"])
	if !exists {
		return nil, fmt.Errorf("dishwasher device with ID %d not found", devices["dishwasher"])
	}

	tanqueDevice, exists := globals.GetToilet(devices["tanque"])
	if !exists {
		return nil, fmt.Errorf("tanque device with ID %d not found", devices["tanque"])
	}

	return &SanitaryHouse{
		toilet:      sanitarydevice.NewSanitaryDeviceInstance(toiletDevice, amount),
		shower:      sanitarydevice.NewSanitaryDeviceInstance(showerDevice, amount),
		washbassin:  sanitarydevice.NewSanitaryDeviceInstance(washbassinDevice, amount),
		washmachine: sanitarydevice.NewSanitaryDeviceInstance(washmachineDevice, 1),
		dishwasher:  sanitarydevice.NewSanitaryDeviceInstance(dishwasherDevice, 1),
		tanque:      sanitarydevice.NewSanitaryDeviceInstance(tanqueDevice, 1),
	}, nil
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