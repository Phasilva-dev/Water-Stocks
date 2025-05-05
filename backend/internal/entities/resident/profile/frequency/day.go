// Package residentprofiles fornece tipos e funções para definir perfis
// relacionados a frequências, possivelmente para simulações ou
// geração de dados, incluindo perfis diários agregados.
package frequency

import (
	"math/rand/v2"
	"simulation/internal/entities/resident/ds/behavioral" // Assumindo que behavioral contém a definição de Frequency e NewFrequency.
)

// FrequencyProfileDay agrupa múltiplos FrequencyProfile para diferentes
// pontos de uso de água (vaso sanitário, chuveiro, pia, etc.),
// representando o perfil de frequência de uso esperado para um dia inteiro.
type FrequencyProfileDay struct {
	freqToilet      *FrequencyProfile // Perfil de frequência para uso do vaso sanitário.
	freqShower      *FrequencyProfile // Perfil de frequência para uso do chuveiro.
	freqWashBassin  *FrequencyProfile // Perfil de frequência para uso da pia do banheiro/lavatório.
	freqWashMachine *FrequencyProfile // Perfil de frequência para uso da máquina de lavar roupa.
	freqDishWasher  *FrequencyProfile // Perfil de frequência para uso da máquina de lavar louça.
	freqTanque      *FrequencyProfile // Perfil de frequência para uso do tanque.
}

// NewFrequencyProfileDay cria e retorna uma nova instância de FrequencyProfileDay.
//
// Recebe um mapa onde as chaves são strings identificando o tipo de uso
// (espera-se "toilet", "shower", "washBassin", "washMachine", "dishWasher", "tanque")
// e os valores são os ponteiros para os *FrequencyProfile correspondentes.
//
// Se uma chave esperada não estiver presente no mapa `profiles` fornecido,
// o campo correspondente na struct FrequencyProfileDay resultante será `nil`.
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

// FreqToilet retorna o perfil de frequência associado ao uso do vaso sanitário.
// Pode retornar nil se nenhum perfil foi definido para este uso.
func (f *FrequencyProfileDay) FreqToilet() *FrequencyProfile {
	return f.freqToilet
}

// FreqShower retorna o perfil de frequência associado ao uso do chuveiro.
// Pode retornar nil se nenhum perfil foi definido para este uso.
func (f *FrequencyProfileDay) FreqShower() *FrequencyProfile {
	return f.freqShower
}

// FreqWashBassin retorna o perfil de frequência associado ao uso da pia do banheiro/lavatório.
// Pode retornar nil se nenhum perfil foi definido para este uso.
func (f *FrequencyProfileDay) FreqWashBassin() *FrequencyProfile {
	return f.freqWashBassin
}

// FreqWashMachine retorna o perfil de frequência associado ao uso da máquina de lavar roupa.
// Pode retornar nil se nenhum perfil foi definido para este uso.
func (f *FrequencyProfileDay) FreqWashMachine() *FrequencyProfile {
	return f.freqWashMachine
}

// FreqDishWasher retorna o perfil de frequência associado ao uso da máquina de lavar louça.
// Pode retornar nil se nenhum perfil foi definido para este uso.
func (f *FrequencyProfileDay) FreqDishWasher() *FrequencyProfile {
	return f.freqDishWasher
}

// FreqTanque retorna o perfil de frequência associado ao uso do tanque.
// Pode retornar nil se nenhum perfil foi definido para este uso.
func (f *FrequencyProfileDay) FreqTanque() *FrequencyProfile {
	return f.freqTanque
}

// validateFrequencyProfile é uma função auxiliar não exportada para gerar dados
// de um FrequencyProfile de forma segura, tratando casos onde o perfil pode ser nil.
//
// Recebe um ponteiro para um FrequencyProfile e um gerador de números aleatórios (*rand.Rand).
// Retorna 0 se o perfil fornecido for nil. Caso contrário, chama GenerateData
// no perfil usando o rng fornecido e retorna o resultado (uint8).
func validateFrequencyProfile(profile *FrequencyProfile, rng *rand.Rand) uint8 {
	if profile == nil {
		return 0 // Retorna 0 se não houver perfil definido para este tipo.
	}
	// Gera o dado usando o perfil fornecido.
	val := profile.GenerateData(rng)
	return val
}

// GenerateData gera um conjunto agregado de dados de frequência para todos os
// tipos de uso definidos neste FrequencyProfileDay, representando uma
// ocorrência diária.
//
// Utiliza os perfis internos (freqToilet, freqShower, etc.) e o gerador de
// números aleatórios (rng) fornecido para calcular a frequência de cada uso.
// Se um perfil interno específico for `nil`, a frequência gerada para esse
// tipo de uso será 0 (tratado por `validateFrequencyProfile`).
//
// Retorna um ponteiro para uma estrutura `residentdata.Frequency` contendo os
// valores de frequência (uint8) gerados para cada tipo de uso no dia.
func (f *FrequencyProfileDay) GenerateData(rng *rand.Rand) *behavioral.Frequency {
	// Valida e gera a frequência para cada tipo de uso.
	freqToilet := validateFrequencyProfile(f.freqToilet, rng)
	freqShower := validateFrequencyProfile(f.freqShower, rng)
	freqWashBassin := validateFrequencyProfile(f.freqWashBassin, rng)
	freqWashMachine := validateFrequencyProfile(f.freqWashMachine, rng)
	freqDishWasher := validateFrequencyProfile(f.freqDishWasher, rng)
	freqTanque := validateFrequencyProfile(f.freqTanque, rng)

	// Cria e retorna a estrutura Frequency do pacote residentdata com os valores gerados.
	return behavioral.NewFrequency(freqToilet, freqShower, freqWashBassin, freqWashMachine, freqDishWasher, freqTanque)
}