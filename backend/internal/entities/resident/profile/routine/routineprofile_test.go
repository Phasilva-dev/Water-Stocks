package routine

import (
	"fmt"
	"math/rand/v2"
	"simulation/internal/dists"
	"simulation/internal/entities/resident/ds/behavioral"
	"testing"
	"time"
)

const (
	DAY_IN_SECONDS = 24 * 60 * 60
)

func TestNewRoutineProfile(t *testing.T) {
	// Cria uma distribuição válida para usar nos testes
	mockDist, _ := dists.NewNormalDist(20, 4)

	// Tabela de testes (table-driven tests)
	tests := []struct {
		name     string                // Nome descritivo do teste
		shift    float64               // Input: shift
		events   []dists.Distribution  // Input: distribuições
		wantErr  bool                  // Se esperamos erro
		errMsg   string                // Mensagem de erro esperada
	}{
		{
			name:     "Caso válido",
			shift:    5,
			events:   []dists.Distribution{mockDist, mockDist},
			wantErr:  false,
		},
		{
			name:     "Número ímpar de eventos",
			shift:    5,
			events:   []dists.Distribution{mockDist},
			wantErr:  true,
			errMsg:   "number of elements in events must be even",
		},
		{
			name:     "shift negativa",
			shift:    -1,
			events:   []dists.Distribution{mockDist, mockDist},
			wantErr:  true,
			errMsg:   "shift must be positive",
		},
		{
			name:     "Distribuição nil",
			shift:    5,
			events:   []dists.Distribution{mockDist, nil},
			wantErr:  true,
			errMsg:   "no distribution can be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Executa a função
			got, err := NewRoutineProfile(tt.events, tt.shift)

			// Verifica se o erro é o esperado
			if tt.wantErr {
				if err == nil {
					t.Fatal("Esperava erro, mas não ocorreu")
				}
				if err.Error() != tt.errMsg {
					t.Errorf("Mensagem de erro incorreta\nEsperada: %s\nObtida: %s", tt.errMsg, err.Error())
				}
				return
			}

			// Se não espera erro
			if err != nil {
				t.Fatalf("Erro inesperado: %v", err)
			}

			if got.shift != tt.shift {
				t.Errorf("shift incorreta\nEsperada: %f\nObtida: %f", tt.shift, got.shift)
			}

			if len(got.Events()) != len(tt.events) {
				t.Errorf("Número de eventos incorreto\nEsperado: %d\nObtido: %d", len(tt.events), len(got.Events()))
			}
		})
	}
}

func TestGenerateData(t *testing.T) {
	const (
		numRoutines = 1000   // Quantidade de rotinas a gerar
		shift       = 60 * 15 // Gap mínimo entre entrada/saída
	)

	// 1. Cria uma distribuição determinística
	wakeUpDist, _ := dists.NewNormalDist(5*60*60, 0)
	leaveDist, _ := dists.NewNormalDist(7*60*60, 0)
	returnDist, _ := dists.NewNormalDist(18*60*60, 0)
	sleepDist, _ := dists.NewNormalDist(23*60*60, 0)

	// 2. Cria o perfil
	profile, err := NewRoutineProfile(
		[]dists.Distribution{wakeUpDist, leaveDist, returnDist, sleepDist},
		shift,
	)
	if err != nil {
		t.Fatalf("Falha ao criar perfil: %v", err)
	}

	// 3. Gera as rotinas
	src := rand.NewPCG(42, 1)
	rng := rand.New(src) // Seed fixa para reprodutibilidade
	routines := make([]*behavioral.Routine, 0, numRoutines)

	for i := 0; i < numRoutines; i++ {
		routine := profile.GenerateData(rng)
		routines = append(routines, routine)
	}

	// 4. Verificações
	expectedTimes := []float64{5 * 60 * 60, 7 * 60 * 60, 18 * 60 * 60, 23 * 60 * 60}

	for i, routine := range routines {
		// Verifica quantidade de tempos
		if len(routine.Times()) != len(expectedTimes) {
			t.Fatalf("[Rotina %d] Número de tempos incorreto. Esperado: %d, Obtido: %d",
				i, len(expectedTimes), len(routine.Times()))
		}

		// Verifica valores e gap mínimo
		for j := 0; j < len(expectedTimes); j++ {
			if routine.Times()[j] != expectedTimes[j] {
				t.Errorf("[Rotina %d] Tempo na posição %d incorreto. Esperado: %f, Obtido: %f",
					i, j, expectedTimes[j], routine.Times()[j])
			}
		}

		// Verifica pares entrada/saída
		for k := 0; k < len(routine.Times()); k += 2 {
			// Verifica se há um par completo (entrada e saída)
			if k+1 >= len(routine.Times()) {
				t.Errorf("[Rotina %d] Número ímpar de tempos: %d", i, len(routine.Times()))
				break
			}

			entry := routine.Times()[k]    // Tempo de entrada
			exit := routine.Times()[k+1]  // Tempo de saída

			// Verifica se o gap mínimo foi respeitado
			if exit-entry < float64(shift) {
				t.Errorf("[Rotina %d] Gap inválido entre %f (entrada) e %f (saída). Esperado: mínimo %f",
					i, entry, exit, float64(shift))
			}
		}
	}

	t.Logf("Geradas %d rotinas consistentes", numRoutines)
}

func TestGenerateDataBatch(t *testing.T) {
	// Configuração determinística
	const (
		numRoutines = 1000   // Quantidade de rotinas a gerar
		shift       = 5.0    // Gap mínimo entre entrada/saída
	)

	// 1. Cria uma distribuição determinística (sempre retorna 10)
	mockDist, _ := dists.NewNormalDist(10, 0) // Média 10, desvio 0 = sempre 10

	// 2. Cria o perfil com 2 eventos (entrada e saída)
	profile, err := NewRoutineProfile(
		[]dists.Distribution{mockDist, mockDist},
		shift,
	)
	if err != nil {
		t.Fatalf("Falha ao criar perfil: %v", err)
	}

	// 3. Gera as rotinas
	src := rand.NewPCG(42, 1)
	rng := rand.New(src) // Seed fixa para reprodutibilidade
	routines := make([]*behavioral.Routine, 0, numRoutines)

	for i := 0; i < numRoutines; i++ {
		routine := profile.GenerateData(rng)
		routines = append(routines, routine)
	}

	// 4. Verificações
	expectedTimes := []float64{10, 15} // Entrada: 10, Saída: 10 + shift = 15

	for i, routine := range routines {
		// Verifica quantidade de tempos
		if len(routine.Times()) != len(expectedTimes) {
			t.Fatalf("[Rotina %d] Número de tempos incorreto. Esperado: %d, Obtido: %d",
				i, len(expectedTimes), len(routine.Times()))
		}

		// Verifica valores e gap mínimo
		for j := 0; j < len(expectedTimes); j++ {
			if routine.Times()[j] != expectedTimes[j] {
				t.Errorf("[Rotina %d] Tempo na posição %d incorreto. Esperado: %f, Obtido: %f",
					i, j, expectedTimes[j], routine.Times()[j])
			}
		}

		// Verifica pares entrada/saída
		for k := 0; k < len(routine.Times()); k += 2 {
			if k+1 >= len(routine.Times()) {
				break
			}

			entry := routine.Times()[k]
			exit := routine.Times()[k+1]

			// O gap deve ser exatamente igual à shift
			if exit-entry != shift {
				t.Errorf("[Rotina %d] Gap inválido entre %f (entrada) e %f (saída). Esperado: %f",
					i, entry, exit, shift)
			}
		}
	}

	t.Logf("Geradas %d rotinas consistentes", numRoutines)
}

func TestGenerateDataReal(t *testing.T) {
	const (
		numRoutines = 1_000_000   // Quantidade de rotinas a gerar
		shift       = 60 * 15  // 15 minutos
	)

	// 1. Cria uma distribuição determinística
	wakeUpDist, _ := dists.NewNormalDist(5*60*60, 30*60)
	leaveDist, _ := dists.NewNormalDist(7*60*60, 30*60)
	returnDist, _ := dists.NewNormalDist(18*60*60, 30*60)
	sleepDist, _ := dists.NewNormalDist(22*60*60, 30*60)

	// 2. Cria o perfil
	profile, err := NewRoutineProfile(
		[]dists.Distribution{wakeUpDist, leaveDist, returnDist, sleepDist},
		shift,
	)
	if err != nil {
		t.Fatalf("Falha ao criar perfil: %v", err)
	}

	// 3. Gera as rotinas
	src := rand.NewPCG(42, 1)
	rng := rand.New(src) // Seed fixa para reprodutibilidade
	routines := make([]*behavioral.Routine, 0, numRoutines)

	for i := 0; i < numRoutines; i++ {
		routine := profile.GenerateData(rng)
		routines = append(routines, routine)
	}

	// 4. Verificações
	var length int = 4

	for i, routine := range routines {
		// Verifica quantidade de tempos
		if len(routine.Times()) != length {
			t.Fatalf("[Rotina %d] Número de tempos incorreto. Esperado: %d, Obtido: %d",
				i, length, len(routine.Times()))
		}

		// Verifica pares entrada/saída
		for k := 0; k < len(routine.Times()); k += 2 {
			// Verifica se há um par completo (entrada e saída)
			if k+1 >= len(routine.Times()) {
				t.Errorf("[Rotina %d] Número ímpar de tempos: %d", i, len(routine.Times()))
				break
			}

			entry := routine.Times()[k]    // Tempo de entrada
			exit := routine.Times()[k+1]  // Tempo de saída

			// Verifica se o gap mínimo foi respeitado
			if exit-entry < float64(shift) {
				t.Errorf("[Rotina %d] Gap inválido entre %f (entrada) e %f (saída). Esperado: mínimo %f",
					i, entry, exit, float64(shift))
			}
		}
	}

	// Formatação dos tempos como datetime
	baseDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println("\nTempos da primeira rotina formatados:")
	for i, timeSec := range routines[0].Times() {
		// Converte segundos para duração e adiciona à data base
		duration := time.Duration(int64(timeSec)) * time.Second
		dateTime := baseDate.Add(duration)
		fmt.Printf("%d: %s (%.0f segundos)\n", i+1, dateTime.Format("02/01/2006 15:04:05"), timeSec)
	}
}