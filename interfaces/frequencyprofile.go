package interfaces

/*import (
	"math/rand/v2"
	"dists"
)

type FrequencyData interface {
	// Métodos para uso individual
	FreqToilet() uint8
	FreqShower() uint8
	FreqWashBassin() uint8
	
	// Métodos para uso compartilhado
	FreqWashMachine() uint8
	FreqDishWasher() uint8
	FreqTanque() uint8
}

type FrequencyProfile interface {

	Shift() uint8
	
	StatDist() dists.Distribution
	
	GenerateData(rng *rand.Rand) uint8
	
	// NewFrequencyProfile (opcional - normalmente funções de factory não são parte da interface)
	// NewFrequencyProfile(shift uint8, dist dists.Distribution) (*FrequencyProfile, error)
}

type FrequencyProfileDay interface {
	// Acessores para os perfis individuais
	ToiletProfile() FrequencyProfile
	ShowerProfile() FrequencyProfile
	WashBassinProfile() FrequencyProfile
	WashMachineProfile() FrequencyProfile
	DishWasherProfile() FrequencyProfile
	TanqueProfile() FrequencyProfile

	// Gera todos os dados de frequência para um dia
	GenerateDayData(rng *rand.Rand) FrequencyData
}*/