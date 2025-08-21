// Package frequency define perfis para geração de dados de frequência de uso,
// aplicando distribuições estatísticas e regras de valores mínimos (minValues).
package frequency

import (
	"errors"
	"math"
	"math/rand/v2"
	"simulation/internal/dists" // Importa o pacote de distribuições estatísticas, que define a interface Distribution.
)

// DeviceProfile representa um perfil estatístico usado para gerar valores de frequência.
// Ele combina uma distribuição estatística com um valor mínimo (minValue), 
// garantindo que as amostras geradas respeitem uma frequência mínima.
type householdDeviceProfile struct {
	statDist dists.Distribution // A distribuição estatística utilizada para gerar os valores.
	minValue    uint8              // O valor mínimo (inclusive) que as amostras podem assumir.
}

// minValue retorna o valor mínimo de frequência (minValue) configurado no perfil.
// Esse getter permite acessar o minValue de forma segura e encapsulada.
func (hdp *householdDeviceProfile) MinValue() uint8 {
	return hdp.minValue
}

// StatDist retorna a distribuição estatística associada ao perfil.
// Permite acessar a distribuição usada internamente, útil para diagnósticos ou reutilização.
func (hdp *householdDeviceProfile) StatDist() dists.Distribution {
	return hdp.statDist
}

// NewhouseholdDeviceProfile cria uma nova instância de householdDeviceProfile, validando se a distribuição é válida.
//
// Parâmetros:
// - dist: distribuição estatística usada para gerar valores. Não pode ser nil.
// - minValue: valor mínimo permitido para a frequência gerada.
//
// Retorna:
// - Um ponteiro para uma nova instância de householdDeviceProfile.
// - Um erro, caso a distribuição seja nil.
func newhouseholdDeviceProfile(dist dists.Distribution, minValue uint8) (DeviceProfile, error) {
	if dist == nil {
		return nil, errors.New("invalid householdDeviceProfile: distribution cannot be nil \n ")
	}

	return &householdDeviceProfile{
		statDist: dist,
		minValue:    minValue,
	}, nil
}

// generateFrequency é uma função auxiliar para gerar um valor de frequência controlado,
// baseado na distribuição e no minValue mínimo.
//
// Parâmetros:
// - rng: gerador de números aleatórios (rand.Rand).
// - minValue: valor mínimo a ser respeitado (uint8).
// - statDist: distribuição da qual o valor será amostrado.
//
// Comportamento:
// - Garante que o valor seja não-negativo.
// - Limita o valor máximo a 255 para caber no tipo uint8.
// - Aplica o minValue, retornando pelo menos esse valor.
func (hdp *householdDeviceProfile)generateFrequency(rng *rand.Rand, minValue uint8, statDist dists.Distribution) uint8 {
	val := statDist.Sample(rng) // Amostra um valor da distribuição

	// Transforma valores negativos em positivos
	if val < 0 {
		val = math.Abs(val)
	}
	// Limita o valor ao máximo possível em uint8
	if val > 255 {
		val = 255
	}

	freq := uint8(val)
	// Garante que o valor não seja menor que o minValue mínimo
	if freq < minValue {
		return minValue
	}
	return freq
}

// GenerateData é o método público da struct householdDeviceProfile para gerar
// uma frequência usando a distribuição e o minValue configurados.
//
// Parâmetros:
// - rng: gerador de números aleatórios.
//
// Retorna:
// - Um valor uint8 de frequência, respeitando o minValue e limitado a 255.
func (hdp *householdDeviceProfile) GenerateData(rng *rand.Rand) uint8 {
	return hdp.generateFrequency(rng, hdp.minValue, hdp.statDist)
}

func (hdp *householdDeviceProfile) IsIndividual() bool {
	return false
}
