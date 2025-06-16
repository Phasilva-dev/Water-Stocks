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

type UsageToilet struct {
}

func GenerateToiletUsage(routine *behavioral.Routine, device sanitarydevice.SanitaryDevice,
	 rng *rand.Rand,) (*log.Usage, error) {
	p := rng.Float64() //Isso é a mesma coisa que uma 0 a 1 uniform

	workTime := routine.WorkTime()
	wakeUpTime := routine.WakeupTime()
	sleepTime := routine.SleepTime()
	returnHome := routine.ReturnHome()
	
	var min, max float64
	var dist dists.UniformDist
	var err error
	var d int
	
	switch {

	case p < 0.025: 
		min, max = sleepTime - 86400, 86400 //Se possivel, seria bom não usar valores fixos
    	d = 1
	case p < 0.05:
		min, max = 0, wakeUpTime //Se possivel, seria bom não usar valores fixos
    	d = -1
	case p < 0.15:
		min, max = wakeUpTime, wakeUpTime+1800
		d = 2
	case p < 0.20:
		min, max = wakeUpTime+900, workTime-900 //min, max = wakeUpTime+1800, workTime-1800
		d = 3
	case p < 0.325:
		min, max = workTime-1800, workTime
		d = 4
	case p < 0.45:
		min, max = returnHome, returnHome+1800
		d = 5
	case p < 0.55:
		min, max = sleepTime-1800, sleepTime
		d = 6
	default:
		min, max = returnHome, sleepTime-1800
		d = 7
	}

	dist, err = dists.UniformDistNew(min, max)
	if err != nil {
		return nil, fmt.Errorf("erro ao gerar distribuição de uso do toilet (p = %.4f), (decisao = %d): %w", p, d, err)
	}

	startUsage := int32(dist.Sample(rng))

	durationUsage := device.GenerateDuration(rng)

	//Deve ter um tratamento de Colisão aqui

	endUsage := startUsage + durationUsage

	return log.NewUsage(startUsage,endUsage,device.GenerateFlowLeak(rng)), nil

}
