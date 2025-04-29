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
	if len(values) > 7 || len(values) == 0 {
		return nil, errors.New("profile list must contain between 1 and 7 entries")
	}

	return &ResidentWeeklyProfile{
		profiles: values,
	}, nil
}

func (r *ResidentWeeklyProfile) GenerateFrequency(day uint8, rng *rand.Rand) *residentdata.Frequency {
	day = r.normalizeDay(day)
	return r.profiles[day].GenerateFrequency(rng)
}

func (r *ResidentWeeklyProfile) GenerateRoutine(day uint8, rng *rand.Rand) *residentdata.Routine {
	day = r.normalizeDay(day)
	return r.profiles[day].GenerateRoutine(rng)
}

func (r *ResidentWeeklyProfile) GenerateUsage(day uint8, freq *residentdata.Frequency, rng *rand.Rand) (*residentdata.Usage, error) {
	day = r.normalizeDay(day)
	return r.profiles[day].GenerateUsage(rng, freq)

}

func (r *ResidentWeeklyProfile) normalizeDay(day uint8) uint8 {
	return day % uint8(len(r.profiles))
}
/* Acredito que não vai ser mais util
func (r *ResidentWeeklyProfile) GenerateDailyData(day uint16,rng *rand.Rand) *residentdata.DailyData{
	return r.profiles[day].GenerateData(rng)
}*/