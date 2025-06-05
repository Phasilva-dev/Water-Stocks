package controller

import (
	"simulation/internal/dists"
	"simulation/internal/entities/house/profile/count"
	"simulation/internal/entities/house/profile/demographics"
	"simulation/internal/entities/resident"
	"simulation/internal/entities/resident/profile/frequency"
	"simulation/internal/entities/resident/profile/habits"
	"simulation/internal/entities/resident/profile/routine"
	"simulation/internal/misc"
)

func defaultResidentProfiles()  map[uint32]*resident.ResidentProfile{

freqProfile := frequency.NewFrequencyProfileDay(map[string]*frequency.FrequencyProfile{
	"toilet":      must(frequency.NewFrequencyProfile(must(dists.CreateDistribution("poisson", 2.75)), 0)),
	"shower":      must(frequency.NewFrequencyProfile(must(dists.CreateDistribution("poisson", 1.08)), 0)),
	"washBassin":  must(frequency.NewFrequencyProfile(must(dists.CreateDistribution("poisson", 5.93)), 0)),
	"washMachine": must(frequency.NewFrequencyProfile(must(dists.CreateDistribution("poisson", 0.37)), 0)),
	"dishWasher":  must(frequency.NewFrequencyProfile(must(dists.CreateDistribution("poisson", 24.88)), 0)),
	"tanque":      must(frequency.NewFrequencyProfile(must(dists.CreateDistribution("poisson", 1.15)), 0)),
})

// Adulto Empregado Caso 1
adultDailyRoutine := must(routine.NewRoutineProfile([]dists.Distribution{
	must(dists.CreateDistribution("normal", 5.5*3600, 3600)),      // Acordar
	must(dists.CreateDistribution("normal", 7.5*3600, 1800)),      // Trabalhar
	must(dists.CreateDistribution("normal", 15.4*3600, 1.8*3600)), // Voltar pra casa
	must(dists.CreateDistribution("normal", 22*3600, 1.8*3600)),   // Dormir
}, 1800))



// Perfil diário
	adultDailyHabits := habits.NewResidentDayProfile(adultDailyRoutine, freqProfile)

	// Perfil semanal (lista de perfis diários)
	adultWeeklyHabits := must(habits.NewResidentWeeklyProfile([]*habits.ResidentDayProfile{
		adultDailyHabits, // Pode replicar ou customizar dias diferentes
	}))

	// Perfil completo do residente

	adultProfile := must(resident.NewResidentProfile(adultWeeklyHabits, 1))


	// Crianca Matutino Caso 1
	ChildrenMorningRoutine := must(routine.NewRoutineProfile([]dists.Distribution{
	must(dists.CreateDistribution("normal", 5.75*3600,3600)),      // Acordar
	must(dists.CreateDistribution("normal", 7*60*60,1800)),      // Trabalhar
	must(dists.CreateDistribution("normal", 13*3600, 1800)), // Voltar pra casa
	must(dists.CreateDistribution("normal", 21.5*3600, 1.8*3600)),   // Dormir
}, 1800))


	// Perfil diário
	ChildrenMorningHabits := habits.NewResidentDayProfile(ChildrenMorningRoutine, freqProfile)

	// Perfil semanal (lista de perfis diários)
	ChildrenMorningWeeklyHabits := must(habits.NewResidentWeeklyProfile([]*habits.ResidentDayProfile{
		ChildrenMorningHabits, // Pode replicar ou customizar dias diferentes
	}))

	// Perfil completo do residente

	ChildrenMorningProfile := must(resident.NewResidentProfile(ChildrenMorningWeeklyHabits, 2))



		// Idoso Caso 3
	AgedRoutine := must(routine.NewRoutineProfile([]dists.Distribution{
	must(dists.CreateDistribution("normal", 5.5*3600,3600)),      // Acordar
	must(dists.CreateDistribution("normal", 10*3600,3*3600)),      // Trabalhar
	must(dists.CreateDistribution("normal", 14*3600, 4*3600)), // Voltar pra casa
	must(dists.CreateDistribution("normal", 22*3600, 1800)),   // Dormir
}, 1800))


	// Perfil diário
	AgedHabits := habits.NewResidentDayProfile(AgedRoutine, freqProfile)

	// Perfil semanal (lista de perfis diários)
	AgedWeeklyHabits := must(habits.NewResidentWeeklyProfile([]*habits.ResidentDayProfile{
		AgedHabits, // Pode replicar ou customizar dias diferentes
	}))

	// Perfil completo do residente

	AgedProfile := must(resident.NewResidentProfile(AgedWeeklyHabits, 3))



	// Crianca Vespertino Caso 4
	ChildrenAfternoonRoutine := must(routine.NewRoutineProfile([]dists.Distribution{
	must(dists.CreateDistribution("normal", 8*3600,3600)),      // Acordar
	must(dists.CreateDistribution("normal", 12.5*3600,1800)),      // Trabalhar
	must(dists.CreateDistribution("normal", 18.5*3600, 1800)), // Voltar pra casa
	must(dists.CreateDistribution("normal", 24.5*3600, 3600)),   // Dormir
}, 1800))

	// Perfil diário
	ChildrenAfternoonHabits := habits.NewResidentDayProfile(ChildrenAfternoonRoutine, freqProfile)

	// Perfil semanal (lista de perfis diários)
	ChildrenAfternoonWeeklyHabits := must(habits.NewResidentWeeklyProfile([]*habits.ResidentDayProfile{
		ChildrenAfternoonHabits, // Pode replicar ou customizar dias diferentes
	}))

	// Perfil completo do residente

	ChildrenAfternoonProfile := must(resident.NewResidentProfile(ChildrenAfternoonWeeklyHabits, 4))


	// Adulto Desempregado Caso 5
	AdultUnemployedRoutine := must(routine.NewRoutineProfile([]dists.Distribution{
	must(dists.CreateDistribution("normal", 8*3600,3600)),      // Acordar
	must(dists.CreateDistribution("normal", 10*3600,10800)),      // Trabalhar
	must(dists.CreateDistribution("normal", 14*3600,4*3600)), // Voltar pra casa
	must(dists.CreateDistribution("normal", 24.5*3600, 3600)),   // Dormir
}, 1800)) 

	// Perfil diário
	AdultUnemployedHabits := habits.NewResidentDayProfile(AdultUnemployedRoutine, freqProfile)

	// Perfil semanal (lista de perfis diários)
	AdultUnemployedWeeklyHabits := must(habits.NewResidentWeeklyProfile([]*habits.ResidentDayProfile{
		AdultUnemployedHabits, // Pode replicar ou customizar dias diferentes
	}))

	// Perfil completo do residente

	AdultUnemployedProfile := must(resident.NewResidentProfile(AdultUnemployedWeeklyHabits, 5))


	// ResidentProfiles
	ResidentProfiles := map[uint32]*resident.ResidentProfile{
	1: adultProfile,
	2: ChildrenMorningProfile,
	3: AgedProfile,
	4: ChildrenAfternoonProfile,
	5: AdultUnemployedProfile,
}

	return ResidentProfiles

}

func defaultHouseProfile() {
	residentCountProfile := count.NewResidentCount(must(dists.CreateDistribution("gamma",4.09588,0.636582)))
	sanitaryCountProfile := count.NewSanitaryCount()
	residentAgeProfile := demographics.NewAge(must(dists.CreateDistribution("weibull",35.8311, 1.58364)))


	//residentOccupation
	under18Selector := []misc.Tuple[uint32, float64]{
		*misc.NewTuple(1, 80), // Adulto Empregado
		*misc.NewTuple(2, 15), // Estudante
		*misc.NewTuple(3, 5),  // Desempregado
	}

	adultSelector := []misc.Tuple[uint32, float64]{
		*misc.NewTuple(1, 90), // Adulto Empregado
		*misc.NewTuple(5, 10), // Adulto Desempregado
	}

	over64Selector := []misc.Tuple[uint32, float64]{
		*misc.NewTuple(3, 100), // Idoso
	}

}