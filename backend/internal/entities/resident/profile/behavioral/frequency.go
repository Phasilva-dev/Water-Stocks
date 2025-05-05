// Package behavioral fornece tipos e funções para definir perfis
// relacionados ao comportamento base do residente como frequências,para simulações ou
// geração de dados.
package behavioral

import (
	"simulation/internal/dists" // Assumindo que "dists" é um pacote local ou de terceiros contendo a interface Distribution.
	"errors"
	"math/rand/v2"
)

// FrequencyProfile define um perfil para gerar valores de frequência.
// Utiliza uma distribuição estatística e aplica um deslocamento (valor mínimo)
// ao resultado gerado.
type FrequencyProfile struct {
	statDist dists.Distribution // A distribuição estatística usada para gerar valores base. (Não exportado)
	shift    uint8              // O valor mínimo que a frequência gerada pode ter. (Não exportado)
}

// Shift retorna o valor de deslocamento (mínimo) configurado para o perfil.
// Este valor é o limite inferior para qualquer frequência gerada pelo perfil.
func (f *FrequencyProfile) Shift() uint8 {
	return f.shift
}

// StatDist retorna a distribuição estatística configurada para o perfil.
// Esta distribuição é usada como base para gerar os valores de frequência.
func (f *FrequencyProfile) StatDist() dists.Distribution {
	return f.statDist
}

// NewFrequencyProfile inicializa e retorna um novo FrequencyProfile com as
// configurações fornecidas.
//
// Recebe um valor de deslocamento (shift) que atuará como limite inferior
// e uma distribuição estatística (dist) para gerar os valores base.
// Retorna um ponteiro para o FrequencyProfile criado e um erro nil se a
// distribuição for válida.
//
// Retorna nil e um erro se a distribuição (dist) fornecida for nula.
//
// NOTA: Embora definido como um método em FrequencyProfile, esta função
// opera como um construtor, criando e retornando uma *nova* instância,
// ignorando o receptor 'f' (se houver).
func (f *FrequencyProfile) NewFrequencyProfile(shift uint8, dist dists.Distribution) (*FrequencyProfile, error) {
	if dist == nil {
		return nil, errors.New("a distribuição não pode ser nula")
	}

	return &FrequencyProfile{
		statDist: dist,
		shift:    shift,
	}, nil
}

// generateFrequency é um método interno responsável por gerar um único valor
// de frequência bruto.
// Ele amostra um valor da distribuição estatística do perfil, garante que
// o valor esteja dentro do intervalo [0, 255] e aplica o deslocamento (shift)
// como um limite inferior.
// Utiliza o gerador de números aleatórios (rng) fornecido.
func (f *FrequencyProfile) generateFrequency(rng *rand.Rand) uint8 {
	// Amostra um valor da distribuição.
	freq := f.statDist.Sample(rng)

	// Garante que a frequência não seja negativa.
	if freq < 0 {
		freq = 0
	}
	// Garante que a frequência não exceda o limite de uint8.
	if freq > 255 {
		freq = 255
	}

	// Converte para uint8 (arredondando para baixo implicitamente).
	roundedFreq := uint8(freq)

	// Aplica o shift como limite inferior.
	// Se a frequência arredondada for menor que o shift, retorna o shift.
	if roundedFreq < f.shift {
		return f.shift
	}

	// Caso contrário, retorna a frequência arredondada.
	return roundedFreq
}

// GenerateData gera e retorna um único valor de frequência (uint8) com base
// nas configurações do perfil (distribuição e shift).
//
// Este é o método público principal para obter um valor de frequência do perfil.
// Utiliza o gerador de números aleatórios (rng) fornecido para a geração.
func (f *FrequencyProfile) GenerateData(rng *rand.Rand) uint8 {
	// Delega a geração e aplicação das regras para o método interno.
	return f.generateFrequency(rng)
}