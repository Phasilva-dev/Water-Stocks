// Package frequency define perfis de comportamento para simulação de
// frequências de uso de aparelhos sanitários. Utiliza distribuições estatísticas
// para gerar dados e aplicar limites mínimos (shifts) aos valores.
package frequency

import (
	"errors"
	"math/rand/v2"
	"math"
	"simulation/internal/dists" // Interface de distribuição estatística.
)

// FrequencyProfile representa um perfil estatístico com um valor mínimo (shift)
// aplicado aos valores gerados por uma distribuição estatística.
type FrequencyProfile struct {
	statDist dists.Distribution // Distribuição base.
	shift    uint8              // Valor mínimo aplicável ao resultado.
}

// Shift retorna o valor mínimo (shift) configurado no perfil.
func (f *FrequencyProfile) Shift() uint8 {
	return f.shift
}

// StatDist retorna a distribuição estatística usada no perfil.
func (f *FrequencyProfile) StatDist() dists.Distribution {
	return f.statDist
}

// NewFrequencyProfile cria um novo perfil de frequência com o shift mínimo e
// uma distribuição estatística base. Retorna erro se a distribuição for nula.
func NewFrequencyProfile(dist dists.Distribution, shift uint8) (*FrequencyProfile, error) {
	if dist == nil {
		return nil, errors.New("distribution cannot be null")
	}

	return &FrequencyProfile{
		statDist: dist,
		shift:    shift,
	}, nil
}

// generateFrequency gera um valor bruto de frequência, aplicando limites e o shift.
// Utiliza um gerador de números aleatórios (rng) fornecido.
func generateFrequency(rng *rand.Rand, shift uint8, statDist dists.Distribution) uint8 {
	val := statDist.Sample(rng)
	if val < 0 {
		val = math.Abs(val)
	} 
	if val > 255 {
		val = 255
	}

	freq := uint8(val)
	if freq < shift {
		return shift
	}
	return freq
}

// GenerateData gera um valor de frequência com base na distribuição e no shift
// definidos no perfil, utilizando o gerador aleatório fornecido.
func (f *FrequencyProfile) GenerateData(rng *rand.Rand) uint8 {
	return generateFrequency(rng, f.shift , f.statDist)
}
