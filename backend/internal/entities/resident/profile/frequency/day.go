// Package frequency define perfis de frequência para uso de aparelhos sanitários,
// permitindo gerar dados diários agregados para simulações ou análises.
package frequency

import (
	"math/rand/v2"
	"simulation/internal/entities/resident/ds/behavioral"
)

// ResidentDeviceProfiles agrega múltiplos FrequencyProfile, cada um representando
// a frequência de uso de um tipo específico de aparelho sanitário em um domicílio.
type ResidentDeviceProfiles struct {
	freqDevices map[string]*DeviceProfile
}

// NewResidentDeviceProfiles cria um novo ResidentDeviceProfiles.
//
// Recebe um mapa onde as chaves são nomes de tipos de uso (ex: "toilet", "shower")
// e os valores são os perfis de frequência correspondentes.
// Perfis não fornecidos no mapa (chaves ausentes) serão definidos como nil.
func NewResidentDeviceProfiles(frequencyDeviceProfiles map[string]*DeviceProfile) *ResidentDeviceProfiles {
	return &ResidentDeviceProfiles{
		freqDevices: frequencyDeviceProfiles,
	}
}

func (f *ResidentDeviceProfiles) FreqDevice(deviceType string) *DeviceProfile {
	if profile, ok := f.freqDevices[deviceType]; ok {
		return profile
	}
	return nil
}

// validateFrequencyProfile é uma função auxiliar que gera dados de um perfil de frequência.
// Se o perfil for nil, retorna 0; caso contrário, usa o perfil para gerar os dados.
func generateFrequencyDeviceProfile(profile *DeviceProfile, rng *rand.Rand) uint8 {
	if profile == nil {
		return 0
	}
	return profile.GenerateData(rng)
}

// GenerateData gera e retorna uma nova estrutura behavioral.Frequency.
//
// Popula a estrutura com os valores de uso diário gerados a partir de cada perfil
// de frequência configurado no ResidentDeviceProfiles.
// Se um perfil específico não estiver definido (nil), seu valor correspondente na estrutura
// behavioral.Frequency será 0.
//
// rng: O gerador de números aleatórios a ser usado para a geração dos dados.
func (f *ResidentDeviceProfiles) GenerateData(rng *rand.Rand) (*behavioral.Frequency, error) {
	data := make(map[string]uint8)

	for deviceType, profile := range f.freqDevices {
		data[deviceType] = generateFrequencyDeviceProfile(profile, rng)
	}

	return behavioral.NewFrequency(data)

}