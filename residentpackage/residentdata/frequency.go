package residentdata

import (
)

type Frequency struct {
	//Uso individual
	freqToilet uint8
	freqShower uint8
	freqWashBassin uint8

	//Uso compartilhado
	freqWashMachine uint8
	freqDishWasher uint8
	freqTanque uint8

}

/*func NewFrequency() *Frequency{
	return &Frequency{
		freqToilet: 0,
		freqShower: 0,
		freqWashBassin: 0,
		freqWashMachine: 0,
		freqDishWasher: 0,
		freqTanque: 0,
	}
}*/

func (f *Frequency) FreqToilet() uint8 {
	return f.freqToilet
}

func (f *Frequency) FreqShower() uint8 {
	return f.freqShower
}

func (f *Frequency) FreqWashBassin() uint8 {
	return f.freqWashBassin
}

//Compartilhado

func (f *Frequency) FreqWashMachine() uint8 {
	return f.freqWashMachine
}

func (f *Frequency) FreqDishWasher() uint8 {
	return f.freqDishWasher
}

func (f *Frequency) FreqTanque() uint8 {
	return f.freqTanque
}

func NewFrequency(
	toilet, shower, washBassin, washMachine, dishWasher, tanque uint8,
) *Frequency {
	return &Frequency{
		freqToilet:      toilet,
		freqShower:      shower,
		freqWashBassin:  washBassin,
		freqWashMachine: washMachine,
		freqDishWasher:  dishWasher,
		freqTanque:      tanque,
	}
}