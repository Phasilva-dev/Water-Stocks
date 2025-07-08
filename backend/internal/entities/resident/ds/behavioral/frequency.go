// Package behavioral define estruturas relacionadas ao comportamento de uso de água
// por residentes, incluindo frequências diárias por tipo de ponto de consumo.
package behavioral

// (O import vazio foi removido, pois goimports ou o compilador o fariam de qualquer forma.
//  É boa prática deixá-lo vazio se não houver imports reais para evitar warnings.)

// Frequency armazena a frequência diária de uso de diferentes pontos de consumo de água
// dentro de uma residência. Os campos são privados para garantir acesso controlado via métodos.
type Frequency struct {
	// Uso individual
	toilet     uint8 // Frequência de uso do vaso sanitário.
	shower     uint8 // Frequência de uso do chuveiro.
	washBassin uint8 // Frequência de uso da pia do banheiro/lavatório.

	// Uso compartilhado (ex: entre membros da família ou para atividades coletivas)
	washMachine uint8 // Frequência de uso da máquina de lavar roupa.
	dishWasher  uint8 // Frequência de uso da lava-louças.
	tanque      uint8 // Frequência de uso do tanque.
}

// Toilet retorna a frequência de uso do vaso sanitário.
func (f *Frequency) Toilet() uint8 { return f.toilet }

// Shower retorna a frequência de uso do chuveiro.
func (f *Frequency) Shower() uint8 { return f.shower }

// WashBassin retorna a frequência de uso da pia do banheiro/lavatório.
func (f *Frequency) WashBassin() uint8 { return f.washBassin }

// WashMachine retorna a frequência de uso da máquina de lavar roupa.
func (f *Frequency) WashMachine() uint8 { return f.washMachine }

// DishWasher retorna a frequência de uso da lava-louças.
func (f *Frequency) DishWasher() uint8 { return f.dishWasher }

// Tanque retorna a frequência de uso do tanque.
func (f *Frequency) Tanque() uint8 { return f.tanque }

// NewFrequency cria e retorna uma nova instância de Frequency.
//
// Recebe como parâmetros as frequências de uso para cada ponto de consumo:
//   - toilet: Frequência para vaso sanitário.
//   - shower: Frequência para chuveiro.
//   - washBassin: Frequência para pia do banheiro/lavatório.
//   - washMachine: Frequência para máquina de lavar roupa.
//   - dishWasher: Frequência para lava-louças.
//   - tanque: Frequência para tanque.
//
// Retorna um ponteiro para a estrutura Frequency preenchida com os valores fornecidos.
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