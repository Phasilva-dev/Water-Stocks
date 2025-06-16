// Pacote routine define perfis de rotina diária baseados em distribuições estatísticas.
package routine

import (
	"fmt"
	"math"
	"math/rand/v2"
	"simulation/internal/dists"
	"simulation/internal/entities/resident/ds/behavioral"
)

// RoutineProfile representa um perfil de rotina diária.
// Contém as distribuições para os eventos da rotina, um tempo mínimo entre eventos
// e um percentil máximo para limitar os valores amostrados das distribuições.
type RoutineProfile struct {
	events     []dists.Distribution
	minShift   float64
	maxPercent float64 // Percentil máximo (entre 0 e 1, ex.: 0.99) para limitar os valores amostrados.
}

// Events retorna as distribuições dos eventos da rotina.
func (p *RoutineProfile) Events() []dists.Distribution {
	return p.events
}

// MinShift retorna o tempo mínimo obrigatório entre eventos consecutivos.
func (p *RoutineProfile) MinShift() float64 {
	return p.minShift
}

// MaxPercent retorna o percentil máximo usado para limitar os valores amostrados dos eventos.
func (p *RoutineProfile) MaxPercent() float64 {
	return p.maxPercent
}

// NewRoutineProfile cria um novo perfil de rotina.
//
// Parâmetros:
//   - events: Lista de distribuições que definem os horários dos eventos da rotina.
//     Deve ter um tamanho par (eventos de "entrada" e "saída") e não conter elementos nulos.
//   - minShift: O espaçamento mínimo em segundos que deve existir entre quaisquer dois eventos consecutivos (≥ 0).
//   - maxPercent: O percentil (entre 0 e 1) usado para limitar o valor máximo de cada evento amostrado.
//     Um valor de 1.0 ou 0.0 (ou fora do intervalo (0,1)) significa que nenhum limite de percentil é aplicado (o valor amostrado é mantido).
//
// Retorna um ponteiro para RoutineProfile ou um erro se a validação falhar.
func NewRoutineProfile(events []dists.Distribution, minShift, maxPercent float64) (*RoutineProfile, error) {
	if len(events) == 0 {
		return nil, fmt.Errorf("invalid routine profile: events must not be empty (got length %d)", len(events))
	}
	if len(events)%2 != 0 {
		return nil, fmt.Errorf("invalid routine profile: number of events must be even (got %d)", len(events))
	}
	if minShift < 0 {
		return nil, fmt.Errorf("invalid routine profile: minShift must be non-negative (got %.4f)", minShift)
	}
	if maxPercent < 0 || maxPercent > 1 {
		return nil, fmt.Errorf("invalid routine profile: maxPercent must be between 0 and 1 (got %.4f)", maxPercent)
	}

	eventsCopy := make([]dists.Distribution, len(events))
	copy(eventsCopy, events)

	for i, dist := range events {
		if dist == nil {
			return nil, fmt.Errorf("invalid routine profile: distribution at index %d is nil", i)
		}
	}

	return &RoutineProfile{
		events:     eventsCopy,
		minShift:   minShift,
		maxPercent: maxPercent,
	}, nil
}

// generateTime amostra um valor de uma distribuição e o trunca para um número inteiro.
func generateTime(dist dists.Distribution, rng *rand.Rand) float64 {
	return math.Trunc(dist.Sample(rng))
}

// enforceMinShift ajusta o tempo 'current' para garantir que haja um espaçamento mínimo de 'minShift'
// em relação ao tempo 'prev'.
// Se 'current' for menor que 'prev + minShift', ele é ajustado para um valor que respeite o mínimo,
// considerando a diferença original entre 'prev' e 'current'.
func (r *RoutineProfile) enforceMinShift(prev, current float64) float64 {
	if r.minShift == 0 {
		return current
	}
	if current < prev+r.minShift {
		diff := math.Abs(current - prev)
		if diff < r.minShift {
			// Se a diferença é pequena, adiciona a diferença e o shift mínimo
			return prev + diff + r.minShift
		}
		// Se a diferença é maior ou igual ao shift, apenas garante que current é pelo menos prev + diff
		return prev + diff
	}
	return current
}

// enforceMaxValue limita o valor 'sample' ao percentil 'maxPercent' da distribuição.
// Se r.maxPercent for 0 ou 1, ou fora do intervalo (0, 1), nenhum limite de percentil é aplicado,
// e o valor 'sample' original é retornado.
// Caso contrário, 'sample' é truncado para o valor do percentil se for maior.
func (r *RoutineProfile) enforceMaxValue(dist dists.Distribution, sample float64) float64 {
	if r.maxPercent >= 1 || r.maxPercent <= 0 {
		// Sem limite por percentil (ou com percentil que não faz sentido limitar)
		return sample
	}

	max := dist.Percentile(r.maxPercent)
	if sample > max {
		return max
	}
	return sample
}

// GenerateData gera uma nova rotina comportamental com base no perfil atual e
// em um gerador de números aleatórios fornecido.
// Os tempos são amostrados, limitados por maxPercent e ajustados para respeitar minShift.
func (r *RoutineProfile) GenerateData(rng *rand.Rand) *behavioral.Routine {
	times := make([]float64, len(r.events))
	dists := r.events // Renomeado para evitar conflito com 'dist' no loop

	// Amostra e aplica o limite de valor máximo a cada evento
	for i, dist := range r.events {
		times[i] = generateTime(dist, rng)
		times[i] = r.enforceMaxValue(dists[i], times[i])
	}

	// Aplica o shift mínimo entre eventos sequencialmente
	for i := 1; i < len(times); i++ {
		times[i] = r.enforceMinShift(times[i-1], times[i])
	}

	return behavioral.NewRoutine(times)
}