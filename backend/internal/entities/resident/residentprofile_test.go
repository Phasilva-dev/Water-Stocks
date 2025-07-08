package resident

import (
	"math/rand/v2"

	"simulation/internal/dists"
	"testing"
	"reflect"

	"simulation/internal/entities/resident/ds/behavioral"
	"simulation/internal/entities/resident/profile/frequency"
	"simulation/internal/entities/resident/profile/routine"
	"simulation/internal/entities/resident/profile/habits"

)

const (
	freqToilet      = "toilet"
	freqShower      = "shower"
	freqWashBassin  = "washBassin"
	freqWashMachine = "washMachine"
	freqDishWasher  = "dishWasher"
	freqTanque      = "tanque"
)
//Variaveis para testar weekly

var mockFreqDist, _ = dists.CreateDistribution("normal", 1, 0)
var mockWakeUpDist, _ = dists.CreateDistribution("normal", 6 * 3600, 0)
var mockWorkTimeDist, _ = dists.CreateDistribution("normal", 8 * 3600, 0)
var mockReturnHomeDist, _ = dists.CreateDistribution("normal", 18 * 3600, 0)
var mockSleepTimeDist, _ = dists.CreateDistribution("normal", 22 * 3600, 0)

var frequencyProfileMock, _ = frequency.NewFrequencyProfile(mockFreqDist, 0)


var routineProfileMock, _ = routine.NewRoutineProfile(
	[]dists.Distribution{mockWakeUpDist, mockWorkTimeDist, mockReturnHomeDist, mockSleepTimeDist},
	0,0,
)

var rdp = habits.NewResidentDayProfile(routineProfileMock, frequency.NewFrequencyProfileDay(map[string]*frequency.FrequencyProfile{
	freqToilet:     frequencyProfileMock,
	freqShower:     frequencyProfileMock,
	freqWashBassin: frequencyProfileMock,
	freqWashMachine: frequencyProfileMock,
	freqDishWasher: frequencyProfileMock,
	freqTanque:     frequencyProfileMock,
}))

var expectedRoutine = behavioral.NewRoutine([]float64{6 * 3600, 8 * 3600, 18 * 3600, 22 * 3600})
var expectedFrequency = behavioral.NewFrequency(1,1,1,1,1,1)

var profiles = []*habits.ResidentDayProfile{rdp}
var wp, _ = habits.NewResidentWeeklyProfile(profiles)
var id uint32 = 1


func TestNewResidentProfile_Success(t *testing.T) {
	rp, err := NewResidentProfile(wp, id)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if rp == nil {
		t.Fatal("expected non-nil ResidentProfile")
	}
	if rp.OccupationID != id {
		t.Errorf("expected OccupationID %d, got %d", id, rp.OccupationID)
	}
}

func TestNewResidentProfile_NilWeeklyProfile(t *testing.T) {
	rp, err := NewResidentProfile(nil, id)
	if err == nil {
		t.Error("expected error when weekly profile is nil, got nil")
	}
	if rp != nil {
		t.Error("expected nil ResidentProfile, got non-nil")
	}
	expectedErr := "invalid ResidentProfile: weekly profile cannot be nil"
	if err.Error() != expectedErr {
		t.Errorf("expected error '%s', got '%s'", expectedErr, err.Error())
	}
}

func TestGenerateRoutine(t *testing.T) {
	rp, _ := NewResidentProfile(wp, id)
	rng := rand.New(rand.NewPCG(42, 54)) // determinístico

	routine := rp.GenerateRoutine(0, rng)
	if routine == nil {
		t.Fatal("expected non-nil Routine")
	}

	if !reflect.DeepEqual(routine.Times(), expectedRoutine.Times()) {
		t.Errorf("expected routine events %v, got %v", expectedRoutine.Times(), routine.Times())
	}
}

func TestGenerateFrequency(t *testing.T) {
	rp, _ := NewResidentProfile(wp, id)
	rng := rand.New(rand.NewPCG(42, 54))

	freq := rp.GenerateFrequency(0, rng)
	if freq == nil {
		t.Fatal("expected non-nil Frequency")
	}

	compareFrequency := func(fieldName string, got, want uint8) {
		if got != want {
			t.Errorf("expected %s = %d, got %d", fieldName, want, got)
		}
	}

	compareFrequency("Toilet", freq.Toilet(), expectedFrequency.Toilet())
	compareFrequency("Shower", freq.Shower(), expectedFrequency.Shower())
	compareFrequency("WashBassin", freq.WashBassin(), expectedFrequency.WashBassin())
	compareFrequency("WashMachine", freq.WashMachine(), expectedFrequency.WashMachine())
	compareFrequency("DishWasher", freq.DishWasher(), expectedFrequency.DishWasher())
	compareFrequency("Tanque", freq.Tanque(), expectedFrequency.Tanque())
}