package accumulator

import (
	"simulation/internal/entities"
	logData "simulation/internal/log"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type PulseDevice struct {
	device string
	pulses []float64
}

func NewPulseDevice(DeviceName string) *PulseDevice {
	pulses := make([]float64, 86400) // 86400 segundos = 24h * 60min * 60s
	return &PulseDevice{
		device:    DeviceName,
		pulses: pulses,
	}
}

func (p *PulseDevice) updateDevice(usage *logData.Usage) {

	for i := usage.StartUsage(); i < usage.EndUsage(); i++ {
		p.pulses[i] = usage.FlowRate()
	}
}

type PulseHouse struct {
	day uint8
	pulsesDevice map[string]*PulseDevice
}

func NewPulseHouse(day uint8, deviceNames []string) *PulseHouse {
	pulses := make(map[string]*PulseDevice)
	for _, name := range deviceNames {
		pulses[name] = NewPulseDevice(name)
	}
	return &PulseHouse{
		day:          day,
		pulsesDevice: pulses,
	}
}

func (p *PulseHouse) GetIndexAndDay(second int32, day uint8) (int, uint8) {
	switch {
	case second >= 0 && second < 86400:
		return int(second), day
	case second >= 86400:
		return int(second) - 86400, day + 1
	default:
		return int(second) + 86400, day - 1
	}
}

func (p *PulseHouse) UpdatePulseWithWindow(day uint8, house *entities.House, dayWindow map[uint8]*PulseHouse) {
	residentsLogs := house.ResidentLogs()

	for i := 0; i < len(residentsLogs); i++ {
		sanitaryLogs := residentsLogs[i].SanitaryLogs()

		sanitaryMap := map[string]*logData.Sanitary{
			"toilet":       sanitaryLogs.ToiletLog(),
			"shower":       sanitaryLogs.ShowerLog(),
			"wash_bassin":  sanitaryLogs.WashBassinLog(),
			"wash_machine": sanitaryLogs.WashMachineLog(),
			"dish_washer":  sanitaryLogs.DishWasherLog(),
			"tanque":       sanitaryLogs.TanqueLog(),
		}

		for name, sanitary := range sanitaryMap {
			usageLogs, ok := sanitary.UsageLogs()
			if !ok {
				continue
			}

			for _, usage := range usageLogs {
				start := usage.StartUsage()
				end := usage.EndUsage()

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
	// Abrir/criar o arquivo CSV
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo CSV: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Cabeçalho
	header := []string{
		"horario_segundos",
		"toilet",
		"shower",
		"wash_bassin",
		"wash_machine",
		"dish_washer",
		"tanque",
		"total",
	}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("erro ao escrever cabeçalho CSV: %w", err)
	}

	devices := []string{"toilet", "shower", "wash_bassin", "wash_machine", "dish_washer", "tanque"}

	// Percorrer todos os segundos do dia
	for sec := 0; sec < 86400; sec++ {
		row := make([]string, len(devices)+2) // +2 para horario e total
		row[0] = strconv.Itoa(sec)

		var total float64
		for i, device := range devices {
			val := 0.0
			if pd, ok := p.pulsesDevice[device]; ok && sec < len(pd.pulses) {
				val = pd.pulses[sec]
			}
			total += val
			row[i+1] = fmt.Sprintf("%.6f", val) // 6 casas decimais, ajuste se quiser
		}

		row[len(devices)+1] = fmt.Sprintf("%.6f", total)

		if err := writer.Write(row); err != nil {
			return fmt.Errorf("erro ao escrever linha CSV: %w", err)
		}
	}

	return nil
}