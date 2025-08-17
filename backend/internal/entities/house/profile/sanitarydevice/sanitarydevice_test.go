// sanitarydevice_test.go
package sanitarydevice

import (
	"math/rand/v2"
	"testing"
	"simulation/internal/dists"

)

// Teste de sucesso: cria uma instância válida de SanitaryDevice
// e verifica se todos os campos foram armazenados corretamente.
func TestNewSanitaryDeviceInstance_Success(t *testing.T) {
	mockTypo := "generic"
	mockID := uint32(123)
	mockFlow, _ := dists.CreateDistribution("deterministic", 1)
	mockDuration, _ := dists.CreateDistribution("deterministic", 2)

	instance, err := CreateSanitaryDevice(mockTypo, mockFlow, mockDuration, mockID)

	if err != nil {
		t.Fatalf("CreateSanitaryDevice retornou erro inesperado para entrada válida: %v", err)
	}
	if instance == nil {
		t.Fatal("CreateSanitaryDevice retornou nil para entrada válida")
	}
	if instance.DurationDist() != mockDuration {
		t.Errorf("DurationDist incorreto. Esperado: %v, Obtido: %v", mockDuration, instance.DurationDist())
	}
	if instance.FlowLeakDist() != mockFlow {
		t.Errorf("FlowLeakDist incorreto. Esperado: %v, Obtido: %v", mockFlow, instance.FlowLeakDist())
	}
	if instance.SanitaryDeviceID() != mockID {
		t.Errorf("ID incorreto. Esperado: %v, Obtido: %v", mockID, instance.SanitaryDeviceID())
	}
}

// Teste de erro: deviceType inválido deve retornar erro e instância nil
func TestCreateSanitaryDevice_WrongDeviceType(t *testing.T) {
	mockTypo := "p" // tipo desconhecido
	mockFlow, _ := dists.CreateDistribution("deterministic", 1)
	mockDuration, _ := dists.CreateDistribution("deterministic", 2)
	mockID := uint32(1)
	expectedErrorMsg := "invalid SanitaryDevice Factory: unknown device type: p"

	instance, err := CreateSanitaryDevice(mockTypo, mockFlow, mockDuration, mockID)

	if err == nil {
		t.Fatal("CreateSanitaryDevice não retornou erro para deviceType inválido")
	}
	if instance != nil {
		t.Errorf("Instância não-nil retornada com deviceType inválido. Instância: %v", instance)
	}
	if err.Error() != expectedErrorMsg {
		t.Errorf("Mensagem de erro inesperada. Esperado: \"%s\", Obtido: \"%s\"", expectedErrorMsg, err.Error())
	}
}

// Testa a geração de valores a partir das distribuições do dispositivo
func TestNewSanitaryDeviceInstance_GenerateDist(t *testing.T) {
	rng := rand.New(rand.NewPCG(12345, 67890))
	mockTypo := "generic"
	mockID := uint32(123)
	mockFlow, _ := dists.CreateDistribution("deterministic", 1)
	mockDuration, _ := dists.CreateDistribution("deterministic", 2)

	instance, err := CreateSanitaryDevice(mockTypo, mockFlow, mockDuration, mockID)
	if err != nil {
		t.Fatalf("CreateSanitaryDevice retornou erro inesperado: %v", err)
	}

	// Testa GenerateDuration usando RNG
	duration := instance.GenerateDuration(rng)
	if duration != 2 {
		t.Errorf("GenerateDuration incorreto. Esperado: 2, Obtido: %v", duration)
	}

	// Testa GenerateFlowLeak usando RNG
	flowLeak := instance.GenerateFlowLeak(rng)
	if flowLeak != 1 {
		t.Errorf("GenerateFlowLeak incorreto. Esperado: 1, Obtido: %v", flowLeak)
	}
}
