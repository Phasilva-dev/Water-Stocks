package residentprofiles

import (
	"dists"
	"math/rand/v2"
	"residentdata"
)
/*get_up = WakeUpTime
work_time = WorkTime
sleep_time = get_up - sleep_duration = SleepTime
return_home = time_out + work_time = ReturnHome*/
type UsageToilet struct {
}

func GenerateUsage(routine *residentdata.Routine, rng *rand.Rand) {
	dist, _ := dists.UniformDistNew(0,1)
	p := dist.Sample(rng)

	if p < 0.025 {
		usageTime = dists.UniformDistNew(float64(routine.SleepTime())+86400, 86400)
	}


	if n<0.025
                horario(x)=random('Uniform',get_up-sleep_duration+86400,86400);
            elseif (0.025<=n) && (n<0.05)
                horario(x)=random('Uniform',0,get_up);

            elseif (0.05<=n) && (n<0.15)
                horario(x)=random('Uniform',get_up,get_up+1800);
            elseif (0.15<=n) && (n<0.20)
                horario(x)=random('Uniform',get_up+1800,work_time-1800);
            elseif (0.2<=n) && (n<0.325)
                horario(x)=random('Uniform',work_time-1800,work_time);
            elseif (0.325<=n) && (n<0.45)
                horario(x)=random('Uniform',return_home,return_home+1800);

            elseif (0.45<=n) && (n<0.55)
                horario(x)=random('Uniform',sleep_time-1800,sleep_time);

            else
                horario(x)=random('Uniform',return_home+1800,sleep_time-1800);
            end
	
}