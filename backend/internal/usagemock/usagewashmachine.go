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

func GenerateWashMachineUsage(routine *behavioral.Routine, house *entities.House, rng *rand.Rand) (*log.Usage, error) {
	
	dist, err := dists.NewLogLogisticDist(10.448, 0.167418) //Problema, não entendi oq foi feito aqui

	if err != nil {
		return nil, err
	}

	startUsage := int32(dist.Sample(rng))
	device := house.SanitaryHouse().WashMachine().Device()
	durationUsage := device.GenerateDuration(rng)

	//Deve ter um tratamento de Colisão aqui

	endUsage := startUsage + durationUsage

	return log.NewUsage(startUsage,endUsage,device.GenerateFlowLeak(rng)), nil


}