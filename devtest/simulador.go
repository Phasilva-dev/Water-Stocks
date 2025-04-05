package devtest

import (
	"dists"
	"math/rand/v2"
	"log"
	//"time"
)



func SimularRotina(means []float64, stds []float64, tamanho int64) []float64 {
	// Validação
	if len(means) < 4 || len(stds) < 4 {
		log.Println("Erro: Means e stds devem ter 4 elementos")
		return nil
	}

	// Fonte de aleatoriedade
	src1 := rand.NewPCG(42,1)
	src := rand.New(src1)

	// Cria distribuições
	dist0, err := dists.NewNormalDist(means[0], stds[0])
	if err != nil {
		log.Println("Erro ao criar dist0:", err)
		return nil
	}
	dist1, err := dists.NewNormalDist(means[1], stds[1])
	if err != nil {
		log.Println("Erro ao criar dist1:", err)
		return nil
	}
	dist2, err := dists.NewNormalDist(means[2], stds[2])
	if err != nil {
		log.Println("Erro ao criar dist2:", err)
		return nil
	}
	dist3, err := dists.NewNormalDist(means[3], stds[3])
	if err != nil {
		log.Println("Erro ao criar dist3:", err)
		return nil
	}

	// Gera amostras
	amostras := make([]float64, 0)
	for i := int64(0); i < tamanho; i++ {
		amostras = append(amostras, dist0.Sample(src))
		amostras = append(amostras, dist1.Sample(src))
		amostras = append(amostras, dist2.Sample(src))
		amostras = append(amostras, dist3.Sample(src))
	}

	log.Println("Simulação concluída com sucesso!")
	return amostras
}
