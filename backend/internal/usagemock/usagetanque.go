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

func GenerateTanqueUsage(routine *behavioral.Routine, house *entities.House, rng *rand.Rand) (*log.Usage, error) {
	//p := rng.Float64() //Isso é a mesma coisa que uma 0 a 1 uniform

	device := house.SanitaryHouse().Tanque().Device()
	durationUsage := device.GenerateDuration(rng)

	workTime := routine.WorkTime()
	wakeUpTime := routine.WakeupTime()
	sleepTime := routine.SleepTime()
	returnHome := routine.ReturnHome()

	var dist dists.UniformDist
	var err error

	if wakeUpTime + 3600 > workTime + 1800 { // Isso é uma condição que não faz sentido, pode ser falsa
		if sleepTime > returnHome { // OUTRA CONDIÇÃO QUE É SEMPRE VERADE 
			dist, err = dists.UniformDistNew(returnHome+1800, sleepTime-1800)
		} else {
			dist, err = dists.UniformDistNew(returnHome+1800, 86400) // Isso permite a pessao lavar roupa enquanto dorme
		}
	} else {
		dist, err = dists.UniformDistNew(wakeUpTime+3600, workTime-1800) // Isso permite a pessoa potencialmente lavar roupa enquanto trabalha
	}

	if err != nil {
		return nil, err
	}

	startUsage := int32(dist.Sample(rng))
	endUsage := startUsage + durationUsage

	return log.NewUsage(startUsage, endUsage, device.GenerateFlowLeak(rng)), nil
}