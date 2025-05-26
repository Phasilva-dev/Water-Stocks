package sanitarydevice

import (
	"math/rand/v2"
	"errors"
)

type Toilet struct {
	sanitaryDeviceID uint32
	flowLeak float64
	duration int32

}

func NewToilet(flowLeak float64, duration int32, id uint32) (*Toilet,error) {

	if flowLeak <= 0 || duration <= 0 {
		return nil, errors.New("flowleak and duration cannot be 0 or negative")
	}
	if id == 0 {
		return nil, errors.New("zero is invalid id")
	}

	return &Toilet{
		sanitaryDeviceID: id,
		flowLeak: flowLeak,
		duration: duration,
	},nil
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