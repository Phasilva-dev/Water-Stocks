package controller

import (
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"path/filepath"
	"simulation/internal/analysis"
	"simulation/internal/entities"
	"time"
)


func RunSimulation(size, day, toiletType, showerType int, filename string, progressCallback func(currentDay, totalDays int)) error {

	if filename == "" {
		return fmt.Errorf("nome do arquivo não pode estar vazio")
	}

	const outputDir = "simulations_output"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Printf("Erro ao criar o diretório de saída '%s': %v", outputDir, err)
		return fmt.Errorf("falha ao criar o diretório de saída: %w", err)
	}

	analysisCsvFilename := filepath.Join(outputDir, fmt.Sprintf("%s_analysis.csv", filename))
	log.Printf("Iniciando simulação. O arquivo de saída será: %s ", analysisCsvFilename)

	rng := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), 0))
	profile := defaultHouseProfile(toiletType, showerType)
	houses := make([]*entities.House, size)
	setHouses(profile, houses, size, rng)
	
	populationData := analysis.NewPopulationData(houses)
	pulseData := analysis.NewPulseHouse(0, analysis.OrderedDeviceKeys())
	simulationAnalysis := analysis.NewSimulationAnalysis(pulseData, populationData, day)

	for i := uint8(0); i < uint8(day); i++ { // i = day
		for j := 0; j < size; j++ { // j = house
			if err := houses[j].GenerateLogs(i+1, rng); err != nil {
				log.Printf("Erro ao gerar logs da casa %d no dia %d: %v", j, i, err)
				return fmt.Errorf("erro ao gerar logs para casa %d no dia %d: %w", j, i, err)
			}
			pulseData.UpdatePulse(houses[j])
		}
		
		// MODIFICAÇÃO 2: Chame o callback após cada dia simulado.
		// Verificamos se ele não é nulo antes de chamar.
		if progressCallback != nil {
			progressCallback(int(i)+1, day)
		}
	}

	var err error
	populationData.ViewPopulationData()
	pulseData.PrintUsageStatistics()

	err = simulationAnalysis.ExportAllDataToCSV(analysisCsvFilename)
	if err != nil {
		log.Printf("Erro exportando csv: %v", err)
		return fmt.Errorf("erro ao exportar dados para CSV: %w", err)
	}
	
	return nil
}