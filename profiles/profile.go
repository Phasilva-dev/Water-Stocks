package profiles

import(
	"math/rand/v2"

)

type Profile interface {
	GenerateData(rng *rand.Rand, constante int32, tipo string)
}