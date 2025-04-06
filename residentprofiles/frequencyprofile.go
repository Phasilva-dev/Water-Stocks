package residentprofiles

import (
	"residentdata"
	"dists"
	"errors"
	"math/rand/v2"

)

type FrequencyProfile struct {
	statDist dists.Distribution
	shift uint8
}


func (f *FrequencyProfile) Shift() uint8 {
	return f.shift
}

func (f *FrequencyProfile) StatDist() dists.Distribution {
	return f.statDist
}

func (f *FrequencyProfile) NewFrequencyProfile(shift uint8, dist dists.Distribution) (*FrequencyProfile, error){

	if dist == nil {
		return nil, errors.New("distribution is null")
	}

	return &FrequencyProfile {
		statDist: dist,
		shift: shift,
	}, nil
}

func (f *FrequencyProfile) generateFrequency(rng *rand.Rand) uint8 {
	freq := f.statDist.Sample(rng)

	// Garantir que freq não seja negativo
	if freq < 0 {
			freq = 0
	}

	roundedFreq := uint8(freq)

	if roundedFreq < f.shift {
			return f.shift
	}
	
	return roundedFreq
}

func (f *FrequencyProfile) GenerateData(rng *rand.Rand) *residentdata.Frequency{
	freq := f.generateFrequency(rng)
	return residentdata.NewFrequency(freq)
}