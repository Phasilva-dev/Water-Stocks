package dists

import (
	"errors"
	"fmt"
	"math/rand/v2"

	"gonum.org/v1/gonum/stat/distuv"
)

// LogNormalDist representa uma distribuição log-normal com parâmetros mean (média da normal subjacente) e std (desvio padrão da normal subjacente).
type LogNormalDist struct {
	mean float64 // Média da distribuição normal subjacente
	std  float64 // Desvio padrão da distribuição normal subjacente
}

// Mean retorna a média da distribuição normal subjacente.
func (l *LogNormalDist) Mean() float64 {
	return l.mean
}

// Std retorna o desvio padrão da distribuição normal subjacente.
func (l *LogNormalDist) Std() float64 {
	return l.std
}

// NewLogNormalDist cria uma nova distribuição log-normal com os parâmetros mean (média) e std (desvio padrão) da normal subjacente.
func NewLogNormalDist(mean, std float64) (*LogNormalDist, error) {
	if std <= 0 {
		return nil, errors.New("std must be > 0")
	}
	return &LogNormalDist{
		mean: mean,
		std:  std,
	}, nil
}

// Sample gera um valor aleatório da distribuição log-normal usando a fonte de números aleatórios fornecida.
func (l *LogNormalDist) Sample(rng *rand.Rand) float64 {
	dist := distuv.LogNormal{
		Mu:    l.mean, // Média da normal subjacente
		Sigma: l.std,  // Desvio padrão da normal subjacente
		Src:   rng,    // Fonte de aleatoriedade
	}
	return dist.Rand()
}

// String retorna uma representação em string da distribuição log-normal.
func (l *LogNormalDist) String() string {
	return fmt.Sprintf("LogNormalDist{mean: %.2f, std: %.2f}", l.mean, l.std)
}