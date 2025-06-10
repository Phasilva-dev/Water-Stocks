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
	toilet     uint8 // Vaso sanitário
	shower     uint8 // Chuveiro
	washBassin uint8 // Pia do banheiro

	// Uso compartilhado
	washMachine uint8 // Máquina de lavar roupa
	dishWasher  uint8 // Lava-louças
	tanque      uint8 // Tanque
}

// Métodos de acesso às frequências de uso individual.

func (f *Frequency) Toilet() uint8     { return f.toilet }
func (f *Frequency) Shower() uint8     { return f.shower }
func (f *Frequency) WashBassin() uint8 { return f.washBassin }

// Métodos de acesso às frequências de uso compartilhado.

func (f *Frequency) WashMachine() uint8 { return f.washMachine }
func (f *Frequency) DishWasher() uint8  { return f.dishWasher }
func (f *Frequency) Tanque() uint8      { return f.tanque }

// NewFrequency cria uma nova instância de Frequency com os valores fornecidos.
func NewFrequency(
	toilet, shower, washBassin, washMachine, dishWasher, tanque uint8,
) *Frequency {
	return &Frequency{
		toilet:      toilet,
		shower:      shower,
		washBassin:  washBassin,
		washMachine: washMachine,
		dishWasher:  dishWasher,
		tanque:      tanque,
	}
}
