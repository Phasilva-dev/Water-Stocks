package residentprofiles

import (
	"math/rand/v2"
	"residentdata"
)

type ResidentProfile struct {
	OccupationID uint32
	weeklyProfile *ResidentWeeklyProfile

}

//func NewResidentProfile() *ResidentProfile {
//
//}

func (r *ResidentProfile) GenerateFrequency(day uint8, rng *rand.Rand) *residentdata.Frequency {
	return r.weeklyProfile.GenerateFrequency(day,rng)
}

func (r *ResidentProfile) GenerateRoutine(day uint8, rng *rand.Rand) *residentdata.Routine {
	return r.weeklyProfile.GenerateRoutine(day,rng)
}



