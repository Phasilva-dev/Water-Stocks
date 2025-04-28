package sanitarydevice

import (
	"dists"
	"math/rand/v2"
	"math"
)

type DishWasher struct {
	sanitaryDeviceID uint16
	flowLeakDist dists.Distribution
	durationDist dists.Distribution
	

}

func NewDishWasher(flowLeakDist, durationDist dists.Distribution, id uint16) *DishWasher {
	return &DishWasher{
		sanitaryDeviceID: id,
		flowLeakDist: flowLeakDist,
		durationDist: durationDist,
	}
}

func (t *DishWasher) FlowLeakDist() dists.Distribution {
	return t.flowLeakDist
}

func (t *DishWasher) DurationDist() dists.Distribution {
	return t.durationDist
}

func (t *DishWasher) GenerateDuration(rng *rand.Rand) int32 {
	sample := t.durationDist.Sample(rng)
	absSample := math.Abs(sample)

	if absSample > math.MaxInt32 {
		absSample = math.MaxInt32
	}

	return int32(absSample)
}

func (t *DishWasher) GenerateFlowLeak(rng *rand.Rand) float64 {
	sample := t.flowLeakDist.Sample(rng)
	absSample := math.Abs(sample)

	return absSample
}