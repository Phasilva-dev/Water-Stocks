package profiles

import (
	"datastruct"
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

func (f *FrequencyProfile) NewFrequencyProfile(symbol string,shift uint8, dist dists.Distribution) (*FrequencyProfile, error){

	if symbol == "" {
		return nil, errors.New("Simbolo nao pode estar vazio")
	}
	if dist == nil {
		return nil, errors.New("Distribuicao esta vazia")
	}
	if shift < 0{
		return nil, errors.New("Constante deve ser positiva")
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

func (f *FrequencyProfile) GenerateData(rng *rand.Rand) *datastruct.Frequency{
	freq := f.generateFrequency(rng)
	return datastruct.NewFrequency(freq)
}

