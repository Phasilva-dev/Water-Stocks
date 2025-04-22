package sanitarydevice

import (
	"dists"
	"math/rand/v2"
	"math"
)

type Shower struct {
	flowLeakDist dists.Distribution
	durationDist dists.Distribution
	amount uint8

}

func NewShower(flowLeakDist, durationDist dists.Distribution, amount uint8) *Shower {
	return &Shower{
		flowLeakDist: flowLeakDist,
		durationDist: durationDist,
		amount: amount,
	}
}

func (t *Shower) FlowLeakDist() dists.Distribution {
	return t.flowLeakDist
}

func (t *Shower) DurationDist() dists.Distribution {
	return t.durationDist
}

func (t *Shower) Amount() uint8 {
	return t.amount
}

func (t *Shower) GenerateDuration(rng *rand.Rand) int32 {
	sample := t.durationDist.Sample(rng)
	absSample := math.Abs(sample)

	if absSample > math.MaxInt32 {
		absSample = math.MaxInt32
	}

	return int32(absSample)
}

func (t *Shower) GenerateFlowLeak(rng *rand.Rand) int32 {
	sample := t.flowLeakDist.Sample(rng)
	absSample := math.Abs(sample)

	if absSample > math.MaxInt32 {
		absSample = math.MaxInt32
	}

	return int32(absSample)
}