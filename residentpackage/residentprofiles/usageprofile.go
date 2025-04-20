package residentprofiles

import (
	"misc"
	"dists"
	"math/rand/v2"
)

type UsageProfile struct {
	usagesSelector *misc.PercentSelector[dists.Distribution]
}

func NewUsageProfile(usagesSlice []misc.Tuple[dists.Distribution, float64]) (*UsageProfile, error) {
	selector, err := misc.NewPercentSelector(usagesSlice)
	if err != nil {
		return nil, err
	}
	
	return &UsageProfile{
		usagesSelector: selector,
	}, nil
}

func (u *UsageProfile) GenerateData(rng *rand.Rand) (int32, error) {
	dist, err := u.usagesSelector.Sample(rng)
	if err != nil {
		return 0, err
	}
	return int32(dist.Sample(rng)),nil

}