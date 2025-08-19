package sanitarydevice

import (
	"simulation/internal/dists"
	"math/rand/v2"
	"math"
	"fmt"
)

type WashBasin struct {
	sanitaryDeviceID uint32
	flowLeakDist dists.Distribution
	durationDist dists.Distribution
	

}
func newWashBasin(flowLeakDist, durationDist dists.Distribution,
	 id uint32) (*WashBasin, error) {
	if flowLeakDist == nil || durationDist == nil {
		return nil, fmt.Errorf("distributions cannot be nil")
	}
	if id == 0 {
		return nil, fmt.Errorf("zero is invalid id")
	}
	return &WashBasin{
		sanitaryDeviceID: id,
		flowLeakDist: flowLeakDist,
		durationDist: durationDist,
	}, nil
}

func (ks *WashBasin) IsCountable() bool {
	return true
}

func (sdi *WashBasin) SanitaryDeviceID() uint32 {
	return sdi.sanitaryDeviceID
}

func (sdi *WashBasin) FlowLeakDist() dists.Distribution {
	return sdi.flowLeakDist
}

func (sdi *WashBasin) DurationDist() dists.Distribution {
	return sdi.durationDist
}

func (sdi *WashBasin) GenerateDuration(rng *rand.Rand) int32 {
	sample := sdi.durationDist.Sample(rng)
	absSample := math.Abs(sample)

	if absSample > math.MaxInt32 {
		absSample = math.MaxInt32
	}

	return int32(absSample)
}

func (sdi *WashBasin) GenerateFlowLeak(rng *rand.Rand) float64 {
	sample := sdi.flowLeakDist.Sample(rng)
	absSample := math.Abs(sample)

	return absSample
}