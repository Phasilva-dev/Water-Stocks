package residentprofiles

import (
	"math/rand/v2"
	"residentdata"
)

type FrequencyProfileDay struct {
	freqToilet      *FrequencyProfile
	freqShower      *FrequencyProfile
	freqWashBassin  *FrequencyProfile
	freqWashMachine *FrequencyProfile
	freqDishWasher  *FrequencyProfile
	freqTanque      *FrequencyProfile
}

func NewFrequencyProfileDay(profiles map[string]*FrequencyProfile) *FrequencyProfileDay {
	return &FrequencyProfileDay{
		freqToilet:      profiles["toilet"],
		freqShower:      profiles["shower"],
		freqWashBassin:  profiles["washBassin"],
		freqWashMachine: profiles["washMachine"],
		freqDishWasher:  profiles["dishWasher"],
		freqTanque:      profiles["tanque"],
	}
}

func (f *FrequencyProfileDay) FreqToilet() *FrequencyProfile {
	return f.freqToilet
}

func (f *FrequencyProfileDay) FreqShower() *FrequencyProfile {
	return f.freqShower
}

func (f *FrequencyProfileDay) FreqWashBassin() *FrequencyProfile {
	return f.freqWashBassin
}

func (f *FrequencyProfileDay) FreqWashMachine() *FrequencyProfile {
	return f.freqWashMachine
}

func (f *FrequencyProfileDay) FreqDishWasher() *FrequencyProfile {
	return f.freqDishWasher
}

func (f *FrequencyProfileDay) FreqTanque() *FrequencyProfile {
	return f.freqTanque
}

func validateFrequencyProfile(profile *FrequencyProfile, rng *rand.Rand) uint8 {
	if profile == nil {
		return 0
	}
	val := profile.GenerateData(rng)

	return val
}

func (f *FrequencyProfileDay) GenerateData(rng *rand.Rand) *residentdata.Frequency {
		freqToilet :=      validateFrequencyProfile(f.freqToilet, rng)
		freqShower :=      validateFrequencyProfile(f.freqShower, rng)
		freqWashBassin :=  validateFrequencyProfile(f.freqWashBassin, rng)
		freqWashMachine := validateFrequencyProfile(f.freqWashMachine, rng)
		freqDishWasher :=  validateFrequencyProfile(f.freqDishWasher, rng)
		freqTanque :=      validateFrequencyProfile(f.freqTanque, rng)
		return residentdata.NewFrequency(freqToilet, freqShower, freqWashBassin, freqWashMachine, freqDishWasher, freqTanque)
}