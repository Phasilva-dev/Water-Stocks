package sanitarydevice

import (
	"dists"
	"math/rand/v2"
	"math"
)

type WashBassin struct {
	flowLeakDist dists.Distribution
	durationDist dists.Distribution
	amount uint8

}

func NewWashBassin(flowLeakDist, durationDist dists.Distribution, amount uint8) *WashBassin {
	return &WashBassin{
		flowLeakDist: flowLeakDist,
		durationDist: durationDist,
	}
}

func (t *WashBassin) FlowLeakDist() dists.Distribution {
	return t.flowLeakDist
}

func (t *WashBassin) DurationDist() dists.Distribution {
	return t.durationDist
}

func (t *WashBassin) GenerateDuration(rng *rand.Rand) int32 {
	sample := t.durationDist.Sample(rng)
	absSample := math.Abs(sample)

	if absSample > math.MaxInt32 {
		absSample = math.MaxInt32
	}

	return int32(absSample)
}

func (t *WashBassin) GenerateFlowLeak(rng *rand.Rand) int32 {
	sample := t.flowLeakDist.Sample(rng)
	absSample := math.Abs(sample)

	if absSample > math.MaxInt32 {
		absSample = math.MaxInt32
	}

	return int32(absSample)
}