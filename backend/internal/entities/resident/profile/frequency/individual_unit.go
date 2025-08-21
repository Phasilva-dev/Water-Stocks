// Package frequency define perfis para geração de dados de frequência de uso,
// aplicando distribuições estatísticas e regras de valores mínimos (shifts).
package frequency

import (
	"errors"
	"math"
	"math/rand/v2"
	"simulation/internal/dists" // Importa o pacote de distribuições estatísticas, que define a interface Distribution.
)

// DeviceProfile representa um perfil estatístico usado para gerar valores de frequência.
// Ele combina uma distribuição estatística com um valor mínimo (shift), 
// garantindo que as amostras geradas respeitem uma frequência mínima.
type individualDeviceProfile struct {
	statDist dists.Distribution // A distribuição estatística utilizada para gerar os valores.
	minValue    uint8              // O valor mínimo (inclusive) que as amostras podem assumir.
}

// Shift retorna o valor mínimo de frequência (shift) configurado no perfil.
// Esse getter permite acessar o shift de forma segura e encapsulada.
func (idp *individualDeviceProfile) MinValue() uint8 {
	return idp.minValue
}

// StatDist retorna a distribuição estatística associada ao perfil.
// Permite acessar a distribuição usada internamente, útil para diagnósticos ou reutilização.
func (idp *individualDeviceProfile) StatDist() dists.Distribution {
	return idp.statDist
}

// NewindividualDeviceProfile cria uma nova instância de individualDeviceProfile, validando se a distribuição é válida.
//
// Parâmetros:
// - dist: distribuição estatística usada para gerar valores. Não pode ser nil.
// - shift: valor mínimo permitido para a frequência gerada.
//
// Retorna:
// - Um ponteiro para uma nova instância de individualDeviceProfile.
// - Um erro, caso a distribuição seja nil.
func newindividualDeviceProfile(dist dists.Distribution, minValue uint8) (DeviceProfile, error) {
	if dist == nil {
		return nil, errors.New("invalid individualDeviceProfile: distribution cannot be nil \n ")
	}

	return &individualDeviceProfile{
		statDist: dist,
		minValue:    minValue,
	}, nil
}

// generateFrequency é uma função auxiliar para gerar um valor de frequência controlado,
// baseado na distribuição e no shift mínimo.
//
// Parâmetros:
// - rng: gerador de números aleatórios (rand.Rand).
// - shift: valor mínimo a ser respeitado (uint8).
// - statDist: distribuição da qual o valor será amostrado.
//
// Comportamento:
// - Garante que o valor seja não-negativo.
// - Limita o valor máximo a 255 para caber no tipo uint8.
// - Aplica o shift mínimo, retornando pelo menos esse valor.
func (idp *individualDeviceProfile)generateFrequency(rng *rand.Rand, shift uint8, statDist dists.Distribution) uint8 {
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
	// Garante que o valor não seja menor que o shift mínimo
	if freq < shift {
		return shift
	}
	return freq
}

// GenerateData é o método público da struct individualDeviceProfile para gerar
// uma frequência usando a distribuição e o shift configurados.
//
// Parâmetros:
// - rng: gerador de números aleatórios.
//
// Retorna:
// - Um valor uint8 de frequência, respeitando o shift e limitado a 255.
func (idp *individualDeviceProfile) GenerateData(rng *rand.Rand) uint8 {
	return idp.generateFrequency(rng, idp.minValue, idp.statDist)
}

func (idp *individualDeviceProfile) IsIndividual() bool {
	return true
}
