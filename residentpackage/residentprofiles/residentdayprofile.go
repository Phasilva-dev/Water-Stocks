package residentprofiles

import (
	"math/rand/v2"
	"residentdata"
)


type ResidentDayProfile struct {
	routineProfile *RoutineProfile
	frequencyProfileDay *FrequencyProfileDay
	usageProfileDay *UsageProfileDay
	
}

func NewResidentDayProfile(routine *RoutineProfile, frequency *FrequencyProfileDay, usage *UsageProfileDay) *ResidentDayProfile {
	return &ResidentDayProfile{
		routineProfile: routine,
		frequencyProfileDay: frequency,
		usageProfileDay: usage,
	}
}

func (r *ResidentDayProfile) RoutineProfile() *RoutineProfile {
	return r.routineProfile
}

func (r *ResidentDayProfile) FrequencyProfileDay() *FrequencyProfileDay {
	return r.frequencyProfileDay
}

func (r *ResidentDayProfile) UsageProfileDay() *UsageProfileDay {
	return r.usageProfileDay
}

func (r *ResidentDayProfile) GenerateRoutine(rng *rand.Rand) *residentdata.Routine {
	return r.routineProfile.GenerateData(rng)
}

func (r *ResidentDayProfile) GenerateFrequency(rng *rand.Rand) *residentdata.Frequency {
	return r.frequencyProfileDay.GenerateData(rng)
}

func (r *ResidentDayProfile) GenerateUsage(freq *residentdata.Frequency,rng *rand.Rand) *residentdata.Usage {
	return r.usageProfileDay.GenerateData(rng, freq)
}
