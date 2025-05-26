package sanitarydevice

import (
	"math/rand/v2"
	"errors"
)

type WashMachine struct {
	sanitaryDeviceID uint32
	flowLeak float64
	duration int32

}

func NewWashMachine(flowLeak float64, duration int32, id uint32) (*WashMachine, error) {

	if flowLeak <= 0 || duration <= 0 {
		return nil, errors.New("flowleak and duration cannot be 0 or negative")
	}
	if id == 0 {
		return nil, errors.New("zero is invalid id")
	}

	return &WashMachine{
		sanitaryDeviceID: id,
		flowLeak: flowLeak,
		duration: duration,
	},nil
}

func (t *WashMachine) SanitaryDeviceID() uint32 {
	return t.sanitaryDeviceID
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