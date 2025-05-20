package habits // Mesmo pacote que o código a ser testado

import (
	"math/rand/v2" // Usando rand/v2
	"reflect"      // Importa o pacote reflect
	"simulation/internal/dists"
	"testing"

	"simulation/internal/entities/resident/ds/behavioral"
	"simulation/internal/entities/resident/profile/frequency"
	"simulation/internal/entities/resident/profile/routine"
)

const (
	freqToilet      = "toilet"
	freqShower      = "shower"
	freqWashBassin  = "washBassin"
	freqWashMachine = "washMachine"
	freqDishWasher  = "dishWasher"
	freqTanque      = "tanque"
	// Adicione outras chaves de frequência se existirem na implementação real
	// Ex: freqOutros = "outros"
)

// Configurações de mocks determinísticos usando instâncias reais
var mockDist, _ = dists.CreateDistribution("normal", 1, 0) // Uma Normal(1, 0) sempre retorna 1.0

// Slice com 4 cópias do mockDist para o RoutineProfile
// Nota: O comportamento de GenerateData do RoutineProfile com shift=1 e 4 Normal(1,0)
// deve produzir predictable values. Se Normal(1,0) sempre retorna 1.0,
// e o shift 1 é adicionado, o resultado mais provável seria related to [2.0, 2.0, 2.0, 2.0].
// Verifique a lógica exata de RoutineProfile.GenerateData se o teste falhar.
var routineProfileMock, _ = routine.NewRoutineProfile(
	[]dists.Distribution{mockDist, mockDist, mockDist, mockDist},
	1, // Shift de 1
)

var frequencyProfileMock, _ = frequency.NewFrequencyProfile(mockDist, 0) // Shift de 0

// Map de perfis de frequência para FrequencyProfileDay
// Cada profile no map usa mockDist (sempre retorna 1.0) e shift 0.
// FrequencyProfileDay.GenerateData deve chamar cada um e combiná-los.
var frequencyDayProfileMock = frequency.NewFrequencyProfileDay(map[string]*frequency.FrequencyProfile{
	freqToilet:     frequencyProfileMock,
	freqShower:     frequencyProfileMock,
	freqWashBassin: frequencyProfileMock,
	freqWashMachine: frequencyProfileMock,
	freqDishWasher: frequencyProfileMock,
	freqTanque:     frequencyProfileMock,
	// Adicione a sétima chave se aplicável:
	// "outraChave": frequencyProfileMock,
})

// --- Testes ---

func TestResidentDayProfile(t *testing.T) {

	// Teste o construtor e os getters
	t.Run("ConstructorAndGetters", func(t *testing.T) {
		// Usa as instâncias reais (determinísticas)
		realRoutineProfile := routineProfileMock
		realFrequencyProfileDay := frequencyDayProfileMock

		// Cria a instância de ResidentDayProfile usando o construtor
		rdp := NewResidentDayProfile(realRoutineProfile, realFrequencyProfileDay)

		// Verifica se os campos internos foram definidos corretamente (comparação de ponteiros)
		if rdp.routineProfile != realRoutineProfile {
			t.Errorf("NewResidentDayProfile: campo routineProfile incorreto. Esperado %v, Obtido %v", realRoutineProfile, rdp.routineProfile)
		}
		if rdp.frequencyProfileDay != realFrequencyProfileDay {
			t.Errorf("NewResidentDayProfile: campo frequencyProfileDay incorreto. Esperado %v, Obtido %v", realFrequencyProfileDay, rdp.frequencyProfileDay)
		}

		// Verifica se os métodos getters retornam os perfis corretos (comparação de ponteiros)
		if gotRoutine := rdp.RoutineProfile(); gotRoutine != realRoutineProfile {
			t.Errorf("RoutineProfile() getter retornou perfil incorreto. Esperado %v, Obtido %v", realRoutineProfile, gotRoutine)
		}
		if gotFrequency := rdp.FrequencyProfileDay(); gotFrequency != realFrequencyProfileDay {
			t.Errorf("FrequencyProfileDay() getter retornou perfil incorreto. Esperado %v, Obtido %v", realFrequencyProfileDay, gotFrequency)
		}
	})

	// Teste a delegação do método GenerateRoutine
	t.Run("GenerateRoutineDelegation", func(t *testing.T) {
		// Usa as instâncias reais (determinísticas)
		realRoutineProfile := routineProfileMock
		realFrequencyProfileDay := frequencyDayProfileMock // Não usado neste sub-teste

		// Define o valor que *deve ser produzido* por routineProfileMock.GenerateData()
		// com a configuração Normal(1,0) e shift=1.
		expectedRoutine := behavioral.NewRoutine([]float64{1.0, 2.0, 3.0, 5.0})

		// Cria a instância de ResidentDayProfile com as instâncias reais
		rdp := NewResidentDayProfile(realRoutineProfile, realFrequencyProfileDay)

		// Cria uma instância de RNG para passar (embora não afete o mock dist determinístico)
		rng := rand.New(rand.NewPCG(123, 456))

		// Chama o método que queremos testar
		generatedRoutine := rdp.GenerateRoutine(rng)

		// --- Verificação da delegação e retorno ---
		// Não podemos verificar flags like `generateDataCalled` nos perfis reais.
		// Assumimos que chamar o método na instância real o executa.
		// Verificamos se o valor retornado corresponde ao esperado.

		// Compara o conteúdo das structs behavioral.Routine usando reflect.DeepEqual
		if !reflect.DeepEqual(generatedRoutine, expectedRoutine) {
			t.Errorf("GenerateRoutine: retornou valor incorreto.\nEsperado: %v\nObtido:   %v", expectedRoutine, generatedRoutine)
		}

        // Opcional: Verificar se o RNG foi passado corretamente (se a GenerateData interna usar o RNG)
        // Isso requer introspecção no RoutineProfile, o que é complexo.
        // Para este nível de teste unitário no ResidentDayProfile, verificar o valor retornado é o foco principal.
        // Se quiser *garantir* que o RNG é passado, o método anterior com mocks manuais era melhor.
        // Mas se o objetivo é testar com perfis reais determinísticos, a verificação do valor basta.
	})

	// Teste a delegação do método GenerateFrequency
	t.Run("GenerateFrequencyDelegation", func(t *testing.T) {
		// Usa as instâncias reais (determinísticas)
		realRoutineProfile := routineProfileMock // Não usado neste sub-teste
		realFrequencyProfileDay := frequencyDayProfileMock

		// Define o valor que *deve ser produzido* por frequencyDayProfileMock.GenerateData()
		// com a configuração de 6 perfis, cada um gerando 1.0 (convertido para int).
		// Assumindo que behavioral.NewFrequency espera 6 ints (para 6 chaves de frequência),
		// e cada profile.GenerateData retorna 1, o resultado esperado seria 6x 1.
		expectedFrequency := behavioral.NewFrequency(1, 1, 1, 1, 1, 1) // Ajuste o valor esperado aqui se necessário

		// Cria a instância de ResidentDayProfile com as instâncias reais
		rdp := NewResidentDayProfile(realRoutineProfile, realFrequencyProfileDay)

		// Cria uma instância de RNG para passar
		rng := rand.New(rand.NewPCG(789, 1011))

		// Chama o método que queremos testar
		generatedFrequency := rdp.GenerateFrequency(rng)

		// --- Verificação da delegação e retorno ---
		// Compara o conteúdo das structs behavioral.Frequency usando reflect.DeepEqual
		if !reflect.DeepEqual(generatedFrequency, expectedFrequency) {
			t.Errorf("GenerateFrequency: retornou valor incorreto.\nEsperado: %v\nObtido:   %v", expectedFrequency, generatedFrequency)
		}

        // Opcional: Verificar se o RNG foi passado corretamente (mesma nota que acima)
	})
}