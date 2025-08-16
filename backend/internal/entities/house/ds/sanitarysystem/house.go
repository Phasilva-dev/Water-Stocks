package sanitarysystem

import (
	"simulation/internal/entities/house/profile/sanitarydevice"
	"errors"
	"fmt"
)

type SanitaryHouse struct {
	devices map[string]*sanitarydevice.SanitaryDeviceInstance
}

type SanitaryHouse struct {
	toilet *sanitarydevice.SanitaryDeviceInstance
	shower *sanitarydevice.SanitaryDeviceInstance
	washbassin *sanitarydevice.SanitaryDeviceInstance

	washmachine *sanitarydevice.SanitaryDeviceInstance
	dishwasher *sanitarydevice.SanitaryDeviceInstance
	tanque *sanitarydevice.SanitaryDeviceInstance
	amount uint8
}

func NewSanitaryHouse(
	devices map[string]sanitarydevice.SanitaryDevice, amount uint8) (*SanitaryHouse, error) {

	// --- Recupera e cria instâncias para cada tipo de dispositivo, verificando erros ---

	// Sanitário (Toilet)
	toiletDevice, ok := devices["toilet"]
	// Verifica se a chave "toilet" existe NO mapa E se o valor associado não é nil
	if !ok || toiletDevice == nil {
		return nil, errors.New("dispositivo 'toilet' faltando ou é nil no mapa fornecido")
	}
	// Chama NewSanitaryDeviceInstance e verifica o erro retornado
	toiletInstance, err := sanitarydevice.NewSanitaryDeviceInstance(toiletDevice, amount)
	if err != nil {
		// Se NewSanitaryDeviceInstance falhar, retorna o erro, possivelmente empacotando-o
		return nil, fmt.Errorf("falha ao criar instância para 'toilet': %w", err)
	}


	// Chuveiro (Shower)
	showerDevice, ok := devices["shower"]
	if !ok || showerDevice == nil {
		return nil, errors.New("dispositivo 'shower' faltando ou é nil no mapa fornecido")
	}
	showerInstance, err := sanitarydevice.NewSanitaryDeviceInstance(showerDevice, amount)
	if err != nil {
		return nil, fmt.Errorf("falha ao criar instância para 'shower': %w", err)
	}

	// Lavatório (Washbassin)
	washbassinDevice, ok := devices["wash_bassin"] // Chave no mapa é "wash_bassin"
	if !ok || washbassinDevice == nil {
		return nil, errors.New("dispositivo 'wash_bassin' faltando ou é nil no mapa fornecido")
	}
	washbassinInstance, err := sanitarydevice.NewSanitaryDeviceInstance(washbassinDevice, amount)
	if err != nil {
		return nil, fmt.Errorf("falha ao criar instância para 'wash_bassin': %w", err)
	}

	// Máquina de Lavar Roupa (Washmachine) - Quantidade fixa em 1
	washmachineDevice, ok := devices["wash_machine"] // Chave no mapa é "wash_machine"
	if !ok || washmachineDevice == nil {
		return nil, errors.New("dispositivo 'wash_machine' faltando ou é nil no mapa fornecido")
	}
	washmachineInstance, err := sanitarydevice.NewSanitaryDeviceInstance(washmachineDevice, 1)
	if err != nil {
		return nil, fmt.Errorf("falha ao criar instância para 'wash_machine': %w", err)
	}

	// Máquina de Lavar Louça (Dishwasher) - Quantidade fixa em 1
	dishwasherDevice, ok := devices["dish_washer"] // Chave no mapa é "dish_washer"
	if !ok || dishwasherDevice == nil {
		return nil, errors.New("dispositivo 'dish_washer' faltando ou é nil no mapa fornecido")
	}
	dishwasherInstance, err := sanitarydevice.NewSanitaryDeviceInstance(dishwasherDevice, 1)
	if err != nil {
		return nil, fmt.Errorf("falha ao criar instância para 'dish_washer': %w", err)
	}

	// Tanque - Quantidade fixa em 1
	tanqueDevice, ok := devices["tanque"]
	if !ok || tanqueDevice == nil {
		return nil, errors.New("dispositivo 'tanque' faltando ou é nil no mapa fornecido")
	}
	tanqueInstance, err := sanitarydevice.NewSanitaryDeviceInstance(tanqueDevice, 1)
	if err != nil {
		return nil, fmt.Errorf("falha ao criar instância para 'tanque': %w", err)
	}


	// Se todas as instâncias de dispositivo foram criadas com sucesso, retorna a casa
	return &SanitaryHouse{
		toilet:      toiletInstance,
		shower:      showerInstance,
		washbassin:  washbassinInstance,
		washmachine: washmachineInstance,
		dishwasher:  dishwasherInstance,
		tanque:      tanqueInstance,
		amount: amount, 
	}, nil // Retorna nil para o erro, indicando sucesso
}

func (h *SanitaryHouse) Toilet() *sanitarydevice.SanitaryDeviceInstance {
	return h.toilet
}

func (h *SanitaryHouse) Shower() *sanitarydevice.SanitaryDeviceInstance {
	return h.shower
}

func (h *SanitaryHouse) WashBassin() *sanitarydevice.SanitaryDeviceInstance {
	return h.washbassin
}

func (h *SanitaryHouse) WashMachine() *sanitarydevice.SanitaryDeviceInstance {
	return h.washmachine
}

func (h *SanitaryHouse) DishWasher() *sanitarydevice.SanitaryDeviceInstance {
	return h.dishwasher
}

func (h *SanitaryHouse) Tanque() *sanitarydevice.SanitaryDeviceInstance {
	return h.tanque
}

