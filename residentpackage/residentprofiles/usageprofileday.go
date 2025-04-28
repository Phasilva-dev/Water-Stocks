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

func (f *UsageProfileDay) GenerateData(rng *rand.Rand, freq *residentdata.Frequency) (*residentdata.Usage, error) {
	usageToilet, err := f.usageToilet.GenerateData(rng, freq.FreqToilet())
	if err != nil {
		return nil, err
	}

	usageShower, err := f.usageShower.GenerateData(rng, freq.FreqShower())
	if err != nil {
		return nil, err
	}

	usageWashBassin, err := f.usageWashBassin.GenerateData(rng, freq.FreqWashBassin())
	if err != nil {
		return nil, err
	}

	usageWashMachine, err := f.usageWashMachine.GenerateData(rng, freq.FreqWashMachine())
	if err != nil {
		return nil, err
	}

	usageDishWasher, err := f.usageDishWasher.GenerateData(rng, freq.FreqDishWasher())
	if err != nil {
		return nil, err
	}

	usageTanque, err := f.usageTanque.GenerateData(rng, freq.FreqTanque())
	if err != nil {
		return nil, err
	}

	return residentdata.NewUsage(
		usageToilet,
		usageShower,
		usageWashBassin,
		usageWashMachine,
		usageDishWasher,
		usageTanque,
	), nil
}
