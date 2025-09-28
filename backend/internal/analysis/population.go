package analysis


import (
	"fmt"
	"log"
	"sort"
	"simulation/internal/entities"
)

type populationData struct {

	residentsTypeCount map[string]uint32
}

func NewPopulationData(houses []*entities.House) *populationData{

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

func (p *populationData) ViewPopulationData() {
	fmt.Println("O total de Residents é: ", p.residentsTotalCount())
	p.printAllResidentsCounts()
}

