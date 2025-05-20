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
	shift  float64
}

// Events retorna as distribuições dos eventos da rotina.
func (p *RoutineProfile) Events() []dists.Distribution {
	return p.events
}

// Shift retorna o tempo mínimo entre pares de eventos.
func (p *RoutineProfile) Shift() float64 {
	return p.shift
}

// NewRoutineProfile cria um novo perfil de rotina.
// O slice de eventos deve ter tamanho par, sem elementos nulos e shift ≥ 0.
func NewRoutineProfile(events []dists.Distribution, shift float64) (*RoutineProfile, error) {
	if len(events) == 0 {
		return nil, errors.New("events needs to be positive and even")
	}
	if len(events)%2 != 0 {
		return nil, errors.New("number of elements in events must be even")
	}
	if shift < 0 {
		return nil, errors.New("shift must be positive")
	}

	eventsCopy := make([]dists.Distribution, len(events))
	copy(eventsCopy, events)
	for _, dist := range eventsCopy {
		if dist == nil {
			return nil, errors.New("no distribution can be empty")
		}
	}

	return &RoutineProfile{
		events: eventsCopy,
		shift:  shift,
	}, nil
}

// generateTime amostra um valor da distribuição e o trunca.
func generateTime(dist dists.Distribution, rng *rand.Rand) float64 {
	return math.Trunc(dist.Sample(rng))
}

/*
// enforceMinimunGap garante que o tempo de saída respeite o intervalo mínimo.
func enforceMinimunGap(entryTime, exitTime, shift float64) float64 {
	if exitTime < entryTime + shift {
		diff := math.Abs(exitTime - entryTime)
		return exitTime + diff + shift
	}
	return exitTime
}*/

// GenerateData gera uma rotina com base no perfil atual e em um gerador de números aleatórios.
func (p *RoutineProfile) GenerateData(rng *rand.Rand) *behavioral.Routine {
	times := make([]float64, len(p.events))
	for i, dist := range p.events {
		times[i] = generateTime(dist, rng)
	}
	//    Aplica a regra do intervalo mínimo (`shift`)
	/*	A regra do intervalo minimo é a seguinte, um valor de tempo sempre deve ser maior que o proximo + o shift
		Caso não seja, coletamos a diferença absoluta do entre o tempo anterior (i-1) e o atual (i)
		Caso a diferença seja menor que o shift, definimos que o times[i] é igual ao tempo anterior somado a
		a diferença absoluta + o desvio, caso seja maior ou igual ao desvio, será simplesmente
		o tempo atual é igual ao tempo anterior + a diferença
		Essa é a forma que encontrei de representar melhor o código do módelo antigo, tornando possivel usar
		distribuições muito espaçadas sem necessariamente definir ela sempre como tempo anterior + shift caso fosse menor
		*/
	for i := 1; i < len(p.events); i++ {
		if times[i] < times [i-1] + p.shift {
			diff := times[i] - times[i-1]
			absDiff := math.Abs(diff)
			if absDiff < p.shift {
				times[i] = times[i-1] + absDiff + p.shift
			} else {
				times[i] = times[i-1] + absDiff
			}
		}
	}
	return behavioral.NewRoutine(times)
}
