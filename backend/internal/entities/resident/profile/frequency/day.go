// Package frequency define perfis de frequência para uso de aparelhos sanitários,
// permitindo gerar dados diários agregados para simulações ou análises.
package frequency

import (
	"math/rand/v2"                                     // Usado para gerar números aleatórios modernos (mais seguros que rand antigo).
	"simulation/internal/entities/resident/ds/behavioral" // Importa a estrutura `Frequency`, que representa o uso diário de dispositivos sanitários.
	"fmt" // Usado para formatar mensagens de erro.
)

// ResidentDeviceProfiles agrega múltiplos `DeviceProfile`s, 
// cada um representando a frequência de uso de um tipo específico 
// de aparelho sanitário em um domicílio (ex: chuveiro, vaso sanitário, etc).
type ResidentDeviceProfiles struct {
	freqDevices map[string]*DeviceProfile // Mapa com os perfis associados a nomes de dispositivos.
}

// NewResidentDeviceProfiles cria uma nova instância de `ResidentDeviceProfiles`,
// validando se todos os perfis no mapa estão corretamente definidos (não nulos).
//
// Parâmetros:
// - frequencyDeviceProfiles: mapa onde as chaves são nomes de dispositivos sanitários
//   (ex.: "toilet", "shower") e os valores são os perfis de frequência correspondentes.
//
// Retorno:
// - *ResidentDeviceProfiles: estrutura criada contendo os perfis válidos.
// - error: se algum valor do mapa for nil (perfil ausente).
func NewResidentDeviceProfiles(frequencyDeviceProfiles map[string]*DeviceProfile) (*ResidentDeviceProfiles, error) {
	for deviceType, profile := range frequencyDeviceProfiles {
		if profile == nil {
			// Se um dos dispositivos estiver sem perfil definido, retorna erro.
			return nil, fmt.Errorf("invalid frequency resident device profile: missing DeviceProfile for device type '%s'", deviceType)
		}
	}

	// Retorna a estrutura válida com os perfis carregados.
	return &ResidentDeviceProfiles{
		freqDevices: frequencyDeviceProfiles,
	}, nil
}

// FreqDevice retorna o `DeviceProfile` associado a um tipo de aparelho.
//
// Parâmetros:
// - deviceType: nome do dispositivo sanitário (ex.: "toilet").
//
// Retorno:
// - *DeviceProfile: o perfil associado.
// - error: caso não exista um perfil para o tipo fornecido.
func (f *ResidentDeviceProfiles) FreqDevice(deviceType string) (*DeviceProfile, error) {
	profile, ok := f.freqDevices[deviceType]
	if !ok {
		return nil, fmt.Errorf("invalid frequency resident device profile: no DeviceProfile found for device type '%s'", deviceType)
	}
	return profile, nil
}

// generateFrequencyDeviceProfile é uma função auxiliar que gera uma frequência de uso
// a partir de um `DeviceProfile`. Caso o perfil seja nil, retorna 0 como valor padrão.
//
// Parâmetros:
// - profile: perfil de frequência do dispositivo.
// - rng: gerador de números aleatórios usado para amostragem.
//
// Retorno:
// - uint8: valor da frequência gerada (entre 0 e 255).
func generateFrequencyDeviceProfile(profile *DeviceProfile, rng *rand.Rand) uint8 {
	if profile == nil {
		// Como fallback (não mais usado se validações forem feitas), retorna 0.
		return 0
	}
	return profile.GenerateData(rng)
}

// GenerateData gera os dados de frequência de uso diário de todos os dispositivos
// definidos em `ResidentDeviceProfiles`, retornando uma estrutura `behavioral.Frequency`.
//
// Parâmetros:
// - rng: gerador de números aleatórios usado para amostrar as frequências.
//
// Retorno:
// - *behavioral.Frequency: estrutura com os dados agregados do dia.
// - error: se ocorrer algum erro na construção da estrutura final (por exemplo, dados inconsistentes).
func (f *ResidentDeviceProfiles) GenerateData(rng *rand.Rand) (*behavioral.Frequency, error) {
	data := make(map[string]uint8) // Mapa para armazenar as frequências geradas por dispositivo.

	// Para cada dispositivo definido, gera a frequência correspondente.
	for deviceType, profile := range f.freqDevices {
		data[deviceType] = generateFrequencyDeviceProfile(profile, rng)
	}

	// Cria a estrutura final do tipo behavioral.Frequency com os dados gerados.
	return behavioral.NewFrequency(data)
}
