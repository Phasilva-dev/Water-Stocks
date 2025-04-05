package dists

import (
	"errors"
  "gonum.org/v1/gonum/stat/distuv"
	"fmt"
	"math/rand/v2"

)

type NormalDist struct {
	mean float64
	stdDev float64
}

func (n *NormalDist) Mean() float64 { 
	return n.mean
}

func (n *NormalDist) StdDev() float64 {
	return n.stdDev
}

func NewNormalDist (mean, stdDev float64) (*NormalDist, error) {

	if stdDev < 0 {
		return nil, errors.New("stdDev must be greater than 0")
	}
	return &NormalDist{
		mean:   mean,
    stdDev: stdDev,
	}, nil

}

func (n *NormalDist) Sample(rng *rand.Rand) float64 {
	dist := distuv.Normal{Mu: n.mean, Sigma: n.stdDev, Src: rng}
	return dist.Rand()
}

func (n *NormalDist) String() string {
	return fmt.Sprintf("NormalDist{mean: %.2f, stdDev: %.2f}", n.mean, n.stdDev)
}