package controller

import (
	"fmt"
	"log"
	"math/rand/v2"
	"simulation/internal/analysis"
	"simulation/internal/entities"
	"time"
)


func RunSimulation(size, day, toiletType, showerType int, filename string) {

	
	// NOVO: Adiciona robustez. Se o nome do arquivo estiver vazio, usa um padrão.
	if filename == "" {
		log.Println("Aviso: Nome do arquivo não fornecido, usando 'default_simulation'")
		filename = "default_simulation"
	}

	// NOVO: Constrói os nomes dos arquivos de saída dinamicamente.
	//pulseCsvFilename := fmt.Sprintf("%s_pulses.csv", filename)
	analysisCsvFilename := fmt.Sprintf("%s_analysis.csv", filename)

	log.Printf("Iniciando simulação. O arquivo de saída será: %s ", analysisCsvFilename)
	
	//Iniciando Simulação
	rng := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), 0))

	profile := defaultHouseProfile(toiletType, showerType)
	houses := make([]*entities.House, size)

	setHouses(profile, houses, size, rng)

	
	//Variaveis de dados
	populationData := analysis.NewPopulationData(houses)
	pulseData := analysis.NewPulseHouse(0,analysis.OrderedDeviceKeys())
	simulationAnalysis := analysis.NewSimulationAnalysis(pulseData,populationData,day)

	/*
	dailyUsagesDataWindow := make(map[uint8]*analysis.AccumulatorDay)
	dailyPulseDataWindow := make(map[uint8]*analysis.PulseHouse)


	for i := uint8(0); i < uint8(day+2); i++ {
		dailyUsagesDataWindow[uint8(i)] = analysis.NewAccumulatorDay(uint8(i),analysis.OrderedDeviceKeys())
		dailyPulseDataWindow[uint8(i)] = analysis.NewPulseHouse(uint8(i),analysis.OrderedDeviceKeys())
	}*/
	
	for i := uint8(0); i < uint8(day); i++ { // i = day
		for j := 0; j < size; j++ { // j = house
			if err := houses[j].GenerateLogs(i+1, rng); err != nil {
				log.Fatalf("Erro ao gerar logs da casa %d no dia %d: %v", j, i, err)
			}
			pulseData.UpdatePulse(houses[j])
			//dailyUsagesDataWindow[1].UpdateAccumulator(i+1,houses[j],dailyUsagesDataWindow)
			//dailyPulseDataWindow[1].UpdatePulseWithWindow(i+1, houses[j],dailyPulseDataWindow)
		}
	}

	var err error
	populationData.ViewPopulationData()
	pulseData.PrintUsageStatistics()

	/*
	analysis.PrintUsagesOverview(dailyUsagesDataWindow, analysis.OrderedDeviceKeys())
	for k := uint8(1); k < uint8(day+2); k++ {
		dailyUsagesDataWindow[uint8(k)].RoundAccumulatorDayValues()
		fmt.Println("Consumo do dia ",k)
		dailyUsagesDataWindow[uint8(k)].PrintDailyTotals()
		dailyPulseDataWindow[k].PrintUsageStatistics()
		dailyUsagesDataWindow[uint8(k)].PrintHourlyWaterConsumption()
	}*/


	/*
	err = dailyUsagesDataWindow[1].ExportToExcel("consumo_diario.xlsx")
		if err != nil {
		log.Fatalf("Erro exportando Excel: %v", err)
	}*/

	/*
	err = pulseData.ExportPulsesToCSV(pulseCsvFilename)
		if err != nil {
		log.Fatalf("Erro exportando csv: %v", err)
	}*/
	
	err = simulationAnalysis.ExportAllDataToCSV(analysisCsvFilename)
	if err != nil {
		log.Fatalf("Erro exportando csv: %v", err)
	}

	
}