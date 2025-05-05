package sanitarydevice

import (
	"math/rand/v2"
)

type Toilet struct {
	sanitaryDeviceID uint32
	flowLeak float64
	duration int32

}

func NewToilet(flowLeak float64, duration int32, id uint32) *Toilet {
	return &Toilet{
		sanitaryDeviceID: id,
		flowLeak: flowLeak,
		duration: duration,
	}
}

func (t *Toilet) SanitaryDeviceID() uint32 {
	return t.sanitaryDeviceID
}

func (t *Toilet) FlowLeak() float64 {
	return t.flowLeak
}

func (t *Toilet) Duration() int32 {
	return t.duration
}

func (t *Toilet) GenerateDuration(rng *rand.Rand) int32 {
	return t.duration
}

func (t *Toilet) GenerateFlowLeak(rng *rand.Rand) float64 {
	return t.flowLeak
}