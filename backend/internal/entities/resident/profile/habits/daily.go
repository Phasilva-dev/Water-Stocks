package habits

import (
	"math/rand/v2"
	"simulation/internal/entities/resident/ds/behavioral"
	"simulation/internal/entities/resident/profile/frequency"
	"simulation/internal/entities/resident/profile/routine"
)


type ResidentDayProfile struct {
	routineProfile *routine.RoutineProfile
	frequencyProfileDay *frequency.FrequencyProfileDay
	
}

func NewResidentDayProfile(routine *routine.RoutineProfile, frequency *frequency.FrequencyProfileDay) *ResidentDayProfile {
	return &ResidentDayProfile{
		routineProfile: routine,
		frequencyProfileDay: frequency,
	}
}

func (r *ResidentDayProfile) RoutineProfile() *routine.RoutineProfile {
	return r.routineProfile
}

func (r *ResidentDayProfile) FrequencyProfileDay() *frequency.FrequencyProfileDay {
	return r.frequencyProfileDay
}

func (r *ResidentDayProfile) GenerateRoutine(rng *rand.Rand) *behavioral.Routine {
	return r.routineProfile.GenerateData(rng)
}

func (r *ResidentDayProfile) GenerateFrequency(rng *rand.Rand) *behavioral.Frequency {
	return r.frequencyProfileDay.GenerateData(rng)
}

