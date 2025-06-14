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

func GenerateTanqueUsage(routine *behavioral.Routine, device sanitarydevice.SanitaryDevice,
	 rng *rand.Rand,) (*log.Usage, error) {
	//p := rng.Float64() //Isso é a mesma coisa que uma 0 a 1 uniform


	durationUsage := device.GenerateDuration(rng)

	workTime := routine.WorkTime()
	wakeUpTime := routine.WakeupTime()
	sleepTime := routine.SleepTime()
	returnHome := routine.ReturnHome()

	var min, max float64
	var dist dists.UniformDist
	var err error
	var d int

	if wakeUpTime + 3600 > workTime + 1800 { // Isso é uma condição que não faz sentido, pode ser falsa
		if sleepTime > returnHome { // OUTRA CONDIÇÃO QUE É SEMPRE VERADE 
			min, max = returnHome+1800, sleepTime-1800
			d = 1
		} else {
			min, max = returnHome+1800, 86400
			d = 2
			// Isso permite a pessao lavar roupa enquanto dorme
		}
	} else {
		min, max = wakeUpTime+3600, workTime-1800 
		d = 3
		// Isso permite a pessoa potencialmente lavar roupa enquanto trabalha
	}

	dist, err = dists.UniformDistNew(min, max)
	if err != nil {
		return nil, fmt.Errorf("erro ao gerar distribuição de uso do tanque (decisao = %d): %w", d, err)
	}

	startUsage := int32(dist.Sample(rng))
	endUsage := startUsage + durationUsage

	return log.NewUsage(startUsage, endUsage, device.GenerateFlowLeak(rng)), nil
}