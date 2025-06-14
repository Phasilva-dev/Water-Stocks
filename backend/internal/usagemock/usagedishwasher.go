package usagemock

import (
	"simulation/internal/dists"
	"simulation/internal/log"
	"simulation/internal/entities/resident/ds/behavioral"
	"simulation/internal/entities/house/profile/sanitarydevice"

	"fmt"
	"math/rand/v2"
)

/*get_up = WakeUpTime
work_time = WorkTime
sleep_time = get_up - sleep_duration = SleepTime
return_home = time_out + work_time = ReturnHome*/

func GenerateDishWasherUsage(routine *behavioral.Routine, device sanitarydevice.SanitaryDevice,
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

	if sleepTime > returnHome { //Mas isso sempre é verdade .-.
		if p < 0.05 {
			min, max = float64(inverteHorarioCiclico(int32(sleepTime))), workTime
			if min > max {
				fmt.Printf("sleep time = min é %1.f \n",sleepTime)
				fmt.Printf("work time = min é %1.f \n",workTime)
			}
			d = 1
		} else if p < 0.3 {
			min, max = wakeUpTime, workTime
			d = 2
		} else {
			min, max = returnHome, sleepTime
			d = 3
		}
	} else { //Essa condição é sempre falsa
		if p < 0.025 {
			min, max = sleepTime, wakeUpTime // Isso literalmente retorna erro, pois Min > Max
			d = 4
		} else if p < 0.3 {
			min, max = wakeUpTime, workTime
			d = 5
		} else {
			if sleepTime < returnHome { //Isso é literalmente impossivel
				if p < ((86400-returnHome) / (86400-returnHome+sleepTime)) {
					min, max = returnHome, sleepTime
					d = 6
				} else {
					min, max = 0, sleepTime // Isso literalmente ignora a rotina
					d = 7
				}
			} else {
				min, max = returnHome, sleepTime
				d = 8
			}
		}
	}

	dist, err = dists.UniformDistNew(min, max)
	if err != nil {
		return nil, fmt.Errorf("erro ao gerar distribuição de uso do dish_washer (p = %.4f), (decisão = %d): %w", p,d, err)
	}

	startUsage := int32(dist.Sample(rng))
	endUsage := startUsage + durationUsage

	return log.NewUsage(startUsage, endUsage, device.GenerateFlowLeak(rng)), nil

}