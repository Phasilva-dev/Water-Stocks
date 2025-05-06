package log

type Resident struct {
	day                  uint8
	houseClassID         uint32
	residentOccupationID uint32
	age                  uint8
	sanitaryLogs         *ResidentSanitary
}

// NewResident creates a new initialized Resident instance
func NewResident(
	day uint8,
	houseClassID uint32,
	residentOccupationID uint32,
	age uint8,
) *Resident {
	return &Resident{
		day:                  day,
		houseClassID:         houseClassID,
		residentOccupationID: residentOccupationID,
		age:                  age,
		sanitaryLogs:         nil, //FALTA AQUI
	}
}

// Getters
func (r *Resident) GetDay() uint8 {
	return r.day
}

func (r *Resident) GetHouseClassID() uint32 {
	return r.houseClassID
}

func (r *Resident) GetResidentOccupationID() uint32 {
	return r.residentOccupationID
}

func (r *Resident) GetAge() uint8 {
	return r.age
}

func (r *Resident) GetSanitaryLogs() *ResidentSanitary {
	return r.sanitaryLogs
}