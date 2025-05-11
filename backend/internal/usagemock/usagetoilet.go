package usagemock

import (
	"simulation/internal/dists"
	"simulation/internal/log"
	"simulation/internal/entities"
	"simulation/internal/entities/resident/ds/behavioral"

	//"errors"
	"math/rand/v2"
)
/*get_up = WakeUpTime
work_time = WorkTime
sleep_time = get_up - sleep_duration = SleepTime
return_home = time_out + work_time = ReturnHome*/

type UsageToilet struct {
}

func GenerateToiletUsage(routine *behavioral.Routine, house *entities.House, rng *rand.Rand) (*log.Usage, error) {
	probDist, _ := dists.UniformDistNew(0, 1)
	p := probDist.Sample(rng)

	workTime := routine.WorkTime()
	wakeUpTime := routine.WakeupTime()
	sleepTime := routine.SleepTime()
	returnHome := routine.ReturnHome()
	
	var dist dists.UniformDist
	var err error
	var startUsage int32
	
	switch {
	case p < 0.05:
		dist, err = dists.UniformDistNew(float64(inverteHorarioCiclico(int32(wakeUpTime))), 86400)
	case p < 0.15:
		dist, err = dists.UniformDistNew(wakeUpTime, wakeUpTime+1800)
	case p < 0.20:
		dist, err = dists.UniformDistNew(wakeUpTime+1800, workTime-1800)
	case p < 0.325:
		dist, err = dists.UniformDistNew(workTime-1800, workTime)
	case p < 0.45:
		dist, err = dists.UniformDistNew(returnHome, returnHome+1800)
	case p < 0.55:
		dist, err = dists.UniformDistNew(sleepTime-1800, sleepTime)
	default:
		dist, err = dists.UniformDistNew(returnHome, sleepTime-1800)
	}

	if err != nil {
		return nil, err
	}

	startUsage = int32(dist.Sample(rng))
	device := house.SanitaryHouse().Toilet().Device()
	durationUsage := device.GenerateDuration(rng)

	//Deve ter um tratamento de Colisão aqui

	endUsage := startUsage + durationUsage

	return log.NewUsage(startUsage,endUsage,device.GenerateFlowLeak(rng)), nil

}
