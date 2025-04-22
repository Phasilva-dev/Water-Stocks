package residentdata

import (
)

type Usage struct {
	//Uso individual
	usageToilet []int32
	usageShower []int32
	usageWashBassin []int32

	//Uso compartilhado
	usageWashMachine []int32
	usageDishWasher []int32
	usageTanque []int32

}

/*func NewUsage() *Usage{
	return &Usage{
		usageToilet: 0,
		usageShower: 0,
		usageWashBassin: 0,
		usageWashMachine: 0,
		usageDishWasher: 0,
		usageTanque: 0,
	}
}*/

func (f *Usage) UsageToilet() []int32 {
	return f.usageToilet
}

func (f *Usage) UsageShower() []int32 {
	return f.usageShower
}

func (f *Usage) UsageWashBassin() []int32 {
	return f.usageWashBassin
}

//Compartilhado

func (f *Usage) UsageWashMachine() []int32 {
	return f.usageWashMachine
}

func (f *Usage) UsageDishWasher() []int32 {
	return f.usageDishWasher
}

func (f *Usage) UsageTanque() []int32 {
	return f.usageTanque
}

func NewUsage(
	toilet, shower, washBassin, washMachine, dishWasher, tanque []int32,
) *Usage {
	return &Usage{
		usageToilet:      toilet,
		usageShower:      shower,
		usageWashBassin:  washBassin,
		usageWashMachine: washMachine,
		usageDishWasher:  dishWasher,
		usageTanque:      tanque,
	}
}