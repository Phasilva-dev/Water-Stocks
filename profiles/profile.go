package profiles

import(
	"golang.org/x/exp/rand"

)

type Profile interface {
	GenerateData(rng rand.Source, constante int32, tipo string)
}