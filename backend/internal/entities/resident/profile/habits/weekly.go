package habits

import (
	"fmt"
	"math/rand/v2"
	"simulation/internal/entities/resident/ds/behavioral"
)

type ResidentWeeklyProfile struct {
	profiles []*ResidentDailyProfile
}



func NewResidentWeeklyProfile(values []*ResidentDailyProfile) (*ResidentWeeklyProfile, error) {
	if len(values) > 7 || len(values) == 0 {
		return nil, fmt.Errorf("invalid ResidentWeeklyProfile: 'values' slice must contain between 1 and 7 entries (got entries = %d) \n ", len(values))
	}

	return &ResidentWeeklyProfile{
		profiles: values,
	}, nil
}

func (r *ResidentWeeklyProfile) Profiles() []*ResidentDailyProfile{
	return r.profiles
}

func (r *ResidentWeeklyProfile) GenerateFrequency(day uint8, rng *rand.Rand) (*behavioral.Frequency, error) {
	day = r.normalizeDay(day)
	return r.profiles[day].GenerateFrequency(rng)
}

func (r *ResidentWeeklyProfile) GenerateRoutine(day uint8, rng *rand.Rand) (*behavioral.Routine, error) {
	day = r.normalizeDay(day)
	return r.profiles[day].GenerateRoutine(rng)
}


func (r *ResidentWeeklyProfile) normalizeDay(day uint8) uint8 {
	return day % uint8(len(r.profiles))
}
