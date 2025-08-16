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
	amount uint8
	

}
func newWashBasin(flowLeakDist, durationDist dists.Distribution,
	amount uint8, id uint32) (*WashBasin, error) {
	if flowLeakDist == nil || durationDist == nil {
		return nil, fmt.Errorf("distributions cannot be nil")
	}
	if id == 0 {
		return nil, fmt.Errorf("zero is invalid id")
	}
	if amount <= 0 {
		return nil, fmt.Errorf("")
	}
	return &WashBasin{
		sanitaryDeviceID: id,
		flowLeakDist: flowLeakDist,
		durationDist: durationDist,
		amount: amount,
	}, nil
}

func (sdi *WashBasin) SanitaryDeviceID() uint32 {
	return sdi.sanitaryDeviceID
}

func (d *WashBasin) Amount() uint8 {
	return d.amount
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