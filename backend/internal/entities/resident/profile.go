package resident

import (
	"math/rand/v2"
	"simulation/internal/entities/resident/ds/behavioral"
	"simulation/internal/entities/resident/profile/habits"
	"errors"
)

type ResidentProfile struct {
	OccupationID uint32
	weeklyProfile *habits.ResidentWeeklyProfile

}

func NewResidentProfile(profile *habits.ResidentWeeklyProfile, id uint32) (*ResidentProfile, error) {

	if profile == nil {
		return nil, errors.New("weekly profile cannot be nil")
	}
	return &ResidentProfile{
		OccupationID: id,
		weeklyProfile: profile,
	}, nil
}

func (r *ResidentProfile) GenerateFrequency(day uint8, rng *rand.Rand) *behavioral.Frequency {
	return r.weeklyProfile.GenerateFrequency(day,rng)
}

func (r *ResidentProfile) GenerateRoutine(day uint8, rng *rand.Rand) *behavioral.Routine {
	return r.weeklyProfile.GenerateRoutine(day,rng)
}



