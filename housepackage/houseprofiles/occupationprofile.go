package houseprofiles

import (
	"misc"
	"math/rand/v2"
)

type OccupationProfile struct {
	under18Selector   *misc.PercentSelector[uint32]
	adultSelector     *misc.PercentSelector[uint32]
	over65Selector    *misc.PercentSelector[uint32]
}

func NewOccupationProfile(
	under18 []misc.Tuple[uint32, float64],
	adult []misc.Tuple[uint32, float64],
	over65 []misc.Tuple[uint32, float64],
) (*OccupationProfile, error) {
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

	return &OccupationProfile{
		under18Selector: u18Sel,
		adultSelector:   adultSel,
		over65Selector:  over65Sel,
	}, nil
}

func (o *OccupationProfile) Under18Selector() *misc.PercentSelector[uint32] {
	return o.under18Selector
}

func (o *OccupationProfile) AdultSelector() *misc.PercentSelector[uint32] {
	return o.adultSelector
}

func (o *OccupationProfile) Over65Selector() *misc.PercentSelector[uint32] {
	return o.over65Selector
}

func (o *OccupationProfile) GenerateUnderSelector(rng *rand.Rand) uint32 {
	id, err := o.under18Selector.Sample(rng)
	if err != nil {
		return 0
	}
	return id
}

func (o *OccupationProfile) GenerateAdultSelector(rng *rand.Rand) uint32 {
	id, err := o.adultSelector.Sample(rng)
	if err != nil {
		return 0
	}
	return id
}

func (o *OccupationProfile) GenerateOverSelector(rng *rand.Rand) uint32 {
	id, err := o.over65Selector.Sample(rng)
	if err != nil {
		return 0
	}
	return id
}
