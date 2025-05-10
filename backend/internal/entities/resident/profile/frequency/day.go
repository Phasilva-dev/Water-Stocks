// Package frequency define perfis de frequência para uso de aparelhos sanitários,
// permitindo gerar dados diários agregados para simulações ou análises.
package frequency

import (
	"math/rand/v2"
	"simulation/internal/entities/resident/ds/behavioral"
)

// FrequencyProfileDay agrega múltiplos FrequencyProfile para diferentes pontos
// de consumo de água em um domicílio (vaso, chuveiro, lavatório, etc.).
type FrequencyProfileDay struct {
	freqToilet      *FrequencyProfile
	freqShower      *FrequencyProfile
	freqWashBassin  *FrequencyProfile
	freqWashMachine *FrequencyProfile
	freqDishWasher  *FrequencyProfile
	freqTanque      *FrequencyProfile
}

// NewFrequencyProfileDay cria um novo FrequencyProfileDay a partir de um mapa
// associando tipos de uso a perfis. Chaves esperadas: "toilet", "shower",
// "washBassin", "washMachine", "dishWasher", "tanque".
// Perfis ausentes são tratados como nil.
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

// Getters para os perfis de frequência por tipo de uso.
// Podem retornar nil se o perfil correspondente não estiver definido.

func (f *FrequencyProfileDay) FreqToilet() *FrequencyProfile      { return f.freqToilet }
func (f *FrequencyProfileDay) FreqShower() *FrequencyProfile      { return f.freqShower }
func (f *FrequencyProfileDay) FreqWashBassin() *FrequencyProfile  { return f.freqWashBassin }
func (f *FrequencyProfileDay) FreqWashMachine() *FrequencyProfile { return f.freqWashMachine }
func (f *FrequencyProfileDay) FreqDishWasher() *FrequencyProfile  { return f.freqDishWasher }
func (f *FrequencyProfileDay) FreqTanque() *FrequencyProfile      { return f.freqTanque }

// validateFrequencyProfile é uma função auxiliar para gerar dados de um perfil,
// retornando 0 se o perfil for nil.
func validateFrequencyProfile(profile *FrequencyProfile, rng *rand.Rand) uint8 {
	if profile == nil {
		return 0
	}
	return profile.GenerateData(rng)
}

// GenerateData retorna uma estrutura Frequency com os valores de uso gerados
// a partir dos perfis definidos. Perfis não definidos resultam em valor 0.
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
