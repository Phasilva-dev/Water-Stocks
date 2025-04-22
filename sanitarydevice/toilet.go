package sanitarydevice


type Toilet struct {
	flowLeak int32
	duration int32
	amount uint8

}

func NewToilet(flowLeak, duration int32, amount uint8) *Toilet {
	return &Toilet{
		flowLeak: flowLeak,
		duration: duration,
		amount: amount,
	}
}

func (t *Toilet) FlowLeak() int32 {
	return t.flowLeak
}

func (t *Toilet) Duration() int32 {
	return t.duration
}

/*

Caso precise criar distribuições para isso, aqui está uma forma de fazer isso

func (t *Toilet) GenerateDuration(rng *rand.Rand) int32 {
	sample := t.durationDist.Sample(rng)
	absSample := math.Abs(sample)

	if absSample > math.MaxInt32 {
		absSample = math.MaxInt32
	}

	return int32(absSample)
}

func (t *Toilet) GenerateFlowLeak(rng *rand.Rand) int32 {
	sample := t.flowLeakDist.Sample(rng)
	absSample := math.Abs(sample)

	if absSample > math.MaxInt32 {
		absSample = math.MaxInt32
	}

	return int32(absSample)
}*/