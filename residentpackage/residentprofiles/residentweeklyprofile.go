package residentprofiles

import (
	"errors"
	"math/rand/v2"
	"residentdata"
)

type ResidentWeeklyProfile struct {
	profiles []*ResidentDayProfile
}

func NewResidentWeeklyProfile(values []*ResidentDayProfile) (*ResidentWeeklyProfile, error) {
	if len(values) > 7 && len(values) <= 0 {
		return nil, errors.New("Deu erro")
	} 

	return &ResidentWeeklyProfile{
		profiles: values,
	}, nil
}

func (r *ResidentWeeklyProfile) GenerateDailyData(day int64,rng *rand.Rand) *residentdata.DailyData{
	return r.profiles[day].GenerateData(rng)
}