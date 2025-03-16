package dists

import (
	"errors"
	"golang.org/x/exp/rand"
	
)

type Distribution interface {
	Sample(rng rand.Source) float64
	String() string
}

// CreateDistribution cria uma distribuição com base no tipo e parâmetros
func CreateDistribution(distType string, params ...float64) (Distribution, error) {
	switch distType {
	case "normal":
			if len(params) != 2 {
					return nil, errors.New("normal requires exactly 2 parameters: mean and stdDev")
			}
			return NewNormalDist(params[0], params[1])
	case "poisson":
			if len(params) != 1 {
					return nil, errors.New("poisson requires exactly 1 parameter: lambda")
			}
			return NewPoissonDist(params[0])
	case "uniform":
			if len(params) != 2 {
					return nil, errors.New("uniform requires exactly 2 parameters: min and max")
			}
			return NewUniformDist(params[0], params[1])
	default:
			return nil, errors.New("unknown distribution type")
	}
}