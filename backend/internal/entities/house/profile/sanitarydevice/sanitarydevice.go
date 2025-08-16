package sanitarydevice

import (
	"math/rand/v2"
	"simulation/internal/dists"
	"fmt"
	"math"
)

type SanitaryDevice interface {
	GenerateFlowLeak(rng *rand.Rand) float64
	GenerateDuration(rng *rand.Rand) int32
	SanitaryDeviceID() uint32
	FlowLeakDist() dists.Distribution
	DurationDist() dists.Distribution
	Amount() uint8
}

// DeviceFactory cria qualquer dispositivo sanitário com base no tipo
func DeviceFactory(deviceType string, flowLeakDist,
	 durationDist dists.Distribution, amount uint8, id uint32) (SanitaryDevice, error) {
	if flowLeakDist == nil || durationDist == nil {
		return nil, fmt.Errorf("distributions cannot be nil")
	}
	if id == 0 {
		return nil, fmt.Errorf("zero is invalid id")
	}
	if amount == 0 {
		return nil, fmt.Errorf("amount must be > 0")
	}

	switch deviceType {
	case "toilet":
		return NewToilet(flowLeakDist, durationDist, amount, id)
	case "shower":
		return NewShower(flowLeakDist, durationDist, amount, id)
	case "wash_basin":
		return NewWashBasin(flowLeakDist, durationDist, amount, id)
	case "wash_machine":
		return NewWashMachine(flowLeakDist, durationDist, amount, id)
	case "kitchen_sink":
		return NewKitchenSink(flowLeakDist, durationDist, amount, id)
	case "laundry_sink":
		return NewLaundrySink(flowLeakDist, durationDist, amount, id)
	case "generic":
		return NewSanitaryDeviceInstance(flowLeakDist, durationDist, amount, id)
	default:
		return nil, fmt.Errorf("unknown device type: %s", deviceType)
	}
}

type SanitaryDeviceInstance struct {
	sanitaryDeviceID uint32
	flowLeakDist dists.Distribution
	durationDist dists.Distribution
	amount uint8
	

}
func NewSanitaryDeviceInstance(flowLeakDist, durationDist dists.Distribution,
	amount uint8, id uint32) (*SanitaryDeviceInstance, error) {
	if flowLeakDist == nil || durationDist == nil {
		return nil, fmt.Errorf("distributions cannot be nil")
	}
	if id == 0 {
		return nil, fmt.Errorf("zero is invalid id")
	}
	if amount <= 0 {
		return nil, fmt.Errorf("")
	}
	return &SanitaryDeviceInstance{
		sanitaryDeviceID: id,
		flowLeakDist: flowLeakDist,
		durationDist: durationDist,
		amount: amount,
	}, nil
}

func (sdi *SanitaryDeviceInstance) SanitaryDeviceID() uint32 {
	return sdi.sanitaryDeviceID
}

func (d *SanitaryDeviceInstance) Amount() uint8 {
	return d.amount
}

func (sdi *SanitaryDeviceInstance) FlowLeakDist() dists.Distribution {
	return sdi.flowLeakDist
}

func (sdi *SanitaryDeviceInstance) DurationDist() dists.Distribution {
	return sdi.durationDist
}

func (sdi *SanitaryDeviceInstance) GenerateDuration(rng *rand.Rand) int32 {
	sample := sdi.durationDist.Sample(rng)
	absSample := math.Abs(sample)

	if absSample > math.MaxInt32 {
		absSample = math.MaxInt32
	}

	return int32(absSample)
}

func (sdi *SanitaryDeviceInstance) GenerateFlowLeak(rng *rand.Rand) float64 {
	sample := sdi.flowLeakDist.Sample(rng)
	absSample := math.Abs(sample)

	return absSample
}

