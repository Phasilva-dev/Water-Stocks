// Pacote dists fornece implementações de distribuições de probabilidade,
// como Gamma, Log-Normal, Normal, Poisson, Triangular, Uniforme e Weibull.
package dists

import (
	"errors"
	"fmt"
	"math/rand/v2" // Utiliza a versão 2 do pacote math/rand

	"gonum.org/v1/gonum/stat/distuv" // Biblioteca Gonum para distribuições
)

// WeibullDist representa uma distribuição de probabilidade de Weibull contínua.
// É frequentemente usada para modelar tempos de vida ou falha.
// É definida por dois parâmetros positivos: forma (shape, k) e escala (scale, λ).
type WeibullDist struct {
	// shape é o parâmetro de forma (k) da distribuição. Controla a forma da curva.
	// Deve ser > 0.
	shape float64 // k
	// scale é o parâmetro de escala (λ) da distribuição. Estica ou comprime a curva.
	// Deve ser > 0.
	scale float64 // lambda
}

// NewWeibullDist cria e retorna um ponteiro para uma nova instância de WeibullDist.
//
// Recebe os parâmetros de forma (shape, k) e escala (scale, λ).
// Retorna um erro se shape ou scale não forem estritamente positivos (> 0).
func NewWeibullDist(shape, scale float64) (*WeibullDist, error) {
	// Verifica se ambos os parâmetros são positivos.
	if shape <= 0 || scale <= 0 {
		return nil, errors.New("parâmetros shape (forma) e scale (escala) devem ser > 0")
	}
	return &WeibullDist{
		shape: shape,
		scale: scale,
	}, nil
}

func (w *WeibullDist) Params() []float64 {
	return []float64{w.shape, w.scale}
}

// Shape retorna o parâmetro de forma (k) da distribuição de Weibull.
func (w *WeibullDist) Shape() float64 {
	return w.shape
}

// Scale retorna o parâmetro de escala (λ) da distribuição de Weibull.
func (w *WeibullDist) Scale() float64 {
	return w.scale
}

// Sample gera uma amostra aleatória (um valor) da distribuição de Weibull.
//
// Utiliza a fonte de números aleatórios (rng *rand.Rand) fornecida.
// Internamente, usa a implementação distuv.Weibull do Gonum, que utiliza
// os parâmetros K (correspondente a shape) e Lambda (correspondente a scale).
func (w *WeibullDist) Sample(rng *rand.Rand) float64 {
	// Cria uma instância da distribuição Weibull do Gonum.
	dist := distuv.Weibull{
		K:      w.shape, // K (forma) corresponde ao nosso shape.
		Lambda: w.scale, // Lambda (escala) corresponde ao nosso scale.
		Src:    rng,     // Define a fonte de aleatoriedade.
	}
	// Gera e retorna um número aleatório da distribuição configurada.
	return dist.Rand()
}

func (w *WeibullDist) Percentile(p float64) float64 {
	dist := distuv.Weibull{
		K:      w.shape, // K (forma) corresponde ao nosso shape.
		Lambda: w.scale, // Lambda (escala) corresponde ao nosso scale.
	}
	return dist.Quantile(p)
}

// String retorna uma representação textual da distribuição de Weibull,
// formatada como "WeibullDist{shape: X.XX, scale: Y.YY}".
func (w *WeibullDist) String() string {
	// Formata a string de saída com os valores de shape e scale.
	return fmt.Sprintf("WeibullDist{shape: %.2f, scale: %.2f}", w.shape, w.scale)
}