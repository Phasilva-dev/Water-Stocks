// Pacote dists fornece implementações de distribuições de probabilidade
// que seguem a interface Distribution.
package dists

import (
	"fmt"
	"math/rand/v2" // Utiliza a versão 2 do pacote math/rand

	"gonum.org/v1/gonum/stat/distuv" // Biblioteca Gonum para distribuições
)

// ExponentialDist representa uma distribuição de probabilidade Exponencial.
// É definida pelo parâmetro de taxa (rate, λ).
type ExponentialDist struct {
	// rate é o parâmetro de taxa (λ) da distribuição. Deve ser > 0.
	rate float64
}

// NewExponentialDist cria e retorna uma nova instância de ExponentialDist com o
// parâmetro de taxa (rate) fornecido.
//
// Retorna um erro se a taxa não for um valor positivo (> 0).
func NewExponentialDist(rate float64) (*ExponentialDist, error) {
	if rate <= 0 {
		return nil, fmt.Errorf("rate parameter must be positive, but got %f", rate)
	}
	return &ExponentialDist{
		rate: rate,
	}, nil
}

// Rate retorna o parâmetro de taxa (λ) da distribuição Exponencial.
// Este é um método auxiliar específico para ExponentialDist, não faz parte da interface.
func (e *ExponentialDist) Rate() float64 {
	return e.rate
}

// --- Implementação da interface Distribution ---

// Params retorna os parâmetros da distribuição. Para a Exponencial, contém apenas a taxa.
func (e *ExponentialDist) Params() []float64 {
	return []float64{e.rate}
}

// Sample gera uma amostra aleatória da distribuição Exponencial.
func (e *ExponentialDist) Sample(rng *rand.Rand) float64 {
	dist := distuv.Exponential{
		Rate: e.rate,
		Src:  rng,
	}
	return dist.Rand()
}

// Percentile calcula o valor correspondente ao percentil p.
func (e *ExponentialDist) Percentile(p float64) float64 {
	dist := distuv.Exponential{
		Rate: e.rate,
	}
	return dist.Quantile(p)
}

// String retorna uma representação textual da distribuição Exponencial.
func (e *ExponentialDist) String() string {
	return fmt.Sprintf("ExponentialDist{rate: %.2f}", e.rate)
}