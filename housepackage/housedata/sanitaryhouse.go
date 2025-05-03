package housedata

import (
	"sanitarydevice"
	"globals"
	"fmt"
	"interfaces"
)

type SanitaryHouse struct {
	toilet interfaces.SanitaryDeviceInstance
	shower interfaces.SanitaryDeviceInstance
	washbassin interfaces.SanitaryDeviceInstance

	washmachine interfaces.SanitaryDeviceInstance
	dishwasher interfaces.SanitaryDeviceInstance
	tanque interfaces.SanitaryDeviceInstance
	amount uint8
}

func NewSanitaryHouse(
	devices map[string]uint32, amount uint8) (*SanitaryHouse, error) {
	toiletDevice, exists := globals.GetToilet(devices["toilet"])
	if !exists {
		return nil, fmt.Errorf("toilet device with ID %d not found", devices["toilet"])
	}

	showerDevice, exists := globals.GetShower(devices["shower"])
	if !exists {
		return nil, fmt.Errorf("shower device with ID %d not found", devices["shower"])
	}

	washbassinDevice, exists := globals.GetWashBasin(devices["washbassin"])
	if !exists {
		return nil, fmt.Errorf("washbassin device with ID %d not found", devices["washbassin"])
	}

	washmachineDevice, exists := globals.GetWashMachine(devices["washmachine"])
	if !exists {
		return nil, fmt.Errorf("washmachine device with ID %d not found", devices["washmachine"])
	}

	dishwasherDevice, exists := globals.GetDishWasher(devices["dishwasher"])
	if !exists {
		return nil, fmt.Errorf("dishwasher device with ID %d not found", devices["dishwasher"])
	}

	tanqueDevice, exists := globals.GetTanque(devices["tanque"])
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
		amount: amount,
	}, nil
}

func (h *SanitaryHouse) Toilet() interfaces.SanitaryDeviceInstance {
	return h.toilet
}

func (h *SanitaryHouse) Shower() interfaces.SanitaryDeviceInstance {
	return h.shower
}

func (h *SanitaryHouse) WashBassin() interfaces.SanitaryDeviceInstance {
	return h.washbassin
}

func (h *SanitaryHouse) WashMachine() interfaces.SanitaryDeviceInstance {
	return h.washmachine
}

func (h *SanitaryHouse) DishWasher() interfaces.SanitaryDeviceInstance {
	return h.dishwasher
}

func (h *SanitaryHouse) Tanque() interfaces.SanitaryDeviceInstance {
	return h.tanque
}