package residentdata

import "housedata"

type ResidentSanitaryLog struct {
	toiletLog      *SanitaryLog
	showerLog      *SanitaryLog
	washBassinLog  *SanitaryLog
	washMachineLog *SanitaryLog
	dishWasherLog  *SanitaryLog
	tanqueLog      *SanitaryLog
}

// NewResidentSanitaryLog creates a new ResidentSanitaryLog with initialized SanitaryLog instances
func NewResidentSanitaryLog(sanitaryHouse *housedata.SanitaryHouse) *ResidentSanitaryLog {

	return &ResidentSanitaryLog{
		toiletLog:      NewSanitaryLog("toilet", sanitaryHouse.Toilet().Device().SanitaryDeviceID()),
		showerLog:      NewSanitaryLog("shower", sanitaryHouse.Shower().Device().SanitaryDeviceID()),
		washBassinLog:  NewSanitaryLog("wash_bassin", sanitaryHouse.WashBassin().Device().SanitaryDeviceID()),
		washMachineLog: NewSanitaryLog("wash_machine", sanitaryHouse.WashMachine().Device().SanitaryDeviceID()),
		dishWasherLog:  NewSanitaryLog("dish_washer", sanitaryHouse.DishWasher().Device().SanitaryDeviceID()),
		tanqueLog:      NewSanitaryLog("tanque", sanitaryHouse.Tanque().Device().SanitaryDeviceID()),
	}
}

// Getters for each log type
func (r *ResidentSanitaryLog) GetToiletLog() *SanitaryLog {
	return r.toiletLog
}

func (r *ResidentSanitaryLog) GetShowerLog() *SanitaryLog {
	return r.showerLog
}

func (r *ResidentSanitaryLog) GetWashBassinLog() *SanitaryLog {
	return r.washBassinLog
}

func (r *ResidentSanitaryLog) GetWashMachineLog() *SanitaryLog {
	return r.washMachineLog
}

func (r *ResidentSanitaryLog) GetDishWasherLog() *SanitaryLog {
	return r.dishWasherLog
}

func (r *ResidentSanitaryLog) GetTanqueLog() *SanitaryLog {
	return r.tanqueLog
}