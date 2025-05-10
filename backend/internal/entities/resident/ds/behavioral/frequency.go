// Package behavioral define estruturas relacionadas ao comportamento de uso de água
// por residentes, incluindo frequências diárias por tipo de ponto de consumo.
package behavioral

import () // Pode ser ajustado automaticamente por ferramentas como goimports.

// Frequency armazena a frequência diária de uso de diferentes pontos de consumo
// de água, divididos entre usos individuais e compartilhados.
//
// Use os métodos públicos para acessar os dados.
type Frequency struct {
	// Uso individual
	freqToilet     uint8 // Vaso sanitário
	freqShower     uint8 // Chuveiro
	freqWashBassin uint8 // Pia do banheiro

	// Uso compartilhado
	freqWashMachine uint8 // Máquina de lavar roupa
	freqDishWasher  uint8 // Lava-louças
	freqTanque      uint8 // Tanque
}

// Métodos de acesso às frequências de uso individual.

func (f *Frequency) FreqToilet() uint8     { return f.freqToilet }
func (f *Frequency) FreqShower() uint8     { return f.freqShower }
func (f *Frequency) FreqWashBassin() uint8 { return f.freqWashBassin }

// Métodos de acesso às frequências de uso compartilhado.

func (f *Frequency) FreqWashMachine() uint8 { return f.freqWashMachine }
func (f *Frequency) FreqDishWasher() uint8  { return f.freqDishWasher }
func (f *Frequency) FreqTanque() uint8      { return f.freqTanque }

// NewFrequency cria uma nova instância de Frequency com os valores fornecidos.
func NewFrequency(
	toilet, shower, washBassin, washMachine, dishWasher, tanque uint8,
) *Frequency {
	return &Frequency{
		freqToilet:      toilet,
		freqShower:      shower,
		freqWashBassin:  washBassin,
		freqWashMachine: washMachine,
		freqDishWasher:  dishWasher,
		freqTanque:      tanque,
	}
}
