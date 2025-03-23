package resident

import (
	"profiles"
	"unique"
	"dists"

	"golang.org/x/exp/rand"
)

type Routine struct {
	Symbol unique.Handle[string]
	Times []ProfileTuple
}

func generateTime (dist dists.Distribution ,rng rand.Source) int32 {
	var time float64 = dist.Sample(rng)
	truncatedTime := int32(time)
	return truncatedTime
}

func enforceMinimunGap(entryTime, exitTime int32, gap int32) int32 {
	if exitTime - (entryTime+gap) < gap {
		exitTime = entryTime + gap
	}
	return exitTime
}

func NewRoutine (profile profiles.RoutineProfileDist, rng rand.Source, gap int32) *Routine {
	events = profile.Events()
	id = profile.Symbol()
	times = []int32
	for i in range(events) {
		times.append(generateTime(events[i].LeaveTime(), rng))
		times.append(generateTime(events[i].ReturnTime(), rng))
		
	}

}

