// Package residentdata define estruturas de dados para armazenar informações
// relacionadas a residentes ou residências, como frequências de uso de água.
package residentdata

import () // Bloco de import mantido, pode ser ajustado por goimports.

// Frequency armazena as contagens de frequência de uso diário para diferentes
// pontos de consumo de água em uma residência.
//
// Os campos internos (não exportados) registram as frequências para usos
// individuais (vaso sanitário, chuveiro, pia) e compartilhados (máquina
// de lavar roupa, lava-louças, tanque). O acesso a esses valores deve ser
// feito através dos métodos exportados (ex: FreqToilet()).
type Frequency struct {
	// Uso individual
	freqToilet     uint8 // Frequência de uso do vaso sanitário.
	freqShower     uint8 // Frequência de uso do chuveiro.
	freqWashBassin uint8 // Frequência de uso da pia do banheiro/lavatório.

	// Uso compartilhado
	freqWashMachine uint8 // Frequência de uso da máquina de lavar roupa.
	freqDishWasher  uint8 // Frequência de uso da máquina de lavar louça.
	freqTanque      uint8 // Frequência de uso do tanque.
}

// FreqToilet retorna a frequência de uso diário do vaso sanitário.
func (f *Frequency) FreqToilet() uint8 {
	return f.freqToilet
}

// FreqShower retorna a frequência de uso diário do chuveiro.
func (f *Frequency) FreqShower() uint8 {
	return f.freqShower
}

// FreqWashBassin retorna a frequência de uso diário da pia do banheiro/lavatório.
func (f *Frequency) FreqWashBassin() uint8 {
	return f.freqWashBassin
}

// --- Uso Compartilhado ---

// FreqWashMachine retorna a frequência de uso diário da máquina de lavar roupa.
func (f *Frequency) FreqWashMachine() uint8 {
	return f.freqWashMachine
}

// FreqDishWasher retorna a frequência de uso diário da máquina de lavar louça.
func (f *Frequency) FreqDishWasher() uint8 {
	return f.freqDishWasher
}

// FreqTanque retorna a frequência de uso diário do tanque.
func (f *Frequency) FreqTanque() uint8 {
	return f.freqTanque
}

// NewFrequency cria e retorna uma nova instância de Frequency com os valores
// de frequência fornecidos para cada tipo de uso.
//
// Parâmetros:
//   - toilet: Frequência (uint8) para o vaso sanitário.
//   - shower: Frequência (uint8) para o chuveiro.
//   - washBassin: Frequência (uint8) para a pia/lavatório.
//   - washMachine: Frequência (uint8) para a máquina de lavar roupa.
//   - dishWasher: Frequência (uint8) para a máquina de lavar louça.
//   - tanque: Frequência (uint8) para o tanque.
//
// Retorna:
//   - Um ponteiro (*Frequency) para a struct recém-criada e inicializada.
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