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
	var d int

	if workTime-wakeUpTime < 3600 {
		switch {
		case p < ((workTime - wakeUpTime - float64(durationUsage)) / (workTime - wakeUpTime + 3600)):
			min, max = wakeUpTime, workTime
			d = 1
		case p < ((workTime - wakeUpTime + 1800) / (workTime - wakeUpTime + 3600)):
			min, max = returnHome, returnHome+1800
			d = 2
		default: // Esse bloco default está MUITO ERRADO
			if sleepTime < 1800 { //Condição impossivel de ser atingida
				switch {
				case p < (sleepTime / 1800):
					min, max = 0, sleepTime //Isso aqui literalmente diz que a pessoa pode lavar a mão mesmo fora de casa
					d = 3
				default:
					min, max = sleepTime-1800, 86400
					d = 4
				}
			} else {
				min, max = sleepTime-1800, sleepTime
				d = 5
			}
		}
	} else {
		switch {
		case p < 0.15:
			min, max = wakeUpTime, wakeUpTime+1800
			d = 6
		case p < 0.35:
			min, max = workTime-1800, workTime
			d = 7
		case p < 0.75:
			min, max = returnHome, returnHome+1800
			d = 8
		default:
			min, max = sleepTime-1800, sleepTime
			d = 9
		}
	}

	dist, err = dists.UniformDistNew(min, max)
	if err != nil {
		return nil, fmt.Errorf("erro ao gerar distribuição de uso do wash_bassin (p = %.4f), (decisao = %d): %w", p, d, err)
	}

	startUsage := int32(dist.Sample(rng))
	endUsage := startUsage + durationUsage

	usage, err := log.NewUsage(startUsage, endUsage, device.GenerateFlowLeak(rng))

	//warningUsage(usage,"wash_bassin",d,p, 0,0)

	return usage, err
}