// Pacote dists fornece implementações de distribuições de probabilidade,
// como Gamma, Log-Normal, Normal, Poisson, Triangular e Uniforme.
package dists

import (
	"fmt"
	"math/rand/v2" // Utiliza a versão 2 do pacote math/rand

	"gonum.org/v1/gonum/stat/distuv" // Biblioteca Gonum para distribuições
)

// UniformDist representa uma distribuição de probabilidade Uniforme contínua.
// É definida por um intervalo [min, max), onde todos os valores dentro
// desse intervalo têm a mesma probabilidade de ocorrência.
// O parâmetro min deve ser estritamente menor que max.
type UniformDist struct {
	// min é o limite inferior (inclusivo) do intervalo da distribuição.
	min float64
	// max é o limite superior (exclusivo) do intervalo da distribuição.
	max float64
}

func (u *UniformDist) Params() []float64 {
	return []float64{u.min, u.max}
}

// Min retorna o limite inferior (min) da distribuição Uniforme.
func (u *UniformDist) Min() float64 {
	return u.min
}

// Max retorna o limite superior (max) da distribuição Uniforme.
func (u *UniformDist) Max() float64 {
	return u.max
}

// UniformDistnew cria e retorna uma nova instância de UniformDist *pelo valor*.
//
// Recebe o limite inferior (min) e o limite superior (max) como parâmetros.
// Retorna um erro se min for maior ou igual a max.
// Note: Esta função retorna a struct diretamente, não um ponteiro.
// A função NewUniformDist é a alternativa que retorna um ponteiro.
func UniformDistNew(min, max float64) (UniformDist, error) {

	if min > max {
		return UniformDist{}, fmt.Errorf(
		"invalid Uniform Distribution Parameters: min must be less than max (min=%.2f, max=%.2f)",
		min, max,
	)
	}

	return UniformDist{
		min: min,
		max: max,
	}, nil
}

// newUniformDist cria e retorna um *ponteiro* para uma nova instância de UniformDist.
//
// Recebe o limite inferior (min) e o limite superior (max) como parâmetros.
// Retorna um erro se min for maior ou igual a max.
// Este é o padrão mais comum em Go para construtores que podem falhar.
func newUniformDist(min, max float64) (*UniformDist, error) {
	// Verifica se min é menor que max.
	if min > max {
		return nil, fmt.Errorf(
		"invalid Uniform Distribution Parameters: min must be less than max (min=%.2f, max=%.2f)",
		min, max,
	)
	}
	// Cria e retorna um ponteiro para a instância da distribuição.
	return &UniformDist{
		min: min,
		max: max,
	}, nil
}

// Sample gera uma amostra aleatória (um valor) da distribuição Uniforme.
//
// Utiliza a fonte de números aleatórios (rng *rand.Rand) fornecida.
// Internamente, usa a implementação distuv.Uniform do Gonum.
// O valor gerado estará no intervalo [min, max).
func (u *UniformDist) Sample(rng *rand.Rand) float64 {
	// Cria uma instância da distribuição Uniforme do Gonum.
	dist := distuv.Uniform{
		Min: u.min, // Define o limite inferior.
		Max: u.max, // Define o limite superior.
		Src: rng,   // Define a fonte de aleatoriedade.
	}
	// Gera e retorna um número aleatório da distribuição configurada.
	return dist.Rand()
}

func (u *UniformDist) Percentile(p float64) float64 {
	dist := distuv.Uniform{
		Min: u.min, // Define o limite inferior.
		Max: u.max, // Define o limite superior.
	}
	return dist.Quantile(p)
}


// String retorna uma representação textual da distribuição Uniforme,
// formatada como "UniformDist{min: X.XX, max: Y.YY}".
func (u *UniformDist) String() string {
	// Formata a string de saída com os valores de min e max.
	return fmt.Sprintf("UniformDist{min: %.2f, max: %.2f}", u.min, u.max)
}