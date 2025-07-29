package controller


import (
	"fmt"
	"log"
	"sort"
	"simulation/internal/entities"
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

// ResidentsTypeCount retorna o número de tipos distintos de residentes registrados no mapa.
func (p *populationData) ResidentsTypeCount() uint32 {
	if p.residentsTypeCount == nil {
		return 0
	}
	return uint32(len(p.residentsTypeCount))
}


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





// Supondo que os tipos são os mesmos usados anteriormente
func printUsagesOverview(dayWindow map[uint8]*AccumulatorDay, sanitaryTypes []string) {
	for day, overview := range dayWindow {
		fmt.Printf("Usos no dia: %d\n", day)

		// Criar contador por dispositivo
		countPerDevice := make(map[string]int)

		for _, hour := range overview.accumulatorHour {
			for _, device := range sanitaryTypes {
				countPerDevice[device] += hour.IndividualDeviceCounter(device)
			}
		}

		// Imprimir os contadores
		for _, device := range sanitaryTypes {
			label := formatDeviceLabel(device)
			fmt.Printf("  %s: %d\n", label, countPerDevice[device])
		}

		fmt.Println()
	}
}

// Apenas para deixar os nomes mais legíveis
func formatDeviceLabel(key string) string {
	switch key {
	case "wash_bassin":
		return "Wash Bassin"
	case "wash_machine":
		return "Wash Machine"
	case "dish_washer":
		return "Dish Washer"
	default:
		return capitalize(key)
	}
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(s[0]-32) + s[1:]
}
