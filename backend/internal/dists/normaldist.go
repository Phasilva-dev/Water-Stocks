// Pacote dists fornece implementações de distribuições de probabilidade,
// como Gamma, Log-Normal e Normal.
package dists

import (
	"errors"
	"fmt"
	"math/rand/v2" // Utiliza a versão 2 do pacote math/rand

	"gonum.org/v1/gonum/stat/distuv" // Biblioteca Gonum para distribuições
)

// NormalDist representa uma distribuição de probabilidade Normal (ou Gaussiana).
// É definida pela média (mean, µ) e pelo desvio padrão (stdDev, σ).
type NormalDist struct {
	// mean é a média (µ), ou valor esperado, da distribuição.
	mean float64
	// stdDev é o desvio padrão (σ) da distribuição.
	// Deve ser um valor não negativo (>= 0). Um valor de 0 representa
	// uma distribuição degenerada (um ponto único na média).
	stdDev float64
}

// Mean retorna a média (µ) da distribuição Normal.
func (n *NormalDist) Mean() float64 {
	return n.mean
}

// StdDev retorna o desvio padrão (σ) da distribuição Normal.
func (n *NormalDist) StdDev() float64 {
	return n.stdDev
}

// NewNormalDist cria e retorna uma nova instância de NormalDist.
//
// Recebe a média (mean) e o desvio padrão (stdDev) como parâmetros.
// Retorna um erro se o desvio padrão (stdDev) for negativo (< 0).
// Note que um stdDev igual a 0 é permitido, resultando em uma
// distribuição degenerada.
func NewNormalDist(mean, stdDev float64) (*NormalDist, error) {
	// Verifica se o desvio padrão é negativo.
	if stdDev < 0 {
		// Retorna erro se o desvio padrão for inválido.
		// A mensagem de erro original menciona "> 0", mas o código verifica "< 0".
		// Documentando a verificação feita no código.
		return nil, errors.New("parâmetro stdDev (desvio padrão) não pode ser negativo")
	}
	// Cria e retorna a instância da distribuição se os parâmetros são válidos.
	return &NormalDist{
		mean:   mean,
		stdDev: stdDev,
	}, nil
}

// Sample gera uma amostra aleatória (um valor) da distribuição Normal.
//
// Utiliza a fonte de números aleatórios (rng *rand.Rand) fornecida.
// Internamente, usa a implementação distuv.Normal do Gonum, que
// utiliza os parâmetros Mu (correspondente a mean) e Sigma (correspondente a stdDev).
func (n *NormalDist) Sample(rng *rand.Rand) float64 {
	// Cria uma instância da distribuição Normal do Gonum.
	dist := distuv.Normal{
		Mu:    n.mean,   // Mu (média) corresponde ao nosso mean.
		Sigma: n.stdDev, // Sigma (desvio padrão) corresponde ao nosso stdDev.
		Src:   rng,      // Define a fonte de aleatoriedade.
	}
	// Gera e retorna um número aleatório da distribuição configurada.
	return dist.Rand()
}

// String retorna uma representação textual da distribuição Normal,
// formatada como "NormalDist{mean: X.XX, stdDev: Y.YY}".
func (n *NormalDist) String() string {
	// Formata a string de saída com os valores de mean e stdDev.
	return fmt.Sprintf("NormalDist{mean: %.2f, stdDev: %.2f}", n.mean, n.stdDev)
}