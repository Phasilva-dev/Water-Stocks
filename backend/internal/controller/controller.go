package controller

import (
	"fmt"
	"log"
	"simulation/internal/entities"
	"math/rand/v2"
	"time"
	"simulation/internal/accumulator"
)


func RunSimulation(size, day, toiletType, showerType int) {

	
	//Iniciando Simulação
	rng := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), 0))

	profile := defaultHouseProfile(toiletType, showerType)
	houses := make([]*entities.House, size)

	setHouses(profile, houses, size, rng)

	
	//Variaveis de dados
	populationData := accumulator.NewPopulationData(houses)
	dailyUsagesDataWindow := make(map[uint8]*accumulator.AccumulatorDay)

	for i := uint8(0); i < uint8(day+2); i++ {
		dailyUsagesDataWindow[uint8(i)] = accumulator.NewAccumulatorDay(uint8(i),accumulator.OrderedDeviceKeys())
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

	populationData.ViewPopulationData()
	fmt.Println()

	accumulator.PrintUsagesOverview(dailyUsagesDataWindow, accumulator.OrderedDeviceKeys())

	for k := uint8(1); k < uint8(day+2); k++ {
		dailyUsagesDataWindow[uint8(k)].RoundAccumulatorDayValues()
		fmt.Println("Consumo do dia ",k)
		dailyUsagesDataWindow[uint8(k)].PrintHourlyWaterConsumption()
	}

	err := dailyUsagesDataWindow[1].ExportToExcel("consumo_diario.xlsx")
		if err != nil {
	log.Fatalf("Erro exportando Excel: %v", err)
	}

	
}