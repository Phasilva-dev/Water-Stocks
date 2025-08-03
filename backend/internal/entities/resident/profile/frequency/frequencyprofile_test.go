package frequency

import (
	"math"
	"math/rand/v2"
	"simulation/internal/dists"
	//"simulation/internal/entities/resident/ds/behavioral"
	"strings"
	"testing"
)

// mockDistribution é uma implementação da interface dists.Distribution para testes.
// Sendo simples, nos permite controlar a saída do método Sample.
type mockDistribution struct {
	SampleValue    float64
	PercentileValue float64
	MeanValue       float64
	StdDevValue     float64
	StringValue     string
}

func (m *mockDistribution) Params() []float64 {
	return []float64 {m.MeanValue}
}

func (m *mockDistribution) Sample(rng *rand.Rand) float64 {
	return m.SampleValue
}

func (m *mockDistribution) Percentile(p float64) float64 {
	return m.PercentileValue
}

func (m *mockDistribution) Mean() float64 {
	return m.MeanValue
}

func (m *mockDistribution) StdDev() float64 {
	return m.StdDevValue
}

func (m *mockDistribution) String() string {
	return m.StringValue
}

// --- Testes para a Unidade (unit.go) ---

// TestNewDeviceProfile consolida os testes de criação do DeviceProfile,
// incluindo a validação dos getters Shift() e StatDist().
func TestNewDeviceProfile(t *testing.T) {
	dist := &mockDistribution{SampleValue: 50.0}

	testCases := []struct {
		name        string
		dist        dists.Distribution
		shift       uint8
		expectError bool
	}{
		{
			name:        "Success: valid distribution and shift",
			dist:        dist,
			shift:       10,
			expectError: false,
		},
		{
			name:        "Failure: nil distribution",
			dist:        nil,
			shift:       20,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			profile, err := NewDeviceProfile(tc.dist, tc.shift)

			if tc.expectError {
				if err == nil {
					t.Fatal("Expected an error for nil distribution, but got nil")
				}
				if profile != nil {
					t.Errorf("Profile should be nil on error, but got a value: %+v", profile)
				}
				// Verificação opcional da mensagem de erro
				if !strings.Contains(err.Error(), "distribution cannot be nil") {
					t.Errorf("Error message does not contain expected text. Got: %s", err.Error())
				}
			} else {
				if err != nil {
					t.Fatalf("Did not expect an error, but got: %v", err)
				}
				if profile == nil {
					t.Fatal("Profile should not be nil on success")
				}
				if gotShift := profile.Shift(); gotShift != tc.shift {
					t.Errorf("Shift() got = %d, want = %d", gotShift, tc.shift)
				}
				if gotDist := profile.StatDist(); gotDist != tc.dist {
					t.Errorf("StatDist() did not return the correct distribution")
				}
			}
		})
	}
}

// TestGenerateFrequency testa a função auxiliar não exportada 'generateFrequency'.
// Isso é possível porque o teste está no mesmo pacote 'frequency'.
func TestGenerateFrequency(t *testing.T) {
	rng := rand.New(rand.NewPCG(1, 0)) // Seed fixa para reprodutibilidade

	testCases := []struct {
		name         string
		sampleValue  float64
		shift        uint8
		expectedFreq uint8
	}{
		{"Normal value", 100.0, 10, 100},
		{"Value below shift", 5.0, 10, 10},
		{"Negative value converted", -50.0, 10, 50},
		{"Negative value converted and shifted", -5.0, 10, 10},
		{"Value above uint8 max", 300.0, 10, 255},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockDist := &mockDistribution{SampleValue: tc.sampleValue}
			got := generateFrequency(rng, tc.shift, mockDist)

			if got != tc.expectedFreq {
				t.Errorf("generateFrequency() with sample %.1f and shift %d: got = %d, want = %d", tc.sampleValue, tc.shift, got, tc.expectedFreq)
			}
		})
	}
}

// --- Testes para Agregação Diária (day.go) ---

// TestNewResidentDeviceProfiles testa o construtor do agregador de perfis.
func TestNewResidentDeviceProfiles(t *testing.T) {
	// Um perfil válido para ser reutilizado nos casos de teste
	validProfile, _ := NewDeviceProfile(&mockDistribution{}, 0)

	testCases := []struct {
		name        string
		inputMap    map[string]*DeviceProfile
		expectError bool
	}{
		{
			name: "Success: All profiles provided",
			inputMap: map[string]*DeviceProfile{
				"toilet": validProfile,
				"shower": validProfile,
			},
			expectError: false,
		},
		{
			name: "Failure: A profile is nil",
			inputMap: map[string]*DeviceProfile{
				"toilet": validProfile,
				"shower": nil, // O perfil inválido
			},
			expectError: true,
		},
		{"Success: Empty map", map[string]*DeviceProfile{}, false},
		{"Success: Nil map", nil, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			profiles, err := NewResidentDeviceProfiles(tc.inputMap)

			if tc.expectError {
				if err == nil {
					t.Fatal("Expected an error but got none")
				}
				if !strings.Contains(err.Error(), "missing DeviceProfile") {
					t.Errorf("Error message mismatch, got: %v", err)
				}
				if profiles != nil {
					t.Error("Profiles object should be nil on error")
				}
			} else {
				if err != nil {
					t.Fatalf("Did not expect an error, but got: %v", err)
				}
				if profiles == nil {
					t.Fatal("Profiles object should not be nil on success")
				}
				// Como estamos no mesmo pacote, podemos verificar o campo interno
				if len(profiles.freqDevices) != len(tc.inputMap) {
					t.Errorf("Internal map length mismatch. got = %d, want = %d", len(profiles.freqDevices), len(tc.inputMap))
				}
			}
		})
	}
}

// TestResidentDeviceProfiles_GenerateData testa a geração de dados agregados.
func TestResidentDeviceProfiles_GenerateData(t *testing.T) {
	rng := rand.New(rand.NewPCG(1, 0))

	// Arrange: Cria perfis com mocks previsíveis
	profileToilet, _ := NewDeviceProfile(&mockDistribution{SampleValue: 5.0}, 0)
	profileShower, _ := NewDeviceProfile(&mockDistribution{SampleValue: 2.0}, 0)
	profileFaucetOver, _ := NewDeviceProfile(&mockDistribution{SampleValue: 300.0}, 0)
	profileTanqueShift, _ := NewDeviceProfile(&mockDistribution{SampleValue: 5.0}, 10)

	testCases := []struct {
		name         string
		profiles     map[string]*DeviceProfile
		expectedData map[string]uint8
	}{
		{
			name: "All profiles present",
			profiles: map[string]*DeviceProfile{
				"toilet": profileToilet,
				"shower": profileShower,
			},
			expectedData: map[string]uint8{"toilet": 5, "shower": 2},
		},
		{
			name:         "Empty map of profiles",
			profiles:     map[string]*DeviceProfile{},
			expectedData: map[string]uint8{},
		},
		{
			name: "Edge cases (clamping and shift)",
			profiles: map[string]*DeviceProfile{
				"faucet": profileFaucetOver,
				"tanque": profileTanqueShift,
			},
			expectedData: map[string]uint8{"faucet": 255, "tanque": 10},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			residentProfiles, err := NewResidentDeviceProfiles(tc.profiles)
			if err != nil {
				t.Fatalf("Test setup failed, could not create ResidentDeviceProfiles: %v", err)
			}

			// Act
			resultFreq, err := residentProfiles.GenerateData(rng)
			if err != nil {
				t.Fatalf("GenerateData returned an unexpected error: %v", err)
			}
			if resultFreq == nil {
				t.Fatal("GenerateData returned a nil result")
			}

			// Assert
			// 1. Verificar se o número de chaves é o mesmo
			// Acesso ao campo não exportado é possível por estar no mesmo pacote.
			if len(resultFreq.DevicesFrequency()) != len(tc.expectedData) {
				t.Fatalf("Result has wrong number of devices. got=%d, want=%d", len(resultFreq.DevicesFrequency()), len(tc.expectedData))
			}

			// 2. Verificar o valor de cada chave esperada
			for device, expectedVal := range tc.expectedData {
				if gotVal := resultFreq.DeviceFrequency(device); gotVal != expectedVal {
					t.Errorf("Frequency for device %q is wrong. got=%d, want=%d", device, gotVal, expectedVal)
				}
			}
		})
	}
}

// --- Teste de Integração/Estatístico ---

// TestDeviceProfile_GenerateData_Statistical é o seu teste estatístico original.
// É uma boa prática mantê-lo separado por uma tag de build para não rodar com os testes de unidade.
// Para executá-lo: go test -v -tags=integration
func TestDeviceProfile_GenerateData_Statistical(t *testing.T) {
	// Arrange
	poissonDist, err := dists.CreateDistribution("poisson", 1.38)
	if err != nil {
		t.Fatalf("Failed to create Poisson distribution: %v", err)
	}

	profile, err := NewDeviceProfile(poissonDist, 0)
	if err != nil {
		t.Fatalf("Failed to create DeviceProfile: %v", err)
	}

	rng := rand.New(rand.NewPCG(1, 0))
	numSimulations := 1_000_000
	frequencies := make(map[uint8]int)

	// Act
	for i := 0; i < numSimulations; i++ {
		value := profile.GenerateData(rng)
		frequencies[value]++
	}

	// Assert
	lambda := 1.38
	tolerance := 0.01

	// Função auxiliar para calcular a probabilidade de Poisson (usando Gamma para fatorial)
	poissonProb := func(k int) float64 {
		return math.Pow(lambda, float64(k)) * math.Exp(-lambda) / math.Gamma(float64(k+1))
	}

	for k := 0; k <= 10; k++ { // Verificando os valores mais prováveis
		theoreticalProb := poissonProb(k)
		observedProb := float64(frequencies[uint8(k)]) / float64(numSimulations)

		if diff := math.Abs(theoreticalProb - observedProb); diff > tolerance {
			t.Errorf("Probability for k=%d is out of tolerance. got=%.4f, want=%.4f (diff=%.4f)", k, observedProb, theoreticalProb, diff)
		}
	}

	t.Log("Statistical test passed: observed frequencies are within tolerance.")
}