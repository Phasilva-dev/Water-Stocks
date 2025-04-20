package dists

import (
	"errors"
	"fmt"
	"math/rand/v2"

	"gonum.org/v1/gonum/stat/distuv"
)

// GammaDist representa uma distribuição Gamma com parâmetros shape (forma) e scale (escala).
type GammaDist struct {
	shape float64 // Parâmetro de forma (α)
	scale float64 // Parâmetro de escala (θ)
}

// Shape retorna o parâmetro de forma da distribuição.
func (g *GammaDist) Shape() float64 {
	return g.shape
}

// Scale retorna o parâmetro de escala da distribuição.
func (g *GammaDist) Scale() float64 {
	return g.scale
}

// NewGammaDist cria uma nova distribuição Gamma com os parâmetros shape (forma) e scale (escala).
func NewGammaDist(shape, scale float64) (*GammaDist, error) {
	if shape <= 0 {
		return nil, errors.New("shape must be > 0")
	}
	if scale <= 0 {
		return nil, errors.New("scale must be > 0")
	}
	return &GammaDist{
		shape: shape,
		scale: scale,
	}, nil
}

// Sample gera um valor aleatório da distribuição Gamma usando a fonte de números aleatórios fornecida.
func (g *GammaDist) Sample(rng *rand.Rand) float64 {
	// Na Gonum, Beta é o parâmetro rate (1/scale)
	dist := distuv.Gamma{
		Alpha: g.shape,       // Parâmetro de forma
		Beta:  1 / g.scale,   // Beta é rate = 1/scale
		Src:   rng,           // Fonte de aleatoriedade
	}
	return dist.Rand()
}

// String retorna uma representação em string da distribuição Gamma.
func (g *GammaDist) String() string {
	return fmt.Sprintf("GammaDist{shape: %.2f, scale: %.2f}", g.shape, g.scale)
}