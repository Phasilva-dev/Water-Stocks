package controller

import (
	"log"
	"fmt"
)

// Helper para ignorar erro (não recomendado em produção sem cuidado)
func must[T any](val T, err error) T {
	if err != nil {
		log.Fatal(err) // Ou log.Fatal(err) se preferir
	}
	return val
}


func PrintLogLines(lines []string) {
	
	fmt.Println("Dia | HouseID | ResidentOccupationID | Idade | DispositivoSanitario | TipoSanitario | Inicio | Fim | Vazao")
    for _, line := range lines {
        fmt.Println(line)
    }
}