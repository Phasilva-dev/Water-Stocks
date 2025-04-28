package sanitarydevice

import (
	"math/rand/v2"
)

type WashMachine struct {
	sanitaryDeviceID uint16
	flowLeak float64
	duration int32

}

func NewWashMachine(flowLeak float64, duration int32, id uint16) *WashMachine {
	return &WashMachine{
		sanitaryDeviceID: id,
		flowLeak: flowLeak,
		duration: duration,
	}
}

func (t *WashMachine) FlowLeak() float64 {
	return t.flowLeak
}

func (t *WashMachine) Duration() int32 {
	return t.duration
}

func (t *WashMachine) GenerateDuration(rng *rand.Rand) int32 {
	return t.duration
}

func (t *WashMachine) GenerateFlowLeak(rng *rand.Rand) float64 {
	return t.flowLeak
}