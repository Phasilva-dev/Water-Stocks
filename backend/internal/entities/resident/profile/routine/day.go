// Pacote routine define perfis de rotina diária baseados em distribuições estatísticas.
package routine

import (
	"errors"
	"math"
	"math/rand/v2"
	"simulation/internal/dists"
	"simulation/internal/entities/resident/ds/behavioral"
)

// RoutineProfile representa um perfil de rotina com eventos (como acordar, sair, dormir)
// definidos por distribuições e um tempo mínimo (`shift`) entre pares de eventos.
type RoutineProfile struct {
	events []dists.Distribution
	minShift  float64
	maxPercent  float64 // percentil máximo (ex.: 0.99)
}

// Events retorna as distribuições dos eventos da rotina.
func (p *RoutineProfile) Events() []dists.Distribution {
	return p.events
}

// Shift retorna o tempo mínimo entre pares de eventos.
func (p *RoutineProfile) MinShift() float64 {
	return p.minShift
}

// Shift retorna o tempo Maximo entre pares de eventos.
func (p *RoutineProfile) MaxPercent() float64 {
	return p.maxPercent
}

// NewRoutineProfile cria um novo perfil de rotina.
//
// Parâmetros:
// - events: lista de distribuições, obrigatoriamente de tamanho par e sem elementos nulos.
// - minShift: espaçamento mínimo entre eventos consecutivos (≥ 0).
// - maxShift: valor máximo absoluto permitido para qualquer evento (≥ 0, ou 0 se não quiser aplicar limite).
//
// Erros são retornados se:
// - O slice de eventos estiver vazio ou com tamanho ímpar.
// - Alguma distribuição for nula.
// - minShift ou maxShift forem negativos.
func NewRoutineProfile(events []dists.Distribution, minShift, maxPercent float64) (*RoutineProfile, error) {
	if len(events) == 0 {
		return nil, errors.New("events must be non-empty and even in length")
	}
	if len(events)%2 != 0 {
		return nil, errors.New("number of events must be even")
	}
	if minShift < 0 {
		return nil, errors.New("minShift must be non-negative")
	}
	if maxPercent < 0 || maxPercent > 1 {
		return nil, errors.New("maxPercent must be between 1 and 0")
	}

	eventsCopy := make([]dists.Distribution, len(events))
	copy(eventsCopy, events)

	for _, dist := range eventsCopy {
		if dist == nil {
			return nil, errors.New("no distribution can be nil")
		}
	}

	return &RoutineProfile{
		events:   eventsCopy,
		minShift: minShift,
		maxPercent: maxPercent,
	}, nil
}

// generateTime amostra um valor da distribuição e o trunca.
func generateTime(dist dists.Distribution, rng *rand.Rand) float64 {
	return math.Trunc(dist.Sample(rng))
}


/*	A regra do intervalo minimo é a seguinte, um valor de tempo sempre deve ser maior que o proximo + o shift
Caso não seja, coletamos a diferença absoluta do entre o tempo anterior (i-1) e o atual (i)
Caso a diferença seja menor que o shift, definimos que o times[i] é igual ao tempo anterior somado a
a diferença absoluta + o desvio, caso seja maior ou igual ao desvio, será simplesmente
o tempo atual é igual ao tempo anterior + a diferença
Essa é a forma que encontrei de representar melhor o código do módelo antigo, tornando possivel usar
distribuições muito espaçadas sem necessariamente definir ela sempre como tempo anterior + shift caso fosse menor
		*/
func (r *RoutineProfile) enforceMinShift(prev, current float64) float64 {
	if r.minShift == 0 {
		return current
	}
	if current < prev+r.minShift {
		diff := math.Abs(current - prev)
		if diff < r.minShift {
			return prev + diff + r.minShift
		}
		return prev + diff
	}
	return current
}

func (r *RoutineProfile) enforceMaxValue(dist dists.Distribution, sample float64) float64 {
    if r.maxPercent >= 1 || r.maxPercent <= 0 {
        // Sem limite — permite qualquer valor amostrado
        return sample
    }

    max := dist.Percentile(r.maxPercent)
    if sample > max {
        return max
    }
    return sample
}

// GenerateData gera uma rotina com base no perfil atual e em um gerador de números aleatórios.
func (r *RoutineProfile) GenerateData(rng *rand.Rand) *behavioral.Routine {
	times := make([]float64, len(r.events))
	dists := r.events

	for i, dist := range r.events {
		times[i] = generateTime(dist, rng)
		times[i] = r.enforceMaxValue(dists[i], times[i])
	}



	// Aplicação do shift mínimo entre eventos
	for i := 1; i < len(times); i++ {
		times[i] = r.enforceMinShift(times[i-1], times[i])
	}

	return behavioral.NewRoutine(times)
}
