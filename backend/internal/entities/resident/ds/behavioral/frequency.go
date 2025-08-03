// Package behavioral define estruturas relacionadas ao comportamento de uso de água
// por residentes, incluindo frequências diárias por tipo de ponto de consumo.
package behavioral

import "fmt"

// Frequency armazena a frequência diária de uso de diferentes pontos de consumo de água
// dentro de uma residência. Os campos são privados para garantir acesso controlado via métodos.
type Frequency struct {
	deviceFrequency map[string]uint8
}

// DeviceFrequency retorna o valor referente ao dispositivo sanitario.
func (f *Frequency) DeviceFrequency(deviceType string) uint8 { 
	if freq, ok := f.deviceFrequency[deviceType]; ok {
		return freq
	}
	return 0
}

// NewFrequency cria e retorna uma nova instância de Frequency.
func NewFrequency(frequencyDevices map[string]uint8) (*Frequency, error) {
	if frequencyDevices == nil {
		return nil, fmt.Errorf("invalid frequency behavioral: frequency values map cannot be nil")
	}
	return &Frequency{
		deviceFrequency: frequencyDevices,
	}, nil
}