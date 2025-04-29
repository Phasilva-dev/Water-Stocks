package models

import (
	"math/rand/v2"
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