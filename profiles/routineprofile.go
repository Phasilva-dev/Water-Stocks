package profiles

import (
	"dists"
	"unique"
	"datastruct"

	"golang.org/x/exp/rand"
)




// RoutineProfileDist contém uma lista de ProfileTupleDist
type RoutineProfileDist struct {
	symbol unique.Handle[string]
	events []dists.Distribution
}

func (p *RoutineProfileDist) Symbol() string {
	return p.symbol.Value()
}

/*Esse getter é fere o encapsulamento, porém, devido ao fato que esse código vai rodar milhares ou sentenas de vezes
permancerá assim, para não gerar um overhead desnecessario copiando o slice inteiro todas as vezes. */
func (p *RoutineProfileDist) Events() []dists.Distribution {
	return p.events
}

func NewRoutineProfileDist(symbol string, events ...dists.Distribution) *RoutineProfileDist {
	return &RoutineProfileDist{
			symbol: unique.Make(symbol),
			events: events,
	}
}

func generateTime (dist dists.Distribution ,rng rand.Source) int32 {
	var time float64 = dist.Sample(rng)
	truncatedTime := int32(time)
	return truncatedTime
}

func enforceMinimunGap(entryTime, exitTime int32, gap int32) int32 {
	if exitTime - (entryTime+gap) < gap {
		exitTime = entryTime + gap
	}
	return exitTime
}

func (p *RoutineProfileDist) GenerateData(rng rand.Source, constante int32, tipo string) *datastruct.Routine {
	times := make([]int32, 0, len(p.events)*2)
	//tuples := make([]dists.Distribution, 0, len(p.events))

	for i := 0; i < len(p.events); i++ {
		times = append(times, generateTime(p.events[i], rng))
	}
	for i := 1; i < len(p.events); i = i+2 {
		times[i] = enforceMinimunGap(times[i-1],times[i],constante)
	}
	return datastruct.NewRoutine(p.symbol.Value(),times)
}
