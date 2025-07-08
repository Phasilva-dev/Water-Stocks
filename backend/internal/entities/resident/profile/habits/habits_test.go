package habits // Mesmo pacote que o código a ser testado

import (
	"math/rand/v2" // Usando rand/v2

	"simulation/internal/dists"
	"testing"

	"simulation/internal/entities/resident/ds/behavioral"
	"simulation/internal/entities/resident/profile/frequency"
	"simulation/internal/entities/resident/profile/routine"

	"reflect" 
	"strconv"
)

const (
	freqToilet      = "toilet"
	freqShower      = "shower"
	freqWashBassin  = "washBassin"
	freqWashMachine = "washMachine"
	freqDishWasher  = "dishWasher"
	freqTanque      = "tanque"
)

// Constantes para testar daily

// Configurações de mocks determinísticos usando instâncias reais
var mockDist, _ = dists.CreateDistribution("normal", 1, 0) // Uma Normal(1, 0) sempre retorna 1.0

// Slice com 4 cópias do mockDist para o RoutineProfile
// Nota: O comportamento de GenerateData do RoutineProfile com shift=1 e 4 Normal(1,0)
// deve produzir predictable values. Se Normal(1,0) sempre retorna 1.0,
// e o shift 1 é adicionado, o resultado mais provável seria related to [2.0, 2.0, 2.0, 2.0].
// Verifique a lógica exata de RoutineProfile.GenerateData se o teste falhar.
var routineProfileMock, _ = routine.NewRoutineProfile(
	[]dists.Distribution{mockDist, mockDist, mockDist, mockDist},
	1,0, // Shift de 1
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

//Variaveis para testar weekly

var mockDist1, _ = dists.CreateDistribution("normal", 1, 0)
var mockDist2, _ = dists.CreateDistribution("normal", 2, 0)
var mockDist3, _ = dists.CreateDistribution("normal", 3, 0)

var frequencyProfileMock1, _ = frequency.NewFrequencyProfile(mockDist1, 0)
var frequencyProfileMock2, _ = frequency.NewFrequencyProfile(mockDist2, 0)
var frequencyProfileMock3, _ = frequency.NewFrequencyProfile(mockDist3, 0)


var routineProfileMock1, _ = routine.NewRoutineProfile(
	[]dists.Distribution{mockDist1, mockDist1, mockDist1, mockDist1},
	0,0,
)


var routineProfileMock2, _ = routine.NewRoutineProfile(
	[]dists.Distribution{mockDist2, mockDist2, mockDist2, mockDist2},
	0,0,
)


var routineProfileMock3, _ = routine.NewRoutineProfile(
	[]dists.Distribution{mockDist3, mockDist3, mockDist3, mockDist3},
	0,0,
)


var expectedRoutine1 = behavioral.NewRoutine([]float64{1.0, 1.0, 1.0, 1.0})
var expectedRoutine2 = behavioral.NewRoutine([]float64{2.0, 2.0, 2.0, 2.0})
var expectedRoutine3 = behavioral.NewRoutine([]float64{3.0, 3.0, 3.0, 3.0})

var expectedFrequency1 = behavioral.NewFrequency(1,1,1,1,1,1)
var expectedFrequency2 = behavioral.NewFrequency(2,2,2,2,2,2)
var expectedFrequency3 = behavioral.NewFrequency(3,3,3,3,3,3)


// ResidentDayProfile instances using the specific mocks
// Vamos criar ResidentDayProfiles que usamos os diferentes mocks
var rdp1 = NewResidentDayProfile(routineProfileMock1, frequency.NewFrequencyProfileDay(map[string]*frequency.FrequencyProfile{
	freqToilet:     frequencyProfileMock1,
	freqShower:     frequencyProfileMock1,
	freqWashBassin: frequencyProfileMock1,
	freqWashMachine: frequencyProfileMock1,
	freqDishWasher: frequencyProfileMock1,
	freqTanque:     frequencyProfileMock1,
})) // Este RDP deve gerar rotinas [1,1,1,1] e frequências (1,1,1,1,1,1)

var rdp2 = NewResidentDayProfile(routineProfileMock2, frequency.NewFrequencyProfileDay(map[string]*frequency.FrequencyProfile{
	freqToilet:     frequencyProfileMock2,
	freqShower:     frequencyProfileMock2,
	freqWashBassin: frequencyProfileMock2,
	freqWashMachine: frequencyProfileMock2,
	freqDishWasher: frequencyProfileMock2,
	freqTanque:     frequencyProfileMock2,
})) // Este RDP deve gerar rotinas [2,2,2,2] e frequências (2,2,2,2,2,2)

var rdp3 = NewResidentDayProfile(routineProfileMock3, frequency.NewFrequencyProfileDay(map[string]*frequency.FrequencyProfile{
	freqToilet:     frequencyProfileMock3,
	freqShower:     frequencyProfileMock3,
	freqWashBassin: frequencyProfileMock3,
	freqWashMachine: frequencyProfileMock3,
	freqDishWasher: frequencyProfileMock3,
	freqTanque:     frequencyProfileMock3,
})) // Este RDP deve gerar rotinas [3,3,3,3] e frequências (3,3,3,3,3,3)



// Use as variáveis expectedRoutineX e expectedFrequencyX já declaradas, pois elas
// correspondem às saídas esperadas dos rdpX correspondentes.

// --- Testes para ResidentWeeklyProfile ---

func TestResidentWeeklyProfile(t *testing.T) {

	// Teste o construtor (NewResidentWeeklyProfile)
	t.Run("NewResidentWeeklyProfile", func(t *testing.T) {
		// Teste caso válido: 1 perfil
		profiles1 := []*ResidentDayProfile{rdp1}
		wp1, err1 := NewResidentWeeklyProfile(profiles1)
		if err1 != nil {
			t.Errorf("Esperado nenhum erro para 1 perfil, obtido %v", err1)
		}
		if wp1 == nil {
			t.Error("Esperado perfil semanal não nulo para 1 perfil")
		}
		if len(wp1.Profiles()) != 1 {
			t.Errorf("Esperado tamanho de profiles 1, obtido %d", len(wp1.profiles))
		}
		if wp1.profiles[0] != rdp1 {
			t.Error("Referência do ResidentDayProfile no índice 0 incorreta")
		}


		// Teste caso válido: 3 perfis (para testar o wrap around depois)
		profiles3 := []*ResidentDayProfile{rdp1, rdp2, rdp3}
		wp3, err3 := NewResidentWeeklyProfile(profiles3)
		if err3 != nil {
			t.Errorf("Esperado nenhum erro para 3 perfis, obtido %v", err3)
		}
		if wp3 == nil {
			t.Error("Esperado perfil semanal não nulo para 3 perfis")
		}
		if len(wp3.Profiles()) != 3 {
			t.Errorf("Esperado tamanho de profiles 3, obtido %d", len(wp3.profiles))
		}
		if wp3.profiles[0] != rdp1 || wp3.profiles[1] != rdp2 || wp3.profiles[2] != rdp3 {
            t.Error("Referências dos ResidentDayProfiles para 3 perfis incorretas")
        }


		// Teste caso válido: 7 perfis
		profiles7 := []*ResidentDayProfile{rdp1, rdp2, rdp3, rdp1, rdp2, rdp3, rdp1} // Exemplo de 7 perfis
		wp7, err7 := NewResidentWeeklyProfile(profiles7)
		if err7 != nil {
			t.Errorf("Esperado nenhum erro para 7 perfis, obtido %v", err7)
		}
		if wp7 == nil {
			t.Error("Esperado perfil semanal não nulo para 7 perfis")
		}
		if len(wp7.Profiles()) != 7 {
			t.Errorf("Esperado tamanho de profiles 7, obtido %d", len(wp7.profiles))
		}


		// Teste caso inválido: 0 perfis
		wp0, err0 := NewResidentWeeklyProfile([]*ResidentDayProfile{})
		if err0 == nil {
			t.Error("Esperado erro para 0 perfis, obtido nil")
		}
		if wp0 != nil {
			t.Error("Esperado perfil semanal nulo para 0 perfis")
		}


		// Teste caso inválido: 8 perfis
		profiles8 := make([]*ResidentDayProfile, 8) // Cria um slice de 8 elementos (não importa o conteúdo para este teste)
		wp8, err8 := NewResidentWeeklyProfile(profiles8)
		if err8 == nil {
			t.Error("Esperado erro para 8 perfis, obtido nil")
		}
		if wp8 != nil {
			t.Error("Esperado perfil semanal nulo para 8 perfis")
		}
	})

	// Teste o método GenerateRoutine, incluindo a lógica de normalizeDay
	t.Run("GenerateRoutine", func(t *testing.T) {
		// Crie um weekly profile com um número de dias específico (ex: 3) para testar o wrap around
		profiles := []*ResidentDayProfile{rdp1, rdp2, rdp3} // Tamanho 3
		wp, err := NewResidentWeeklyProfile(profiles)
		if err != nil {
			t.Fatalf("Falha ao criar o perfil semanal: %v", err)
		}

		rng := rand.New(rand.NewPCG(2023, 11)) // RNG determinístico

		// Teste vários dias, verificando qual perfil diário deve ser usado (day % 3)
		tests := []struct {
			day            uint8
			expectedRoutine *behavioral.Routine
			expectedDayIdx uint8 // O índice esperado no slice de profiles
		}{
			{day: 0, expectedRoutine: expectedRoutine1, expectedDayIdx: 0}, // 0 % 3 = 0 -> rdp1
			{day: 1, expectedRoutine: expectedRoutine2, expectedDayIdx: 1}, // 1 % 3 = 1 -> rdp2
			{day: 2, expectedRoutine: expectedRoutine3, expectedDayIdx: 2}, // 2 % 3 = 2 -> rdp3
			{day: 3, expectedRoutine: expectedRoutine1, expectedDayIdx: 0}, // 3 % 3 = 0 -> rdp1 (wrap)
			{day: 4, expectedRoutine: expectedRoutine2, expectedDayIdx: 1}, // 4 % 3 = 1 -> rdp2 (wrap)
			{day: 5, expectedRoutine: expectedRoutine3, expectedDayIdx: 2}, // 5 % 3 = 2 -> rdp3 (wrap)
			{day: 6, expectedRoutine: expectedRoutine1, expectedDayIdx: 0}, // 6 % 3 = 0 -> rdp1 (wrap)
			{day: 7, expectedRoutine: expectedRoutine2, expectedDayIdx: 1}, // 7 % 3 = 1 -> rdp2 (wrap)
			{day: 8, expectedRoutine: expectedRoutine3, expectedDayIdx: 2}, // 8 % 3 = 2 -> rdp3 (wrap)
			// Adicione mais dias se desejar, a lógica de % deve se manter
		}

		for _, tt := range tests {
            // Usamos t.Run para cada dia para melhor granularidade dos resultados
			t.Run("Day"+strconv.Itoa(int(tt.day)), func(t *testing.T) {
				generatedRoutine := wp.GenerateRoutine(tt.day, rng)

				// Verifica se a rotina gerada é a esperada para aquele dia (e o perfil diário correspondente)
				if !reflect.DeepEqual(generatedRoutine, tt.expectedRoutine) {
					t.Errorf("GenerateRoutine para o dia %d (índice %d): rotina incorreta.\nEsperado: %v\nObtido:   %v",
						tt.day, tt.expectedDayIdx, tt.expectedRoutine, generatedRoutine)
				}
                // Não podemos facilmente verificar *se* o GenerateRoutine do rdp correto foi chamado,
                // mas verificar a saída garante que a lógica de seleção de perfil (normalizeDay) funcionou corretamente.
			})
		}
	})


	// Teste o método GenerateFrequency, incluindo a lógica de normalizeDay
	t.Run("GenerateFrequency", func(t *testing.T) {
		// Use o mesmo weekly profile de 3 dias
		profiles := []*ResidentDayProfile{rdp1, rdp2, rdp3} // Tamanho 3
		wp, err := NewResidentWeeklyProfile(profiles)
		if err != nil {
			t.Fatalf("Falha ao criar o perfil semanal: %v", err)
		}

		rng := rand.New(rand.NewPCG(2023, 11)) // Outro RNG determinístico

		// Teste vários dias, verificando qual perfil diário deve ser usado (day % 3)
		tests := []struct {
			day              uint8
			expectedFrequency *behavioral.Frequency
			expectedDayIdx   uint8 // O índice esperado no slice de profiles
		}{
			{day: 0, expectedFrequency: expectedFrequency1, expectedDayIdx: 0}, // 0 % 3 = 0 -> rdp1
			{day: 1, expectedFrequency: expectedFrequency2, expectedDayIdx: 1}, // 1 % 3 = 1 -> rdp2
			{day: 2, expectedFrequency: expectedFrequency3, expectedDayIdx: 2}, // 2 % 3 = 2 -> rdp3
			{day: 3, expectedFrequency: expectedFrequency1, expectedDayIdx: 0}, // 3 % 3 = 0 -> rdp1 (wrap)
			{day: 4, expectedFrequency: expectedFrequency2, expectedDayIdx: 1}, // 4 % 3 = 1 -> rdp2 (wrap)
			{day: 5, expectedFrequency: expectedFrequency3, expectedDayIdx: 2}, // 5 % 3 = 2 -> rdp3 (wrap)
			{day: 6, expectedFrequency: expectedFrequency1, expectedDayIdx: 0}, // 6 % 3 = 0 -> rdp1 (wrap)
			{day: 7, expectedFrequency: expectedFrequency2, expectedDayIdx: 1}, // 7 % 3 = 1 -> rdp2 (wrap)
            {day: 8, expectedFrequency: expectedFrequency3, expectedDayIdx: 2}, // 8 % 3 = 2 -> rdp3 (wrap)
		}

		for _, tt := range tests {
            // Usamos t.Run para cada dia
			t.Run("Day"+strconv.Itoa(int(tt.day)), func(t *testing.T) {
				generatedFrequency := wp.GenerateFrequency(tt.day, rng)

				// Verifica se a frequência gerada é a esperada para aquele dia
				if !reflect.DeepEqual(generatedFrequency, tt.expectedFrequency) {
					t.Errorf("GenerateFrequency para o dia %d (índice %d): frequência incorreta.\nEsperado: %v\nObtido:   %v",
						tt.day, tt.expectedDayIdx, tt.expectedFrequency, generatedFrequency)
				}
			})
		}
	})
}