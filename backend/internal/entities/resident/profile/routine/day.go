// Pacote routine define perfis de rotina diária baseados em distribuições estatísticas.
package routine

import (
	"fmt"
	"math"
	"math/rand/v2"
	"simulation/internal/dists"
	"simulation/internal/entities/resident/ds/behavioral"
)

// DayProfile representa um perfil de rotina diária.
// Contém as distribuições para os eventos da rotina, um tempo mínimo entre eventos
// e um percentil máximo para limitar os valores amostrados das distribuições.
type DayProfile struct {
	events     []dists.Distribution
	minShift   float64
	maxPercent float64 // Percentil máximo (entre 0 e 1, ex.: 0.99) para limitar os valores amostrados.
}

// Events retorna as distribuições dos eventos da rotina.
func (p *DayProfile) Events() []dists.Distribution {
	return p.events
}

// MinShift retorna o tempo mínimo obrigatório entre eventos consecutivos.
func (p *DayProfile) MinShift() float64 {
	return p.minShift
}

// MaxPercent retorna o percentil máximo usado para limitar os valores amostrados dos eventos.
func (p *DayProfile) MaxPercent() float64 {
	return p.maxPercent
}

// NewDayProfile cria um novo perfil de rotina.
//
// Parâmetros:
//   - events: Lista de distribuições que definem os horários dos eventos da rotina.
//     Deve ter um tamanho par (eventos de "entrada" e "saída") e não conter elementos nulos.
//   - minShift: O espaçamento mínimo em segundos que deve existir entre quaisquer dois eventos consecutivos (≥ 0).
//   - maxPercent: O percentil (entre 0 e 1) usado para limitar o valor máximo de cada evento amostrado.
//     Um valor de 1.0 ou 0.0 (ou fora do intervalo (0,1)) significa que nenhum limite de percentil é aplicado (o valor amostrado é mantido).
//
// Retorna um ponteiro para DayProfile ou um erro se a validação falhar.
func NewDayProfile(events []dists.Distribution, minShift, maxPercent float64) (*DayProfile, error) {
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

	return &DayProfile{
		events:     eventsCopy,
		minShift:   minShift,
		maxPercent: maxPercent,
	}, nil
}

// generateTime amostra um valor de uma distribuição e o trunca para um número inteiro.
//
// Parâmetros:
//   - dist: Distribuição da qual o valor será amostrado.
//   - rng: Gerador de números aleatórios.
//
// Retorna um valor truncado da distribuição.
func generateTime(dist dists.Distribution, rng *rand.Rand) float64 {
	return math.Trunc(dist.Sample(rng))
}

// enforceMinShift ajusta o tempo 'current' para garantir que haja um espaçamento mínimo
// de 'minShift' em relação ao tempo 'prev'.
//
// A função tenta preservar a diferença original entre os tempos amostrados,
// mas, se necessário, empurra 'current' para frente a fim de respeitar o intervalo mínimo.
//
// Regras:
//   - Se minShift for 0, retorna 'current' sem alterações.
//   - Se 'current' for menor que 'prev + minShift', calcula a diferença absoluta entre eles:
//       - Se essa diferença for menor que minShift, retorna 'prev + diff + minShift'.
//         (empurra o evento ainda mais para evitar sobreposição)
//       - Caso contrário, retorna 'prev + diff'.
//   - Se 'current' já respeita o espaçamento mínimo, retorna o valor original.
//
// Parâmetros:
//   - prev: Tempo do evento anterior (em segundos).
//   - current: Tempo amostrado para o próximo evento (em segundos).
//   - minShift: Espaçamento mínimo entre eventos consecutivos (em segundos).
//
// Retorna:
//   - O tempo ajustado de 'current' que respeita o espaçamento mínimo.
func enforceMinShift(prev, current, minShift float64) float64 {
	if minShift == 0 {
		return current
	}

	limit := prev + minShift

	if current < limit {
		diff := math.Abs(current - prev)
		if diff < minShift {
			return prev + diff + minShift
		}
		return prev + diff
	}

	return current
}

// enforceMaxValue limita o valor 'sample' ao percentil 'maxPercent' da distribuição.
//
// Parâmetros:
//   - dist: Distribuição associada ao evento.
//   - sample: Valor amostrado a ser verificado.
//
// Retorna o valor original ou limitado pelo percentil da distribuição.
func enforceMaxValue(dist dists.Distribution, sample, maxPercent float64) float64 {
	if maxPercent >= 1 || maxPercent <= 0 {
		return sample
	}

	max := dist.Percentile(maxPercent)
	if sample > max {
		return max
	}
	return sample
}

// GenerateData gera uma nova rotina comportamental com base no perfil atual e
// em um gerador de números aleatórios fornecido.
// Os tempos são amostrados, limitados por maxPercent e ajustados para respeitar minShift.
//
// Parâmetros:
//   - rng: Gerador de números aleatórios.
//
// Retorna um ponteiro para Routine com os tempos gerados, ou um erro se a geração falhar após múltiplas tentativas.
func (r *DayProfile) GenerateData(rng *rand.Rand) (*behavioral.Routine, error) {
	dists := r.events // Renomeado para evitar conflito com 'dist' no loop
	times := make([]float64, len(dists))

	const maxAttempts = 10
	const maxTimeValue = 172799.0 // Último segundo do dia (23h59m59s)
	const minTimeValue = 0.0
	attempts := 0

	for {
		// Amostra e aplica o limite de valor máximo a cada evento
		for i, dist := range dists {
			times[i] = generateTime(dist, rng)
			times[i] = enforceMaxValue(dists[i], times[i], r.maxPercent)
		}

		// Aplica o shift mínimo entre eventos sequencialmente
		for i := 1; i < len(times); i++ {
			times[i] = enforceMinShift(times[i-1], times[i],r.minShift)
			times[i] = math.Trunc(times[i])
		}

		// Verificar se todos os tempos estão no intervalo válido
		valid := true
		for i := 0; i < len(times); i++ {
			if times[i] < minTimeValue || times[i] > maxTimeValue {
				valid = false
				break
			}
		}

		if valid {
			break // Todos os valores estão dentro do intervalo, sair do loop
		}

		attempts++
		if attempts >= maxAttempts {
			return nil, fmt.Errorf("invalid routine profile simulation: time should be between 0 and 172799.0 in at least 1 out of 10 attempts")
		}
	}

	return behavioral.NewRoutine(times), nil
}
