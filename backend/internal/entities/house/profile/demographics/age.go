// Package demographics define estruturas para gerar dados demográficos,
// como a idade dos residentes, com base em distribuições estatísticas.
package demographics

import (
	"math"
	"math/rand/v2"
	"simulation/internal/dists"
)

// Age representa um perfil para a geração de dados de idade.
// Ele utiliza uma distribuição estatística para amostrar valores de idade.
type Age struct {
	dist dists.Distribution // A distribuição estatística usada para gerar as idades.
}

// NewAge cria e retorna uma nova instância de Age.
//
// dist: A distribuição estatística que será usada para amostrar as idades.
//       Espera-se que esta distribuição produza valores apropriados para representar idades.
//
// Retorna um ponteiro para a estrutura Age.
func NewAge(dist dists.Distribution) *Age {
	return &Age{
		dist: dist,
	}
}

// AgeDist retorna a distribuição estatística usada para gerar as idades.
func (a *Age) AgeDist() dists.Distribution {
	return a.dist
}

// GenerateData gera um valor de idade com base na distribuição do perfil.
// O valor amostrado é processado para garantir que seja um número inteiro não-negativo
// e dentro dos limites de um uint8 (0 a 255).
//
// rng: O gerador de números aleatórios a ser usado para a amostragem.
//
// Retorna a idade gerada como um uint8, garantindo que seja um valor
// válido (mínimo de 0, máximo de 255).
func (a *Age) GenerateData(rng *rand.Rand) uint8 {
	sample := a.dist.Sample(rng) // Amostra um valor da distribuição.
	absSample := math.Abs(sample) // Pega o valor absoluto para garantir que seja não-negativo.
	ceilSample := math.Ceil(absSample) // Arredonda para cima para o inteiro mais próximo.

	// Garante que o valor esteja dentro dos limites válidos para uint8.
	// Nota: Com math.Abs e math.Ceil, ceilSample já será >= 0,
	// então a primeira condição `ceilSample < 0` é redundante aqui,
	// mas mantida por robustez ou intenção de projeto original.
	if ceilSample < 0 {
		ceilSample = 0
	}
	if ceilSample > math.MaxUint8 { // math.MaxUint8 é 255
		ceilSample = math.MaxUint8
	}

	return uint8(ceilSample) // Converte o valor processado para uint8.
}