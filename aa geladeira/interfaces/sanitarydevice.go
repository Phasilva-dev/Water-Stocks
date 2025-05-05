package interfaces

import (
	"math/rand/v2"
	"dists"
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

type DishWasher interface {
	SanitaryDeviceID() uint32
	FlowLeakDist() dists.Distribution
	DurationDist() dists.Distribution
	GenerateDuration(rng *rand.Rand) int32
	GenerateFlowLeak(rng *rand.Rand) float64
}

type Shower interface {
	SanitaryDeviceID() uint32
	FlowLeakDist() dists.Distribution
	DurationDist() dists.Distribution
	GenerateDuration(rng *rand.Rand) int32
	GenerateFlowLeak(rng *rand.Rand) float64
}

type Tanque interface {
	SanitaryDeviceID() uint32
	FlowLeakDist() dists.Distribution
	DurationDist() dists.Distribution
	GenerateDuration(rng *rand.Rand) int32
	GenerateFlowLeak(rng *rand.Rand) float64
}


type Toilet interface {
	SanitaryDeviceID() uint32
	FlowLeak() float64
	Duration() int32
	GenerateDuration(rng *rand.Rand) int32
	GenerateFlowLeak(rng *rand.Rand) float64
}

type WashBassin interface {
	SanitaryDeviceID() uint32
	FlowLeakDist() dists.Distribution
	DurationDist() dists.Distribution
	GenerateDuration(rng *rand.Rand) int32
	GenerateFlowLeak(rng *rand.Rand) float64
}

type WashMachine interface {
	SanitaryDeviceID() uint32
	FlowLeak() float64
	Duration() int32
	GenerateDuration(rng *rand.Rand) int32
	GenerateFlowLeak(rng *rand.Rand) float64
}