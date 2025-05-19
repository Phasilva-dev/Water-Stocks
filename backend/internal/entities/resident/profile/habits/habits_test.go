package habits // Mesmo pacote que o código a ser testado

import (
	"math/rand/v2" // Usando rand/v2 para consistência com testes anteriores
	"testing"
	"simulation/internal/dists"

	"simulation/internal/entities/resident/ds/behavioral" // Importa os tipos de retorno esperados
	// Importa os pacotes dos perfis, mas não as structs diretamente no mock
	"simulation/internal/entities/resident/profile/frequency"
	"simulation/internal/entities/resident/profile/routine"
)

var mockdist, _ = dists.CreateDistribution("normal", 1 , 0)
var routineProfileMock, _ = routine.NewRoutineProfile([mockdist], 1)
var frequencyProfileMock, _ = frequency.NewFrequencyProfile()


// --- Testes ---

func TestResidentDayProfile(t *testing.T) {

	// Teste o construtor e os getters
	t.Run("ConstructorAndGetters", func(t *testing.T) {
		// Cria instâncias mock
		mockRoutine := &mockRoutineProfile{}
		mockFrequency := &mockFrequencyProfileDay{}

		// Cria a instância de ResidentDayProfile usando o construtor
		rdp := NewResidentDayProfile(mockRoutine, mockFrequency)

		// Verifica se os campos internos foram definidos corretamente
		if rdp.routineProfile != mockRoutine {
			t.Errorf("NewResidentDayProfile: campo routineProfile incorreto. Esperado %v, Obtido %v", mockRoutine, rdp.routineProfile)
		}
		if rdp.frequencyProfileDay != mockFrequency {
			t.Errorf("NewResidentDayProfile: campo frequencyProfileDay incorreto. Esperado %v, Obtido %v", mockFrequency, rdp.frequencyProfileDay)
		}

		// Verifica se os métodos getters retornam os perfis corretos
		if gotRoutine := rdp.RoutineProfile(); gotRoutine != mockRoutine {
			t.Errorf("RoutineProfile() getter retornou perfil incorreto. Esperado %v, Obtido %v", mockRoutine, gotRoutine)
		}
		if gotFrequency := rdp.FrequencyProfileDay(); gotFrequency != mockFrequency {
			t.Errorf("FrequencyProfileDay() getter retornou perfil incorreto. Esperado %v, Obtido %v", mockFrequency, gotFrequency)
		}
	})

	// Teste a delegação do método GenerateRoutine
	t.Run("GenerateRoutineDelegation", func(t *testing.T) {
		// Cria mocks. Precisamos de um mock de rotina configurado.
		mockRoutine := &mockRoutineProfile{}
		mockFrequency := &mockFrequencyProfileDay{} // Mock de frequência não será usado neste teste específico

		// Define o valor que o mock de rotina deve retornar
		expectedRoutine := &behavioral.Routine{} // Cria uma instância de behavioral.Routine (pode ser vazia para o teste)
		mockRoutine.returnValue = expectedRoutine

		// Cria a instância de ResidentDayProfile com o mock de rotina
		rdp := NewResidentDayProfile(mockRoutine, mockFrequency)

		// Cria uma instância de RNG para passar
		rng := rand.New(rand.NewPCG(123, 456)) // Seed fixa para reprodutibilidade

		// Chama o método que queremos testar
		generatedRoutine := rdp.GenerateRoutine(rng)

		// Verifica se o método GenerateData do mock de rotina foi chamado
		if !mockRoutine.generateDataCalled {
			t.Error("GenerateRoutine: mockRoutineProfile.GenerateData não foi chamado")
		}

		// Verifica se o RNG correto foi passado para o mock
		if mockRoutine.receivedRNG != rng {
			t.Errorf("GenerateRoutine: mockRoutineProfile.GenerateData recebeu o RNG incorreto. Esperado %v, Obtido %v", rng, mockRoutine.receivedRNG)
		}

		// Verifica se o valor retornado por GenerateRoutine é o mesmo que o mock retornou
		if generatedRoutine != expectedRoutine {
			t.Errorf("GenerateRoutine: retornou valor incorreto. Esperado %v, Obtido %v", expectedRoutine, generatedRoutine)
		}
	})

	// Teste a delegação do método GenerateFrequency
	t.Run("GenerateFrequencyDelegation", func(t *testing.T) {
		// Cria mocks. Precisamos de um mock de frequência configurado.
		mockRoutine := &mockRoutineProfile{} // Mock de rotina não será usado neste teste específico
		mockFrequency := &mockFrequencyProfileDay{}

		// Define o valor que o mock de frequência deve retornar
		expectedFrequency := &behavioral.Frequency{} // Cria uma instância de behavioral.Frequency (pode ser vazia)
		mockFrequency.returnValue = expectedFrequency

		// Cria a instância de ResidentDayProfile com o mock de frequência
		rdp := NewResidentDayProfile(mockRoutine, mockFrequency)

		// Cria uma instância de RNG para passar
		rng := rand.New(rand.NewPCG(789, 1011)) // Outra seed fixa

		// Chama o método que queremos testar
		generatedFrequency := rdp.GenerateFrequency(rng)

		// Verifica se o método GenerateData do mock de frequência foi chamado
		if !mockFrequency.generateDataCalled {
			t.Error("GenerateFrequency: mockFrequencyProfileDay.GenerateData não foi chamado")
		}

		// Verifica se o RNG correto foi passado para o mock
		if mockFrequency.receivedRNG != rng {
			t.Errorf("GenerateFrequency: mockFrequencyProfileDay.GenerateData recebeu o RNG incorreto. Esperado %v, Obtido %v", rng, mockFrequency.receivedRNG)
		}

		// Verifica se o valor retornado por GenerateFrequency é o mesmo que o mock retornou
		if generatedFrequency != expectedFrequency {
			t.Errorf("GenerateFrequency: retornou valor incorreto. Esperado %v, Obtido %v", expectedFrequency, generatedFrequency)
		}
	})
}