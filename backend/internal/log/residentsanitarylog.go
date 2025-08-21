package log



type ResidentSanitary struct {
	toiletLog      *Sanitary
	showerLog      *Sanitary
	washBassinLog  *Sanitary
	washMachineLog *Sanitary
	dishWasherLog  *Sanitary
	tanqueLog      *Sanitary
}

// NewResidentSanitary creates a new ResidentSanitary with initialized Sanitary instances
func NewResidentSanitary(toilet,shower,washBassin,washMachine,dishWasher,tanque *Sanitary) *ResidentSanitary {

	return &ResidentSanitary{
		toiletLog:      toilet,
		showerLog:      shower,
		washBassinLog:  washBassin,
		washMachineLog: washMachine,
		dishWasherLog:  dishWasher,
		tanqueLog:      tanque,
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

func (r *ResidentSanitary) SetWashMachineLog(log *Sanitary) {
	r.washMachineLog = log
}

func (r *ResidentSanitary) SetDishWasherLog(log *Sanitary) {
	r.dishWasherLog = log
}

func (r *ResidentSanitary) SetTanqueLog(log *Sanitary) {
	r.tanqueLog = log
}