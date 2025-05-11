package dists

import (
	"errors"
	"fmt"
	"math"
	"math/rand/v2"
)

// LogLogisticDist representa uma distribuição de probabilidade LogLogistic.
type LogLogisticDist struct {
	shape float64 // α, parâmetro de forma
	scale float64 // β, parâmetro de escala
}

// Shape retorna o parâmetro de forma (α).
func (ll *LogLogisticDist) Shape() float64 {
	return ll.shape
}

// Scale retorna o parâmetro de escala (β).
func (ll *LogLogisticDist) Scale() float64 {
	return ll.scale
}

// NewLogLogisticDist cria uma nova instância de LogLogisticDist.
func NewLogLogisticDist(shape, scale float64) (*LogLogisticDist, error) {
	if shape <= 0 {
		return nil, errors.New("parâmetro shape (forma) deve ser > 0")
	}
	if scale <= 0 {
		return nil, errors.New("parâmetro scale (escala) deve ser > 0")
	}
	return &LogLogisticDist{
		shape: shape,
		scale: scale,
	}, nil
}

// Sample gera uma amostra aleatória da distribuição LogLogistic.
func (ll *LogLogisticDist) Sample(rng *rand.Rand) float64 {
	// A distribuição LogLogistic pode ser gerada a partir de uma uniforme.
	// Se U ~ Uniform(0,1), então X = β * (U/(1-U))^(1/α) segue LogLogistic(α, β).
	u := rng.Float64()
	if u == 0 { // Evita divisão por zero
		u = 1e-10
	}
	return ll.scale * math.Pow(u/(1-u), 1/ll.shape)
}

// String retorna uma representação textual da distribuição.
func (ll *LogLogisticDist) String() string {
	return fmt.Sprintf("LogLogisticDist{shape: %.2f, scale: %.2f}", ll.shape, ll.scale)
}