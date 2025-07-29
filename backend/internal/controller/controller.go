package controller

import (
	"fmt"
	"log"

	"simulation/internal/entities"
	"simulation/internal/entities/house"


	"math/rand/v2"
	"time"
)

var devices = []string{
	"toilet",
	"shower",
	"wash_bassin",
	"wash_machine",
	"dish_washer",
	"tanque",
}



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
	populationData := newPopulationData(houses)
	dailyUsagesDataWindow := make(map[uint8]*AccumulatorDay)

	for i := uint8(0); i < uint8(day+2); i++ {
		dailyUsagesDataWindow[uint8(i)] = NewAccumulatorDay(uint8(i),devices)
	}

	for i := uint8(0); i < uint8(day); i++ { // i = day
		for j := 0; j < size; j++ { // j = house
			if err := houses[j].GenerateLogs(i+1, rng); err != nil {
				log.Fatalf("Erro ao gerar logs da casa %d no dia %d: %v", j, i, err)
			}
			dailyUsagesDataWindow[1].UpdateAccumulator(i+1,houses[j],dailyUsagesDataWindow)
		}
	}

	
	fmt.Println("Passou")

	populationData.viewPopulationData()
	fmt.Println()

	printUsagesOverview(dailyUsagesDataWindow, devices)

	
}