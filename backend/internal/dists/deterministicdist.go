// Pacote dists fornece implementações de distribuições de probabilidade,
// incluindo distribuições contínuas (Normal, Log-Normal, Gamma),
// discretas (Poisson) e também distribuições determinísticas (valor fixo).
package dists

import (
	"fmt"
	"math/rand/v2" // Utiliza a versão 2 do pacote math/rand
)

// DeterministicDist representa uma distribuição determinística,
// isto é, sempre retorna o mesmo valor fixo.
// Esse tipo de distribuição é útil em simulações quando
// não se deseja aleatoriedade.
type DeterministicDist struct {
	// value é o valor fixo que sempre será retornado
	// ao amostrar a distribuição.
	value float64
}

// Value retorna o valor fixo associado à distribuição determinística.
func (d *DeterministicDist) Value() float64 {
	return d.value
}

// Params retorna os parâmetros da distribuição determinística.
// No caso, apenas o valor fixo.
func (d *DeterministicDist) Params() []float64 {
	return []float64{d.value}
}

// newDeterministicDist cria e retorna uma nova instância de DeterministicDist.
//
// Recebe o valor fixo como parâmetro. Não há restrições de validade,
// qualquer valor float64 é aceito (incluindo negativos, zero, infinitos ou NaN).
func newDeterministicDist(value float64) (Distribution, error) {
	// Nenhuma validação necessária, sempre válido.
	return &DeterministicDist{
		value: value,
	}, nil
}

// Sample retorna sempre o mesmo valor fixo, ignorando a fonte de
// aleatoriedade fornecida (rng).
func (d *DeterministicDist) Sample(rng *rand.Rand) float64 {
	return d.value
}

// Percentile retorna o mesmo valor fixo para qualquer percentil solicitado,
// pois em uma distribuição determinística todos os valores são iguais.
func (d *DeterministicDist) Percentile(p float64) float64 {
	return d.value
}

// String retorna uma representação textual da distribuição determinística,
// formatada como "DeterministicDist{value: X.XX}".
func (d *DeterministicDist) String() string {
	return fmt.Sprintf("DeterministicDist{value: %.2f}", d.value)
}
