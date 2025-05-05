package sanitarydevice

import (
	"dists"
	"math/rand/v2"
	"math"
	"interfaces"
)

type Shower struct {
	sanitaryDeviceID uint32
	flowLeakDist dists.Distribution
	durationDist dists.Distribution

}

func NewShower(flowLeakDist, durationDist dists.Distribution, id uint32) interfaces.Shower {
	return &Shower{
		sanitaryDeviceID: id,
		flowLeakDist: flowLeakDist,
		durationDist: durationDist,
	}
}

func (t *Shower) SanitaryDeviceID() uint32 {
	return t.sanitaryDeviceID
}

func (t *Shower) FlowLeakDist() dists.Distribution {
	return t.flowLeakDist
}

func (t *Shower) DurationDist() dists.Distribution {
	return t.durationDist
}


func (t *Shower) GenerateDuration(rng *rand.Rand) int32 {
	sample := t.durationDist.Sample(rng)
	absSample := math.Abs(sample)

	if absSample > math.MaxInt32 {
		absSample = math.MaxInt32
	}

	return int32(absSample)
}

func (t *Shower) GenerateFlowLeak(rng *rand.Rand) float64 {
	sample := t.flowLeakDist.Sample(rng)
	absSample := math.Abs(sample)

	return absSample
}