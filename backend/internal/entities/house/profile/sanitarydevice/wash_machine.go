package sanitarydevice

import (
	"simulation/internal/dists"
	"math/rand/v2"
	"math"
	"fmt"
)

type WashMachine struct {
	sanitaryDeviceID uint32
	flowLeakDist dists.Distribution
	durationDist dists.Distribution
	amount uint8
	

}
func newWashMachine(flowLeakDist, durationDist dists.Distribution,
	amount uint8, id uint32) (*WashMachine, error) {
	if flowLeakDist == nil || durationDist == nil {
		return nil, fmt.Errorf("distributions cannot be nil")
	}
	if id == 0 {
		return nil, fmt.Errorf("zero is invalid id")
	}
	if amount <= 0 {
		return nil, fmt.Errorf("")
	}
	return &WashMachine{
		sanitaryDeviceID: id,
		flowLeakDist: flowLeakDist,
		durationDist: durationDist,
		amount: amount,
	}, nil
}

func (sdi *WashMachine) SanitaryDeviceID() uint32 {
	return sdi.sanitaryDeviceID
}

func (d *WashMachine) Amount() uint8 {
	return d.amount
}

func (sdi *WashMachine) FlowLeakDist() dists.Distribution {
	return sdi.flowLeakDist
}

func (sdi *WashMachine) DurationDist() dists.Distribution {
	return sdi.durationDist
}

func (sdi *WashMachine) GenerateDuration(rng *rand.Rand) int32 {
	sample := sdi.durationDist.Sample(rng)
	absSample := math.Abs(sample)

	if absSample > math.MaxInt32 {
		absSample = math.MaxInt32
	}

	return int32(absSample)
}

func (sdi *WashMachine) GenerateFlowLeak(rng *rand.Rand) float64 {
	sample := sdi.flowLeakDist.Sample(rng)
	absSample := math.Abs(sample)

	return absSample
}
