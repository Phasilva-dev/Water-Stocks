// Pacote demographics fornece estruturas e funções para modelar
// seleção de ocupações baseada em faixas etárias.
// Ele utiliza PercentSelector para realizar a escolha probabilística
// de ocupações dentro de cada faixa etária definida.
package demographics

import (
	"fmt"
	"math/rand/v2"

	"simulation/internal/misc"
)

// AgeRangeSelector representa um seletor de valores probabilísticos
// associado a uma faixa etária (minAge..maxAge). Ele define qual
// PercentSelector será utilizado para idades dentro dessa faixa.
type AgeRangeSelector struct {
	minAge   uint8
	maxAge   uint8
	selector *misc.PercentSelector[uint32]
}

// MinAge retorna a idade mínima associada ao seletor.
func (a *AgeRangeSelector) MinAge() uint8 {
	return a.minAge
}

// MaxAge retorna a idade máxima associada ao seletor.
func (a *AgeRangeSelector) MaxAge() uint8 {
	return a.maxAge
}

// Selector retorna o PercentSelector associado à faixa etária.
func (a *AgeRangeSelector) Selector() *misc.PercentSelector[uint32] {
	return a.selector
}

// NewAgeRangeSelector cria uma nova instância de AgeRangeSelector.
// Retorna erro caso o selector seja nulo ou caso minAge seja maior que maxAge.
func NewAgeRangeSelector(minAge, maxAge uint8, selector *misc.PercentSelector[uint32]) (*AgeRangeSelector, error) {
	if selector == nil {
		return nil, fmt.Errorf("failed to create AgeRangeSelector: selector cannot be nil")
	}
	if minAge > maxAge {
		return nil, fmt.Errorf("failed to create AgeRangeSelector: minAge (%d) cannot be greater than maxAge (%d)", minAge, maxAge)
	}
	return &AgeRangeSelector{
		minAge:   minAge,
		maxAge:   maxAge,
		selector: selector,
	}, nil
}

// Occupation representa um perfil de ocupação que combina múltiplos
// AgeRangeSelectors, garantindo que não haja sobreposição de idades
// entre eles.
type Occupation struct {
	selectors []*AgeRangeSelector
}

// NewOccupation cria uma nova instância de Occupation.
// Retorna erro se a lista de seletores estiver vazia, contiver elementos nulos
// ou se houver sobreposição de faixas etárias entre seletores.
func NewOccupation(selectors []*AgeRangeSelector) (*Occupation, error) {
	if len(selectors) == 0 {
		return nil, fmt.Errorf("failed to create Occupation: selectors cannot be empty")
	}

	// Verificar validade e sobreposição
	for i := 0; i < len(selectors); i++ {
		selA := selectors[i]
		if selA == nil {
			return nil, fmt.Errorf("failed to create Occupation: selector at index %d is nil", i)
		}
		for j := i + 1; j < len(selectors); j++ {
			selB := selectors[j]
			if selB == nil {
				return nil, fmt.Errorf("failed to create Occupation: selector at index %d is nil", j)
			}

			if rangesOverlap(selA.MinAge(), selA.MaxAge(), selB.MinAge(), selB.MaxAge()) {
				return nil, fmt.Errorf("failed to create Occupation: age ranges overlap between selectors [%d-%d] and [%d-%d]",
					selA.MinAge(), selA.MaxAge(), selB.MinAge(), selB.MaxAge())
			}
		}
	}

	return &Occupation{selectors: selectors}, nil
}

// rangesOverlap verifica se duas faixas etárias se sobrepõem.
func rangesOverlap(minA, maxA, minB, maxB uint8) bool {
	return minA <= maxB && minB <= maxA
}

// Selectors retorna a lista de AgeRangeSelectors associados à ocupação.
func (o *Occupation) Selectors() []*AgeRangeSelector {
	return o.selectors
}

// Sample seleciona uma ocupação com base na idade fornecida e
// na fonte de aleatoriedade (rng).
//
// Se a idade for maior que o maior MaxAge definido, ela é limitada ao valor máximo.
// Retorna erro se não for encontrado nenhum seletor correspondente.
func (o *Occupation) Sample(age uint8, rng *rand.Rand) (uint32, error) {
	// Encontrar o maior MaxAge
	var maxAllowedAge uint8 = 0
	for _, sel := range o.selectors {
		if sel.MaxAge() > maxAllowedAge {
			maxAllowedAge = sel.MaxAge()
		}
	}

	// Limitar a idade ao maior MaxAge
	if age > maxAllowedAge {
		age = maxAllowedAge
	}

	// Procurar seletor compatível
	for _, sel := range o.selectors {
		if age >= sel.MinAge() && age <= sel.MaxAge() {
			return sel.Selector().Sample(rng)
		}
	}

	return 0, fmt.Errorf("failed to sample occupation: no selector found for age %d", age)
}
