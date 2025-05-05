// Pacote dists fornece implementações de distribuições de probabilidade.
// Atualmente, inclui a distribuição Gamma.
package dists

import (
	"errors"
	"fmt"
	"math/rand/v2" // Utiliza a versão 2 do pacote math/rand

	"gonum.org/v1/gonum/stat/distuv" // Biblioteca Gonum para distribuições
)

// GammaDist representa uma distribuição de probabilidade Gamma.
// É definida pelos parâmetros de forma (shape, α) e escala (scale, θ).
type GammaDist struct {
	// shape é o parâmetro de forma (α) da distribuição. Deve ser > 0.
	shape float64
	// scale é o parâmetro de escala (θ) da distribuição. Deve ser > 0.
	scale float64
}

// Shape retorna o parâmetro de forma (α) da distribuição Gamma.
func (g *GammaDist) Shape() float64 {
	return g.shape
}

// Scale retorna o parâmetro de escala (θ) da distribuição Gamma.
func (g *GammaDist) Scale() float64 {
	return g.scale
}

// NewGammaDist cria e retorna uma nova instância de GammaDist com os
// parâmetros shape (forma) e scale (escala) fornecidos.
//
// Retorna um erro se shape ou scale não forem valores positivos (> 0).
func NewGammaDist(shape, scale float64) (*GammaDist, error) {
	if shape <= 0 {
		// Retorna erro específico se a forma for inválida.
		return nil, errors.New("parâmetro shape (forma) deve ser > 0")
	}
	if scale <= 0 {
		// Retorna erro específico se a escala for inválida.
		return nil, errors.New("parâmetro scale (escala) deve ser > 0")
	}
	// Cria e retorna a instância da distribuição se os parâmetros são válidos.
	return &GammaDist{
		shape: shape,
		scale: scale,
	}, nil
}

// Sample gera uma amostra aleatória (um valor) da distribuição Gamma.
//
// Utiliza a fonte de números aleatórios (rng *rand.Rand) fornecida.
// Internamente, usa a implementação distuv.Gamma do Gonum, que espera
// o parâmetro Beta (rate), que é o inverso do parâmetro scale (Beta = 1/scale).
func (g *GammaDist) Sample(rng *rand.Rand) float64 {
	// Cria uma instância da distribuição Gamma do Gonum.
	dist := distuv.Gamma{
		Alpha: g.shape,     // Alpha (forma) corresponde ao nosso shape.
		Beta:  1 / g.scale, // Beta (rate) é o inverso do nosso scale.
		Src:   rng,         // Define a fonte de aleatoriedade.
	}
	// Gera e retorna um número aleatório da distribuição configurada.
	return dist.Rand()
}

// String retorna uma representação textual da distribuição Gamma,
// formatada como "GammaDist{shape: X.XX, scale: Y.YY}".
func (g *GammaDist) String() string {
	// Formata a string de saída com os valores de shape e scale.
	return fmt.Sprintf("GammaDist{shape: %.2f, scale: %.2f}", g.shape, g.scale)
}