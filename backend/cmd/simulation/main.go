package main

import (
	"flag"
	"fmt"
	"runtime"
	"simulation/internal/controller" // seu pacote controller
	"time"
)

func main() {
	// Definir as flags de entrada via terminal
	size := flag.Int("size", 10, "Número de casas a simular")
	day := flag.Int("day", 1, "Dia da simulação")
	toiletType := flag.Int("toiletType", 1, "Tipo de Toilet (1 a 4)")
	showerType := flag.Int("showerType", 1, "Tipo de Shower (1 ou 2)")

	// Parse das flags
	flag.Parse()

	// Exibir os parâmetros recebidos
	fmt.Printf("Executando simulação com %d casas, dia %d, toiletType %d, showerType %d\n", *size, *day, *toiletType, *showerType)

	start := time.Now()
	// Executar a simulação
	controller.RunSimulation(*size, *day, *toiletType, *showerType)

	// Calcular duração
	elapsed := time.Since(start)

	fmt.Println("Simulação concluída.")
	fmt.Printf("Tempo total: %s\n", elapsed)
	runtime.GC()
}
