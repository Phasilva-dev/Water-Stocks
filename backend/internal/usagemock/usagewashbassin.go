package usagemock
/*
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
/*
func GenerateWashBassinUsage(routine *behavioral.Routine, house *entities.House, rng *rand.Rand) (*log.Usage, error) {
	p := rng.Float64() //Isso é a mesma coisa que uma 0 a 1 uniform

	device := house.SanitaryHouse().WashBassin().Device()
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

	if workTime - wakeUpTime < 3600 {
		if p < ((work))
	}



	if work_time-get_up<3600
            if ph_pia<((work_time-get_up-duracao(f))/(work_time-get_up+3600))
                horario(f)=random('Uniform',get_up,work_time);
            elseif (((work_time-get_up)/(work_time-get_up+3600))<=ph_pia) && (ph_pia<(work_time-get_up+1800)/(work_time-get_up+3600))
                horario(f)=random('Uniform',return_home,return_home+1800);
            else
                if sleep_time<1800
                    p=random('Uniform',0,1);
                    if p<(sleep_time/1800)
                        horario(f)=random('Uniform',0,sleep_time);
                    else
                        horario(f)=random('Uniform',86400-1800+sleep_time,86400);
                    end
                else
                    horario(f)=random('Uniform',sleep_time-1800,sleep_time);
                end
            end

}*/