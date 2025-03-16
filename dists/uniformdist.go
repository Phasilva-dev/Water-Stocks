package dists

import (
	"errors"
  "gonum.org/v1/gonum/stat/distuv"
  "golang.org/x/exp/rand"
	"fmt"

)

type UniformDist struct {
	min float64
	max float64
}

func (n *UniformDist) Min() float64 {
	return n.min
}

func (n *UniformDist) Max() float64 {
	return n.max
}

func NewUniformDist (min, max float64) (*UniformDist, error) {

	if min >= max {
		return nil, errors.New("min must be less than max")
	}
	return &UniformDist{
			min:  min,
			max:  max,
	}, nil
}

func (n *UniformDist) Sample(src rand.Source) float64 {
	dist := distuv.Uniform{Min: n.min, Max: n.max, Src: src}
	return dist.Rand()
}

func (n *UniformDist) String() string {
	return fmt.Sprintf("UniformDist{min: %.2f, max: %.2f}", n.min, n.max)
}