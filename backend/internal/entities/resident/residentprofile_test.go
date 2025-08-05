package resident

import (
	"math/rand/v2"

	"reflect"
	"simulation/internal/configs"
	"simulation/internal/dists"
	"testing"

	"simulation/internal/entities/resident/ds/behavioral"
	"simulation/internal/entities/resident/profile/frequency"
	"simulation/internal/entities/resident/profile/habits"
	"simulation/internal/entities/resident/profile/routine"
)


//Variaveis para testar weekly

var mockFreqDist, _ = dists.CreateDistribution("normal", 1, 0)
var mockWakeUpDist, _ = dists.CreateDistribution("normal", 6 * 3600, 0)
var mockWorkTimeDist, _ = dists.CreateDistribution("normal", 8 * 3600, 0)
var mockReturnHomeDist, _ = dists.CreateDistribution("normal", 18 * 3600, 0)
var mockSleepTimeDist, _ = dists.CreateDistribution("normal", 22 * 3600, 0)

var frequencyProfileMock, _ = frequency.NewDeviceProfile(mockFreqDist, 0)


var routineProfileMock, _ = routine.NewDayProfile(
	[]dists.Distribution{mockWakeUpDist, mockWorkTimeDist, mockReturnHomeDist, mockSleepTimeDist},
	0,0,
)

func buildFrequencyProfileDay(profile *frequency.DeviceProfile) *frequency.ResidentDeviceProfiles {
	m := make(map[string]*frequency.DeviceProfile)
	for _, key := range configs.OrderedDeviceKeys() {
		m[key] = profile
	}
	profiles, _ := frequency.NewResidentDeviceProfiles(m)
	return profiles
}

var rdp = habits.NewResidentDailyProfile(routineProfileMock,
	buildFrequencyProfileDay(frequencyProfileMock))

var expectedRoutine = behavioral.NewRoutine([]float64{6 * 3600, 8 * 3600, 18 * 3600, 22 * 3600})
var expectedFrequency, _ = behavioral.NewFrequency(map[string]uint8{
		"toilet": 1, "shower": 1, "wash_bassin": 1, "wash_machine": 1, "dish_washer": 1, "tanque": 1,
	})

var profiles = []*habits.ResidentDailyProfile{rdp}
var wp, _ = habits.NewResidentWeeklyProfile(profiles)
var id uint32 = 1


func TestNewResidentProfile_Success(t *testing.T) {
	rp, err := NewProfile(wp, id)
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
	rp, err := NewProfile(nil, id)
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
	rp, _ := NewProfile(wp, id)
	rng := rand.New(rand.NewPCG(42, 54)) // determinístico

	routine, _ := rp.GenerateRoutine(0, rng)
	if routine == nil {
		t.Fatal("expected non-nil Routine")
	}

	if !reflect.DeepEqual(routine.Times(), expectedRoutine.Times()) {
		t.Errorf("expected routine events %v, got %v", expectedRoutine.Times(), routine.Times())
	}
}

func TestGenerateFrequency(t *testing.T) {
	rp, _ := NewProfile(wp, id)
	rng := rand.New(rand.NewPCG(42, 54))

	freq, _ := rp.GenerateFrequency(0, rng)
	if freq == nil {
		t.Fatal("expected non-nil Frequency")
	}

	compareFrequency := func(fieldName string, got, want uint8) {
		if got != want {
			t.Errorf("expected %s = %d, got %d", fieldName, want, got)
		}
	}

	compareFrequency("Toilet", freq.DeviceFrequency("toilet"), expectedFrequency.DeviceFrequency("toilet"))
	compareFrequency("Shower", freq.DeviceFrequency("shower"), expectedFrequency.DeviceFrequency("shower"))
	compareFrequency("WashBassin", freq.DeviceFrequency("wash_bassin"), expectedFrequency.DeviceFrequency("wash_bassin"))
	compareFrequency("WashMachine", freq.DeviceFrequency("wash_machine"), expectedFrequency.DeviceFrequency("wash_machine"))
	compareFrequency("DishWasher", freq.DeviceFrequency("dish_washer"), expectedFrequency.DeviceFrequency("dish_washer"))
	compareFrequency("Tanque", freq.DeviceFrequency("tanque"), expectedFrequency.DeviceFrequency("tanque"))
}