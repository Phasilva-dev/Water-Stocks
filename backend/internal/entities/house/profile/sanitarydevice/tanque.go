package sanitarydevice

import (
	"simulation/internal/dists"
	"math/rand/v2"
	"math"
	"errors"
)

type Tanque struct {
	sanitaryDeviceID uint32
	flowLeakDist dists.Distribution
	durationDist dists.Distribution

}

func NewTanque(flowLeakDist, durationDist dists.Distribution, id uint32) (*Tanque,error) {
	if flowLeakDist == nil || durationDist == nil {
		return nil, errors.New("distributions cannot be nil")
	}
	if id == 0 {
		return nil, errors.New("zero is invalid id")
	}
	return &Tanque{
		sanitaryDeviceID: id,
		flowLeakDist: flowLeakDist,
		durationDist: durationDist,
	},nil
}

func (t *Tanque) SanitaryDeviceID() uint32 {
	return t.sanitaryDeviceID
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

func (t *Tanque) GenerateFlowLeak(rng *rand.Rand) float64 {
	sample := t.flowLeakDist.Sample(rng)
	absSample := math.Abs(sample)


	return absSample
}