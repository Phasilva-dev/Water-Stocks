package analysis

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// SimulationAnalysis é uma estrutura que encapsula todos os dados
// e o contexto de uma simulação para facilitar a análise e a exportação.
type SimulationAnalysis struct {
	consumptionData     *PulseHouse
	demographicsData        *populationData
	totalSimulationDays int
}

func (sa *SimulationAnalysis) ConsumptionData() *PulseHouse {
	return sa.consumptionData
}

func (sa *SimulationAnalysis) DemographicsData() *populationData {
	return sa.demographicsData
}

func (sa *SimulationAnalysis) TotalSimulationDays() int {
	return sa.totalSimulationDays
}

// NewSimulationAnalysis cria uma nova instância de SimulationAnalysis.
func NewSimulationAnalysis(consumption *PulseHouse, demographics *populationData, totalDays int) *SimulationAnalysis {
	return &SimulationAnalysis{
		consumptionData:     consumption,
		demographicsData:        demographics,
		totalSimulationDays: totalDays,
	}
}

// ExportAllDataToCSV gera um único arquivo CSV contendo todos os dados brutos e
// metadados necessários para realizar análises completas no frontend.
// O arquivo é estruturado em seções para facilitar o parsing.
func (sa *SimulationAnalysis) ExportAllDataToCSV(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo CSV de dados brutos: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// --- SEÇÃO 1: METADATA ---
	// Escreve os parâmetros da simulação.
	writer.Write([]string{"[METADATA]"})
	writer.Write([]string{"Parameter", "Value"})
	writer.Write([]string{"Total Population", strconv.Itoa(int(sa.demographicsData.residentsTotalCount()))})
	writer.Write([]string{"Total Simulation Days", strconv.Itoa(sa.totalSimulationDays)})
	writer.Write([]string{""}) // Linha em branco para separação

	// --- SEÇÃO 2: DEVICE_SUMMARY ---
	// Escreve os totais agregados por aparelho. Isso é crucial para a eficiência,
	// pois o frontend não precisará somar os 86.400 segundos para obter esses totais.
	writer.Write([]string{"[DEVICE_SUMMARY]"})
	writer.Write([]string{"Device", "TotalLitersConsumed", "TotalUses"})
	deviceKeys := OrderedDeviceKeys()
	for _, deviceName := range deviceKeys {
		if stats, ok := sa.consumptionData.usageStats[deviceName]; ok {
			writer.Write([]string{
				deviceName,
				fmt.Sprintf("%.2f", stats.TotalLitersConsumed),
				strconv.Itoa(stats.UsageCount),
			})
		}
	}
	writer.Write([]string{""}) // Linha em branco para separação

	// --- SEÇÃO 3: PULSE_DATA ---
	// Escreve os dados de vazão segundo a segundo.
	writer.Write([]string{"[PULSE_DATA]"})
	pulseHeader := []string{"horario_segundos"}
	pulseHeader = append(pulseHeader, deviceKeys...)
	pulseHeader = append(pulseHeader, "total")
	writer.Write(pulseHeader)

	// Percorre todos os segundos do "dia médio"
	for sec := 0; sec < 86400; sec++ {
		row := make([]string, len(deviceKeys)+2)
		row[0] = strconv.Itoa(sec)

		var totalFlowPerSecond float64
		for i, deviceName := range deviceKeys {
			var val float64
			if device, ok := sa.consumptionData.pulsesDevice[deviceName]; ok {
				val = device.pulses[sec]
			}
			row[i+1] = fmt.Sprintf("%.2f", val)
			totalFlowPerSecond += val
		}
		row[len(deviceKeys)+1] = fmt.Sprintf("%.2f", totalFlowPerSecond)
		writer.Write(row)
	}

	return writer.Error()
}