package dists

import (
	"errors"
  "gonum.org/v1/gonum/stat/distuv"
  "golang.org/x/exp/rand"
	"fmt"
)

type PoissonDist struct {
	lambda float64
}

func (n *PoissonDist) Lambda() float64 {
	return n.lambda
}

func NewPoissonDist (lambda float64) (*PoissonDist, error) {

	if lambda <= 0 {
		return nil, errors.New("lambda must be greater than 0")
	}
	return &PoissonDist {
		lambda: lambda,
	}, nil
}

func (n *PoissonDist) Sample(src rand.Source) float64 {
	dist := distuv.Poisson{Lambda: n.lambda, Src: src}
	return dist.Rand()
}

func (n *PoissonDist) String() string {
	return fmt.Sprintf("PoissonDist{lambda: %.2f}", n.lambda)
}