package sanitarydevice

import (
	"dists"
	"math/rand/v2"
	"math"
)

type Tanque struct {
	flowLeakDist dists.Distribution
	durationDist dists.Distribution

}

func NewTanque(flowLeakDist, durationDist dists.Distribution, amount uint8) *Tanque {
	return &Tanque{
		flowLeakDist: flowLeakDist,
		durationDist: durationDist,
	}
}

func (t *Tanque) FlowLeakDist() dists.Distribution {
	return t.flowLeakDist
}

func (t *Tanque) DurationDist() dists.Distribution {
	return t.durationDist
}

func (t *Tanque) GenerateDuration(rng *rand.Rand) int32 {
	sample := t.durationDist.Sample(rng)
	absSample := math.Abs(sample)

	if absSample > math.MaxInt32 {
		absSample = math.MaxInt32
	}

	return int32(absSample)
}

func (t *Tanque) GenerateFlowLeak(rng *rand.Rand) int32 {
	sample := t.flowLeakDist.Sample(rng)
	absSample := math.Abs(sample)

	if absSample > math.MaxInt32 {
		absSample = math.MaxInt32
	}

	return int32(absSample)
}