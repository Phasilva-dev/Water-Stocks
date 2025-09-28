package simulation

import (
	"log"
	"simulation/internal/controller" // seu pacote controller
)

func RunSimulation(size, day, toiletType, showerType int, filename string) {

	// Exibir os parâmetros recebidos
	log.Printf("Executando simulacao com %d casas, dia %d, toiletType %d, showerType %d\n filename %s:", size, day, toiletType, showerType, filename)

	controller.RunSimulation(size, day, toiletType, showerType, filename)

	log.Println("Simulacao concluida.")
}
