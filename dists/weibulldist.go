package dists

import (
	"errors"
	"fmt"
	"gonum.org/v1/gonum/stat/distuv"
	"math/rand/v2"
)

type WeibullDist struct {
	shape float64 // k
	scale float64 // lambda
}

func NewWeibullDist(shape, scale float64) (*WeibullDist, error) {
	if shape <= 0 || scale <= 0 {
		return nil, errors.New("shape and scale must be greater than 0")
	}
	return &WeibullDist{
		shape: shape,
		scale: scale,
	}, nil
}

func (w *WeibullDist) Shape() float64 {
	return w.shape
}

func (w *WeibullDist) Scale() float64 {
	return w.scale
}

func (w *WeibullDist) Sample(rng *rand.Rand) float64 {
	dist := distuv.Weibull{K: w.shape, Lambda: w.scale, Src: rng}
	return dist.Rand()
}

func (w *WeibullDist) String() string {
	return fmt.Sprintf("WeibullDist{shape: %.2f, scale: %.2f}", w.shape, w.scale)
}
