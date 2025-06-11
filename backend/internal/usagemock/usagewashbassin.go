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

func GenerateWashBassinUsage(routine *behavioral.Routine, device sanitarydevice.SanitaryDevice,
	 rng *rand.Rand,) (*log.Usage, error) {
	p := rng.Float64()

	durationUsage := device.GenerateDuration(rng)

	workTime := routine.WorkTime()
	wakeUpTime := routine.WakeupTime()
	sleepTime := routine.SleepTime()
	returnHome := routine.ReturnHome()

	var min, max float64
	var dist dists.UniformDist
	var err error

	if workTime-wakeUpTime < 3600 {
		switch {
		case p < ((workTime - wakeUpTime - float64(durationUsage)) / (workTime - wakeUpTime + 3600)):
			min, max = wakeUpTime, workTime
		case p < ((workTime - wakeUpTime + 1800) / (workTime - wakeUpTime + 3600)):
			min, max = returnHome, returnHome+1800
		default: // Esse bloco default está MUITO ERRADO
			if sleepTime < 1800 { //Condição impossivel de ser atingida
				switch {
				case p < (sleepTime / 1800):
					min, max = 0, sleepTime //Isso aqui literalmente diz que a pessoa pode lavar a mão mesmo fora de casa
				default:
					min, max = sleepTime-1800, 86400
				}
			} else {
				min, max = sleepTime-1800, sleepTime
			}
		}
	} else {
		switch {
		case p < 0.15:
			min, max = wakeUpTime, wakeUpTime+1800
		case p < 0.35:
			min, max = workTime-1800, workTime
		case p < 0.75:
			min, max = returnHome, returnHome+1800
		default:
			min, max = sleepTime-1800, sleepTime
		}
	}

	dist, err = dists.UniformDistNew(min, max)
	if err != nil {
		return nil, fmt.Errorf("erro ao gerar distribuição de uso do wash_bassin (min = %.2f, max = %.2f): %w", min, max, err)
	}

	startUsage := int32(dist.Sample(rng))
	endUsage := startUsage + durationUsage

	return log.NewUsage(startUsage, endUsage, device.GenerateFlowLeak(rng)), nil
}