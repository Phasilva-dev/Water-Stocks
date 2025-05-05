package interfaces

import (
	"math/rand/v2"
	"dists"
)

// ResidentProfile define o comportamento esperado para perfis de residentes
type ResidentProfile interface {
	GenerateFrequency(day uint8, rng *rand.Rand) Frequency
	GenerateRoutine(day uint8, rng *rand.Rand) Routine
}

type FrequencyProfile interface {
	// Métodos que representam as funções da struct FrequencyProfile
	Shift() uint8
	StatDist() dists.Distribution
	//NewFrequencyProfile(shift uint8, dist dists.Distribution) (FrequencyProfile, error)
	GenerateData(rng *rand.Rand) uint8
}

type FrequencyProfileDay interface {
	// Métodos que representam as funções da struct FrequencyProfileDay
	FreqToilet() FrequencyProfile
	FreqShower() FrequencyProfile
	FreqWashBassin() FrequencyProfile
	FreqWashMachine() FrequencyProfile
	FreqDishWasher() FrequencyProfile
	FreqTanque() FrequencyProfile
	GenerateData(rng *rand.Rand) Frequency
}

type RoutineProfile interface {

	Events() []dists.Distribution
	Shift() int32
	GenerateData(rng *rand.Rand) Routine
}

type ResidentDayProfile interface {
	RoutineProfile() *RoutineProfile
	FrequencyProfileDay() *FrequencyProfileDay
	GenerateRoutine(rng *rand.Rand) Routine
	GenerateFrequency(rng *rand.Rand) Frequency

}

type ResidentWeeklyProfile interface {
	GenerateFrequency(day uint8, rng *rand.Rand) Frequency
	GenerateRoutine(day uint8, rng *rand.Rand) Routine
}