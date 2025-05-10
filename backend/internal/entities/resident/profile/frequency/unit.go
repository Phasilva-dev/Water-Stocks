// Package frequency define perfis de comportamento para simulação de
// frequências de uso de aparelhos sanitários. Utiliza distribuições estatísticas
// para gerar dados e aplicar limites mínimos (shifts) aos valores.
package frequency

import (
	"errors"
	"math/rand/v2"
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
func (f *FrequencyProfile) NewFrequencyProfile(shift uint8, dist dists.Distribution) (*FrequencyProfile, error) {
	if dist == nil {
		return nil, errors.New("a distribuição não pode ser nula")
	}
	if shift < 0 {
		return nil, errors.New("shift não pode ser negativo")
	}
	return &FrequencyProfile{
		statDist: dist,
		shift:    shift,
	}, nil
}

// generateFrequency gera um valor bruto de frequência, aplicando limites e o shift.
// Utiliza um gerador de números aleatórios (rng) fornecido.
func (f *FrequencyProfile) generateFrequency(rng *rand.Rand) uint8 {
	val := f.statDist.Sample(rng)
	if val < 0 {
		val = 0
	} else if val > 255 {
		val = 255
	}

	freq := uint8(val)
	if freq < f.shift {
		return f.shift
	}
	return freq
}

// GenerateData gera um valor de frequência com base na distribuição e no shift
// definidos no perfil, utilizando o gerador aleatório fornecido.
func (f *FrequencyProfile) GenerateData(rng *rand.Rand) uint8 {
	return f.generateFrequency(rng)
}
