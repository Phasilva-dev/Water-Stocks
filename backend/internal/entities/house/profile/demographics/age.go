package demographics

import (
	"simulation/internal/dists"
	"math/rand/v2"
	"math"
)

type Age struct {
	dist dists.Distribution
}

func NewAge(dist dists.Distribution) *Age {
	return &Age{
		dist: dist,
	}
}

func (a *Age) AgeDist() dists.Distribution {
	return a.dist
}

func (a *Age) GenerateData(rng *rand.Rand) uint8 {
	sample := a.dist.Sample(rng)
	absSample := math.Abs(sample)

	if absSample > math.MaxUint8 {
		absSample = math.MaxUint8
	}

	return uint8(absSample)
}

