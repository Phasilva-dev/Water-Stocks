package sanitarydevice


type WashMachine struct {
	flowLeak int32
	duration int32

}

func NewWashMachine(flowLeak, duration int32) *WashMachine {
	return &WashMachine{
		flowLeak: flowLeak,
		duration: duration,
	}
}

func (t *WashMachine) FlowLeak() int32 {
	return t.flowLeak
}

func (t *WashMachine) Duration() int32 {
	return t.duration
}

/*

Caso precise criar distribuições para isso, aqui está uma forma de fazer isso

func (t *WashMachine) GenerateDuration(rng *rand.Rand) int32 {
	sample := t.durationDist.Sample(rng)
	absSample := math.Abs(sample)

	if absSample > math.MaxInt32 {
		absSample = math.MaxInt32
	}

	return int32(absSample)
}

func (t *WashMachine) GenerateFlowLeak(rng *rand.Rand) int32 {
	sample := t.flowLeakDist.Sample(rng)
	absSample := math.Abs(sample)

	if absSample > math.MaxInt32 {
		absSample = math.MaxInt32
	}

	return int32(absSample)
}*/