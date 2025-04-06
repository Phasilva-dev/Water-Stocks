package residentdata

import (
)

type Frequency struct {
	freq uint8
}

func NewFrequency(freq uint8) *Frequency{
	return &Frequency{
		freq: freq,
	}
}