package controller

import (
	"log"
)

// Helper para ignorar erro (não recomendado em produção sem cuidado)
func must[T any](val T, err error) T {
	if err != nil {
		log.Fatal(err) // Ou log.Fatal(err) se preferir
	}
	return val
}