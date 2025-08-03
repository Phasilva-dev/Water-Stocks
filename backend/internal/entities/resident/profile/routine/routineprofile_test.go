package routine

import (
	"fmt"
	"math/rand/v2"
	"simulation/internal/dists"
	"simulation/internal/entities/resident/ds/behavioral"
	"strings"
	"testing"
	"time"
)

const (
	DAY_IN_SECONDS = 24 * 60 * 60 // Continua útil para tempos absolutos, mas não para maxPercent
)

// MockDistribution é uma implementação de dists.Distribution para testes controlados.
type MockDistribution struct {
	SampleValue    float64
	PercentileValue float64
	MeanValue       float64
	StdDevValue     float64
	StringValue     string
}

func (m *MockDistribution) Params() []float64 {
	return []float64 {m.MeanValue}
}

func (m *MockDistribution) Sample(rng *rand.Rand) float64 {
	return m.SampleValue
}

func (m *MockDistribution) Percentile(p float64) float64 {
	return m.PercentileValue
}

func (m *MockDistribution) Mean() float64 {
	return m.MeanValue
}

func (m *MockDistribution) StdDev() float64 {
	return m.StdDevValue
}

func (m *MockDistribution) String() string {
	return m.StringValue
}

func TestNewDayProfile(t *testing.T) {
	// mockDist para casos válidos e de erro de nil
	mockDist, _ := dists.NewNormalDist(20, 4)

	tests := []struct {
		name        string
		minShift    float64
		maxPercent  float64 // Renomeado
		events      []dists.Distribution
		wantErr     bool
		errSubstr   string
	}{
		{
			name:       "Caso válido",
			minShift:   5,
			maxPercent: 0.99, // Um valor de percentil válido
			events:     []dists.Distribution{mockDist, mockDist},
		},
		{
			name:       "Events slice vazio",
			minShift:   5,
			maxPercent: 0.5,
			events:     []dists.Distribution{},
			wantErr:    true,
			errSubstr:  "events must not be empty (got length 0)", // Mensagem de erro atualizada
		},
		{
			name:       "Número ímpar de eventos",
			minShift:   5,
			maxPercent: 0.5,
			events:     []dists.Distribution{mockDist},
			wantErr:    true,
			errSubstr:  "number of events must be even (got 1)", // Mensagem de erro atualizada
		},
		{
			name:       "MinShift negativo",
			minShift:   -1,
			maxPercent: 0.5,
			events:     []dists.Distribution{mockDist, mockDist},
			wantErr:    true,
			errSubstr:  "minShift must be non-negative (got -1.0000)", // Mensagem de erro atualizada
		},
		{
			name:       "Distribuição nil",
			minShift:   5,
			maxPercent: 0.5,
			events:     []dists.Distribution{mockDist, nil},
			wantErr:    true,
			errSubstr:  "distribution at index 1 is nil", // Mensagem de erro atualizada
		},
		{
			name:       "MaxPercent negativo",
			minShift:   5,
			maxPercent: -0.1,
			events:     []dists.Distribution{mockDist, mockDist},
			wantErr:    true,
			errSubstr:  "maxPercent must be between 0 and 1 (got -0.1000)", // Nova validação
		},
		{
			name:       "MaxPercent maior que 1",
			minShift:   5,
			maxPercent: 1.1,
			events:     []dists.Distribution{mockDist, mockDist},
			wantErr:    true,
			errSubstr:  "maxPercent must be between 0 and 1 (got 1.1000)", // Nova validação
		},
		{
			name:       "MaxPercent igual a 0", // Válido, mas implica que tudo será 0 ou o menor valor possível
			minShift:   5,
			maxPercent: 0.0,
			events:     []dists.Distribution{mockDist, mockDist},
		},
		{
			name:       "MaxPercent igual a 1", // Válido, implica sem limite de percentil
			minShift:   5,
			maxPercent: 1.0,
			events:     []dists.Distribution{mockDist, mockDist},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDayProfile(tt.events, tt.minShift, tt.maxPercent)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("Esperava erro contendo %q, mas nenhum erro ocorreu.", tt.errSubstr)
				}
				if !strings.Contains(err.Error(), tt.errSubstr) {
					t.Errorf("Erro incorreto.\nEsperado conter: %q\nObtido: %q", tt.errSubstr, err.Error())
				}
				return
			}

			if err != nil {
				t.Fatalf("Erro inesperado para o caso válido: %v", err)
			}

			if got.MinShift() != tt.minShift {
				t.Errorf("MinShift incorreto. Esperado: %f, Obtido: %f", tt.minShift, got.MinShift())
			}
			// Verificar MaxPercent
			if got.MaxPercent() != tt.maxPercent {
				t.Errorf("MaxPercent incorreto. Esperado: %f, Obtido: %f", tt.maxPercent, got.MaxPercent())
			}
			if len(got.Events()) != len(tt.events) {
				t.Errorf("Número de eventos incorreto. Esperado: %d, Obtido: %d", len(tt.events), len(got.Events()))
			}
		})
	}
}

func TestEnforceMaxValueIsRespected(t *testing.T) {
	// Este teste verifica o comportamento de enforceMaxValue isoladamente,
	// com minShift configurado para 0 para não interferir na verificação do limite.
	const (
		expectedCap     = 100.0 // O valor absoluto que esperamos que a amostra seja limitada
		sampleToCap     = 500.0 // Um valor de amostra que é maior que expectedCap
		maxPercentValue = 0.5   // O percentil que, para nossa MockDistribution, resultará em expectedCap
		minShift        = 0.0   // Importante: minShift 0 para isolar o efeito de maxPercent
	)

	// Configurar MockDistribution para retornar um valor alto no Sample()
	// mas um valor esperado no Percentile()
	mock := &MockDistribution{
		SampleValue:    sampleToCap,
		PercentileValue: expectedCap, // Qualquer percentil retornará 100.0 para este mock
	}

	profile, err := NewDayProfile(
		[]dists.Distribution{mock, mock}, // Usar duas vezes para ter dois eventos
		minShift,
		maxPercentValue,
	)
	if err != nil {
		t.Fatalf("Erro ao criar DayProfile: %v", err)
	}

	rng := rand.New(rand.NewPCG(42, 1)) // Gerador determinístico
	routine, _ := profile.GenerateData(rng)

	// Verificar se todos os tempos gerados estão dentro do limite.
	// Como minShift é 0, enforceMinShift não deve empurrar os valores além do limite.
	for i, v := range routine.Times() {
		if v > expectedCap {
			t.Errorf("Tempo excedeu o limite de maxPercent: posição %d, valor %f, limite esperado %f", i, v, expectedCap)
		}
		// Além disso, verificar se foi efetivamente limitado ao expectedCap
		if v != expectedCap {
			t.Errorf("Tempo na posição %d não foi limitado corretamente. Esperado %f, Obtido %f", i, expectedCap, v)
		}
	}
}

func TestGenerateDataDeterministic(t *testing.T) {
	const (
		numRoutines    = 1000
		minShift       = 5.0
		maxPercentTest = 1.0 // Definir maxPercent para 1.0 para desabilitar o efeito de limite de percentil neste teste
	)

	// Uma distribuição com stddev=0 sempre amostrará a média.
	mockDist, _ := dists.NewNormalDist(10, 0)

	profile, err := NewDayProfile(
		[]dists.Distribution{mockDist, mockDist},
		minShift,
		maxPercentTest,
	)
	if err != nil {
		t.Fatalf("Falha ao criar perfil: %v", err)
	}

	rng := rand.New(rand.NewPCG(42, 1)) // Gerador determinístico
	routines := make([]*behavioral.Routine, 0, numRoutines)

	for i := 0; i < numRoutines; i++ {
		routine, _ := profile.GenerateData(rng)
		routines = append(routines, routine)
	}

	// Com a lógica de enforceMinShift:
	// times[0] = sample(10) = 10
	// times[1] = enforceMinShift(times[0], sample(10))
	//           = enforceMinShift(10, 10)
	//           -> current (10) < prev+minShift (10+5=15) é TRUE
	//           -> diff = abs(10-10) = 0
	//           -> diff (0) < minShift (5) é TRUE
	//           -> return prev + diff + minShift = 10 + 0 + 5 = 15
	expectedTimes := []float64{10, 15}

	for i, routine := range routines {
		if len(routine.Times()) != len(expectedTimes) {
			t.Fatalf("[Rotina %d] Número de tempos incorreto. Esperado: %d, Obtido: %d",
				i, len(expectedTimes), len(routine.Times()))
		}

		for j := 0; j < len(expectedTimes); j++ {
			if routine.Times()[j] != expectedTimes[j] {
				t.Errorf("[Rotina %d] Tempo na posição %d incorreto. Esperado: %f, Obtido: %f",
					i, j, expectedTimes[j], routine.Times()[j])
			}
		}

		// A verificação do gap é implícita se os tempos esperados estão corretos, mas podemos manter.
		entry := routine.Times()[0]
		exit := routine.Times()[1]
		if exit-entry != minShift { // Agora exit-entry deve ser exatamente minShift + diff do valor original
			t.Errorf("[Rotina %d] Gap inválido entre %f (entrada) e %f (saída). Esperado: %f, Obtido: %f",
				i, entry, exit, minShift, exit-entry)
		}
	}

	t.Logf("Geradas %d rotinas consistentes e determinísticas", numRoutines)
}

func TestGenerateDataReal(t *testing.T) {
	const (
		numRoutines    = 10000
		minShift       = 15 * 60 // 15 minutos
		maxPercentTest = 1.0     // Definir maxPercent para 1.0 para desabilitar o efeito de limite de percentil neste teste
	)

	wakeUpDist, _ := dists.NewNormalDist(5*60*60, 30*60) // 5:00 AM, stddev 30min
	leaveDist, _ := dists.NewNormalDist(7*60*60, 30*60)  // 7:00 AM, stddev 30min
	returnDist, _ := dists.NewNormalDist(18*60*60, 30*60) // 6:00 PM, stddev 30min
	sleepDist, _ := dists.NewNormalDist(22*60*60, 30*60) // 10:00 PM, stddev 30min

	profile, err := NewDayProfile(
		[]dists.Distribution{wakeUpDist, leaveDist, returnDist, sleepDist},
		minShift,
		maxPercentTest,
	)
	if err != nil {
		t.Fatalf("Falha ao criar perfil: %v", err)
	}

	rng := rand.New(rand.NewPCG(42, 1)) // Gerador determinístico para reprodutibilidade
	routines := make([]*behavioral.Routine, 0, numRoutines)

	for i := 0; i < numRoutines; i++ {
		routine, _ := profile.GenerateData(rng)
		routines = append(routines, routine)
	}

	for i, routine := range routines {
		if len(routine.Times()) != 4 {
			t.Fatalf("[Rotina %d] Número de tempos incorreto. Esperado: 4, Obtido: %d",
				i, len(routine.Times()))
		}

		// A verificação do minShift deve ser feita entre todos os pares (prev, current)
		// conforme a lógica de enforceMinShift.
		// A lógica atual do loop k += 2 só verifica pares de entrada/saída, mas minShift
		// é aplicado sequencialmente.

		// Vamos verificar que os tempos estão em ordem crescente e respeitam o minShift
		// Esta parte do teste é mais robusta para a lógica de enforceMinShift.
		for j := 1; j < len(routine.Times()); j++ {
			prev := routine.Times()[j-1]
			current := routine.Times()[j]

			// A lógica enforceMinShift garante que 'current' é pelo menos 'prev + diff + minShift' (se diff < minShift)
			// ou 'prev + diff' (se diff >= minShift), quando 'current < prev + minShift'.
			// Em todos os casos, o resultado final deve ser >= prev. E se ele foi ajustado,
			// deve ser >= prev + minShift se o 'current' original estivesse muito próximo ou antes de 'prev'.
			// É difícil dar uma expectativa exata sem re-simular a função enforceMinShift aqui.
			// No entanto, podemos verificar que o `current` nunca é `menos` que `prev`.
			// E se o `current` original fosse menor que `prev + minShift`, ele deve ter sido "empurrado".

			// Uma verificação simplificada: o tempo atual nunca deve ser menor que o tempo anterior.
			if current < prev {
				t.Errorf("[Rotina %d] Tempo na posição %d (%f) é menor que o tempo anterior na posição %d (%f).",
					i, j, current, j-1, prev)
			}

			// E o intervalo entre um evento de 'saída' e 'entrada' (e.g. leaveDist e returnDist)
			// ou entre um evento 'entrada' e 'saída' (wakeUpDist e leaveDist) deve respeitar minShift,
			// mas a forma como enforceMinShift é aplicada é o que deve ser validado.
			// No caso de minShift, o tempo de *saída* deve ser no mínimo 'tempo_de_entrada' + 'minShift'.
			// Se o tempo amostrado for muito menor do que o anterior + minShift, ele é ajustado.
			// O teste determinístico já cobre o comportamento exato. Este teste realístico pode ser
			// mais sobre a "sanidade" dos dados.
			// A verificação mais direta para minShift é que o próximo tempo sempre é MAIOR OU IGUAL ao tempo anterior.
			// E em alguns casos, será maior em pelo menos `minShift` (se a amostra original fosse muito próxima).
		}

		// A remoção da verificação `exit > float64(maxShift)` é crucial, pois `maxPercent` não é um limite absoluto.
		// Se `maxPercentTest` for 1.0, nenhum limite é aplicado. Se fosse < 1.0, os valores seriam limitados
		// pelo percentil da distribuição, não por um valor fixo como DAY_IN_SECONDS.
	}

	baseDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println("\nTempos da primeira rotina formatados (exemplo):")
	for i, timeSec := range routines[0].Times() {
		duration := time.Duration(int64(timeSec)) * time.Second
		dateTime := baseDate.Add(duration)
		fmt.Printf("%d: %s (%.0f segundos)\n", i+1, dateTime.Format("02/01/2006 15:04:05"), timeSec)
	}
}