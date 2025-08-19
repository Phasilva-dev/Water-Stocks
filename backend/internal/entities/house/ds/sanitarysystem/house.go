package sanitarysystem

import (
	"simulation/internal/entities/house/profile/sanitarydevice"
	"fmt"
)

type SanitaryHouse struct {
	devices map[string]*sanitarydevice.SanitaryDeviceInstance
}

// NewSanitaryHouse cria uma nova instância de SanitaryHouse.
// A função itera sobre os dispositivos fornecidos e cria instâncias para cada um.
// A quantidade (amount) de cada instância é determinada pelo método IsCountable() do dispositivo:
// - Se IsCountable() for true, a instância recebe o 'amount' passado para a função.
// - Se IsCountable() for false, a instância recebe a quantidade fixa de 1.
func NewSanitaryHouse(devices map[string]sanitarydevice.SanitaryDevice,
	amount uint8) (*SanitaryHouse, error) {

	// Validação inicial: a quantidade para dispositivos contáveis não pode ser zero,
	// pois o construtor NewSanitaryDeviceInstance irá retornar um erro.
	if amount == 0 {
		return nil, fmt.Errorf("invalid amount for countable devices: must be greater than 0")
	}
	
	// Inicializa o mapa que armazenará as instâncias finais.
	// É crucial inicializar o mapa antes de adicionar itens a ele.
	houseInstances := make(map[string]*sanitarydevice.SanitaryDeviceInstance)

	// A forma correta de iterar sobre um mapa em Go é com for...range.
	for key, device := range devices {
		var finalAmount uint8

		// AQUI ESTÁ A LÓGICA PRINCIPAL:
		// Perguntamos ao próprio dispositivo qual é sua regra de contagem.
		if device.IsCountable() {
			// Se for contável, usamos o valor passado para a função.
			finalAmount = amount
		} else {
			// Caso contrário, a quantidade é sempre 1.
			finalAmount = 1
		}
        
        // --- Nota sobre o ponteiro para interface ---
		// A assinatura de NewSanitaryDeviceInstance espera um *SanitaryDevice.
		// Como 'device' é uma variável de loop, não podemos usar &device diretamente,
		// pois seu endereço é reutilizado. Criamos uma cópia local para obter um ponteiro seguro.
		// Se a assinatura fosse (device SanitaryDevice, ...), essa linha não seria necessária.
		d := device

		// Criamos a instância usando o construtor fornecido.
		// Ele já contém as validações de `amount > 0` e `device != nil`.
		instance, err := sanitarydevice.NewSanitaryDeviceInstance(&d, finalAmount)
		if err != nil {
			// Se houver um erro ao criar a instância (ex: amount == 0),
			// encapsulamos o erro com mais contexto e o retornamos, interrompendo a criação da casa.
			return nil, fmt.Errorf("failed to create instance for device '%s': %w", key, err)
		}

		// Adiciona a instância criada com sucesso ao nosso mapa final.
		houseInstances[key] = instance
	}

	// Cria a estrutura SanitaryHouse com o mapa de instâncias populado.
	house := &SanitaryHouse{
		devices: houseInstances,
	}

	// Retorna a casa criada e nenhum erro.
	return house, nil
}
