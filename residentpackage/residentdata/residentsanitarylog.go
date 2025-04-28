package residentdata

type ResidentSanitaryLog struct {
	toiletLog      *SanitaryLog
	showerLog      *SanitaryLog
	washBassinLog  *SanitaryLog
	washMachineLog *SanitaryLog
	dishWasherLog  *SanitaryLog
	tanqueLog      *SanitaryLog
}

// NewResidentSanitaryLog creates a new ResidentSanitaryLog with initialized SanitaryLog instances
func NewResidentSanitaryLog() *ResidentSanitaryLog {
	return &ResidentSanitaryLog{
		toiletLog:      NewSanitaryLog("toilet", 0),
		showerLog:      NewSanitaryLog("shower", 0),
		washBassinLog:  NewSanitaryLog("wash_bassin", 0),
		washMachineLog: NewSanitaryLog("wash_machine", 0),
		dishWasherLog:  NewSanitaryLog("dish_washer", 0),
		tanqueLog:      NewSanitaryLog("tanque", 0),
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