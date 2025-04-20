package residentprofiles

import (
	"residentdata"
	"math/rand/v2"
)

type UsageProfileDay struct {
	usageToilet      *UsageProfile
	usageShower      *UsageProfile
	usageWashBassin  *UsageProfile
	usageWashMachine *UsageProfile
	usageDishWasher  *UsageProfile
	usageTanque      *UsageProfile
}

func NewUsageProfileDay(profiles map[string]*UsageProfile) *UsageProfileDay {
	return &UsageProfileDay{
		usageToilet:      profiles["toilet"],
		usageShower:      profiles["shower"],
		usageWashBassin:  profiles["washBassin"],
		usageWashMachine: profiles["washMachine"],
		usageDishWasher:  profiles["dishWasher"],
		usageTanque:      profiles["tanque"],
	}
}

func (f *UsageProfileDay) UsageToilet() *UsageProfile {
	return f.usageToilet
}

func (f *UsageProfileDay) UsageShower() *UsageProfile {
	return f.usageShower
}

func (f *UsageProfileDay) UsageWashBassin() *UsageProfile {
	return f.usageWashBassin
}

func (f *UsageProfileDay) UsageWashMachine() *UsageProfile {
	return f.usageWashMachine
}

func (f *UsageProfileDay) UsageDishWasher() *UsageProfile {
	return f.usageDishWasher
}

func (f *UsageProfileDay) UsageTanque() *UsageProfile {
	return f.usageTanque
}

func validateUsageProfile(profile *UsageProfile, rng *rand.Rand) int32 {
	if profile == nil {
		return 0
	}
	val, err := profile.GenerateData(rng)
	if err != nil {
		return 0
	}
	return val
}

func (f *UsageProfileDay) GenerateData(rng *rand.Rand) *residentdata.Usage {
		usageToilet :=      validateUsageProfile(f.usageToilet, rng)
		usageShower :=      validateUsageProfile(f.usageShower, rng)
		usageWashBassin :=  validateUsageProfile(f.usageWashBassin, rng)
		usageWashMachine := validateUsageProfile(f.usageWashMachine, rng)
		usageDishWasher :=  validateUsageProfile(f.usageDishWasher, rng)
		usageTanque :=      validateUsageProfile(f.usageTanque, rng)
		return residentdata.NewUsage(usageToilet, usageShower, usageWashBassin, usageWashMachine, usageDishWasher, usageTanque)
}