package resident

import (
	"residentprofiles"
	"math/rand/v2"
	"residentdata"

)

type ResidentProfile struct {
	weeklyProfile *residentprofiles.ResidentWeeklyProfile
	occupation *residentprofiles.OccupationProfile
}

func (r *ResidentProfile) GenerateOccupation(age uint8, rng *rand.Rand) uint32 {
	if age >= 18 && age < 65 {
		return r.occupation.GenerateAdultSelector(rng)
	} else if age < 18 {
		return r.occupation.GenerateUnderSelector(rng)
	}
	return r.occupation.GenerateOverSelector(rng)
}

func (r *ResidentProfile) GenerateFrequency(day uint8, rng *rand.Rand) *residentdata.Frequency {
	return r.weeklyProfile.GenerateFrequency(day,rng)
}

func (r *ResidentProfile) GenerateRoutine(day uint8, rng *rand.Rand) *residentdata.Routine {
	return r.weeklyProfile.GenerateRoutine(day,rng)
}

func (r *ResidentProfile) GenerateUsage(day uint8, freq *residentdata.Frequency, rng *rand.Rand) *residentdata.Usage {
	return r.weeklyProfile.GenerateUsage(day,freq, rng)
}


