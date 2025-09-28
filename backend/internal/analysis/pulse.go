package analysis

import (
	"encoding/csv"
	"fmt"
	"os"
	"simulation/internal/entities"
	logData "simulation/internal/log"
	"strconv"
)

type PulseDevice struct {
	device string
	pulses []float64
}

func NewPulseDevice(DeviceName string) *PulseDevice {
	pulses := make([]float64, 86400) // 86400 segundos = 24h * 60min * 60s
	return &PulseDevice{
		device: DeviceName,
		pulses: pulses,
	}
}

type PulseHouse struct {
	day          uint8
	pulsesDevice map[string]*PulseDevice
	usageStats   map[string]*DeviceUsageStats
}

type DeviceUsageStats struct {
	UsageCount          int
	TotalDurationSeconds int64
	TotalFlowRateSum    float64
	TotalLitersConsumed float64
}

func NewPulseHouse(day uint8, deviceNames []string) *PulseHouse {
	pulses := make(map[string]*PulseDevice)
	stats := make(map[string]*DeviceUsageStats)

	for _, name := range deviceNames {
		pulses[name] = NewPulseDevice(name)
		stats[name] = &DeviceUsageStats{}
	}

	return &PulseHouse{
		day:          day,
		pulsesDevice: pulses,
		usageStats:   stats,
	}
}

func (p *PulseHouse) GetIndexAndDay(second int32, day uint8) (int, uint8) {
	switch {
	case second >= 0 && second < 86400:
		return int(second), day
	case second >= 86400:
		return int(second) - 86400, day + 1
	default: // second < 0
		return int(second) + 86400, day - 1
	}
}

// NOVO MÉTODO PRIVADO: O "bloco de construção" que processa um único uso.
// Ele achata o tempo para um único período de 24h usando o operador de módulo.
func (p *PulseHouse) processSingleUsage(sanitaryType string, usage *logData.Usage) {
	if stats, exists := p.usageStats[sanitaryType]; exists {
		start := usage.StartUsage()
		end := usage.EndUsage()
		duration := end - start
		flowRate := usage.FlowRate()

		// 1. Atualiza as estatísticas
		stats.UsageCount++
		stats.TotalDurationSeconds += int64(duration)
		stats.TotalFlowRateSum += flowRate
		stats.TotalLitersConsumed += float64(duration) * flowRate

		// 2. Preenche os pulsos, segundo a segundo
		for t := start; t < end; t++ {
			// Converte qualquer segundo para um índice válido entre 0 e 86399
			index := (int(t)%86400 + 86400) % 86400

			if device, deviceExists := p.pulsesDevice[sanitaryType]; deviceExists {
				device.pulses[index] += flowRate
			}
		}
	}
}

// NOVA FUNÇÃO PÚBLICA: Usa o método privado para agregar todos os logs de uma casa
// em um único PulseHouse, criando um "dia médio".
func (p *PulseHouse) UpdatePulse(house *entities.House) {
	residentsLogs := house.ResidentLogs()

	for i := 0; i < len(residentsLogs); i++ {
		sanitaryLogs := residentsLogs[i].SanitaryLogs()
		sanitaryMap := map[string]*logData.Sanitary{
			"toilet":       sanitaryLogs.ToiletLog(), "shower": sanitaryLogs.ShowerLog(),
			"wash_bassin":  sanitaryLogs.WashBassinLog(), "wash_machine": sanitaryLogs.WashMachineLog(),
			"dish_washer":  sanitaryLogs.DishWasherLog(), "tanque": sanitaryLogs.TanqueLog(),
		}

		for name, sanitary := range sanitaryMap {
			usageLogs, ok := sanitary.UsageLogs()
			if !ok { continue }
			for _, usage := range usageLogs {
				p.processSingleUsage(name, usage)
			}
		}
	}
}


// FUNÇÃO ORIGINAL: Mantém sua lógica complexa para distribuir os pulsos
// corretamente entre os dias da janela.
func (p *PulseHouse) UpdatePulseWithWindow(day uint8, house *entities.House, dayWindow map[uint8]*PulseHouse) {
	residentsLogs := house.ResidentLogs()

	for i := 0; i < len(residentsLogs); i++ {
		sanitaryLogs := residentsLogs[i].SanitaryLogs()
		sanitaryMap := map[string]*logData.Sanitary{
			"toilet":       sanitaryLogs.ToiletLog(), "shower": sanitaryLogs.ShowerLog(),
			"wash_bassin":  sanitaryLogs.WashBassinLog(), "wash_machine": sanitaryLogs.WashMachineLog(),
			"dish_washer":  sanitaryLogs.DishWasherLog(), "tanque": sanitaryLogs.TanqueLog(),
		}

		for name, sanitary := range sanitaryMap {
			usageLogs, ok := sanitary.UsageLogs()
			if !ok { continue }

			for _, usage := range usageLogs {
				start := usage.StartUsage()
				end := usage.EndUsage()

				// --- Lógica de Estatísticas (atribuída ao dia de início) ---
				_, startDay := p.GetIndexAndDay(start, day)
				if targetDayStats, ok := dayWindow[startDay]; ok {
					if stats, exists := targetDayStats.usageStats[name]; exists {
						duration := end - start
						flowRate := usage.FlowRate()
						stats.UsageCount++
						stats.TotalDurationSeconds += int64(duration)
						stats.TotalFlowRateSum += flowRate
						stats.TotalLitersConsumed += float64(duration) * flowRate
					}
				}

				// --- Lógica de Pulso (distribuída segundo a segundo) ---
				for t := start; t < end; t++ {
					index, targetDay := p.GetIndexAndDay(t, day)
					if target, ok := dayWindow[targetDay]; ok {
						if device, exists := target.pulsesDevice[name]; exists && index >= 0 && index < 86400 {
							device.pulses[index] += usage.FlowRate()
						}
					}
				}
			}
		}
	}
}


func (p *PulseHouse) ExportPulsesToCSV(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo CSV: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"horario_segundos", "toilet", "shower", "wash_bassin", "wash_machine", "dish_washer", "tanque", "total"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("erro ao escrever cabeçalho CSV: %w", err)
	}

	devices := OrderedDeviceKeys()
	for sec := 0; sec < 86400; sec++ {
		row := make([]string, len(devices)+2)
		row[0] = strconv.Itoa(sec)
		var total float64
		for i, device := range devices {
			val := 0.0
			if pd, ok := p.pulsesDevice[device]; ok {
				val = pd.pulses[sec]
			}
			total += val
			row[i+1] = fmt.Sprintf("%.6f", val)
		}
		row[len(devices)+1] = fmt.Sprintf("%.6f", total)
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("erro ao escrever linha CSV: %w", err)
		}
	}
	return nil
}

func (p *PulseHouse) PrintUsageStatistics() {
	fmt.Printf("\n--- Estatísticas de Uso Detalhado para o Dia %d ---\n", p.day)

	deviceKeys := OrderedDeviceKeys()
	var grandTotalUsageCount int
	var grandTotalDurationSeconds int64
	var grandTotalFlowRateSum float64
	var grandTotalLitersConsumed float64

	for _, deviceName := range deviceKeys {
		stats, ok := p.usageStats[deviceName]
		fmt.Printf("\nDispositivo: %s\n", deviceName)
		if !ok || stats.UsageCount == 0 {
			fmt.Println("  - Nenhum uso registrado.")
			continue
		}
		averageDuration := float64(stats.TotalDurationSeconds) / float64(stats.UsageCount)
		averageFlowRate := stats.TotalFlowRateSum / float64(stats.UsageCount)
		fmt.Printf("  - Total de Usos:        %d\n", stats.UsageCount)
		fmt.Printf("  - Total Consumido:      %.2f Litros\n", stats.TotalLitersConsumed)
		fmt.Printf("  - Tempo Médio por Uso:  %.2f segundos\n", averageDuration)
		fmt.Printf("  - Vazão Média por Uso:  %.4f L/s\n", averageFlowRate)
		grandTotalUsageCount += stats.UsageCount
		grandTotalDurationSeconds += int64(stats.TotalDurationSeconds)
		grandTotalFlowRateSum += stats.TotalFlowRateSum
		grandTotalLitersConsumed += stats.TotalLitersConsumed
	}
	fmt.Printf("\n----------------------------------------\n")
	fmt.Printf("--- Resumo Total do Dia %d ---\n", p.day)
	if grandTotalUsageCount == 0 {
		fmt.Println("  - Nenhum uso registrado no dia.")
		fmt.Printf("----------------------------------------\n")
		return
	}
	overallAverageDuration := float64(grandTotalDurationSeconds) / float64(grandTotalUsageCount)
	overallAverageFlowRate := grandTotalFlowRateSum / float64(grandTotalUsageCount)
	fmt.Printf("  - Total de Usos (todos):    %d\n", grandTotalUsageCount)
	fmt.Printf("  - Total Consumido (geral):  %.2f Litros\n", grandTotalLitersConsumed)
	fmt.Printf("  - Tempo Médio por Uso (geral): %.2f segundos\n", overallAverageDuration)
	fmt.Printf("  - Vazão Média por Uso (geral): %.4f L/s\n", overallAverageFlowRate)
	fmt.Printf("----------------------------------------\n")
}