package interfaces

import (
	"math/rand/v2"
)

type SanitaryDevice interface {
	GenerateFlowLeak(rng *rand.Rand) float64
	GenerateDuration(rng *rand.Rand) int32
	SanitaryDeviceID() uint32
}

type SanitaryDeviceInstance interface {
	Device() SanitaryDevice
	Amount() uint8
}