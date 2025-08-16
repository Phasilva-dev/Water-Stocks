// sanitarydevice_test.go
package sanitarydevice // O pacote de teste é o mesmo do código sendo testado

import (
	"math/rand/v2"
	"simulation/internal/dists"
	"testing"
)

// Teste para a função construtora NewSanitaryDeviceInstance - Caso de Sucesso.
// Verifica se a instância é criada corretamente quando um dispositivo não-nil é passado,
// e se nenhum erro é retornado.
func TestNewSanitaryDeviceInstance_Success(t *testing.T) {
    // Configura um dispositivo mock com valores de exemplo
    mockID := uint32(123)
    mockFlow := dists.New
    mockDuration := int32(180)
    mockDev := &SanitaryDeviceInstance{
        id: mockID,
        flowLeakVal: mockFlow,
        durationVal: mockDuration,
    }
    expectedAmount := uint8(3)

    // Chama a função que está sendo testada e captura a instância E o erro
    instance, err := NewSanitaryDeviceInstance(mockDev, expectedAmount)

    // 1. Verifica se NENHUM erro foi retornado (caso de sucesso esperado)
    if err != nil {
        t.Fatalf("NewSanitaryDeviceInstance retornou erro inesperado para entrada válida: %v", err)
    }

    // 2. Verifica se a instância foi criada (não deve ser nil neste ponto)
    if instance == nil {
        t.Fatal("NewSanitaryDeviceInstance retornou nil para entrada válida, mas nenhum erro")
    }

    // 3. Verifica se o dispositivo armazenado na instância é o mesmo que foi passado
    if instance.device != mockDev {
         t.Errorf("NewSanitaryDeviceInstance armazenou o dispositivo incorreto. Esperado: %v, Obtido: %v", mockDev, instance.device)
    }

    // 4. Verifica se a quantidade armazenada na instância é a esperada
    if instance.amount != expectedAmount {
        t.Errorf("NewSanitaryDeviceInstance armazenou a quantidade incorreta. Esperado: %d, Obtido: %d", expectedAmount, instance.amount)
    }

    // Verificamos também os métodos acessores
    if instance.Device() != mockDev {
        t.Errorf("Método Device() após NewSanitaryDeviceInstance retornou o dispositivo incorreto. Esperado: %v, Obtido: %v", mockDev, instance.Device())
    }
    if instance.Amount() != expectedAmount {
        t.Errorf("Método Amount() após NewSanitaryDeviceInstance retornou a quantidade incorreta. Esperado: %d, Obtido: %d", expectedAmount, instance.Amount())
    }
}

// Teste para a função construtora NewSanitaryDeviceInstance - Caso de Falha (deviceType nil).
// Verifica se a função retorna nil para a instância e um erro adequado quando nil é passado
// para o parâmetro deviceType.
func TestNewSanitaryDeviceInstance_NilDeviceError(t *testing.T) {
     amount := uint8(10)
     expectedErrorMsg := "devicetype cannot be nil"

     // Chama a função com dispositivo nil e captura a instância E o erro
     instance, err := NewSanitaryDeviceInstance(nil, amount)

     // 1. Verifica se um erro foi retornado (caso de falha esperado)
     if err == nil {
         t.Fatal("NewSanitaryDeviceInstance não retornou erro quando deviceType era nil")
     }

     // 2. Verifica se a instância retornada é nil (caso de falha esperado)
     if instance != nil {
          t.Errorf("NewSanitaryDeviceInstance retornou uma instância não-nil quando deviceType era nil. Instância obtida: %v", instance)
     }

     // 3. Opcional: Verifica se a mensagem do erro é a esperada
     // Isso torna o teste mais específico sobre a causa da falha.
     if err.Error() != expectedErrorMsg {
         t.Errorf("Mensagem de erro inesperada. Esperado: \"%s\", Obtido: \"%s\"", expectedErrorMsg, err.Error())
     }
}


// Teste para o método Device() da estrutura SanitaryDeviceInstance.
// Verifica se ele retorna a referência correta do dispositivo.
// Agora inclui a verificação de erro na criação da instância.
func TestSanitaryDeviceInstance_Device(t *testing.T) {
    mockID := uint32(456)
    mockDev := &mockSanitaryDevice{id: mockID, flowLeakVal: 1.0, durationVal: 10}
    amount := uint8(2)

    // Cria uma instância usando o mock e verifica se não houve erro na criação
    instance, err := NewSanitaryDeviceInstance(mockDev, amount)
    if err != nil {
        // Se a criação falhou inesperadamente, o teste não pode continuar e deve reportar.
        t.Fatalf("Falha na criação da instância para o teste Device(): %v", err)
    }
    // Se chegou aqui, a instância foi criada com sucesso.

    // Chama o método que está sendo testado
    returnedDevice := instance.Device()

    // Verifica se o dispositivo retornado é o mesmo que foi usado na criação (comparando ponteiros)
    if returnedDevice != mockDev {
        t.Errorf("Método Device() retornou o dispositivo incorreto. Esperado: %v, Obtido: %v", mockDev, returnedDevice)
    }

    // Opcional: Verifique se o dispositivo retornado *se comporta* como o mock
    if returnedDevice.SanitaryDeviceID() != mockID {
         t.Errorf("Método Device() retornou um dispositivo com ID incorreto. Esperado: %d, Obtido: %d", mockID, returnedDevice.SanitaryDeviceID())
    }
}

// Teste para o método Amount() da estrutura SanitaryDeviceInstance.
// Verifica se ele retorna a quantidade correta.
// Agora inclui a verificação de erro na criação da instância.
func TestSanitaryDeviceInstance_Amount(t *testing.T) {
    mockDev := &mockSanitaryDevice{id: 789, flowLeakVal: 2.0, durationVal: 20}
    expectedAmount := uint8(8)

    // Cria uma instância com a quantidade esperada e verifica se não houve erro
    instance, err := NewSanitaryDeviceInstance(mockDev, expectedAmount)
     if err != nil {
        // Se a criação falhou inesperadamente, o teste não pode continuar.
        t.Fatalf("Falha na criação da instância para o teste Amount(): %v", err)
    }
    // Se chegou aqui, a instância foi criada com sucesso.

    // Chama o método que está sendo testado
    returnedAmount := instance.Amount()

    // Verifica se a quantidade retornada é a esperada
    if returnedAmount != expectedAmount {
        t.Errorf("Método Amount() retornou a quantidade incorreta. Esperado: %d, Obtido: %d", expectedAmount, returnedAmount)
    }
}