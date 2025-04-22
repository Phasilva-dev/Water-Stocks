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

func (f *UsageProfileDay) GenerateData(rng *rand.Rand, freq *residentdata.Frequency) *residentdata.Usage {
	// Armazena as frequências em variáveis locais
	freqToilet := freq.FreqToilet()
	freqShower := freq.FreqShower()
	freqWashBassin := freq.FreqWashBassin()
	freqWashMachine := freq.FreqWashMachine()
	freqDishWasher := freq.FreqDishWasher()
	freqTanque := freq.FreqTanque()

	// Inicializa slices com capacidade baseada na frequência
	usageToilet := make([]int32, 0, freqToilet)
	usageShower := make([]int32, 0, freqShower)
	usageWashBassin := make([]int32, 0, freqWashBassin)
	usageWashMachine := make([]int32, 0, freqWashMachine)
	usageDishWasher := make([]int32, 0, freqDishWasher)
	usageTanque := make([]int32, 0, freqTanque)

	// Usa loops com uint8
	for i := uint8(0); i < freqToilet; i++ {
		if val, err := f.usageToilet.GenerateData(rng); err == nil {
			usageToilet = append(usageToilet, val)
		}
	}
	for i := uint8(0); i < freqShower; i++ {
		if val, err := f.usageShower.GenerateData(rng); err == nil {
			usageShower = append(usageShower, val)
		}
	}
	for i := uint8(0); i < freqWashBassin; i++ {
		if val, err := f.usageWashBassin.GenerateData(rng); err == nil {
			usageWashBassin = append(usageWashBassin, val)
		}
	}
	for i := uint8(0); i < freqWashMachine; i++ {
		if val, err := f.usageWashMachine.GenerateData(rng); err == nil {
			usageWashMachine = append(usageWashMachine, val)
		}
	}
	for i := uint8(0); i < freqDishWasher; i++ {
		if val, err := f.usageDishWasher.GenerateData(rng); err == nil {
			usageDishWasher = append(usageDishWasher, val)
		}
	}
	for i := uint8(0); i < freqTanque; i++ {
		if val, err := f.usageTanque.GenerateData(rng); err == nil {
			usageTanque = append(usageTanque, val)
		}
	}

	return residentdata.NewUsage(usageToilet,
		usageShower,
		usageWashBassin,
		usageWashMachine,
		usageDishWasher,
		usageTanque,)
	
}
