// Pacote dists fornece implementações de diversas distribuições de probabilidade
// (Normal, Poisson, Uniforme, Gamma, LogNormal, Triangular, Weibull)
// e uma função fábrica (CreateDistribution) para instanciá-las dinamicamente.
package dists

import (
	"fmt"
	"math/rand/v2"
)

// Distribution define a interface comum para todas as distribuições
// implementadas neste pacote.
// Qualquer tipo que implemente os métodos Sample e String satisfaz
// esta interface, permitindo tratamento polimórfico.
type Distribution interface {
	// Sample gera uma amostra aleatória (um valor float64) da distribuição,
	// utilizando a fonte de números aleatórios (rng) fornecida.
	Sample(rng *rand.Rand) float64

	Percentile(p float64) float64

	// String retorna uma representação textual da distribuição, geralmente
	// incluindo seu tipo e os valores de seus parâmetros formatados.
	String() string

	Params() []float64
}

// CreateDistribution é uma função fábrica que cria e retorna uma instância de uma
// distribuição específica, identificada pelo nome em `distType` e configurada
// com os `params` fornecidos.
//
// A função verifica se o `distType` é conhecido e se o número de `params`
// corresponde ao esperado para aquele tipo de distribuição.
//
// Parâmetros esperados por tipo (`distType`):
//   - "normal":      mean, stdDev (2 parâmetros: média e desvio padrão)
//   - "poisson":     lambda (1 parâmetro: taxa média de ocorrência)
//   - "uniform":     min, max (2 parâmetros: limite inferior e superior)
//   - "gamma":       shape, scale (2 parâmetros: forma e escala)
//   - "lognormal":   mean, std (2 parâmetros: média e desvio padrão da normal subjacente)
//   - "triangle":    a, b, c (3 parâmetros: mínimo, moda e máximo)
//   - "weibull":     shape, scale (2 parâmetros: escala λ e forma k)
//   - "loglogistic": shape, scale (2 parâmetros: forma α e escala β)
//   - "exponential": rate (1 parâmetro)
//
// Retorna um valor que implementa a interface `Distribution` e um erro `nil` em caso
// de sucesso. Se `distType` for desconhecido ou o número de `params` estiver
// incorreto para o tipo solicitado, retorna `nil` para a distribuição e um
// erro descritivo.
func CreateDistribution(distType string, params ...float64) (Distribution, error) {
	switch distType {
	case "normal":
		if len(params) != 2 {
			return nil, fmt.Errorf("normal distribution requires 2 parameters (mean, stdDev), but got %d", len(params))
		}
		// Assumes NewNormalDist exists and returns (Distribution, error)
		return NewNormalDist(params[0], params[1])

	case "poisson":
		if len(params) != 1 {
			return nil, fmt.Errorf("poisson distribution requires 1 parameter (lambda), but got %d", len(params))
		}
		// Assumes NewPoissonDist exists and returns (Distribution, error)
		return NewPoissonDist(params[0])

	case "uniform":
		if len(params) != 2 {
			return nil, fmt.Errorf("uniform distribution requires 2 parameters (min, max), but got %d", len(params))
		}
		// Assumes NewUniformDist exists and returns (Distribution, error)
		return NewUniformDist(params[0], params[1])

	case "gamma":
		if len(params) != 2 {
			return nil, fmt.Errorf("gamma distribution requires 2 parameters (shape, scale), but got %d", len(params))
		}
		// Assumes NewGammaDist exists and returns (Distribution, error)
		return NewGammaDist(params[0], params[1])

	case "lognormal":
		if len(params) != 2 {
			return nil, fmt.Errorf("lognormal distribution requires 2 parameters (mean, std), but got %d", len(params))
		}
		// Assumes NewLogNormalDist exists and returns (Distribution, error)
		return NewLogNormalDist(params[0], params[1])

	case "triangle":
		if len(params) != 3 {
			return nil, fmt.Errorf("triangle distribution requires 3 parameters (a, b, c), but got %d", len(params))
		}
		// Assumes NewTriangleDist exists and returns (Distribution, error)
		return NewTriangleDist(params[0], params[1], params[2])

	case "weibull":
		if len(params) != 2 {
			return nil, fmt.Errorf("weibull distribution requires 2 parameters (shape, scale), but got %d", len(params))
		}
		// Assumes NewWeibullDist exists and returns (Distribution, error)
		return NewWeibullDist(params[0], params[1])

	case "loglogistic":
		if len(params) != 2 {
			return nil, fmt.Errorf("loglogistic distribution requires 2 parameters (shape, scale), but got %d", len(params))
		}
		// Assumes NewLogLogisticDist exists and returns (Distribution, error)
		return NewLogLogisticDist(params[0], params[1])

	case "exponential":
		if len(params) != 1 {
			return nil, fmt.Errorf("exponential distribution requires 1 parameter (rate), but got %d", len(params))
		}
		// Assumes NewExponentialDist exists and returns (Distribution, error)
		return NewExponentialDist(params[0])

	default:
		// Using %q adds quotes around the string, which is helpful for debugging.
		return nil, fmt.Errorf("unknown distribution type: %q", distType)
	}
}

// --- NOTA: Este código assume que as funções construtoras ---
// --- (NewNormalDist, NewPoissonDist, NewUniformDist,     ---
// --- NewGammaDist, NewLogNormalDist, NewTriangleDist,    ---
// --- NewWeibullDist,NewLogLogisticDist) e as respectivas structs que         ---
// --- implementam a interface Distribution estão definidas ---
// --- neste mesmo pacote 'dists'.                          ---