package demographics

import (
	"math/rand/v2"
	"simulation/internal/misc"

	"errors"
)

type Occupation struct {
	under18Selector   *misc.PercentSelector[uint32]
	adultSelector     *misc.PercentSelector[uint32]
	over65Selector    *misc.PercentSelector[uint32]
}

func NewOccupation(
	under18 *misc.PercentSelector[uint32],
	adult *misc.PercentSelector[uint32],
	over65 *misc.PercentSelector[uint32],
) (*Occupation, error) {
	// Adicione a validação para seletores nil aqui
	if under18 == nil || adult == nil || over65 == nil {
		// Uma mensagem de erro informativa ajuda a depurar
		return nil, errors.New("failed to create occupation: one or more selector inputs are nil")
	}

	return &Occupation{
		under18Selector: under18,
		adultSelector:   adult,
		over65Selector:  over65,
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
