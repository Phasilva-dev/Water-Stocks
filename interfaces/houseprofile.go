package interfaces

import (
	"math/rand/v2"
	"dists"
	"misc"
)

// Interface que descreve o comportamento de um HouseProfile
type HouseProfile interface {
	GenerateNumbersOfResidents(rng *rand.Rand) uint8
	GenerateAgeofResidents(rng *rand.Rand) uint8
	GenerateOccupation(age uint8, rng *rand.Rand) uint32
	GenerateNumberOfSanitaryDevices(rng *rand.Rand, numberOfResidents uint8) error
	GetNumberOfSanitaryDevices() uint8
	GenerateSanitaryHouse(devices map[string]uint32, amount uint8) error
}

type AgeProfile interface {
	AgeDist() dists.Distribution
	GenerateData(rng *rand.Rand) uint32
}

type OccupationProfile interface {
	Under18Selector() *misc.PercentSelector[uint32]
	AdultSelector() *misc.PercentSelector[uint32]
	Over65Selector() *misc.PercentSelector[uint32]

	GenerateUnderSelector(rng *rand.Rand) uint32
	GenerateAdultSelector(rng *rand.Rand) uint32
	GenerateOverSelector(rng *rand.Rand) uint32
}

type ResidentCountProfile interface {
	GenerateData(rng *rand.Rand) uint8
}

type SanitaryCountProfile interface {
	GenerateData(rng *rand.Rand, numResidents uint8) (uint8, error)
}