// Pacote dists fornece implementações de diversas distribuições de probabilidade
// (Normal, Poisson, Uniforme, Gamma, LogNormal, Triangular, Weibull)
// e uma função fábrica (CreateDistribution) para instanciá-las dinamicamente.
package dists

import (
	"errors"
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
//   - "weibull":     shape, scale (2 parâmetros: forma k e escala λ)
//   - "loglogistic": shape, scale (2 parâmetros: forma α e escala β)  // <-- ADICIONADO AQUI
//
// Retorna um valor que implementa a interface `Distribution` e um erro `nil` em caso
// de sucesso. Se `distType` for desconhecido ou o número de `params` estiver
// incorreto para o tipo solicitado, retorna `nil` para a distribuição e um
// erro descritivo.
func CreateDistribution(distType string, params ...float64) (Distribution, error) {
	// Seleciona a lógica de criação com base no tipo de distribuição.
	switch distType {
	case "normal":
		// Verifica o número correto de parâmetros para a distribuição Normal.
		if len(params) != 2 {
			return nil, errors.New("distribuição normal requer exatamente 2 parâmetros: mean (média) e stdDev (desvio padrão)")
		}
		// Chama o construtor específico da NormalDist (assumido existir).
		return NewNormalDist(params[0], params[1]) // Assume que NewNormalDist existe
	case "poisson":
		// Verifica o número correto de parâmetros para a distribuição de Poisson.
		if len(params) != 1 {
			return nil, errors.New("distribuição poisson requer exatamente 1 parâmetro: lambda (taxa)")
		}
		// Chama o construtor específico da PoissonDist (assumido existir).
		return NewPoissonDist(params[0]) // Assume que NewPoissonDist existe
	case "uniform":
		// Verifica o número correto de parâmetros para a distribuição Uniforme.
		if len(params) != 2 {
			return nil, errors.New("distribuição uniforme requer exatamente 2 parâmetros: min (mínimo) e max (máximo)")
		}
		// Chama o construtor específico da UniformDist (assumido existir).
		return NewUniformDist(params[0], params[1]) // Assume que NewUniformDist existe
	case "gamma":
		// Verifica o número correto de parâmetros para a distribuição Gamma.
		if len(params) != 2 {
			return nil, errors.New("distribuição gamma requer exatamente 2 parâmetros: shape (forma) e scale (escala)")
		}
		// Chama o construtor específico da GammaDist (assumido existir).
		return NewGammaDist(params[0], params[1]) // Assume que NewGammaDist existe
	case "lognormal":
		// Verifica o número correto de parâmetros para a distribuição LogNormal.
		if len(params) != 2 {
			return nil, errors.New("distribuição lognormal requer exatamente 2 parâmetros: mean (média normal subjacente) e std (desvio padrão normal subjacente)")
		}
		// Chama o construtor específico da LogNormalDist (assumido existir).
		return NewLogNormalDist(params[0], params[1]) // Assume que NewLogNormalDist existe
	case "triangle":
		// Verifica o número correto de parâmetros para a distribuição Triangular.
		if len(params) != 3 {
			return nil, errors.New("distribuição triangular requer exatamente 3 parâmetros: a (mínimo), b (moda) e c (máximo)")
		}
		// Chama o construtor específico da TriangleDist (assumido existir).
		return NewTriangleDist(params[0], params[1], params[2]) // Assume que NewTriangleDist existe
	case "weibull":
		// Verifica o número correto de parâmetros para a distribuição Weibull.
		if len(params) != 2 {
			return nil, errors.New("distribuição weibull requer exatamente 2 parâmetros: shape (forma, k) e scale (escala, lambda)")
		}
		// Chama o construtor específico da WeibullDist (assumido existir).
		return NewWeibullDist(params[0], params[1]) // Assume que NewWeibullDist existe
	case "loglogistic": // <-- NOVO CASE ADICIONADO
		// Verifica o número correto de parâmetros para a distribuição LogLogistic.
		if len(params) != 2 {
			return nil, errors.New("distribuição loglogistic requer exatamente 2 parâmetros: shape (forma, α) e scale (escala, β)")
		}
		// Chama o construtor específico da LogLogisticDist.
		return NewLogLogisticDist(params[0], params[1])
	default:
		// Caso o distType fornecido não corresponda a nenhuma distribuição conhecida.
		return nil, fmt.Errorf("tipo de distribuição desconhecido: %s", distType)
	}
}

// --- NOTA: Este código assume que as funções construtoras ---
// --- (NewNormalDist, NewPoissonDist, NewUniformDist,     ---
// --- NewGammaDist, NewLogNormalDist, NewTriangleDist,    ---
// --- NewWeibullDist,NewLogLogisticDist) e as respectivas structs que         ---
// --- implementam a interface Distribution estão definidas ---
// --- neste mesmo pacote 'dists'.                          ---