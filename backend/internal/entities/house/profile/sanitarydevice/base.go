package sanitarydevice

import (
	"fmt"
	"math/rand/v2"
	"simulation/internal/dists"
	"math"
)


type BaseDevice struct {
	sanitaryDeviceID uint32
	flowLeakDist dists.Distribution
	durationDist dists.Distribution

}
func newBaseDevice(flowLeakDist, durationDist dists.Distribution,
	 id uint32) (*BaseDevice, error) {
	if flowLeakDist == nil || durationDist == nil {
		return nil, fmt.Errorf("distributions cannot be nil")
	}
	if id == 0 {
		return nil, fmt.Errorf("zero is invalid id")
	}
	return &BaseDevice{
		sanitaryDeviceID: id,
		flowLeakDist: flowLeakDist,
		durationDist: durationDist,
	}, nil
}

func (ks *BaseDevice) IsCountable() bool {
	return false
}

func (g *BaseDevice) SanitaryDeviceID() uint32 {
	return g.sanitaryDeviceID
}

func (g *BaseDevice) FlowLeakDist() dists.Distribution {
	return g.flowLeakDist
}

func (g *BaseDevice) DurationDist() dists.Distribution {
	return g.durationDist
}

func (g *BaseDevice) GenerateDuration(rng *rand.Rand) int32 {
	sample := g.durationDist.Sample(rng)
	absSample := math.Abs(sample)

	if absSample > math.MaxInt32 {
		absSample = math.MaxInt32
	}

	return int32(absSample)
}

func (g *BaseDevice) GenerateFlowLeak(rng *rand.Rand) float64 {
	sample := g.flowLeakDist.Sample(rng)
	absSample := math.Abs(sample)

	return absSample
}