package simulation

import (
	"log"
	"simulation/internal/controller" // seu pacote controller
)

func RunSimulation(size, day, toiletType, showerType int, filename string, progressCallback func(currentDay, totalDays int)) error {

	log.Printf("Executando simulacao com %d casas, dia %d, toiletType %d, showerType %d\n filename %s:", size, day, toiletType, showerType, filename)

	err := controller.RunSimulation(size, day, toiletType, showerType, filename, progressCallback)
	if err != nil {
		return err // Apenas repasse o erro.
	}

	log.Println("Simulacao concluida.")
	
	return nil
}
