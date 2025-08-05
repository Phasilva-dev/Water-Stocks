package habits

import (
	"math/rand/v2"
	"reflect"
	"strconv"
	"testing"

	"simulation/internal/configs"
	"simulation/internal/dists"
	"simulation/internal/entities/resident/ds/behavioral"
	"simulation/internal/entities/resident/profile/frequency"
	"simulation/internal/entities/resident/profile/routine"
)

// --- Mocks determinísticos para os testes ---

// Distribuição normal com média 1 e desvio 0: sempre retorna 1.0
var mockDist, _ = dists.CreateDistribution("normal", 1, 0)

// Mock de rotina: 4 distribuições idênticas com shift=1
// Deve gerar valores previsíveis como [2.0, 2.0, 2.0, 2.0] após shift
var routineProfileMock, _ = routine.NewDayProfile(
	[]dists.Distribution{mockDist, mockDist, mockDist, mockDist},
	1, 0,
)

// Mock de perfil de frequência com a mesma distribuição e shift 0
var frequencyProfileMock, _ = frequency.NewDeviceProfile(mockDist, 0)

// Gera ResidentDeviceProfiles com todos os dispositivos usando o mesmo profile
func buildFrequencyProfileDay(profile *frequency.DeviceProfile) *frequency.ResidentDeviceProfiles {
	m := make(map[string]*frequency.DeviceProfile)
	for _, key := range configs.OrderedDeviceKeys() {
		m[key] = profile
	}
	profiles, _ := frequency.NewResidentDeviceProfiles(m)
	return profiles
}

var frequencyDayProfileMock = buildFrequencyProfileDay(frequencyProfileMock)


// --- Testes para ResidentDailyProfile ---

func TestResidentDailyProfile(t *testing.T) {

	t.Run("ConstructorAndGetters", func(t *testing.T) {
		rdp := NewResidentDailyProfile(routineProfileMock, frequencyDayProfileMock)

		if rdp.routineProfile != routineProfileMock {
			t.Errorf("routineProfile incorreto. Esperado %v, obtido %v", routineProfileMock, rdp.routineProfile)
		}
		if rdp.frequencyProfileDay != frequencyDayProfileMock {
			t.Errorf("frequencyProfileDay incorreto. Esperado %v, obtido %v", frequencyDayProfileMock, rdp.frequencyProfileDay)
		}

		if got := rdp.RoutineProfile(); got != routineProfileMock {
			t.Errorf("RoutineProfile() retornou errado. Esperado %v, obtido %v", routineProfileMock, got)
		}
		if got := rdp.FrequencyProfileDay(); got != frequencyDayProfileMock {
			t.Errorf("FrequencyProfileDay() retornou errado. Esperado %v, obtido %v", frequencyDayProfileMock, got)
		}
	})

	t.Run("GenerateRoutineDelegation", func(t *testing.T) {
		expectedRoutine := behavioral.NewRoutine([]float64{1.0, 2.0, 3.0, 5.0})
		rdp := NewResidentDailyProfile(routineProfileMock, frequencyDayProfileMock)
		rng := rand.New(rand.NewPCG(123, 456))

		routine, _ := rdp.GenerateRoutine(rng)
		if !reflect.DeepEqual(routine, expectedRoutine) {
			t.Errorf("GenerateRoutine incorreto.\nEsperado: %v\nObtido:   %v", expectedRoutine, routine)
		}
	})

	t.Run("GenerateFrequencyDelegation", func(t *testing.T) {
		expected, _ := behavioral.NewFrequency(map[string]uint8{
			"toilet": 1, "shower": 1, "wash_bassin": 1,
			"wash_machine": 1, "dish_washer": 1, "tanque": 1,
		})
		rdp := NewResidentDailyProfile(routineProfileMock, frequencyDayProfileMock)
		rng := rand.New(rand.NewPCG(789, 1011))

		freq, _ := rdp.GenerateFrequency(rng)
		if !reflect.DeepEqual(freq, expected) {
			t.Errorf("GenerateFrequency incorreto.\nEsperado: %v\nObtido:   %v", expected, freq)
		}
	})
}

// --- Preparação de mocks para ResidentWeeklyProfile ---

var (
	mockDist1, _ = dists.CreateDistribution("normal", 1, 0)
	mockDist2, _ = dists.CreateDistribution("normal", 2, 0)
	mockDist3, _ = dists.CreateDistribution("normal", 3, 0)

	frequencyProfileMock1, _ = frequency.NewDeviceProfile(mockDist1, 0)
	frequencyProfileMock2, _ = frequency.NewDeviceProfile(mockDist2, 0)
	frequencyProfileMock3, _ = frequency.NewDeviceProfile(mockDist3, 0)

	routineProfileMock1, _ = routine.NewDayProfile([]dists.Distribution{mockDist1, mockDist1, mockDist1, mockDist1}, 0, 0)
	routineProfileMock2, _ = routine.NewDayProfile([]dists.Distribution{mockDist2, mockDist2, mockDist2, mockDist2}, 0, 0)
	routineProfileMock3, _ = routine.NewDayProfile([]dists.Distribution{mockDist3, mockDist3, mockDist3, mockDist3}, 0, 0)

	expectedRoutine1 = behavioral.NewRoutine([]float64{1.0, 1.0, 1.0, 1.0})
	expectedRoutine2 = behavioral.NewRoutine([]float64{2.0, 2.0, 2.0, 2.0})
	expectedRoutine3 = behavioral.NewRoutine([]float64{3.0, 3.0, 3.0, 3.0})

	expectedFrequency1, _ = behavioral.NewFrequency(map[string]uint8{
		"toilet": 1, "shower": 1, "wash_bassin": 1, "wash_machine": 1, "dish_washer": 1, "tanque": 1,
	})
	expectedFrequency2, _ = behavioral.NewFrequency(map[string]uint8{
		"toilet": 2, "shower": 2, "wash_bassin": 2, "wash_machine": 2, "dish_washer": 2, "tanque": 2,
	})
	expectedFrequency3, _ = behavioral.NewFrequency(map[string]uint8{
		"toilet": 3, "shower": 3, "wash_bassin": 3, "wash_machine": 3, "dish_washer": 3, "tanque": 3,
	})
)



// Perfis diários mockados com distribuições diferentes
var (
	rdp1 = NewResidentDailyProfile(routineProfileMock1, buildFrequencyProfileDay(frequencyProfileMock1))
	rdp2 = NewResidentDailyProfile(routineProfileMock2, buildFrequencyProfileDay(frequencyProfileMock2))
	rdp3 = NewResidentDailyProfile(routineProfileMock3, buildFrequencyProfileDay(frequencyProfileMock3))
)

// --- Testes para ResidentWeeklyProfile ---

func TestResidentWeeklyProfile(t *testing.T) {

	t.Run("NewResidentWeeklyProfile", func(t *testing.T) {
		// Testes válidos com 1, 3 e 7 perfis
		validCases := []struct {
			profiles []*ResidentDailyProfile
			wantLen  int
		}{
			{[]*ResidentDailyProfile{rdp1}, 1},
			{[]*ResidentDailyProfile{rdp1, rdp2, rdp3}, 3},
			{[]*ResidentDailyProfile{rdp1, rdp2, rdp3, rdp1, rdp2, rdp3, rdp1}, 7},
		}

		for _, tc := range validCases {
			wp, err := NewResidentWeeklyProfile(tc.profiles)
			if err != nil || wp == nil || len(wp.Profiles()) != tc.wantLen {
				t.Errorf("Erro ao criar perfil semanal válido. Esperado %d perfis, erro: %v", tc.wantLen, err)
			}
		}

		// Casos inválidos: 0 e 8 perfis
		if wp, err := NewResidentWeeklyProfile([]*ResidentDailyProfile{}); err == nil || wp != nil {
			t.Error("Esperado erro para 0 perfis, mas foi aceito")
		}
		if wp, err := NewResidentWeeklyProfile(make([]*ResidentDailyProfile, 8)); err == nil || wp != nil {
			t.Error("Esperado erro para 8 perfis, mas foi aceito")
		}
	})

	t.Run("GenerateRoutine", func(t *testing.T) {
		wp, _ := NewResidentWeeklyProfile([]*ResidentDailyProfile{rdp1, rdp2, rdp3})
		rng := rand.New(rand.NewPCG(2023, 11))

		tests := []struct {
			day     uint8
			want    *behavioral.Routine
			wantIdx uint8
		}{
			{0, expectedRoutine1, 0},
			{1, expectedRoutine2, 1},
			{2, expectedRoutine3, 2},
			{3, expectedRoutine1, 0},
			{4, expectedRoutine2, 1},
			{5, expectedRoutine3, 2},
			{6, expectedRoutine1, 0},
		}

		for _, tt := range tests {
			t.Run("Day"+strconv.Itoa(int(tt.day)), func(t *testing.T) {
				got, _ := wp.GenerateRoutine(tt.day, rng)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("GenerateRoutine[%d] incorreto.\nEsperado: %v\nObtido:   %v", tt.day, tt.want, got)
				}
			})
		}
	})

	t.Run("GenerateFrequency", func(t *testing.T) {
		wp, _ := NewResidentWeeklyProfile([]*ResidentDailyProfile{rdp1, rdp2, rdp3})
		rng := rand.New(rand.NewPCG(2023, 11))

		tests := []struct {
			day     uint8
			want    *behavioral.Frequency
			wantIdx uint8
		}{
			{0, expectedFrequency1, 0},
			{1, expectedFrequency2, 1},
			{2, expectedFrequency3, 2},
			{3, expectedFrequency1, 0},
			{4, expectedFrequency2, 1},
			{5, expectedFrequency3, 2},
			{6, expectedFrequency1, 0},
		}

		for _, tt := range tests {
			t.Run("Day"+strconv.Itoa(int(tt.day)), func(t *testing.T) {
				got, _ := wp.GenerateFrequency(tt.day, rng)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("GenerateFrequency[%d] incorreto.\nEsperado: %v\nObtido:   %v", tt.day, tt.want, got)
				}
			})
		}
	})
}
