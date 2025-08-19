package sanitarydevice

import (
	"fmt"
	"math/rand/v2"
	"simulation/internal/dists"
)

type SanitaryDevice interface {
	GenerateFlowLeak(rng *rand.Rand) float64
	GenerateDuration(rng *rand.Rand) int32
	SanitaryDeviceID() uint32
	FlowLeakDist() dists.Distribution
	DurationDist() dists.Distribution
	IsCountable() bool
}

// DeviceFactory cria qualquer dispositivo sanitário com base no tipo
func CreateSanitaryDevice(deviceType string, flowLeakDist,
	 durationDist dists.Distribution, id uint32) (SanitaryDevice, error) {
	switch deviceType {
	case "toilet":
		return newToilet(flowLeakDist, durationDist, id)
	case "shower":
		return newShower(flowLeakDist, durationDist, id)
	case "wash_basin":
		return newWashBasin(flowLeakDist, durationDist, id)
	case "wash_machine":
		return newWashMachine(flowLeakDist, durationDist, id)
	case "kitchen_sink":
		return newKitchenSink(flowLeakDist, durationDist, id)
	case "laundry_sink":
		return newLaundrySink(flowLeakDist, durationDist, id)
	case "base":
		return newBaseDevice(flowLeakDist, durationDist, id)
	default:
		return nil, fmt.Errorf("invalid SanitaryDevice Factory: unknown device type: %s", deviceType)
	}
}



