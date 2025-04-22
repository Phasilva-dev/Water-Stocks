package houseprofiles

import (
	"dists"
	"math/rand/v2"
	"math"
)

type AgeProfile struct {
	ageDist dists.Distribution
}

func (a *AgeProfile) AgeDist() dists.Distribution {
	return a.ageDist
}

func (a *AgeProfile) GenerateData(rng *rand.Rand) uint32 {
	sample := a.ageDist.Sample(rng)
	absSample := math.Abs(sample)

	if absSample > math.MaxInt32 {
		absSample = math.MaxInt32
	}

	return uint32(absSample)
}

