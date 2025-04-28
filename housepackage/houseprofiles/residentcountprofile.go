package houseprofiles

import (
	"dists"
	"math"
	"math/rand/v2"
)

type ResidentCountProfile struct {

	residentCountDist dists.Distribution

}

func NewResidentCountProfile(dist dists.Distribution) *ResidentCountProfile {
	return &ResidentCountProfile{
		residentCountDist: dist,
	}
}

func (r *ResidentCountProfile) GenerateData(rng *rand.Rand) uint8 {
	sample := r.residentCountDist.Sample(rng)
	
	processedValue := math.Ceil(math.Abs(sample))

	if processedValue < 0 {
		processedValue = 0
	} else if processedValue > 255 {
		processedValue = 255
	}

	return uint8(processedValue)
}