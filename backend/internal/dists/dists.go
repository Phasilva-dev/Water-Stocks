// Pacote dists fornece implementações de diversas distribuições de probabilidade
// (Normal, Poisson, Uniforme, Gamma, LogNormal, Triangular, Weibull)
// e uma função fábrica (CreateDistribution) para instanciá-las dinamicamente.
package dists

import (
	"fmt" // Usado para formatar mensagens de erro.
	"math/rand/v2" // Usado para geração de números aleatórios.
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
//   - "weibull":     scale, shape (2 parâmetros: escala λ e forma k)
//   - "loglogistic": shape, scale (2 parâmetros: forma α e escala β)
//   - "deterministic": value (1 parâmetro)
//
// Retorna um valor que implementa a interface `Distribution` e um erro `nil` em caso
// de sucesso. Se `distType` for desconhecido ou o número de `params` estiver
// incorreto para o tipo solicitado, retorna `nil` para a distribuição e um
// erro descritivo.
func CreateDistribution(distType string, params ...float64) (Distribution, error) {
	switch distType {

	case "normal":
		if len(params) != 2 {
			return nil, fmt.Errorf("invalid Distribution Factory: normal requires exactly 2 parameters: mean (média) e stdDev (desvio padrão)")
		}
		return newNormalDist(params[0], params[1])

	case "poisson":
		if len(params) != 1 {
			return nil, fmt.Errorf("invalid Distribution Factory: poisson requires exactly 1 parameter: lambda (taxa)")
		}
		return newPoissonDist(params[0])

	case "uniform":
		if len(params) != 2 {
			return nil, fmt.Errorf("invalid Distribution Factory: uniform requires exactly 2 parameters: min (mínimo) e max (máximo)")
		}
		return newUniformDist(params[0], params[1])

	case "gamma":
		if len(params) != 2 {
			return nil, fmt.Errorf("invalid Distribution Factory: gamma requires exactly 2 parameters: shape (forma) e scale (escala)")
		}
		return newGammaDist(params[0], params[1])

	case "lognormal":
		if len(params) != 2 {
			return nil, fmt.Errorf("invalid Distribution Factory: lognormal requires exactly 2 parameters: mean (média normal subjacente) e std (desvio padrão normal subjacente)")
		}
		return newLogNormalDist(params[0], params[1])

	case "triangle":
		if len(params) != 3 {
			return nil, fmt.Errorf("invalid Distribution Factory: triangle requires exactly 3 parameters: a (mínimo), b (moda) e c (máximo)")
		}
		return newTriangleDist(params[0], params[1], params[2])

	case "weibull":
		if len(params) != 2 {
			return nil, fmt.Errorf("invalid Distribution Factory: weibull requires exactly 2 parameters: shape (forma, k) e scale (escala, lambda)")
		}
		return newWeibullDist(params[0], params[1])

	case "loglogistic":
		if len(params) != 2 {
			return nil, fmt.Errorf("invalid Distribution Factory: loglogistic requires exactly 2 parameters: shape (forma, α) e scale (escala, β)")
		}
		return newLogLogisticDist(params[0], params[1])

	case "deterministic":
		if len(params) != 1 {
			return nil, fmt.Errorf("invalid Distribution Factory: deterministic requires exactly 1 parameter: value (valor fixo)")
		}
		return newDeterministicDist(params[0])

	default:
		return nil, fmt.Errorf("invalid Distribution Factory: unknown distribution type '%s'", distType)
	}
}

// --- NOTA: Este código assume que as funções construtoras e as respectivas structs
// --- que implementam a interface Distribution estão definidas neste mesmo pacote 'dists'.                          ---