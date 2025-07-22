package controller

import (
	"fmt"
	"log"
	"simulation/internal/entities"
	"sort"
	logData "simulation/internal/log"
)

type populationData struct {

	residentsTypeCount map[string]uint32
}

func newPopulationData(houses []*entities.House) *populationData{

	pData := &populationData{
		residentsTypeCount: make(map[string]uint32),
	}

	for i := 0; i < len(houses); i++ {
		residents := houses[i].Residents()
		for j := 0; j < len(residents); j++ {
			occupationIDVal := residents[j].OccupationID()
			occupationIDString := fmt.Sprintf("%d", occupationIDVal) // %d formata um inteiro como decimal

			pData.residentsTypeCount[occupationIDString]++
		}

	}

	return pData

}

// ResidentsTotalCount retorna a soma de todos os residentes de todos os tipos.
func (p *populationData) residentsTotalCount() uint32 {
	var total uint32 // Inicializa o acumulador
	if p.residentsTypeCount == nil {
		return 0 // Se o mapa for nil, não há residentes
	}
	for _, count := range p.residentsTypeCount {
		total += count // Adiciona cada contagem ao total
	}
	return total
}

/*// ResidentsTypeCount retorna o número de tipos distintos de residentes registrados no mapa.
func (p *populationData) ResidentsTypeCount() uint32 {
	if p.residentsTypeCount == nil {
		return 0
	}
	return uint32(len(p.residentsTypeCount))
}*/

// GetAllResidentCounts (Função Original) retorna um slice com todas as contagens de residentes (os valores do map).
func (p *populationData) getAllResidentCounts() []uint32 {
	if p.residentsTypeCount == nil {
		return []uint32{}
	}

	counts := make([]uint32, 0, len(p.residentsTypeCount))
	for _, count := range p.residentsTypeCount {
		counts = append(counts, count)
	}
	return counts
}

// GetAllResidentCountsFormatted printa todos os tipos de residentes e a sua quantidade
// Os resultados são ordenados alfabeticamente pelo nome do tipo.
func (p *populationData) printAllResidentsCounts()  { //[]string
	if p.residentsTypeCount == nil {
		log.Fatal("Incapaz de contabilizar residents")
	}

	// 1. Coleta as chaves (nomes dos tipos) do map.
	// Isso é necessário porque a iteração direta no map não garante ordem.
	keys := make([]string, 0, len(p.residentsTypeCount))
	for key := range p.residentsTypeCount {
		keys = append(keys, key)
	}
	// 2. Ordena as chaves alfabeticamente.
	sort.Strings(keys)

	// 4. Itera sobre as chaves ordenadas para formatar as strings.
	for _, key := range keys {
		count := p.residentsTypeCount[key] // Pega a contagem correspondente à chave
		fmt.Println("Tipo: ", key," Quantidade: ", count)
	}
}

func (p *populationData) viewPopulationData() {
	fmt.Println("O total de Residents é: ", p.residentsTotalCount())
	p.printAllResidentsCounts()
}


type usagesOverview struct {
	day uint8
	usagesCount map[string]uint32
}

func newUsagesOverview(day uint8) *usagesOverview {
	

	// 1. Inicializa o mapa com 'make'
	usagesCount := make(map[string]uint32)

	// 2. Adiciona as 6 chaves predefinidas e inicializa seus valores com 0
	usagesCount["toilet"] = 0
	usagesCount["shower"] = 0
	usagesCount["wash_bassin"] = 0
	usagesCount["wash_machine"] = 0
	usagesCount["dish_washer"] = 0
	usagesCount["tanque"] = 0

	return &usagesOverview{
		usagesCount: usagesCount,
		day: day,
	}
}

func (u *usagesOverview) updateCount(key string, value uint32) {
	u.usagesCount[key] += value
}

func updateUsagesOverview (house *entities.House, dayWindow map[uint8]*usagesOverview, day uint8) {
	residentLogs := house.ResidentLogs()
	for i := 0; i < len(residentLogs); i++ {
		toiletLogs, ok := residentLogs[i].SanitaryLogs().ToiletLog().UsageLogs()
		if ok {
			countUsages(toiletLogs, dayWindow, day, "toilet")
		}

		showerLogs, ok := residentLogs[i].SanitaryLogs().ShowerLog().UsageLogs()
		if ok {
			countUsages(showerLogs, dayWindow, day, "shower")
		}

		washBassinLogs, ok := residentLogs[i].SanitaryLogs().WashBassinLog().UsageLogs()
		if ok {
			countUsages(washBassinLogs, dayWindow, day, "wash_bassin")
		}

		washMachineLogs, ok := residentLogs[i].SanitaryLogs().WashMachineLog().UsageLogs()
		if ok {
			countUsages(washMachineLogs, dayWindow, day, "wash_machine")
		}

		dishWasherLogs, ok := residentLogs[i].SanitaryLogs().DishWasherLog().UsageLogs()
		if ok {
			countUsages(dishWasherLogs, dayWindow, day, "dish_washer")
		}

		tanqueLogs, ok := residentLogs[i].SanitaryLogs().TanqueLog().UsageLogs()
		if ok {
			countUsages(tanqueLogs, dayWindow, day, "tanque")
		}
	}
	

}

func countUsages (usages []*logData.Usage, dayWindow map[uint8]*usagesOverview, day uint8, key string) {
	for i := 0; i < len(usages); i++ {
		if usages[i].StartUsage() >= 0 && usages[i].StartUsage() < 86400 {
			dayWindow[day].updateCount(key, 1)
		} else if usages[i].StartUsage() >= 86400 {
			dayWindow[day+1].updateCount(key, 1)
		} else {
			dayWindow[day-1].updateCount(key, 1)
		}
	}
}

func printUsagesOverview(dayWindow map[uint8]*usagesOverview) {
	for _, overview := range dayWindow {
		fmt.Printf("Usos no dia: %d\n", overview.day)
		fmt.Printf("  Toilet: %d\n", overview.usagesCount["toilet"])
		fmt.Printf("  Shower: %d\n", overview.usagesCount["shower"])
		fmt.Printf("  Wash Bassin: %d\n", overview.usagesCount["wash_bassin"])
		fmt.Printf("  Wash Machine: %d\n", overview.usagesCount["wash_machine"])
		fmt.Printf("  Dish Washer: %d\n", overview.usagesCount["dish_washer"])
		fmt.Printf("  Tanque: %d\n", overview.usagesCount["tanque"])
		fmt.Println()
	}
}
