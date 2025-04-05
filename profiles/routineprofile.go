package profiles

import (
	"dists"
	"unique"
	"datastruct"
	"errors"

	"math/rand/v2"
)

// RoutineProfileDist contém uma lista de ProfileTupleDist
type RoutineProfileDist struct {
	symbol unique.Handle[string]
	events []dists.Distribution
	shift int32
	
}

func (p *RoutineProfileDist) Symbol() string {
	return p.symbol.Value()
}

/*Esse getter é fere o encapsulamento, porém, devido ao fato que esse código vai rodar milhares ou sentenas de vezes
permancerá assim, para não gerar um overhead desnecessario copiando o slice inteiro todas as vezes. */
func (p *RoutineProfileDist) Events() []dists.Distribution {
	return p.events
}

func (p *RoutineProfileDist) Shift() int32 {
	return p.shift
}

func NewRoutineProfileDist(symbol string, shift int32, events []dists.Distribution) (*RoutineProfileDist, error) {
	if symbol == "" {
		return nil, errors.New("símbolo não pode estar vazio")
	}
	if len(events) == 0 {
		return nil, errors.New("events não pode estar vazio")
	}
	if len(events)%2 != 0 {
		return nil, errors.New("número de elementos em events deve ser par")
	}
	if shift < 0 {
		return nil, errors.New("constante deve ser positiva")
	}

	// Criar uma cópia para garantir imutabilidade
	eventsCopy := make([]dists.Distribution, len(events))
	copy(eventsCopy, events)

	for _, dist := range eventsCopy {
		if dist == nil {
			return nil, errors.New("Nenhuma distribuição pode estar vazia")
		}
	}

	return &RoutineProfileDist{
		symbol: unique.Make(symbol),
		events: eventsCopy, // Usando a cópia
		shift:  shift,
	}, nil
}


func generateTime(dist dists.Distribution, rng *rand.Rand) int32 {
	var time float64 = dist.Sample(rng)
	truncatedTime := int32(time)
	return truncatedTime
}

func (p *RoutineProfileDist) enforceMinimunGap(entryTime, exitTime int32) (int32) {
	if exitTime - (entryTime + p.shift) < p.shift {
		return entryTime + p.shift
	}
	return exitTime
}

func (p *RoutineProfileDist) GenerateData(rng *rand.Rand) *datastruct.Routine {
	times := make([]int32, len(p.events)) //aloca diretamente o slice
	for i, dist := range p.events {
    times[i] = generateTime(dist, rng)
	}
	/*A logica desse For é: os eventos impares sempre vão ser o de sair de casa ou dormir
	estamos presupondo que ele sempre vai ter um tempo minimo para sair de casa novamente
	algo como se arrumar para ir para o trabalho após acordar, ou simplesmente não valeria a penar ele computar o
	retorno dele em casa para passar algo como 5 minutos!!!
	Impar = saida
	Par = Entrada
	entre uma saida e uma entrada ele não interage com a casa, inclusive a ultima saida é sempre ele dormindo
	e a primeira entrada é sempre ele acordando*/
	for i := 1; i < len(p.events); i = i+2 {
		times[i] = p.enforceMinimunGap(times[i-1],times[i])
	}
	return datastruct.NewRoutine(p.symbol.Value(),times)
}
