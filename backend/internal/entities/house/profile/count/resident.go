package count

import (
	"simulation/internal/dists"
	"math"
	"math/rand/v2"
)

type ResidentCount struct {

	dist dists.Distribution

}

func NewResidentCount(dist dists.Distribution) *ResidentCount {
	return &ResidentCount{
		dist: dist,
	}
}

func (r *ResidentCount) GenerateData(rng *rand.Rand) uint8 {
	sample := r.dist.Sample(rng)
	
	absolutValue := math.Abs(sample)
	processedValue := math.Ceil(absolutValue)
	

	if processedValue < 0 {
		processedValue = 0
	} else if processedValue > math.MaxUint8 {
		processedValue = 255
	}

	return uint8(processedValue)
}