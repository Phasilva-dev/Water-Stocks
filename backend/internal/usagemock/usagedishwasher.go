package usagemock

import (
	"simulation/internal/dists"
	"simulation/internal/log"
	"simulation/internal/entities/resident/ds/behavioral"
	"simulation/internal/entities/house/profile/sanitarydevice"

	//"errors"
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

	var dist dists.UniformDist
	var err error

	if sleepTime > returnHome { //Mas isso sempre é verade .-.
		if p < 0.05 {
			dist, err = dists.UniformDistNew(float64(inverteHorarioCiclico(int32(sleepTime))), workTime)
		} else if p < 0.3 {
			dist, err = dists.UniformDistNew(wakeUpTime, workTime)
		} else {
			dist, err = dists.UniformDistNew(returnHome, sleepTime)
		}
	} else { //Essa condição é sempre falsa
		if p < 0.025 {
			dist, err = dists.UniformDistNew(sleepTime, wakeUpTime) // Isso literalmente retorna erro, pois Min > Max
		} else if p < 0.3 {
			dist, err = dists.UniformDistNew(wakeUpTime, workTime)
		} else {
			if sleepTime < returnHome { //Isso é literalmente impossivel
				if p < ((86400-returnHome) / (86400-returnHome+sleepTime)) {
					dist, err = dists.UniformDistNew(returnHome, sleepTime)
				} else {
					dist, err = dists.UniformDistNew(0, sleepTime) // Isso literalmente ignora a rotina
				}
			} else {
				dist, err = dists.UniformDistNew(returnHome, sleepTime)
			}
		}
	}

	if err != nil {
		return nil, err
	}

	startUsage := int32(dist.Sample(rng))
	endUsage := startUsage + durationUsage

	return log.NewUsage(startUsage, endUsage, device.GenerateFlowLeak(rng)), nil

}