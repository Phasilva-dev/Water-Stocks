// Pacote dists fornece implementações de distribuições de probabilidade,
// como Gamma, Log-Normal, Normal e Poisson.
package dists

import (
	"fmt"
	"math/rand/v2" // Utiliza a versão 2 do pacote math/rand

	"gonum.org/v1/gonum/stat/distuv" // Biblioteca Gonum para distribuições
)

// PoissonDist representa uma distribuição de probabilidade de Poisson.
// É definida pelo parâmetro lambda (λ), que representa a taxa média
// de ocorrência de eventos em um intervalo fixo de tempo ou espaço.
type PoissonDist struct {
	// lambda é o parâmetro de taxa (λ) da distribuição.
	// Representa o número médio esperado de eventos. Deve ser > 0.
	lambda float64
}

func (p *PoissonDist) Params() []float64 {
	return []float64{p.lambda}
}

// Lambda retorna o parâmetro de taxa (λ) da distribuição de Poisson.
func (p *PoissonDist) Lambda() float64 {
	return p.lambda
}

// newPoissonDist cria e retorna uma nova instância de PoissonDist.
//
// Recebe o parâmetro de taxa (lambda, λ) como argumento.
// Retorna um erro se lambda não for estritamente positivo (> 0).
func newPoissonDist(lambda float64) (Distribution, error) {
	// Verifica se o parâmetro lambda é positivo.
	if lambda <= 0 {
	return nil, fmt.Errorf(
		"invalid Poisson Distribution Parameters: lambda (taxa) must be > 0 (lambda=%.2f)",
		lambda,
	)
}
	// Cria e retorna a instância da distribuição se o parâmetro é válido.
	return &PoissonDist{
		lambda: lambda,
	}, nil
}

// Sample gera uma amostra aleatória (um valor inteiro não negativo, retornado como float64)
// da distribuição de Poisson.
//
// Utiliza a fonte de números aleatórios (rng *rand.Rand) fornecida.
// Internamente, usa a implementação distuv.Poisson do Gonum, que
// utiliza o parâmetro Lambda (correspondente ao nosso lambda).
// Note que a distribuição de Poisson gera valores inteiros (0, 1, 2, ...),
// mas o método Rand() do Gonum retorna float64.
func (p *PoissonDist) Sample(rng *rand.Rand) float64 {
	// Cria uma instância da distribuição Poisson do Gonum.
	dist := distuv.Poisson{
		Lambda: p.lambda, // Lambda (taxa) corresponde ao nosso lambda.
		Src:    rng,      // Define a fonte de aleatoriedade.
	}
	// Gera e retorna um número aleatório (inteiro, como float64) da distribuição.
	return dist.Rand()
}

func (pd *PoissonDist) Percentile(p float64) float64 {
	if p < 0 || p > 1 {
		panic("Percentile must be in [0, 1]")
	}

	dist := distuv.Poisson{
		Lambda: pd.lambda,
	}

	cdf := 0.0
	k := 0.0

	for {
		cdf += dist.Prob(k)
		if cdf >= p {
			return k
		}
		k++
	}
}

// String retorna uma representação textual da distribuição de Poisson,
// formatada como "PoissonDist{lambda: X.XX}".
func (p *PoissonDist) String() string {
	// Formata a string de saída com o valor de lambda.
	return fmt.Sprintf("PoissonDist{lambda: %.2f}", p.lambda)
}