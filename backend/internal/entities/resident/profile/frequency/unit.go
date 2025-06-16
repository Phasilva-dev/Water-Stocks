// Package frequency define perfis para geração de dados de frequência de uso,
// aplicando distribuições estatísticas e regras de valores mínimos (shifts).
package frequency

import (
	"errors"
	"math"
	"math/rand/v2"
	"simulation/internal/dists" // Interface para distribuições estatísticas.
)

// FrequencyProfile representa um perfil estatístico usado para gerar valores de frequência.
// Ele inclui uma distribuição base e um valor mínimo (shift) a ser aplicado ao resultado.
type FrequencyProfile struct {
	statDist dists.Distribution // A distribuição estatística base para amostrar frequências.
	shift    uint8              // O valor mínimo (inclusive) que uma frequência gerada pode ter.
}

// Shift retorna o valor mínimo (shift) configurado no perfil.
func (f *FrequencyProfile) Shift() uint8 {
	return f.shift
}

// StatDist retorna a distribuição estatística base usada pelo perfil.
func (f *FrequencyProfile) StatDist() dists.Distribution {
	return f.statDist
}

// NewFrequencyProfile cria um novo perfil de frequência.
//
// dist: A distribuição estatística base para o perfil. Não pode ser nula.
// shift: O valor mínimo que as frequências geradas devem ter.
//
// Retorna um ponteiro para o FrequencyProfile ou um erro se a distribuição for nula.
func NewFrequencyProfile(dist dists.Distribution, shift uint8) (*FrequencyProfile, error) {
	if dist == nil {
		return nil, errors.New("distribution cannot be nil") // Corrigido para "nil"
	}

	return &FrequencyProfile{
		statDist: dist,
		shift:    shift,
	}, nil
}

// generateFrequency amostra um valor da distribuição, aplica limites (0-255) e o shift mínimo.
//
// rng: O gerador de números aleatórios a ser usado para a amostragem.
// shift: O valor mínimo a ser garantido.
// statDist: A distribuição da qual o valor será amostrado.
//
// Retorna a frequência gerada como uint8, garantindo que esteja entre [shift, 255].
func generateFrequency(rng *rand.Rand, shift uint8, statDist dists.Distribution) uint8 {
	val := statDist.Sample(rng)

	// Garante que o valor seja não-negativo
	if val < 0 {
		val = math.Abs(val)
	}
	// Limita o valor máximo para caber em uint8
	if val > 255 {
		val = 255
	}

	freq := uint8(val)
	// Aplica o shift mínimo
	if freq < shift {
		return shift
	}
	return freq
}

// GenerateData gera um valor de frequência com base nas configurações do perfil.
//
// rng: O gerador de números aleatórios a ser usado.
//
// Retorna um valor de frequência uint8 que respeita o shift mínimo do perfil
// e está dentro dos limites de 0 a 255.
func (f *FrequencyProfile) GenerateData(rng *rand.Rand) uint8 {
	return generateFrequency(rng, f.shift, f.statDist)
}