package entities

import (
	"math/rand/v2"
	"simulation/internal/entities/resident/ds/temporal"
	"simulation/internal/entities/resident"


)

type Resident struct {
	age uint8
	occupationID uint32 //Ocupação, exemplo, estudante
	dayData *temporal.DailyData
	residentProfile *resident.ResidentProfile
	//sanitaryLog *residentdata.ResidentSanitaryLog
	house *House
}

func NewResident(age uint8, occupation uint32, profile *resident.ResidentProfile, house *House) *Resident {
	return &Resident{
		age: age,
		occupationID: occupation,
		dayData: temporal.NewDailyData(nil,nil),
		residentProfile: profile,
		//sanitaryLog: nil,
		house: house,
	}
}

func (r *Resident) Age() uint8 {
	return r.age
}
func (r *Resident) OccupationID() uint32 {
	return r.occupationID
}

func (r *Resident) DayData() *temporal.DailyData {
	return r.dayData
}

func (r *Resident) ResidentProfile() *resident.ResidentProfile {
	return r.residentProfile
}


func (r *Resident) GenerateFrequency(day uint8, rng *rand.Rand) {
	r.dayData.SetFrequency(r.residentProfile.GenerateFrequency(day,rng))
}

func (r *Resident) GenerateRoutine(day uint8, rng *rand.Rand) {
	r.dayData.SetRoutine(r.residentProfile.GenerateRoutine(day,rng))
}


//Terei que trocar
func (r *Resident) GenerateDailyData(day uint8, rng *rand.Rand) {
	r.GenerateRoutine(day,rng)
	r.GenerateFrequency(day,rng)
}
