// Pacote dists fornece implementações de distribuições de probabilidade,
// como Gamma e Log-Normal.
package dists

import (
	"errors"
	"fmt"
	"math/rand/v2" // Utiliza a versão 2 do pacote math/rand

	"gonum.org/v1/gonum/stat/distuv" // Biblioteca Gonum para distribuições
)

// LogNormalDist representa uma distribuição de probabilidade Log-Normal.
// Uma variável aleatória X tem distribuição Log-Normal se log(X) tem
// distribuição Normal. É parametrizada pela média (mean, µ) e
// desvio padrão (std, σ) da distribuição Normal *subjacente*.
type LogNormalDist struct {
	// mean é a média (µ) da distribuição Normal subjacente associada.
	mean float64
	// std é o desvio padrão (σ) da distribuição Normal subjacente associada.
	// Deve ser um valor estritamente positivo (> 0).
	stdDev float64
}

func (n *LogNormalDist) Params() []float64 {
	return []float64{n.mean, n.stdDev}
}

// Mean retorna a média (µ) da distribuição Normal subjacente
// associada a esta distribuição Log-Normal.
func (l *LogNormalDist) Mean() float64 {
	return l.mean
}

// Std retorna o desvio padrão (σ) da distribuição Normal subjacente
// associada a esta distribuição Log-Normal.
func (l *LogNormalDist) StdDev() float64 {
	return l.stdDev
}

// NewLogNormalDist cria e retorna uma nova instância de LogNormalDist.
//
// Recebe a média (mean) e o desvio padrão (std) da distribuição Normal
// subjacente como parâmetros.
// Retorna um erro se o desvio padrão (std) não for positivo (> 0).
func NewLogNormalDist(mean, std float64) (*LogNormalDist, error) {
	if std <= 0 {
		// Retorna erro se o desvio padrão for inválido.
		return nil, errors.New("parâmetro std (desvio padrão) deve ser > 0")
	}
	// Cria e retorna a instância da distribuição se os parâmetros são válidos.
	return &LogNormalDist{
		mean: mean,
		stdDev:  std,
	}, nil
}

// Sample gera uma amostra aleatória (um valor) da distribuição Log-Normal.
//
// Utiliza a fonte de números aleatórios (rng *rand.Rand) fornecida.
// Internamente, usa a implementação distuv.LogNormal do Gonum, que
// utiliza os parâmetros Mu (correspondente a mean) e Sigma (correspondente a std)
// da distribuição Normal subjacente.
func (l *LogNormalDist) Sample(rng *rand.Rand) float64 {
	// Cria uma instância da distribuição LogNormal do Gonum.
	dist := distuv.LogNormal{
		Mu:    l.mean, // Mu (média da Normal subjacente) corresponde ao nosso mean.
		Sigma: l.stdDev,  // Sigma (desvio padrão da Normal subjacente) corresponde ao nosso std.
		Src:   rng,    // Define a fonte de aleatoriedade.
	}
	// Gera e retorna um número aleatório da distribuição configurada.
	return dist.Rand()
}

func (l *LogNormalDist) Percentile(p float64) float64 {
	dist := distuv.LogNormal{
		Mu:    l.mean, // Mu (média da Normal subjacente) corresponde ao nosso mean.
		Sigma: l.stdDev,  // Sigma (desvio padrão da Normal subjacente) corresponde ao nosso std.
	}
	return dist.Quantile(p)
}

// String retorna uma representação textual da distribuição Log-Normal,
// formatada como "LogNormalDist{mean: X.XX, std: Y.YY}".
// Note que 'mean' e 'std' referem-se aos parâmetros da Normal subjacente.
func (l *LogNormalDist) String() string {
	// Formata a string de saída com os valores de mean e std da Normal subjacente.
	return fmt.Sprintf("LogNormalDist{mean: %.2f, std: %.2f}", l.mean, l.stdDev)
}