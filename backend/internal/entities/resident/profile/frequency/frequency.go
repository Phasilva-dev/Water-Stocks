package frequency

import (
	"math/rand/v2"
	"simulation/internal/dists"
	"simulation/internal/entities/resident/ds/behavioral"
	"fmt"
)

type DeviceProfile interface {
	MinValue() uint8
	StatDist() dists.Distribution
	generateFrequency(rng *rand.Rand, shift uint8, statDist dists.Distribution) uint8
	GenerateData(rng *rand.Rand) uint8
	IsIndividual() bool

}

type ResidentDeviceProfiles interface {
	freqDevice() map[string]DeviceProfile
	GenerateData(rng *rand.Rand) (*behavioral.Frequency, error)
	DeviceProfile(typo string) (DeviceProfile, bool)

}

func CreateDeviceProfile(typo string,dist dists.Distribution, minValue uint8) (DeviceProfile, error) {
	switch typo {
	case "individual":
		return newindividualDeviceProfile(dist,minValue)
	case "household":
		return newhouseholdDeviceProfile(dist,minValue)
	default:
		return nil, fmt.Errorf("invalid FrequencyDeviceProfile Factory: unknown distribution type '%s'", typo)
	}
}

func CreateResidentDeviceProfile(typo string,
	frequencyDeviceProfiles map[string]DeviceProfile) (ResidentDeviceProfiles, error) {
	switch typo {
	case "normal":
		return newResidentDeviceProfilesService(frequencyDeviceProfiles)
	default:
		return nil, fmt.Errorf("invalid Resident.FrequencyDeviceProfile Factory: unknown distribution type '%s'", typo)
	}
}
