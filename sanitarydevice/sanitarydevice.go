package sanitarydevice

import (
	"math/rand/v2"
)

type SanitaryDevice interface {
	IntensitySample(rng *rand.Rand) float64
	DurationSample(rng *rand.Rand) int32
}


