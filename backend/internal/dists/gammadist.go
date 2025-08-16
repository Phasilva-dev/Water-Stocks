// Pacote dists fornece implementações de distribuições de probabilidade.
// Atualmente, inclui a distribuição Gamma.
package dists

import (
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

func (g *GammaDist) Params() []float64 {
	return []float64{g.shape, g.scale}
}


// Shape retorna o parâmetro de forma (α) da distribuição Gamma.
func (g *GammaDist) Shape() float64 {
	return g.shape
}

// Scale retorna o parâmetro de escala (θ) da distribuição Gamma.
func (g *GammaDist) Scale() float64 {
	return g.scale
}

// newGammaDist cria e retorna uma nova instância de GammaDist com os
// parâmetros shape (forma) e scale (escala) fornecidos.
//
// Retorna um erro se shape ou scale não forem valores positivos (> 0).
func newGammaDist(shape, scale float64) (*GammaDist, error) {

	if shape <= 0 {
		return nil, fmt.Errorf(
			"invalid Gamma Distribution Parameters: shape (forma) must be > 0 (shape=%.2f)",
			shape,
		)
	}
	if scale <= 0 {
		return nil, fmt.Errorf(
			"invalid Gamma Distribution Parameters: scale (escala) must be > 0 (scale=%.2f)",
			scale,
		)
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

func (g *GammaDist) Percentile(p float64) float64 {
	dist := distuv.Gamma{
		Alpha: g.shape,     // Alpha (forma) corresponde ao nosso shape.
		Beta:  1 / g.scale, // Beta (rate) é o inverso do nosso scale.
	}
	return dist.Quantile(p)
}

// String retorna uma representação textual da distribuição Gamma,
// formatada como "GammaDist{shape: X.XX, scale: Y.YY}".
func (g *GammaDist) String() string {
	// Formata a string de saída com os valores de shape e scale.
	return fmt.Sprintf("GammaDist{shape: %.2f, scale: %.2f}", g.shape, g.scale)
}