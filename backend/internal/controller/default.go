package controller

import (
	
	"simulation/internal/dists"
	"simulation/internal/entities/house"
	"simulation/internal/entities/house/profile/count"
	"simulation/internal/entities/house/profile/demographics"
	"simulation/internal/entities/house/profile/sanitarydevice"
	"simulation/internal/entities/resident"
	"simulation/internal/entities/resident/profile/frequency"
	"simulation/internal/entities/resident/profile/habits"
	"simulation/internal/entities/resident/profile/routine"
	"simulation/internal/misc"
	"simulation/internal/entities"

		"log"
	"math/rand/v2"

)

func setHouses(profile *house.HouseProfile, houses []*entities.House, size int, rng *rand.Rand) {

	for i := 0; i < size; i++ {

		houses[i] = entities.NewHouse(1,profile)
		if err := houses[i].GenerateHouseData(rng); err != nil {
			log.Fatalf("Erro ao criar a casa %d : %v", i, err)
		} 

	}
	//fmt.Println("Casas criadas ")

}

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
}, 1800,0.999999))



// Perfil diário
	adultDailyHabits := habits.NewResidentDayProfile(adultDailyRoutine, freqProfile)

	// Perfil semanal (lista de perfis diários)
	adultWeeklyHabits := must(habits.NewResidentWeeklyProfile([]*habits.ResidentDayProfile{
		adultDailyHabits, // Pode replicar ou customizar dias diferentes
	}))

	// Perfil completo do residente

	adultProfile := must(resident.NewResidentProfile(adultWeeklyHabits, 1))


	// Crianca Matutino Caso 2
	ChildrenMorningRoutine := must(routine.NewRoutineProfile([]dists.Distribution{
	must(dists.CreateDistribution("normal", 5.75*3600, 3600)),      // Acordar
	must(dists.CreateDistribution("normal", 7*60*60, 1800)),      // Trabalhar
	must(dists.CreateDistribution("normal", 13*3600, 1800)), // Voltar pra casa
	must(dists.CreateDistribution("normal", 20.25*3600, 3600)),   // Dormir
}, 1800,0.999999))


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
	must(dists.CreateDistribution("normal", 5.5*3600, 3600)),      // Acordar
	must(dists.CreateDistribution("normal", 10*3600, 3*3600)),      // Trabalhar
	must(dists.CreateDistribution("normal", 14*3600, 4*3600)), // Voltar pra casa
	must(dists.CreateDistribution("normal", 22*3600, 1800)),   // Dormir
}, 1800,0.9772))


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
	must(dists.CreateDistribution("normal", 8*3600, 3600)),      // Acordar
	must(dists.CreateDistribution("normal", 12.5*3600, 1800)),      // Trabalhar
	must(dists.CreateDistribution("normal", 18.5*3600, 1800)), // Voltar pra casa
	must(dists.CreateDistribution("normal", 24.5*3600, 3600)),   // Dormir
}, 1800,0.9772))

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
}, 1800,0.999999)) 

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

func defaultHouseProfile(toiletType, showerType int) *house.HouseProfile {
	residentCountProfile := count.NewResidentCount(must(dists.CreateDistribution("gamma",4.09588,0.636582)))
	sanitaryCountProfile := count.NewSanitaryCount()
	residentAgeProfile := demographics.NewAge(must(dists.CreateDistribution("weibull", 35.8311, 1.58364)))



	//residentOccupation
	// Definir os seletores de ocupação por faixa etária
	childrenUnder15Selector := must(misc.NewPercentSelector([]misc.Tuple[uint32, float64]{
		*misc.NewTuple(uint32(2), 50.0), // Criança Matutino
		*misc.NewTuple(uint32(4), 50.0), // Criança Vespertino
	}))

	children15to17Selector := must(misc.NewPercentSelector([]misc.Tuple[uint32, float64]{
		*misc.NewTuple(uint32(2), 45.0), // Criança Matutino
		*misc.NewTuple(uint32(4), 45.0), // Criança Vespertino
		*misc.NewTuple(uint32(5), 10.0), // Desempregado
	}))

	adultSelector := must(misc.NewPercentSelector([]misc.Tuple[uint32, float64]{
		*misc.NewTuple(uint32(1), 90.0), // Adulto Empregado
		*misc.NewTuple(uint32(5), 10.0), // Desempregado
	}))

	over64Selector := must(misc.NewPercentSelector([]misc.Tuple[uint32, float64]{
		*misc.NewTuple(uint32(3), 100.0), // Idoso
	}))

	childUnder15 := must(demographics.NewAgeRangeSelector(0, 14, childrenUnder15Selector))
	child15to17 := must(demographics.NewAgeRangeSelector(15, 17, children15to17Selector))
	adult := must(demographics.NewAgeRangeSelector(18, 64, adultSelector))
	elder := must(demographics.NewAgeRangeSelector(65, 130, over64Selector)) // 130 como limite superior

	// Criar o Occupation
	occupation := must(demographics.NewOccupation([]*demographics.AgeRangeSelector{
		childUnder15,
		child15to17,
		adult,
		elder,
	}))

	devices := defaultHouseSanitaryDevice(toiletType,showerType)

	residentProfiles := defaultResidentProfiles()

	profile := must(house.NewHouseProfile(
		1,
		residentCountProfile,
		residentAgeProfile,
		occupation,
		sanitaryCountProfile,
		residentProfiles,
		devices))

	return profile

}

func defaultHouseSanitaryDevice(toiletType, showerType int) map[string]sanitarydevice.SanitaryDevice {
	devices := make(map[string]sanitarydevice.SanitaryDevice)

	// Criar os toilets
	toilet1 := must(sanitarydevice.NewToilet(0.4, 5, 1))
	toilet2 := must(sanitarydevice.NewToilet(0.042, 1.8*60, 2))
	toilet3 := must(sanitarydevice.NewToilet(0.25, 60, 3))
	toilet4 := must(sanitarydevice.NewToilet(0.042, 1.2*60, 4))

	// Selecionar o toilet com base no tipo
	var selectedToilet sanitarydevice.SanitaryDevice
	switch toiletType {
	case 1:
		selectedToilet = toilet1
	case 2:
		selectedToilet = toilet2
	case 3:
		selectedToilet = toilet3
	case 4:
		selectedToilet = toilet4
	default:
		selectedToilet = toilet1 // fallback
	}
	devices["toilet"] = selectedToilet

	// Criar os showers
	shower1 := must(sanitarydevice.NewShower(
		must(dists.CreateDistribution("triangle", 3.0 / 60, 4.0 / 60, 5.0 / 60)),
		must(dists.CreateDistribution("triangle", 2 * 60, 3.5 * 60, 5 * 60)),
		1,
	))

	shower2 := must(sanitarydevice.NewShower(
		must(dists.CreateDistribution("lognormal", -2.4205, 0.2014)),
		must(dists.CreateDistribution("gamma", 6.5216 * 60, 0.7668 * 60)),
		2,
	))

	// Selecionar o shower com base no tipo
	var selectedShower sanitarydevice.SanitaryDevice
	switch showerType {
	case 1:
		selectedShower = shower1
	case 2:
		selectedShower = shower2
	default:
		selectedShower = shower1 // fallback
	}
	devices["shower"] = selectedShower

	// Criar os outros dispositivos (sem múltiplas opções)
	devices["wash_bassin"] = must(sanitarydevice.NewWashBassin(
		must(dists.CreateDistribution("lognormal", -2.6677, 0.3275)), 
		must(dists.CreateDistribution("lognormal", 3.3551, 0.8449)),
		1,
	))

	devices["wash_machine"] = must(sanitarydevice.NewWashMachine(
		0.1,
		4 * 6 * 60,
		1,
	))

	devices["dish_washer"] = must(sanitarydevice.NewDishWasher(
		must(dists.CreateDistribution("weibull", 0.0569, 1.5871)),
		must(dists.CreateDistribution("lognormal", 3.1763, 0.785)),
		1,
	))

	devices["tanque"] = must(sanitarydevice.NewTanque(
		must(dists.CreateDistribution("lognormal", -2.3485, 0.3279)),
		must(dists.CreateDistribution("lognormal", 3.2905, 0.8918)),
		1,
	))

	return devices
}