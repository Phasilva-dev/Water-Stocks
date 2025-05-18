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
	//    Aplica a regra do intervalo mínimo (`shift`) aos eventos de índice ímpar.
	//    Itera sobre os índices ímpares (1, 3, 5, ...).
	for i := 1; i < len(p.events); i += 2 {
		// `times[i-1]` é o tempo de "entrada" (par).
		// `times[i]` é o tempo de "saída" (ímpar) bruto.
		// Ajusta `times[i]` (saída) para garantir o gap mínimo após `times[i-1]` (entrada).
		// Essa saida é sempre igual ao modulo da diferença entre a saida e entrada somada ao shift para garantir um tempo minimo
		if times[i] <= times [i-1] + p.shift {
			diff := math.Abs(times[i] - times[i-1])
			times[i] = times[i] + diff + p.shift
		}
	}
	return behavioral.NewRoutine(times)
}
