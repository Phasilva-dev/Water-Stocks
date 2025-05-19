package frequency

import (
	"math/rand/v2"
	"simulation/internal/dists"
	"testing"
	"math"
	"simulation/internal/entities/resident/ds/behavioral"
)

// MockDistribution implementa a interface dists.Distribution para testes
type MockDistribution struct {
	returnValue float64
}

// Sample retorna um valor pré-definido para teste
func (m *MockDistribution) Sample(rng *rand.Rand) float64 {
	return m.returnValue
}

// String retorna uma representação textual da distribuição de teste
func (m *MockDistribution) String() string {
	return "MockDistribution"
}

func TestNewFrequencyProfile(t *testing.T) {
	// Testa criação com distribuição válida
	fp := &FrequencyProfile{}
	shift := uint8(10)
	dist := &MockDistribution{returnValue: 50.0}
	
	profile, err := fp.NewFrequencyProfile(dist, shift)
	if err != nil {
		t.Errorf("Erro inesperado ao criar perfil de frequência: %v", err)
	}
	
	if profile.shift != shift {
		t.Errorf("Shift esperado: %d, obtido: %d", shift, profile.shift)
	}
	
	if profile.statDist != dist {
		t.Errorf("Distribuição estatística não corresponde à fornecida")
	}
	
	// Testa criação com distribuição nula
	profile, err = fp.NewFrequencyProfile(nil, shift)
	if err == nil {
		t.Error("Deveria retornar erro ao criar perfil com distribuição nula")
	}
	
	if profile != nil {
		t.Error("Deveria retornar perfil nulo ao criar com distribuição nula")
	}
}

func TestShift(t *testing.T) {
	shift := uint8(15)
	fp := &FrequencyProfile{
		shift:    shift,
		statDist: &MockDistribution{},
	}
	
	if fp.Shift() != shift {
		t.Errorf("Shift esperado: %d, obtido: %d", shift, fp.Shift())
	}
}

func TestStatDist(t *testing.T) {
	dist := &MockDistribution{returnValue: 100.0}
	fp := &FrequencyProfile{
		shift:    10,
		statDist: dist,
	}
	
	if fp.StatDist() != dist {
		t.Error("StatDist não retornou a distribuição correta")
	}
}

func TestGenerateFrequency(t *testing.T) {
	testCases := []struct {
		name       string
		dist       *MockDistribution
		shift      uint8
		expected   uint8
		description string
	}{
		{
			name:       "Valor normal dentro dos limites",
			dist:       &MockDistribution{returnValue: 100.0},
			shift:      10,
			expected:   100,
			description: "Valor que não precisa ser ajustado",
		},
		{
			name:       "Valor abaixo do shift",
			dist:       &MockDistribution{returnValue: 5.0},
			shift:      10,
			expected:   10,
			description: "Valor abaixo do shift deve ser ajustado para o shift",
		},
		{
			name:       "Valor negativo convertido para positivo",
			dist:       &MockDistribution{returnValue: -50.0},
			shift:      10,
			expected:   50,
			description: "Valor negativo deve ser convertido para positivo",
		},
		{
			name:       "Valor negativo convertido e ajustado pelo shift",
			dist:       &MockDistribution{returnValue: -5.0},
			shift:      10,
			expected:   10,
			description: "Valor negativo convertido que fica abaixo do shift",
		},
		{
			name:       "Valor acima do limite máximo",
			dist:       &MockDistribution{returnValue: 300.0},
			shift:      10,
			expected:   255,
			description: "Valor acima de 255 deve ser limitado a 255",
		},
	}
	
	rng := rand.New(rand.NewPCG(1, 2))
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := generateFrequency(rng, tc.shift, tc.dist)
			if result != tc.expected {
				t.Errorf("%s: esperado %d, obtido %d", tc.description, tc.expected, result)
			}
		})
	}
}

func TestGenerateData(t *testing.T) {
	testCases := []struct {
		name       string
		dist       *MockDistribution
		shift      uint8
		expected   uint8
	}{
		{
			name:       "Gera valor normal",
			dist:       &MockDistribution{returnValue: 100.0},
			shift:      10,
			expected:   100,
		},
		{
			name:       "Gera valor com shift aplicado",
			dist:       &MockDistribution{returnValue: 5.0},
			shift:      10,
			expected:   10,
		},
		{
			name:       "Gera valor limitado",
			dist:       &MockDistribution{returnValue: 300.0},
			shift:      10,
			expected:   255,
		},
	}
	
	rng := rand.New(rand.NewPCG(1, 2))
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fp := &FrequencyProfile{
				statDist: tc.dist,
				shift:    tc.shift,
			}
			
			result := fp.GenerateData(rng)
			if result != tc.expected {
				t.Errorf("esperado %d, obtido %d", tc.expected, result)
			}
		})
	}
}


// TestGenerateRealData testa o comportamento estatístico da função GenerateData
// usando uma distribuição de Poisson real e executando um grande número de amostras
// para verificar se os resultados estão dentro da margem estatística esperada.
func TestGenerateRealData(t *testing.T) {
	// Pula o teste por padrão devido ao alto custo computacional
	//t.Skip("Este teste é computacionalmente intensivo e deve ser executado apenas quando necessário")
	
	// Para criar este teste precisamos importar a implementação real de Poisson
	// Crie uma distribuição de Poisson com lambda = 1.38
	poissonDist, err := dists.CreateDistribution("poisson", 1.38)
	if err != nil {
		t.Fatalf("Erro ao criar distribuição de Poisson: %v", err)
	}
	
	shift := uint8(0)
	fp := &FrequencyProfile{
		statDist: poissonDist,
		shift:    shift,
	}
	
	// Cria um gerador de números aleatórios com seed fixa para reprodutibilidade
	rng := rand.New(rand.NewPCG(12345, 67890))
	
	// Número de simulações
	numSimulations := 1_000_000
	
	// Conta a frequência de cada valor
	frequencies := make(map[uint8]int)
	
	// Executa as simulações
	for i := 0; i < numSimulations; i++ {
		value := fp.GenerateData(rng)
		frequencies[value]++
	}
	
	// Calcula as probabilidades teóricas para uma distribuição de Poisson
	// Verifica uma faixa maior de valores (0 até 10)
	lambda := 1.38
	maxValueToCheck := 255
	
	// Função para calcular fatorial
	factorial := func(n int) float64 {
		if n <= 1 {
			return 1
		}
		result := float64(1)
		for i := 2; i <= n; i++ {
			result *= float64(i)
		}
		return result
	}
	
	// Calcula a probabilidade de Poisson para um valor k: P(X=k) = (lambda^k * e^-lambda) / k!
	poissonProb := func(k int) float64 {
		return math.Pow(lambda, float64(k)) * math.Exp(-lambda) / factorial(k)
	}
	
	// Tolerância para o teste (margem de erro aceitável)
	tolerance := 0.01 // 1% de margem de erro
	
	// Coleta estatísticas para exibir no relatório
	var failCount int
	var succeedCount int
	
	// Verifica a faixa completa de valores
	for k := 0; k <= maxValueToCheck; k++ {
		theoreticalProb := poissonProb(k)
		observedProb := float64(frequencies[uint8(k)]) / float64(numSimulations)
		diff := math.Abs(observedProb - theoreticalProb)
		
		if diff > tolerance {
			t.Errorf("Para o valor %d: probabilidade teórica = %.4f, observada = %.4f, diferença = %.4f, além da tolerância de %.4f",
				k, theoreticalProb, observedProb, diff, tolerance)
			failCount++
		} else {
			t.Logf("Para o valor %d: probabilidade teórica = %.4f, observada = %.4f, dentro da tolerância", 
				k, theoreticalProb, observedProb)
			succeedCount++
		}
	}
	
	// Verifica a probabilidade acumulada para valores maiores que maxValueToCheck
	// Probabilidade teórica acumulada P(X > maxValueToCheck)
	cumulativeTheoreticalProb := 0.0
	for k := 0; k <= maxValueToCheck; k++ {
		cumulativeTheoreticalProb += poissonProb(k)
	}
	remainingTheoreticalProb := 1.0 - cumulativeTheoreticalProb
	
	// Probabilidade observada acumulada P(X > maxValueToCheck)
	cumulativeObservedCount := 0
	for k := 0; k <= maxValueToCheck; k++ {
		cumulativeObservedCount += frequencies[uint8(k)]
	}
	remainingObservedCount := numSimulations - cumulativeObservedCount
	remainingObservedProb := float64(remainingObservedCount) / float64(numSimulations)
	
	// Verifica se a diferença na cauda está dentro da tolerância
	diffInTail := math.Abs(remainingObservedProb - remainingTheoreticalProb)
	if diffInTail > tolerance {
		t.Errorf("Para valores > %d (cauda): probabilidade teórica = %.4f, observada = %.4f, diferença = %.4f, além da tolerância de %.4f",
			maxValueToCheck, remainingTheoreticalProb, remainingObservedProb, diffInTail, tolerance)
		failCount++
	} else {
		t.Logf("Para valores > %d (cauda): probabilidade teórica = %.4f, observada = %.4f, dentro da tolerância", 
			maxValueToCheck, remainingTheoreticalProb, remainingObservedProb)
		succeedCount++
	}
	
	// Resultados gerais do teste
	t.Logf("Resumo: %d verificações dentro da tolerância, %d fora da tolerância", succeedCount, failCount)
	
	// Estatísticas básicas
	var sum int
	var count int
	var max uint8
	for value, freq := range frequencies {
		sum += int(value) * freq
		count += freq
		if value > max {
			max = value
		}
	}
	
	mean := float64(sum) / float64(count)
	t.Logf("Média esperada (lambda): %.4f, Média observada: %.4f", lambda, mean)
	t.Logf("Valor máximo observado: %d", max)
}

// --- Novos testes para FrequencyProfileDay ---

func TestNewFrequencyProfileDay(t *testing.T) {
	// Cria mocks para os perfis - Estes não chamam NewFrequencyProfile, são mocks diretos
	mockToilet := &FrequencyProfile{statDist: &MockDistribution{returnValue: 1}, shift: 1}
	mockShower := &FrequencyProfile{statDist: &MockDistribution{returnValue: 2}, shift: 2}
	mockWashBassin := &FrequencyProfile{statDist: &MockDistribution{returnValue: 3}, shift: 3}
	mockWashMachine := &FrequencyProfile{statDist: &MockDistribution{returnValue: 4}, shift: 4}
	mockDishWasher := &FrequencyProfile{statDist: &MockDistribution{returnValue: 5}, shift: 5}
	mockTanque := &FrequencyProfile{statDist: &MockDistribution{returnValue: 6}, shift: 6}

	testCases := []struct {
		name     string
		profiles map[string]*FrequencyProfile
		expected map[string]*FrequencyProfile // Usamos um mapa esperado para facilitar a verificação
	}{
		{
			name: "Todos os perfis fornecidos",
			profiles: map[string]*FrequencyProfile{
				"toilet": mockToilet,
				"shower": mockShower,
				"washBassin": mockWashBassin,
				"washMachine": mockWashMachine,
				"dishWasher": mockDishWasher,
				"tanque": mockTanque,
			},
			expected: map[string]*FrequencyProfile{
				"toilet": mockToilet,
				"shower": mockShower,
				"washBassin": mockWashBassin,
				"washMachine": mockWashMachine,
				"dishWasher": mockDishWasher,
				"tanque": mockTanque,
			},
		},
		{
			name: "Alguns perfis fornecidos, outros nulos",
			profiles: map[string]*FrequencyProfile{
				"toilet": mockToilet,
				"shower": mockShower,
			},
			expected: map[string]*FrequencyProfile{
				"toilet": mockToilet,
				"shower": mockShower,
				"washBassin": nil,
				"washMachine": nil,
				"dishWasher": nil,
				"tanque": nil,
			},
		},
		{
			name: "Mapa vazio fornecido",
			profiles: map[string]*FrequencyProfile{},
			expected: map[string]*FrequencyProfile{
				"toilet": nil,
				"shower": nil,
				"washBassin": nil,
				"washMachine": nil,
				"dishWasher": nil,
				"tanque": nil,
			},
		},
		{
			name: "Mapa nulo fornecido",
			profiles: nil,
			expected: map[string]*FrequencyProfile{
				"toilet": nil,
				"shower": nil,
				"washBassin": nil,
				"washMachine": nil,
				"dishWasher": nil,
				"tanque": nil,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			profileDay := NewFrequencyProfileDay(tc.profiles)

			// Verifica se o perfil day foi criado (não deve ser nulo)
			if profileDay == nil {
				t.Fatal("NewFrequencyProfileDay retornou nil")
			}

			// Verifica cada campo usando os getters
			if profileDay.FreqToilet() != tc.expected["toilet"] {
				t.Errorf("FreqToilet: esperado %v, obtido %v", tc.expected["toilet"], profileDay.FreqToilet())
			}
			if profileDay.FreqShower() != tc.expected["shower"] {
				t.Errorf("FreqShower: esperado %v, obtido %v", tc.expected["shower"], profileDay.FreqShower())
			}
			if profileDay.FreqWashBassin() != tc.expected["washBassin"] {
				t.Errorf("FreqWashBassin: esperado %v, obtido %v", tc.expected["washBassin"], profileDay.FreqWashBassin())
			}
			if profileDay.FreqWashMachine() != tc.expected["washMachine"] {
				t.Errorf("FreqWashMachine: esperado %v, obtido %v", tc.expected["washMachine"], profileDay.FreqWashMachine())
			}
			if profileDay.FreqDishWasher() != tc.expected["dishWasher"] {
				t.Errorf("FreqDishWasher: esperado %v, obtido %v", tc.expected["dishWasher"], profileDay.FreqDishWasher())
			}
			if profileDay.FreqTanque() != tc.expected["tanque"] {
				t.Errorf("FreqTanque: esperado %v, obtido %v", tc.expected["tanque"], profileDay.FreqTanque())
			}
		})
	}
}

func TestFrequencyProfileDayGetters(t *testing.T) {
	// Teste simples para garantir que os getters retornam os valores corretos
	// já cobertos implicitamente por TestNewFrequencyProfileDay, mas explícito.

	mockProfile1 := &FrequencyProfile{} // Valor real não importa para este teste
	mockProfile2 := &FrequencyProfile{}

	profileDay := &FrequencyProfileDay{
		freqToilet: mockProfile1,
		freqShower: mockProfile2,
		// Outros são nil por default
	}

	if profileDay.FreqToilet() != mockProfile1 {
		t.Errorf("FreqToilet: esperado %v, obtido %v", mockProfile1, profileDay.FreqToilet())
	}
	if profileDay.FreqShower() != mockProfile2 {
		t.Errorf("FreqShower: esperado %v, obtido %v", mockProfile2, profileDay.FreqShower())
	}
	if profileDay.FreqWashBassin() != nil {
		t.Errorf("FreqWashBassin: esperado nil, obtido %v", profileDay.FreqWashBassin())
	}
	// Verificar os outros getters nil também
	if profileDay.FreqWashMachine() != nil {
		t.Errorf("FreqWashMachine: esperado nil, obtido %v", profileDay.FreqWashMachine())
	}
	if profileDay.FreqDishWasher() != nil {
		t.Errorf("FreqDishWasher: esperado nil, obtido %v", profileDay.FreqDishWasher())
	}
	if profileDay.FreqTanque() != nil {
		t.Errorf("FreqTanque: esperado nil, obtido %v", profileDay.FreqTanque())
	}
}

func TestFrequencyProfileDayGenerateData(t *testing.T) {
	rng := rand.New(rand.NewPCG(98765, 43210)) // Seed fixa para reprodutibilidade

	// Cria mocks de distribuição com valores específicos
	mockDistToilet := &MockDistribution{returnValue: 5.0} 
	mockDistShower := &MockDistribution{returnValue: 2.0}
	mockDistWashBassin := &MockDistribution{returnValue: 0.0}
	mockDistWashMachine := &MockDistribution{returnValue: 1.0}
	mockDistDishWasher := &MockDistribution{returnValue: 3.0}
	mockDistTanque := &MockDistribution{returnValue: 8.0}

	// Cria perfis de frequência usando os mocks.
	// Usamos shift=0 para simplificar a expectativa do valor retornado (será o returnValue da mock, limitado a 255).
	// **Verificamos os erros aqui**
	fpToilet, err := (&FrequencyProfile{}).NewFrequencyProfile(mockDistToilet, 0)
	if err != nil { t.Fatalf("Failed to create fpToilet: %v", err) }
	
	fpShower, err := (&FrequencyProfile{}).NewFrequencyProfile(mockDistShower, 0)
	if err != nil { t.Fatalf("Failed to create fpShower: %v", err) }
	
	fpWashBassin, err := (&FrequencyProfile{}).NewFrequencyProfile(mockDistWashBassin, 0)
	if err != nil { t.Fatalf("Failed to create fpWashBassin: %v", err) }
	
	fpWashMachine, err := (&FrequencyProfile{}).NewFrequencyProfile(mockDistWashMachine, 0)
	if err != nil { t.Fatalf("Failed to create fpWashMachine: %v", err) }
	
	fpDishWasher, err := (&FrequencyProfile{}).NewFrequencyProfile(mockDistDishWasher, 0)
	if err != nil { t.Fatalf("Failed to create fpDishWasher: %v", err) }
	
	fpTanque, err := (&FrequencyProfile{}).NewFrequencyProfile(mockDistTanque, 0)
	if err != nil { t.Fatalf("Failed to create fpTanque: %v", err) }

	// Criação de perfis específicos para os casos de edge (limites, negativos, shift)
	// **Verificamos os erros aqui**
	fpToiletOver255, err := (&FrequencyProfile{}).NewFrequencyProfile(&MockDistribution{returnValue: 300.0}, 0)
	if err != nil { t.Fatalf("Failed to create fpToiletOver255: %v", err) }

	fpToiletNegative, err := (&FrequencyProfile{}).NewFrequencyProfile(&MockDistribution{returnValue: -10.0}, 0)
	if err != nil { t.Fatalf("Failed to create fpToiletNegative: %v", err) }

	fpToiletBelowShift, err := (&FrequencyProfile{}).NewFrequencyProfile(&MockDistribution{returnValue: 5.0}, 10)
	if err != nil { t.Fatalf("Failed to create fpToiletBelowShift: %v", err) }


	testCases := []struct {
		name     string
		profiles map[string]*FrequencyProfile
		expected *behavioral.Frequency // Estrutura Frequency esperada
	}{
		{
			name: "Todos os perfis presentes",
			profiles: map[string]*FrequencyProfile{
				"toilet": fpToilet,
				"shower": fpShower,
				"washBassin": fpWashBassin,
				"washMachine": fpWashMachine,
				"dishWasher": fpDishWasher,
				"tanque": fpTanque,
			},
			expected: behavioral.NewFrequency(5, 2, 0, 1, 3, 8), // Valores esperados baseados nos returnValue das mocks (shift 0)
		},
		{
			name: "Alguns perfis nulos",
			profiles: map[string]*FrequencyProfile{
				"toilet": fpToilet,
				"shower": fpShower,
				// washBassin, washMachine, dishWasher, tanque serão nil
			},
			expected: behavioral.NewFrequency(5, 2, 0, 0, 0, 0), // Nulos devem resultar em 0
		},
		{
			name: "Todos os perfis nulos (mapa vazio)",
			profiles: map[string]*FrequencyProfile{},
			expected: behavioral.NewFrequency(0, 0, 0, 0, 0, 0),
		},
		{
			name: "Todos os perfis nulos (mapa nulo)",
			profiles: nil,
			expected: behavioral.NewFrequency(0, 0, 0, 0, 0, 0),
		},
		{
			name: "Valor de mock acima de 255 (deve ser limitado)",
			profiles: map[string]*FrequencyProfile{
				"toilet": fpToiletOver255, // Usando o perfil criado e verificado
			},
			expected: behavioral.NewFrequency(255, 0, 0, 0, 0, 0), // Esperado 255
		},
		{
			name: "Valor de mock negativo (deve ser convertido para positivo)",
			profiles: map[string]*FrequencyProfile{
				"toilet": fpToiletNegative, // Usando o perfil criado e verificado
			},
			expected: behavioral.NewFrequency(10, 0, 0, 0, 0, 0), // Esperado 10 (abs(-10) = 10, shift 0)
		},
		{
			name: "Valor de mock abaixo do shift (deve usar shift)",
			profiles: map[string]*FrequencyProfile{
				"toilet": fpToiletBelowShift, // Usando o perfil criado e verificado
			},
			expected: behavioral.NewFrequency(10, 0, 0, 0, 0, 0), // Esperado 10 (valor 5, shift 10 -> 10)
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			profileDay := NewFrequencyProfileDay(tc.profiles)
			result := profileDay.GenerateData(rng)

			// Compara o resultado com o esperado
			if *result != *tc.expected {
				t.Errorf("Resultado inesperado para o caso '%s'. Esperado %+v, obtido %+v",
					tc.name, *tc.expected, *result)
			}
		})
	}
}