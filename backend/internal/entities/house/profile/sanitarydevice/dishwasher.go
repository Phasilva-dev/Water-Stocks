package sanitarydevice

import (
	"simulation/internal/dists"
	"math/rand/v2"
	"math"
)

type DishWasher struct {
	sanitaryDeviceID uint32
	flowLeakDist dists.Distribution
	durationDist dists.Distribution
	

}

func NewDishWasher(flowLeakDist, durationDist dists.Distribution, id uint32) *DishWasher {
	return &DishWasher{
		sanitaryDeviceID: id,
		flowLeakDist: flowLeakDist,
		durationDist: durationDist,
	}
}

func (t *DishWasher) SanitaryDeviceID() uint32 {
	return t.sanitaryDeviceID
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