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
func (r *Resident) Day() uint8 {
	return r.day
}

func (r *Resident) HouseClassID() uint32 {
	return r.houseClassID
}

func (r *Resident) ResidentOccupationID() uint32 {
	return r.residentOccupationID
}

func (r *Resident) Age() uint8 {
	return r.age
}

func (r *Resident) SanitaryLogs() *ResidentSanitary {
	return r.sanitaryLogs
}