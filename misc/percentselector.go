package misc

import(
	"errors"
	"math/rand/v2"
	"dists"
	"sort"
)

type PercentSelector struct {
	values []Tuple
}

// validateValues verifica se a soma dos valores excede 100%.
func validateValues(entries map[string]float64) error {
	total := float64(0)
	for _, val := range entries {
		if val < 0 { // Validação de valores negativos
			return errors.New("values cannot be negative")
		}
		total += val
	}
	if total > 100 {
		return errors.New("the sum of the chances exceeds 100%")
	}
	return nil
}

func sortedSlice(values map[string]float64) []Tuple {
	tuples := make([]Tuple, 0, len(values))

	for k, v := range values {
		tuples = append(tuples, Tuple{key: k, value: v})
	}
	sort.Slice(tuples, func(i, j int) bool {
		return tuples[i].Value() < tuples[j].Value()
	})
	return tuples
}

func NewPercentSelector(values map[string]float64) (*PercentSelector, error) {
	if err := validateValues(values); err != nil {
		return nil, err
	}

	// Criar o map com somas cumulativas
	cumullativeMap := make(map[string]float64, len(values))
	sum := float64(0)

	for k, v := range values {
		sum += v
		cumullativeMap[k] = sum
	}
	slice := sortedSlice(cumullativeMap)

	return &PercentSelector{
		values: slice,
	}, nil
}


func (p *PercentSelector) Sample(rng *rand.Rand) (string, error) {

	dist, err := dists.NewUniformDist(0.0, 100.0)
	if err != nil {
		return "" ,err
	}
	
	rngValue := dist.Sample(rng)

	// Percorre os valores ordenados e retorna a chave correspondente ao intervalo
	for _, tuple := range p.values {
		if rngValue <= tuple.Value() {
				return tuple.Key(), nil
		}
}

// Caso algo dê errado (o que não deve acontecer se validateValues foi chamado)
return "", errors.New("no value selected")

}