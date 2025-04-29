package models

import (
	"math/rand/v2"
	"residentdata"
)

// ResidentProfile define o comportamento esperado para perfis de residentes
type ResidentProfile interface {
	GenerateFrequency(day uint8, rng *rand.Rand) *residentdata.Frequency
	GenerateRoutine(day uint8, rng *rand.Rand) *residentdata.Routine
	GenerateUsage(day uint8, freq *residentdata.Frequency, rng *rand.Rand) (*residentdata.Usage, error)
}