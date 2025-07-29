package controller

import (
	"fmt"
	"log"

	"simulation/internal/entities"
	"simulation/internal/entities/house"


	"math/rand/v2"
	"time"
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


func RunSimulation(size, day, toiletType, showerType int) {

	
	//Iniciando Simulação
	rng := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), 0))

	profile := defaultHouseProfile(toiletType, showerType)
	houses := make([]*entities.House, size)

	setHouses(profile, houses, size, rng)

	
	//Variaveis de dados
	lines := []SanitaryUsageLine{}
	populationData := newPopulationData(houses)
	dailyUsagesWindow := make(map[uint8]*usagesOverview)

	for i := uint8(0); i < uint8(day+2); i++ {
		dailyUsagesWindow[uint8(i)] = newUsagesOverview(uint8(i))
	}

	for i := uint8(0); i < uint8(day); i++ { // i = day
		for j := 0; j < size; j++ { // j = house
			if err := houses[j].GenerateLogs(i+1, rng); err != nil {
				log.Fatalf("Erro ao gerar logs da casa %d no dia %d: %v", j, i, err)
			}
			updateUsagesOverview(houses[j],dailyUsagesWindow, i+1)
			usageLines := ToSanitaryUsageLines(houses[j])
			lines = append(lines, usageLines...)
		}
	}

	
	//fmt.Println("Passou")

	populationData.viewPopulationData()
	fmt.Println()

	printUsagesOverview(dailyUsagesWindow)


/*
	agg := AggregateSanitaryUsage(lines)
	PrintUsageByDevicePerHour(agg)
	PrintTotalPerHour(agg)
	PrintTotalPerDevice(agg)
	PrintTotalUsage(agg)*/
	
}