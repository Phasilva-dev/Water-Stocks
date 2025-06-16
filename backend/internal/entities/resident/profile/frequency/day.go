// Package frequency define perfis de frequência para uso de aparelhos sanitários,
// permitindo gerar dados diários agregados para simulações ou análises.
package frequency

import (
	"math/rand/v2"
	"simulation/internal/entities/resident/ds/behavioral"
)

// FrequencyProfileDay agrega múltiplos FrequencyProfile, cada um representando
// a frequência de uso de um tipo específico de aparelho sanitário em um domicílio.
type FrequencyProfileDay struct {
	freqToilet      *FrequencyProfile
	freqShower      *FrequencyProfile
	freqWashBassin  *FrequencyProfile
	freqWashMachine *FrequencyProfile
	freqDishWasher  *FrequencyProfile
	freqTanque      *FrequencyProfile
}

// NewFrequencyProfileDay cria um novo FrequencyProfileDay.
//
// Recebe um mapa onde as chaves são nomes de tipos de uso (ex: "toilet", "shower")
// e os valores são os perfis de frequência correspondentes.
// Perfis não fornecidos no mapa (chaves ausentes) serão definidos como nil.
func NewFrequencyProfileDay(profiles map[string]*FrequencyProfile) *FrequencyProfileDay {
	return &FrequencyProfileDay{
		freqToilet:      profiles["toilet"],
		freqShower:      profiles["shower"],
		freqWashBassin:  profiles["washBassin"],
		freqWashMachine: profiles["washMachine"],
		freqDishWasher:  profiles["dishWasher"],
		freqTanque:      profiles["tanque"],
	}
}

// FreqToilet retorna o perfil de frequência para vaso sanitário. Pode ser nil.
func (f *FrequencyProfileDay) FreqToilet() *FrequencyProfile { return f.freqToilet }

// FreqShower retorna o perfil de frequência para chuveiro. Pode ser nil.
func (f *FrequencyProfileDay) FreqShower() *FrequencyProfile { return f.freqShower }

// FreqWashBassin retorna o perfil de frequência para lavatório. Pode ser nil.
func (f *FrequencyProfileDay) FreqWashBassin() *FrequencyProfile { return f.freqWashBassin }

// FreqWashMachine retorna o perfil de frequência para máquina de lavar roupas. Pode ser nil.
func (f *FrequencyProfileDay) FreqWashMachine() *FrequencyProfile { return f.freqWashMachine }

// FreqDishWasher retorna o perfil de frequência para lava-louças. Pode ser nil.
func (f *FrequencyProfileDay) FreqDishWasher() *FrequencyProfile { return f.freqDishWasher }

// FreqTanque retorna o perfil de frequência para tanque. Pode ser nil.
func (f *FrequencyProfileDay) FreqTanque() *FrequencyProfile { return f.freqTanque }

// validateFrequencyProfile é uma função auxiliar que gera dados de um perfil de frequência.
// Se o perfil for nil, retorna 0; caso contrário, usa o perfil para gerar os dados.
func validateFrequencyProfile(profile *FrequencyProfile, rng *rand.Rand) uint8 {
	if profile == nil {
		return 0
	}
	return profile.GenerateData(rng)
}

// GenerateData gera e retorna uma nova estrutura behavioral.Frequency.
//
// Popula a estrutura com os valores de uso diário gerados a partir de cada perfil
// de frequência configurado no FrequencyProfileDay.
// Se um perfil específico não estiver definido (nil), seu valor correspondente na estrutura
// behavioral.Frequency será 0.
//
// rng: O gerador de números aleatórios a ser usado para a geração dos dados.
func (f *FrequencyProfileDay) GenerateData(rng *rand.Rand) *behavioral.Frequency {
	return behavioral.NewFrequency(
		validateFrequencyProfile(f.freqToilet, rng),
		validateFrequencyProfile(f.freqShower, rng),
		validateFrequencyProfile(f.freqWashBassin, rng),
		validateFrequencyProfile(f.freqWashMachine, rng),
		validateFrequencyProfile(f.freqDishWasher, rng),
		validateFrequencyProfile(f.freqTanque, rng),
	)
}