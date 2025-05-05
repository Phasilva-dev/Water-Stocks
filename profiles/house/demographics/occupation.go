package demographics

import (
	"misc"
	"math/rand/v2"
)

type Occupation struct {
	under18Selector   *misc.PercentSelector[uint32]
	adultSelector     *misc.PercentSelector[uint32]
	over65Selector    *misc.PercentSelector[uint32]
}

func NewOccupation(
	under18 []misc.Tuple[uint32, float64],
	adult []misc.Tuple[uint32, float64],
	over65 []misc.Tuple[uint32, float64],
) (*Occupation, error) {
	u18Sel, err := misc.NewPercentSelector(under18)
	if err != nil {
		return nil, err
	}
	adultSel, err := misc.NewPercentSelector(adult)
	if err != nil {
		return nil, err
	}
	over65Sel, err := misc.NewPercentSelector(over65)
	if err != nil {
		return nil, err
	}

	return &Occupation{
		under18Selector: u18Sel,
		adultSelector:   adultSel,
		over65Selector:  over65Sel,
	}, nil
}

func (o *Occupation) Under18Selector() *misc.PercentSelector[uint32] {
	return o.under18Selector
}

func (o *Occupation) AdultSelector() *misc.PercentSelector[uint32] {
	return o.adultSelector
}

func (o *Occupation) Over65Selector() *misc.PercentSelector[uint32] {
	return o.over65Selector
}

func (o *Occupation) GenerateUnder18Selector(rng *rand.Rand) uint32 {
	id, err := o.under18Selector.Sample(rng)
	if err != nil {
		return 0
	}
	return id
}

func (o *Occupation) GenerateAdultSelector(rng *rand.Rand) uint32 {
	id, err := o.adultSelector.Sample(rng)
	if err != nil {
		return 0
	}
	return id
}

func (o *Occupation) GenerateOver65Selector(rng *rand.Rand) uint32 {
	id, err := o.over65Selector.Sample(rng)
	if err != nil {
		return 0
	}
	return id
}
