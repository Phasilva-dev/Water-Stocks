package sanitarysystem

import (
	"simulation/internal/entities/house/profile/sanitarydevice"
	//"globals"
	"fmt"
	//"interfaces" Caso eu for usar globals
)

type SanitaryHouse struct {
	toilet *sanitarydevice.SanitaryDeviceInstance
	shower *sanitarydevice.SanitaryDeviceInstance
	washbassin *sanitarydevice.SanitaryDeviceInstance

	washmachine *sanitarydevice.SanitaryDeviceInstance
	dishwasher *sanitarydevice.SanitaryDeviceInstance
	tanque *sanitarydevice.SanitaryDeviceInstance
	amount uint8
}

func NewSanitaryHouse(
	devices map[string]sanitarydevice.SanitaryDevice, amount uint8) (*SanitaryHouse, error) {
	toiletDevice := devices["toilet"]
	if toiletDevice != nil {
		return nil, fmt.Errorf("toilet device with ID %d not found", devices["toilet"])
	}

	showerDevice := devices["shower"]
	if showerDevice != nil {
		return nil, fmt.Errorf("shower device with ID %d not found", devices["shower"])
	}

	washbassinDevice := devices["wash_bassin"]
	if washbassinDevice != nil {
		return nil, fmt.Errorf("washbassin device with ID %d not found", devices["washbassin"])
	}

	washmachineDevice := devices["wash_machine"]
	if washmachineDevice != nil {
		return nil, fmt.Errorf("washmachine device with ID %d not found", devices["washmachine"])
	}

	dishwasherDevice := devices["dish_washer"]
	if dishwasherDevice != nil {
		return nil, fmt.Errorf("dishwasher device with ID %d not found", devices["dishwasher"])
	}

	tanqueDevice := devices["tanque"]
	if tanqueDevice != nil {
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

