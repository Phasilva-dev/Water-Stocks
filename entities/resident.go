package entities

import (
	"math/rand/v2"
	"residentdata"
	"residentprofiles"

)

type Resident struct {
	age uint8
	occupationID uint32 //Ocupação, exemplo, estudante
	dayData *residentdata.DailyData
	residentProfile *residentprofiles.ResidentProfile
	//residentdata.ResidentSanitaryLog
	house *House
}

func NewResident(age uint8, occupation uint32, profile *residentprofiles.ResidentProfile, house *House) *Resident {
	return &Resident{
		age: age,
		occupationID: occupation,
		dayData: nil,
		residentProfile: profile,
		house: house,
	}
}

func (r *Resident) Age() uint8 {
	return r.age
}
func (r *Resident) OccupationID() uint32 {
	return r.occupationID
}

func (r *Resident) DayData() *residentdata.DailyData {
	return r.dayData
}

func (r *Resident) ResidentProfile() *residentprofiles.ResidentProfile {
	return r.residentProfile
}


func (r *Resident) GenerateFrequency(day uint8, rng *rand.Rand) {
	r.dayData.SetFrequency(r.residentProfile.GenerateFrequency(day,rng))
}

func (r *Resident) GenerateRoutine(day uint8, rng *rand.Rand) {
	r.dayData.SetRoutine(r.residentProfile.GenerateRoutine(day,rng))
}
//Terei que trocar
func (r *Resident) GenerateUsage(day uint8, rng *rand.Rand) error {
	usage, err := r.residentProfile.GenerateUsage(day,r.dayData.Frequency(), rng)
	if err != nil {
		// Propaga o erro para a função chamadora
		return err
	}
	// Se não houver erro, define o uso normalmente
	r.dayData.SetUsage(usage)
	return nil
}

//Terei que trocar
func (r *Resident) GenerateDailyData(day uint8, rng *rand.Rand) {
	r.GenerateRoutine(day,rng)
	r.GenerateFrequency(day,rng)
	r.GenerateUsage(day,rng)
}
