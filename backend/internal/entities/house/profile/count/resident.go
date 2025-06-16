package count

import (
	"math"
	"math/rand/v2"
	"simulation/internal/dists"
)

// ResidentCount representa um perfil para gerar o número de residentes.
// Utiliza uma distribuição estatística para amostrar a contagem de residentes.
type ResidentCount struct {
	dist dists.Distribution // A distribuição estatística usada para gerar a contagem.
}

// NewResidentCount cria e retorna uma nova instância de ResidentCount.
//
// dist: A distribuição estatística que será usada para amostrar o número de residentes.
//       Espera-se que esta distribuição produza valores apropriados para contagem.
//
// Retorna um ponteiro para a estrutura ResidentCount.
func NewResidentCount(dist dists.Distribution) *ResidentCount {
	return &ResidentCount{
		dist: dist,
	}
}

// GenerateData gera uma contagem de residentes com base na distribuição do perfil.
// O valor amostrado é processado para garantir que seja um número inteiro positivo
// e dentro dos limites de um uint8 (0 a 255).
//
// rng: O gerador de números aleatórios a ser usado para a amostragem.
//
// Retorna a contagem de residentes como um uint8, garantindo que seja um valor
// válido (mínimo de 1, máximo de 255).
func (r *ResidentCount) GenerateData(rng *rand.Rand) uint8 {
	sample := r.dist.Sample(rng)

	// Arredonda o valor amostrado para cima para o inteiro mais próximo.
	processedValue := math.Ceil(sample)

	// Garante que a contagem mínima seja 1.
	if processedValue <= 0 {
		processedValue = 1
	} else if processedValue > math.MaxUint8 { // Garante que não exceda o valor máximo de uint8 (255).
		processedValue = math.MaxUint8 // 255
	}

	return uint8(processedValue)
}