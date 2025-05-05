package usagemock
/*
import (
	"dists"
	"math/rand/v2"
	"residentdata"
	"errors"
	"entities"
)*/
/*get_up = WakeUpTime
work_time = WorkTime
sleep_time = get_up - sleep_duration = SleepTime
return_home = time_out + work_time = ReturnHome*/
/*
type UsageToilet struct {
}

func GenerateUsage(routine *residentdata.Routine, house *entities.House, rng *rand.Rand) (*residentdata.UsageLog, error) {
	dist, _ := dists.UniformDistNew(0,1)
	p := dist.Sample(rng)

	if p < 0.05 {
		usageTime, err := dists.UniformDistNew(float64(residentdata.InverteHorarioCiclico(routine.SleepTime())), 86400)
	} else if p < 0.15 {
		usageTime, err := dists.UniformDistNew(float64(routine.WakeupTime()),float64(routine.WakeupTime())+1800)
	} else if p < 0.20 {
		usageTime, err := dists.UniformDistNew(float64(routine.WakeupTime())+1800,float64(routine.WorkTime())-1800)
	} else if p < 0.325 {
		usageTime, err := dists.UniformDistNew(float64(routine.WorkTime())-1800, float64(routine.WorkTime()))
	} else if p < 0.45 {
		usageTime, err := dists.UniformDistNew(float64(routine.ReturnHome()), float64(routine.ReturnHome())+1800)
	} else if p < 0.55 {
		usageTime, err := dists.UniformDistNew(float64(routine.SleepTime())-1800, float64(routine.SleepTime()))
	} else {
		usageTime, err := dists.UniformDistNew(float64(routine.ReturnHome()),float64(routine.SleepTime())-1800)
	}


	
}*/