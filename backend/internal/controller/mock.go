package controller

import (
	"simulation/internal/entities"
	"simulation/internal/log"

	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type SanitaryUsageLine struct {
	Day                 uint8
	HouseID             uint32
	ResidentOccupationID uint32
	Age                 uint8
	SanitaryDeviceID    uint32
	SanitaryType        string
	StartUsage          int32
	EndUsage            int32
	FlowRate            float64
}

func ResidentLogToUsageLines(houseID uint32, resident *log.Resident) []SanitaryUsageLine {
	lines := []SanitaryUsageLine{}

	sanitaryLogs := []*log.Sanitary{
		resident.SanitaryLogs().ToiletLog(),
		resident.SanitaryLogs().ShowerLog(),
		resident.SanitaryLogs().WashBassinLog(),
		resident.SanitaryLogs().WashMachineLog(),
		resident.SanitaryLogs().DishWasherLog(),
		resident.SanitaryLogs().TanqueLog(),
	}

	for _, sanitary := range sanitaryLogs {
		usages, ok := sanitary.UsageLogs()
		if !ok {
			continue
		}

		for _, usage := range usages {
			line := SanitaryUsageLine{
				Day:                  resident.Day(),
				HouseID:              houseID,
				ResidentOccupationID: resident.ResidentOccupationID(),
				Age:                  resident.Age(),
				SanitaryDeviceID:     sanitary.SanitaryDeviceID(),
				SanitaryType:         sanitary.SanitaryType(),
				StartUsage:           usage.StartUsage(),
				EndUsage:             usage.EndUsage(),
				FlowRate:             usage.FlowRate(),
			}
			lines = append(lines, line)
		}
	}

	return lines
}

func ToSanitaryUsageLines(h *entities.House) []SanitaryUsageLine {
	lines := []SanitaryUsageLine{}

	for _, residentLog := range h.ResidentLogs() {
		residentLines := ResidentLogToUsageLines(h.HouseClassID(), residentLog)
		lines = append(lines, residentLines...)
	}

	return lines
}

func ExportUsageLinesToCSV(filename string, lines []SanitaryUsageLine) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Cabeçalho
	writer.Write([]string{
		"Day", "HouseID", "ResidentOccupationID", "Age",
		"SanitaryDeviceID", "SanitaryType",
		"StartUsage", "EndUsage", "FlowRate",
	})

	// Dados
	for _, line := range lines {
		record := []string{
			strconv.Itoa(int(line.Day)),
			strconv.Itoa(int(line.HouseID)),
			strconv.Itoa(int(line.ResidentOccupationID)),
			strconv.Itoa(int(line.Age)),
			strconv.Itoa(int(line.SanitaryDeviceID)),
			line.SanitaryType,
			strconv.Itoa(int(line.StartUsage)),
			strconv.Itoa(int(line.EndUsage)),
			strconv.FormatFloat(line.FlowRate, 'f', 2, 64),
		}
		writer.Write(record)
	}

	return nil
}

// Ordem fixa dos dispositivos
var deviceOrder = []string{
	"toilet",
	"shower",
	"wash_bassin",
	"wash_machine",
	"dish_washer",
	"tanque",
}

// Resultado de agregação
type UsageAggregation struct {
	ByDevicePerHour map[string]map[int]float64 // ["Shower"][hour] = flow
	TotalPerHour    map[int]float64            // [hour] = flow
	TotalPerDevice  map[string]float64         // ["Shower"] = flow
	TotalUsage      float64
}

// Função principal de agregação
func AggregateSanitaryUsage(lines []SanitaryUsageLine) UsageAggregation {
	byDevicePerHour := make(map[string]map[int]float64)
	totalPerHour := make(map[int]float64)
	totalPerDevice := make(map[string]float64)
	var totalUsage float64 = 0

	for _, line := range lines {
		// Hora baseada em segundos, pode ser negativa ou maior que 24h
		hour := int(line.StartUsage) / 3600
		device := line.SanitaryType

		// Inicializa mapa interno para o dispositivo, se não existir
		if _, ok := byDevicePerHour[device]; !ok {
			byDevicePerHour[device] = make(map[int]float64)
		}

		flow := line.FlowRate

		// Soma por dispositivo e hora
		byDevicePerHour[device][hour] += flow

		// Soma total por hora
		totalPerHour[hour] += flow

		// Soma total por dispositivo
		totalPerDevice[device] += flow

		// Soma total geral
		totalUsage += flow
	}

	return UsageAggregation{
		ByDevicePerHour: byDevicePerHour,
		TotalPerHour:    totalPerHour,
		TotalPerDevice:  totalPerDevice,
		TotalUsage:      totalUsage,
	}
}

func PrintUsageByDevicePerHour(agg UsageAggregation) {
	fmt.Println("🔹 Uso por dispositivo hora a hora:")
	for _, device := range deviceOrder {
		hours, ok := agg.ByDevicePerHour[device]
		if !ok {
			continue
		}

		totalDevice := agg.TotalPerDevice[device]
		if totalDevice == 0 {
			totalDevice = 1 // evitar divisão por zero
		}

		fmt.Printf("Dispositivo: %s\n", device)

		// Ordena horas numericamente incluindo negativas e >24
		sortedHours := make([]int, 0, len(hours))
		for h := range hours {
			sortedHours = append(sortedHours, h)
		}
		sort.Ints(sortedHours)

		for _, h := range sortedHours {
			flow := hours[h]
			percent := (flow / totalDevice) * 100
			fmt.Printf("  Hora %+d: %.2f (%.2f%% do total %s)\n", h, flow, percent, device)
		}
		fmt.Println()
	}
}

func PrintTotalPerHour(agg UsageAggregation) {
	fmt.Println("🔸 Uso total da casa hora a hora:")
	totalGeneral := agg.TotalUsage
	if totalGeneral == 0 {
		totalGeneral = 1 // evitar divisão por zero
	}

	sortedHours := make([]int, 0, len(agg.TotalPerHour))
	for h := range agg.TotalPerHour {
		sortedHours = append(sortedHours, h)
	}
	sort.Ints(sortedHours)

	for _, h := range sortedHours {
		flow := agg.TotalPerHour[h]
		percent := (flow / totalGeneral) * 100
		fmt.Printf("  Hora %+d: %.2f (%.2f%% do total geral)\n", h, flow, percent)
	}
	fmt.Println()
}

func PrintTotalPerDevice(agg UsageAggregation) {
	fmt.Println("🔹 Uso total por dispositivo:")
	for _, device := range deviceOrder {
		usage, ok := agg.TotalPerDevice[device]
		if !ok {
			usage = 0
		}
		fmt.Printf("  %s: %.2f\n", device, usage)
	}
	fmt.Println()
}

func PrintTotalUsage(agg UsageAggregation) {
	fmt.Printf("🔸 Uso total da casa: %.2f\n", agg.TotalUsage)
}
