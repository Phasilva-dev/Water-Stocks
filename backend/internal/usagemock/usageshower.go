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

func GenerateShowerUsage(routine *behavioral.Routine, house *entities.House, rng *rand.Rand) (*log.Usage, error) {

	p := rng.Float64() //Isso é a mesma coisa que uma 0 a 1 uniform

	device := house.SanitaryHouse().Shower().Device()
	durationUsage := device.GenerateDuration(rng)

	workTime := routine.WorkTime()
	wakeUpTime := routine.WakeupTime()
	sleepTime := routine.SleepTime()
	returnHome := routine.ReturnHome()

	// Pré-calcula versões em float64
	ftWorkTime := float64(workTime)
	ftWakeUp := float64(wakeUpTime)
	ftSleep := float64(sleepTime)
	ftReturn := float64(returnHome)

	var dist dists.UniformDist
	var err error

	if workTime-wakeUpTime < 3600 {
		// período muito curto entre acordar e sair
		switch {
		case p < ((ftWorkTime-ftWakeUp) / (ftWorkTime-ftWakeUp+3600)):
			dist, err = dists.UniformDistNew(ftWakeUp, ftWorkTime)
		case p < ((ftWorkTime-ftWakeUp+1800) / (ftWorkTime-ftWakeUp+3600)):
			dist, err = dists.UniformDistNew(ftReturn, ftReturn+1800)
		default:
			dist, err = dists.UniformDistNew(ftSleep-1800, ftSleep)
		}

	} else {
		switch {
		case p < 0.5: //Caso mais comum 50%
			if returnHome > 18*3600 {
				dist, err = dists.UniformDistNew(ftReturn, ftReturn+1800)
			} else {
				dist, err = dists.UniformDistNew(ftReturn, ftSleep-1800)
			}
		case p < 0.8: // 30%
			dist, err = dists.UniformDistNew(ftWorkTime-1800, ftWorkTime)
		case p < 0.95: // 15%
			dist, err = dists.UniformDistNew(ftWakeUp, ftWakeUp+1800)
		default: //5%
			dist, err = dists.UniformDistNew(ftSleep-1800, ftSleep)
		}
	}

	if err != nil {
		return nil, err
	}

	startUsage := int32(dist.Sample(rng))
	endUsage := startUsage + durationUsage

	return log.NewUsage(startUsage, endUsage, device.GenerateFlowLeak(rng)), nil
}
