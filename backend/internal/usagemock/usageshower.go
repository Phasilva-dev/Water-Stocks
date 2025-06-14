package usagemock

import (
	"simulation/internal/dists"
	"simulation/internal/log"
	"simulation/internal/entities/house/profile/sanitarydevice"
	"simulation/internal/entities/resident/ds/behavioral"

	"fmt"
	"math/rand/v2"
)

/*get_up = WakeUpTime
work_time = WorkTime
sleep_time = get_up - sleep_duration = SleepTime
return_home = time_out + work_time = ReturnHome*/

func GenerateShowerUsage(routine *behavioral.Routine, device sanitarydevice.SanitaryDevice,
	 rng *rand.Rand,) (*log.Usage, error) {

	p := rng.Float64() //Isso é a mesma coisa que uma 0 a 1 uniform


	durationUsage := device.GenerateDuration(rng)

	workTime := routine.WorkTime()
	wakeUpTime := routine.WakeupTime()
	sleepTime := routine.SleepTime()
	returnHome := routine.ReturnHome()

	var min, max float64
	var dist dists.UniformDist
	var err error
	var d int

	if workTime-wakeUpTime < 3600 {
		// período muito curto entre acordar e sair
		switch {
		case p < ((workTime-wakeUpTime) / (workTime-wakeUpTime+3600)):
			min, max = wakeUpTime, workTime
			d = 1
		case p < ((workTime-wakeUpTime+1800) / (workTime-wakeUpTime+3600)):
			min, max = returnHome, returnHome+1800
			d = 2
		default:
			min, max = sleepTime-1800, sleepTime
			d = 3
		}

	} else {
		switch {
		case p < 0.5: //Caso mais comum 50%
			if returnHome > 18*3600 {
				min, max = returnHome, returnHome+1800
				d = 4
			} else {
				min, max = returnHome, sleepTime-1800
				d = 5
			}
		case p < 0.8: // 30%
			min, max = workTime-1800, workTime
			d = 6
		case p < 0.95: // 15%
			min, max = wakeUpTime, wakeUpTime+1800
			d = 7
		default: //5%
			min, max = sleepTime-1800, sleepTime
			d = 8
		}
	}

	dist, err = dists.UniformDistNew(min, max)
	if err != nil {
		return nil, fmt.Errorf("erro ao gerar distribuição de uso do shower (p = %.4f) (decisao = %d): %w", p, d, err)
	}

	startUsage := int32(dist.Sample(rng))
	endUsage := startUsage + durationUsage

	return log.NewUsage(startUsage, endUsage, device.GenerateFlowLeak(rng)), nil
}
