package misc

import(
	"errors"
	"math/rand/v2"
	"dists"
	"sort"
)

type PercentSelector[K comparable] struct {
	values []*Tuple[K, float64]
}

// validateValues verifica se a soma dos valores excede 100%.
func validateValues[K comparable](entries []Tuple[K, float64]) error {
	total := float64(0)
	for _, entry := range entries {
		if entry.Value() <= 0 { // Validação de valores negativos ou nulos
			return errors.New("values cannot be negative or zero")
		}
		total += entry.Value()
	}
	if total > 100 {
		return errors.New("the sum of the chances exceeds 100%")
	}
	return nil
}

func sortedSlice[K comparable](entries []Tuple[K, float64]) []Tuple[K, float64] {
	tuples := make([]Tuple[K, float64], len(entries))
	copy(tuples, entries)

	sort.Slice(tuples, func(i, j int) bool {
		return tuples[i].Value() < tuples[j].Value()
	})

	return tuples
}

func NewPercentSelector[K comparable](entries []Tuple[K, float64]) (*PercentSelector[K], error) {
	// Valida se as chances são válidas (não negativas e somam até 100%)
	if err := validateValues(entries); err != nil {
		return nil, err
	}

	// Ordenar os Tuples do menor para o maior (para criar o cumulativo)
	sortedEntries := sortedSlice(entries)

	cumulative := make([]*Tuple[K, float64], len(sortedEntries))
	sum := float64(0)

	// Percorre a sortedEntries de trás para frente
	for i := len(sortedEntries) - 1; i >= 0; i-- {
		// Atualiza o valor acumulado e coloca no novo slice
		sum += sortedEntries[i].Value()
		cumulative[len(sortedEntries)-1-i] = NewTuple(sortedEntries[i].Key(), sum)
	}
	
	return &PercentSelector[K]{
		values: cumulative,
	}, nil
}


func (p *PercentSelector[K]) Sample(rng *rand.Rand) (K, error) {

	dist, err := dists.NewUniformDist(0.0, 100.0)
    if err != nil {
        var zero K
        return zero, err
    }
	
	rngValue := dist.Sample(rng)

	// Percorre os valores cumulativos e retorna a chave correspondente
	for _, tuple := range p.values {
		if rngValue <= tuple.Value() {
				return tuple.Key(), nil
		}
	}

	var zero K
	return zero, errors.New("no value selected")
}