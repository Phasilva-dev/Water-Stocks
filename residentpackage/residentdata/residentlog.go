package residentdata

type ResidentLog struct {
	day                  uint8
	houseClassID         uint16
	residentOccupationID uint16
	age                  uint8
	sanitaryLogs         *ResidentSanitaryLog
}

// NewResidentLog creates a new initialized ResidentLog instance
func NewResidentLog(
	day uint8,
	houseClassID uint16,
	residentOccupationID uint16,
	age uint8,
) *ResidentLog {
	return &ResidentLog{
		day:                  day,
		houseClassID:         houseClassID,
		residentOccupationID: residentOccupationID,
		age:                  age,
		sanitaryLogs:         , //FALTA AQUI
	}
}

// Getters
func (r *ResidentLog) GetDay() uint8 {
	return r.day
}

func (r *ResidentLog) GetHouseClassID() uint16 {
	return r.houseClassID
}

func (r *ResidentLog) GetResidentOccupationID() uint16 {
	return r.residentOccupationID
}

func (r *ResidentLog) GetAge() uint8 {
	return r.age
}

func (r *ResidentLog) GetSanitaryLogs() *ResidentSanitaryLog {
	return r.sanitaryLogs
}