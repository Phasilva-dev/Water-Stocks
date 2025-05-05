package habits

import (
	"errors"
	"math/rand/v2"
	"simulation/internal/entities/resident/ds/behavioral"
)

type ResidentWeeklyProfile struct {
	profiles []*ResidentDayProfile
}

func NewResidentWeeklyProfile(values []*ResidentDayProfile) (*ResidentWeeklyProfile, error) {
	if len(values) > 7 || len(values) == 0 {
		return nil, errors.New("profile list must contain between 1 and 7 entries")
	}

	return &ResidentWeeklyProfile{
		profiles: values,
	}, nil
}

func (r *ResidentWeeklyProfile) GenerateFrequency(day uint8, rng *rand.Rand) *behavioral.Frequency {
	day = r.normalizeDay(day)
	return r.profiles[day].GenerateFrequency(rng)
}

func (r *ResidentWeeklyProfile) GenerateRoutine(day uint8, rng *rand.Rand) *behavioral.Routine {
	day = r.normalizeDay(day)
	return r.profiles[day].GenerateRoutine(rng)
}


func (r *ResidentWeeklyProfile) normalizeDay(day uint8) uint8 {
	return day % uint8(len(r.profiles))
}
