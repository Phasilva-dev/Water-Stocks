package dists

import (
	"errors"
	"fmt"
	"math/rand/v2"

	"gonum.org/v1/gonum/stat/distuv"
)

// TriangleDist representa uma distribuição triangular com parâmetros a (mínimo), b (moda) e c (máximo).
type TriangleDist struct {
	a float64 // Limite inferior
	b float64 // Moda
	c float64 // Limite superior
}

// A retorna o limite inferior da distribuição.
func (t *TriangleDist) A() float64 {
	return t.a
}

// B retorna a moda da distribuição.
func (t *TriangleDist) B() float64 {
	return t.b
}

// C retorna o limite superior da distribuição.
func (t *TriangleDist) C() float64 {
	return t.c
}

// NewTriangleDist cria uma nova distribuição triangular com os parâmetros a (mínimo), b (moda) e c (máximo).
func NewTriangleDist(a, b, c float64) (*TriangleDist, error) {
	if a > b || b > c {
		return nil, errors.New("invalid triangular parameters: must satisfy a ≤ b ≤ c")
	}
	if a == c {
		return nil, errors.New("a and c must be different")
	}
	return &TriangleDist{
		a: a,
		b: b,
		c: c,
	}, nil
}

// Sample gera um valor aleatório da distribuição triangular usando a fonte de números aleatórios fornecida.
func (t *TriangleDist) Sample(rng *rand.Rand) float64 {
	// Criar a distribuição triangular do Gonum
	dist := distuv.NewTriangle(t.a, t.c, t.b, rng) // a: limite inferior, c: limite superior, b: moda
	return dist.Rand()
}

// String retorna uma representação em string da distribuição triangular.
func (t *TriangleDist) String() string {
	return fmt.Sprintf("TriangleDist{a: %.2f, b: %.2f, c: %.2f}", t.a, t.b, t.c)
}