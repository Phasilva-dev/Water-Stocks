package demographics

import (
	"errors"
	"math/rand/v2"
	"simulation/internal/misc"
)

type AgeRangeSelector struct {
	minAge   uint8
	maxAge   uint8
	selector *misc.PercentSelector[uint32]
}

func (a *AgeRangeSelector) MinAge() uint8 {
	return a.minAge
}

func (a *AgeRangeSelector) MaxAge() uint8 {
	return a.maxAge
}

func (a *AgeRangeSelector) Selector() *misc.PercentSelector[uint32] {
	return a.selector
}

func NewAgeRangeSelector(minAge, maxAge uint8, selector *misc.PercentSelector[uint32]) (*AgeRangeSelector, error) {
	if selector == nil {
		return nil, errors.New("selector cannot be nil")
	}
	if minAge > maxAge {
		return nil, errors.New("minAge cannot be greater than maxAge")
	}
	return &AgeRangeSelector{
		minAge:   minAge,
		maxAge:   maxAge,
		selector: selector,
	}, nil
}


type Occupation struct {
	selectors []*AgeRangeSelector
}

func NewOccupation(selectors []*AgeRangeSelector) (*Occupation, error) {
	if len(selectors) == 0 {
		return nil, errors.New("selectors cannot be empty")
	}

	// Verificar sobreposição
	for i := 0; i < len(selectors); i++ {
		selA := selectors[i]
		if selA == nil {
			return nil, errors.New("selectors cannot be nil")
		}
		for j := i + 1; j < len(selectors); j++ {
			selB := selectors[j]
			if selB == nil {
				return nil, errors.New("selectors cannot be nil")
			}

			if rangesOverlap(selA.MinAge(), selA.MaxAge(), selB.MinAge(), selB.MaxAge()) {
				return nil, errors.New(
					"age ranges overlap between selectors")
			}
		}
	}

	return &Occupation{selectors: selectors}, nil
}

func rangesOverlap(minA, maxA, minB, maxB uint8) bool {
	return minA <= maxB && minB <= maxA
}

func (o *Occupation) Selectors() []*AgeRangeSelector {
	return o.selectors
}

func (o *Occupation) Sample(age uint8, rng *rand.Rand) (uint32, error) {
	for _, sel := range o.selectors {
		if age >= sel.MinAge() && age <= sel.MaxAge() {
			return sel.Selector().Sample(rng)
		}
	}
	return 0, errors.New("no selector found for age")
}