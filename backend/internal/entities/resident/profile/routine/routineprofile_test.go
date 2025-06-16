package routine

import (
	"fmt"
	"math/rand/v2"
	"simulation/internal/dists"
	"simulation/internal/entities/resident/ds/behavioral"
	"testing"
	"time"
	"strings"
)

const (
	DAY_IN_SECONDS = 24 * 60 * 60
)

func TestNewRoutineProfile(t *testing.T) {
	mockDist, _ := dists.NewNormalDist(20, 4)

	tests := []struct {
		name      string
		minShift  float64
		maxShift  float64
		events    []dists.Distribution
		wantErr   bool
		errSubstr string
	}{
		{
			name:     "Caso válido",
			minShift: 5,
			maxShift: DAY_IN_SECONDS,
			events:   []dists.Distribution{mockDist, mockDist},
		},
		{
			name:      "Número ímpar de eventos",
			minShift:  5,
			maxShift:  DAY_IN_SECONDS,
			events:    []dists.Distribution{mockDist},
			wantErr:   true,
			errSubstr: "number of elements in events must be even",
		},
		{
			name:      "MinShift negativo",
			minShift:  -1,
			maxShift:  DAY_IN_SECONDS,
			events:    []dists.Distribution{mockDist, mockDist},
			wantErr:   true,
			errSubstr: "shift must be positive",
		},
		{
			name:      "Distribuição nil",
			minShift:  5,
			maxShift:  DAY_IN_SECONDS,
			events:    []dists.Distribution{mockDist, nil},
			wantErr:   true,
			errSubstr: "no distribution can be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRoutineProfile(tt.events, tt.minShift, tt.maxShift)

			if tt.wantErr {
				if err == nil {
					t.Fatal("Esperava erro, mas não ocorreu")
				}
				if !contains(err.Error(), tt.errSubstr) {
					t.Errorf("Erro incorreto.\nEsperado conter: %q\nObtido: %q", tt.errSubstr, err.Error())
				}
				return
			}

			if err != nil {
				t.Fatalf("Erro inesperado: %v", err)
			}

			if got.MinShift() != tt.minShift {
				t.Errorf("MinShift incorreto. Esperado: %f, Obtido: %f", tt.minShift, got.MinShift())
			}
			if got.MaxPercent() != tt.maxShift {
				t.Errorf("MaxPercent incorreto. Esperado: %f, Obtido: %f", tt.maxShift, got.MaxPercent())
			}
			if len(got.Events()) != len(tt.events) {
				t.Errorf("Número de eventos incorreto. Esperado: %d, Obtido: %d", len(tt.events), len(got.Events()))
			}
		})
	}
}

func TestMaxShiftIsRespected(t *testing.T) {
	const (
		minShift = 5.0
		maxShift = 100.0 // Limite baixo para forçar o teste
	)

	highDist, _ := dists.NewNormalDist(5000, 0)

	profile, err := NewRoutineProfile(
		[]dists.Distribution{highDist, highDist},
		minShift,
		maxShift,
	)
	if err != nil {
		t.Fatalf("Erro ao criar RoutineProfile: %v", err)
	}

	rng := rand.New(rand.NewPCG(42, 1))
	routine := profile.GenerateData(rng)

	for i, v := range routine.Times() {
		if v > maxShift {
			t.Errorf("Tempo excedeu maxShift: posição %d, valor %f, limite %f", i, v, maxShift)
		}
	}
}

func TestGenerateDataDeterministic(t *testing.T) {
	const (
		numRoutines = 1000
		minShift    = 5.0
		maxShift    = DAY_IN_SECONDS
	)

	mockDist, _ := dists.NewNormalDist(10, 0)

	profile, err := NewRoutineProfile(
		[]dists.Distribution{mockDist, mockDist},
		minShift,
		maxShift,
	)
	if err != nil {
		t.Fatalf("Falha ao criar perfil: %v", err)
	}

	rng := rand.New(rand.NewPCG(42, 1))
	routines := make([]*behavioral.Routine, 0, numRoutines)

	for i := 0; i < numRoutines; i++ {
		routine := profile.GenerateData(rng)
		routines = append(routines, routine)
	}

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

		entry := routine.Times()[0]
		exit := routine.Times()[1]
		if exit-entry != minShift {
			t.Errorf("[Rotina %d] Gap inválido entre %f (entrada) e %f (saída). Esperado: %f",
				i, entry, exit, minShift)
		}
	}

	t.Logf("Geradas %d rotinas consistentes", numRoutines)
}

func TestGenerateDataReal(t *testing.T) {
	const (
		numRoutines = 10000
		minShift    = 15 * 60
		maxShift    = DAY_IN_SECONDS
	)

	wakeUpDist, _ := dists.NewNormalDist(5*60*60, 30*60)
	leaveDist, _ := dists.NewNormalDist(7*60*60, 30*60)
	returnDist, _ := dists.NewNormalDist(18*60*60, 30*60)
	sleepDist, _ := dists.NewNormalDist(22*60*60, 30*60)

	profile, err := NewRoutineProfile(
		[]dists.Distribution{wakeUpDist, leaveDist, returnDist, sleepDist},
		minShift,
		maxShift,
	)
	if err != nil {
		t.Fatalf("Falha ao criar perfil: %v", err)
	}

	rng := rand.New(rand.NewPCG(42, 1))
	routines := make([]*behavioral.Routine, 0, numRoutines)

	for i := 0; i < numRoutines; i++ {
		routine := profile.GenerateData(rng)
		routines = append(routines, routine)
	}

	for i, routine := range routines {
		if len(routine.Times()) != 4 {
			t.Fatalf("[Rotina %d] Número de tempos incorreto. Esperado: 4, Obtido: %d",
				i, len(routine.Times()))
		}

		for k := 0; k < len(routine.Times()); k += 2 {
			if k+1 >= len(routine.Times()) {
				t.Errorf("[Rotina %d] Número ímpar de tempos: %d", i, len(routine.Times()))
				break
			}

			entry := routine.Times()[k]
			exit := routine.Times()[k+1]

			if exit-entry < float64(minShift) {
				t.Errorf("[Rotina %d] Gap inválido entre %f (entrada) e %f (saída). Esperado mínimo: %f",
					i, entry, exit, float64(minShift))
			}
			if exit > float64(maxShift) {
				t.Errorf("[Rotina %d] Tempo de saída excede maxShift: %f > %f", i, exit, float64(maxShift))
			}
		}
	}

	baseDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println("\nTempos da primeira rotina formatados:")
	for i, timeSec := range routines[0].Times() {
		duration := time.Duration(int64(timeSec)) * time.Second
		dateTime := baseDate.Add(duration)
		fmt.Printf("%d: %s (%.0f segundos)\n", i+1, dateTime.Format("02/01/2006 15:04:05"), timeSec)
	}
}

func contains(haystack, needle string) bool {
	return strings.Contains(haystack, needle)
}
