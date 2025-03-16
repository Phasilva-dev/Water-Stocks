package devtest

import (
	"dists"
	"golang.org/x/exp/rand"
	"log"
	"time"
)

type Simulador struct{}

func (s *Simulador) SimularRotina(means []float64, stds []float64, tamanho int64) []float64 {
	// Verifica se os slices têm o tamanho esperado
	if len(means) < 4 || len(stds) < 4 {
		log.Fatal("Os slices 'means' e 'stds' devem ter pelo menos 4 elementos")
	}

	// Cria uma fonte de aleatoriedade
	src := rand.NewSource(uint64(time.Now().UnixNano()))

	// Cria as distribuições normais
	dist0, err := dists.NewNormalDist(means[0], stds[0])
	if err != nil {
		log.Fatal("Erro ao criar dist0:", err)
	}
	dist1, err := dists.NewNormalDist(means[1], stds[1])
	if err != nil {
		log.Fatal("Erro ao criar dist1:", err)
	}
	dist2, err := dists.NewNormalDist(means[2], stds[2])
	if err != nil {
		log.Fatal("Erro ao criar dist2:", err)
	}
	dist3, err := dists.NewNormalDist(means[3], stds[3])
	if err != nil {
		log.Fatal("Erro ao criar dist3:", err)
	}

	// Cria um slice vazio para armazenar as amostras
	amostras := make([]float64, 0)

	// Gera as amostras
	for i := int64(0); i < tamanho; i++ {
		amostras = append(amostras, dist0.Sample(src))
		amostras = append(amostras, dist1.Sample(src))
		amostras = append(amostras, dist2.Sample(src))
		amostras = append(amostras, dist3.Sample(src))
	}

	return amostras
}