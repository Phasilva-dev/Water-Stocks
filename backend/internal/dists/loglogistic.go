package dists

import (
	"fmt"
	"math"
	"math/rand/v2"
)

// LogLogisticDist representa uma distribuição de probabilidade LogLogistic.
type LogLogisticDist struct {
	shape float64 // α, parâmetro de forma
	scale float64 // β, parâmetro de escala
}

func (l *LogLogisticDist) Params() []float64 {
	return []float64{l.shape, l.scale}
}

// Shape retorna o parâmetro de forma (α).
func (ll *LogLogisticDist) Shape() float64 {
	return ll.shape
}

// Scale retorna o parâmetro de escala (β).
func (ll *LogLogisticDist) Scale() float64 {
	return ll.scale
}

// newLogLogisticDist cria uma nova instância de LogLogisticDist.
func newLogLogisticDist(shape, scale float64) (Distribution, error) {
	if shape <= 0 {
		return nil, fmt.Errorf(
			"invalid LogLogistic Distribution Parameters: shape (forma) must be > 0 (shape=%.2f)",
			shape,
		)
	}
	if scale <= 0 {
		return nil, fmt.Errorf(
			"invalid LogLogistic Distribution Parameters: scale (escala) must be > 0 (scale=%.2f)",
			scale,
		)
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

// Percentile calcula o valor x tal que P(X <= x) = p.
// p deve estar no intervalo [0, 1].
func (ll *LogLogisticDist) Percentile(p float64) float64 {
	if p < 0 || p > 1 {
		// Comportamento comum para p fora do domínio [0,1] é retornar NaN.
		return math.NaN()
	}
	if p == 0 {
		// Se p = 0, x = β * (0 / 1)^(1/α) = β * 0 = 0.
		return 0.0
	}
	if p == 1 {
		// Se p = 1, x = β * (1 / 0)^(1/α) = β * ∞^(1/α) = ∞.
		return math.Inf(1) // Retorna +Infinito
	}

	// Fórmula da função quantil: β * (p / (1-p))^(1/α)
	// α = ll.shape
	// β = ll.scale
	return ll.scale * math.Pow(p/(1.0-p), 1.0/ll.shape)
}


// String retorna uma representação textual da distribuição.
func (ll *LogLogisticDist) String() string {
	return fmt.Sprintf("LogLogisticDist{shape: %.2f, scale: %.2f}", ll.shape, ll.scale)
}