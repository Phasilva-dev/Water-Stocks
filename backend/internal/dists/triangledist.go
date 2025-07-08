// Pacote dists fornece implementações de distribuições de probabilidade,
// como Gamma, Log-Normal, Normal, Poisson e Triangular.
package dists

import (
	"errors"
	"fmt"
	"math/rand/v2" // Utiliza a versão 2 do pacote math/rand

	"gonum.org/v1/gonum/stat/distuv" // Biblioteca Gonum para distribuições
)

// TriangleDist representa uma distribuição de probabilidade Triangular contínua.
// É definida por três parâmetros: um limite inferior (a), uma moda (b)
// e um limite superior (c). A densidade de probabilidade forma um triângulo.
// Os parâmetros devem satisfazer a condição a ≤ b ≤ c e a < c.
type TriangleDist struct {
	// a é o limite inferior (valor mínimo) da distribuição.
	a float64
	// b é a moda (valor mais provável) da distribuição.
	b float64
	// c é o limite superior (valor máximo) da distribuição.
	c float64
}

func (t *TriangleDist) Params() []float64 {
	return []float64{t.a, t.b, t.c}
}

// A retorna o limite inferior (a) da distribuição Triangular.
func (t *TriangleDist) A() float64 {
	return t.a
}

// B retorna a moda (b) da distribuição Triangular.
func (t *TriangleDist) B() float64 {
	return t.b
}

// C retorna o limite superior (c) da distribuição Triangular.
func (t *TriangleDist) C() float64 {
	return t.c
}

// NewTriangleDist cria e retorna uma nova instância de TriangleDist.
//
// Recebe o limite inferior (a), a moda (b) e o limite superior (c)
// como parâmetros.
// Retorna um erro se a condição a ≤ b ≤ c não for satisfeita, ou se a == c.
func NewTriangleDist(a, b, c float64) (*TriangleDist, error) {
	// Verifica a ordem dos parâmetros.
	if a > b || b > c {
		// Retorna erro se a ordem a ≤ b ≤ c não for válida.
		return nil, errors.New("parâmetros inválidos: deve satisfazer a ≤ b ≤ c")
	}
	// Verifica se os limites inferior e superior são iguais.
	if a == c {
		// Retorna erro se a distribuição for degenerada (linha única).
		return nil, errors.New("parâmetros a (mínimo) e c (máximo) devem ser diferentes")
	}
	// Cria e retorna a instância da distribuição se os parâmetros são válidos.
	return &TriangleDist{
		a: a,
		b: b,
		c: c,
	}, nil
}

// Sample gera uma amostra aleatória (um valor) da distribuição Triangular.
//
// Utiliza a fonte de números aleatórios (rng *rand.Rand) fornecida.
// Internamente, usa a implementação distuv.NewTriangle do Gonum. Note que
// a ordem dos parâmetros em distuv.NewTriangle é (min, max, mode, src).
func (t *TriangleDist) Sample(rng *rand.Rand) float64 {
	// Cria uma instância da distribuição Triangular do Gonum.
	// Atenção à ordem dos parâmetros: a (min), c (max), b (mode).
	dist := distuv.NewTriangle(t.a, t.c, t.b, rng)
	// Gera e retorna um número aleatório da distribuição configurada.
	return dist.Rand()
}

func (t *TriangleDist) Percentile(p float64) float64 {

	dist := distuv.NewTriangle(t.a, t.c, t.b, nil) // rand.Source é nil
	return dist.Quantile(p)
}

// String retorna uma representação textual da distribuição Triangular,
// formatada como "TriangleDist{a: X.XX, b: Y.YY, c: Z.ZZ}".
func (t *TriangleDist) String() string {
	// Formata a string de saída com os valores de a, b e c.
	return fmt.Sprintf("TriangleDist{a: %.2f, b: %.2f, c: %.2f}", t.a, t.b, t.c)
}