package residentprofiles

import (
	"math/rand/v2"
	"residentdata"
)


type ResidentDayProfile struct {
	routineProfile *RoutineProfile
	frequencyProfileDay *FrequencyProfileDay
	
}

func NewResidentDayProfile(routine *RoutineProfile, frequency *FrequencyProfileDay) *ResidentDayProfile {
	return &ResidentDayProfile{
		routineProfile: routine,
		frequencyProfileDay: frequency,
	}
}

func (r *ResidentDayProfile) RoutineProfile() *RoutineProfile {
	return r.routineProfile
}

func (r *ResidentDayProfile) FrequencyProfileDay() *FrequencyProfileDay {
	return r.frequencyProfileDay
}

func (r *ResidentDayProfile) GenerateRoutine(rng *rand.Rand) *residentdata.Routine {
	return r.routineProfile.GenerateData(rng)
}

func (r *ResidentDayProfile) GenerateFrequency(rng *rand.Rand) *residentdata.Frequency {
	return r.frequencyProfileDay.GenerateData(rng)
}

