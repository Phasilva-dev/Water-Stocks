package log

import "simulation/internal/entities/house/ds/sanitarysystem"

type ResidentSanitary struct {
	toiletLog      *Sanitary
	showerLog      *Sanitary
	washBassinLog  *Sanitary
	washMachineLog *Sanitary
	dishWasherLog  *Sanitary
	tanqueLog      *Sanitary
}

// NewResidentSanitary creates a new ResidentSanitary with initialized Sanitary instances
func NewResidentSanitary(sanitaryHouse *sanitarysystem.SanitaryHouse) *ResidentSanitary {

	return &ResidentSanitary{
		toiletLog:      NewSanitary("toilet", sanitaryHouse.Toilet().Device().SanitaryDeviceID()),
		showerLog:      NewSanitary("shower", sanitaryHouse.Shower().Device().SanitaryDeviceID()),
		washBassinLog:  NewSanitary("wash_bassin", sanitaryHouse.WashBassin().Device().SanitaryDeviceID()),
		washMachineLog: NewSanitary("wash_machine", sanitaryHouse.WashMachine().Device().SanitaryDeviceID()),
		dishWasherLog:  NewSanitary("dish_washer", sanitaryHouse.DishWasher().Device().SanitaryDeviceID()),
		tanqueLog:      NewSanitary("tanque", sanitaryHouse.Tanque().Device().SanitaryDeviceID()),
	}
}

// Getters for each log type
func (r *ResidentSanitary) ToiletLog() *Sanitary {
	return r.toiletLog
}

func (r *ResidentSanitary) ShowerLog() *Sanitary {
	return r.showerLog
}

func (r *ResidentSanitary) WashBassinLog() *Sanitary {
	return r.washBassinLog
}

func (r *ResidentSanitary) WashMachineLog() *Sanitary {
	return r.washMachineLog
}

func (r *ResidentSanitary) DishWasherLog() *Sanitary {
	return r.dishWasherLog
}

func (r *ResidentSanitary) TanqueLog() *Sanitary {
	return r.tanqueLog
}