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

func GenerateWashMachineUsagee(routine *behavioral.Routine, device sanitarydevice.SanitaryDevice,
	 rng *rand.Rand,) (*log.Usage, error) {
	
	
	shape := 10.448
	scale := 0.167418
	dist, err := dists.NewLogLogisticDist(shape, scale) //Problema, não entendi oq foi feito aqui

	if err != nil {
		return nil, fmt.Errorf("erro ao gerar distribuição de uso do wash_machine (shape = %.2f, scale = %.2f): %w", shape, scale, err)
	}

	startUsage := int32(dist.Sample(rng) * 86400) //É necessario a multiplicação para dar sentido

	durationUsage := device.GenerateDuration(rng)

	//Deve ter um tratamento de Colisão aqui

	endUsage := startUsage + durationUsage

	usage, err := log.NewUsage(startUsage,endUsage,device.GenerateFlowLeak(rng))

	warningUsage(usage,"wash_machine", 0, 0, 0, 0)

	return usage, err


}

func GenerateWashMachineUsage(routine *behavioral.Routine, device sanitarydevice.SanitaryDevice,
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

	case p < 0.10: 
		min, max = 0, 86400 //Se possivel, seria bom não usar valores fixos
    	d = 1
	case p < 0.40:
		min, max = wakeUpTime, workTime //min, max = wakeUpTime+1800, workTime-1800
		d = 3
	default:
		min, max = returnHome, sleepTime
		d = 7
	}

	dist, err = dists.UniformDistNew(min, max)
	if err != nil {
		return nil, fmt.Errorf("erro ao gerar distribuição de uso do wash_Machine (p = %.4f), (decisao = %d): %w", p, d, err)
	}

	startUsage := int32(dist.Sample(rng))

	durationUsage := device.GenerateDuration(rng)

	//Deve ter um tratamento de Colisão aqui

	endUsage := startUsage + durationUsage

	usage, err := log.NewUsage(startUsage,endUsage,device.GenerateFlowLeak(rng))

	//warningUsage(usage,"toilet",d,p, 0, 0)

	return usage, err

}