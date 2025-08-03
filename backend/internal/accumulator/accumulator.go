package accumulator

import (
	"log"
	"simulation/internal/entities"
	logData "simulation/internal/log"
	"fmt"
	"math"
)

type AccumulatorInterface interface {
	AddConsumption(liters float64)
	Mean() float64
	EstimatedAverageFlowRate() float64
	Total() float64
	Counter() int
}

type AccumulatorUnit struct {
	WaterConsumption    float64
	counter int
}

func (a *AccumulatorUnit) AddConsumption(liters float64) {
	a.WaterConsumption += liters
	a.counter++
}

// Isso representa a media litros por usos
func (a *AccumulatorUnit) Mean() float64 {
	if a.counter == 0 {
		return 0
	}
	return a.WaterConsumption / float64(a.counter)
}

// Isso representa a media litros por segundo durante um intervalo de 1 hora
func (a *AccumulatorUnit) EstimatedAverageFlowRate() float64 {
	return a.WaterConsumption / 3600.0
}

func (a *AccumulatorUnit) Total() float64 {
	return a.WaterConsumption
}

func (a *AccumulatorUnit) Counter() int {
	return a.counter
}



type AccumulatorHour struct {
	sanitaryDevice map[string]AccumulatorInterface
}

func newAccumulatorHour(typeDevice []string) AccumulatorHour {
	device := make(map[string]AccumulatorInterface)
	for _, tipo := range typeDevice {
		device[tipo] = &AccumulatorUnit{}
	}
	return AccumulatorHour{sanitaryDevice: device}
}

func (a *AccumulatorHour) IndividualDeviceAddConsumption(typeDevice string, liters float64) {
	if device, ok := a.sanitaryDevice[typeDevice]; ok {
		device.AddConsumption(liters)
	}
}

func (a *AccumulatorHour) IndividualDeviceMean(typeDevice string) float64 {
	if device, ok := a.sanitaryDevice[typeDevice]; ok {
		return device.Mean()
	}
	return 0
}

func (a *AccumulatorHour) IndividualDeviceTotal(typeDevice string) float64 {
	if device, ok := a.sanitaryDevice[typeDevice]; ok {
		return device.Total()
	}
	return 0
}

func (a *AccumulatorHour) IndividualDeviceCounter(typeDevice string) int {
	if device, ok := a.sanitaryDevice[typeDevice]; ok {
		return device.Counter()
	}
	return 0
}

func (a *AccumulatorHour) IndividualDeviceAvarageFlowRate(typeDevice string) float64 {
	if device, ok := a.sanitaryDevice[typeDevice]; ok {
		return device.EstimatedAverageFlowRate()
	}
	return 0
}

func (a *AccumulatorHour) Mean() float64 {
	total := a.TotalConsumption()
	count := a.Count()

	if count == 0 {
		return 0
	}

	return total / float64(count)
}

func (a *AccumulatorHour) TotalConsumption() float64 {
	var total float64
	for _, device := range a.sanitaryDevice {
		total += device.Total()
	}
	return total
}

func (a *AccumulatorHour) Count() int {
	var count int
	for _, device := range a.sanitaryDevice {
		count += device.Counter()
	}
	return count
}

func (a *AccumulatorHour) FlowRate() float64 {
	totalFlow := a.TotalConsumption() / 3600

	return totalFlow

}



type AccumulatorDay struct {
	day uint8
	accumulatorHour []*AccumulatorHour
}

func NewAccumulatorDay(day uint8, tiposDispositivo []string) *AccumulatorDay {
	acc := make([]*AccumulatorHour, 24)
	for i := 0; i < 24; i++ {
		ah := newAccumulatorHour(tiposDispositivo)
		acc[i] = &ah
	}
	return &AccumulatorDay{
		day: day,
		accumulatorHour: acc,
	}
}

func (a *AccumulatorDay) getHour(seconds int32) int {
	return int(seconds) / 3600
}

func (a *AccumulatorDay) UpdateAccumulator(day uint8,house *entities.House,
		dayWindow map[uint8]*AccumulatorDay) {

	residentLogs := house.ResidentLogs()
	for i := 0; i < len(residentLogs); i++ {
		sanitaryLog := residentLogs[i].SanitaryLogs()
		a.SanitaryLogAccumulator(*sanitaryLog, dayWindow, day, house)
	}
}

func (a *AccumulatorDay) SanitaryLogAccumulator(sanitaryLog logData.ResidentSanitary,
		dayWindow map[uint8]*AccumulatorDay,day uint8, house *entities.House) {

	toiletLogs, ok := sanitaryLog.ToiletLog().UsageLogs()
		if ok {
			a.UsageLogAccumulator(toiletLogs, dayWindow, day, "toilet",house)
		}

		showerLogs, ok := sanitaryLog.ShowerLog().UsageLogs()
		if ok {
			a.UsageLogAccumulator(showerLogs, dayWindow, day, "shower",house)
		}

		washBassinLogs, ok := sanitaryLog.WashBassinLog().UsageLogs()
		if ok {
			a.UsageLogAccumulator(washBassinLogs, dayWindow, day, "wash_bassin",house)
		}

		washMachineLogs, ok := sanitaryLog.WashMachineLog().UsageLogs()
		if ok {
			a.UsageLogAccumulator(washMachineLogs, dayWindow, day, "wash_machine",house)
		}

		dishWasherLogs, ok := sanitaryLog.DishWasherLog().UsageLogs()
		if ok {
			a.UsageLogAccumulator(dishWasherLogs, dayWindow, day, "dish_washer",house)
		}

		tanqueLogs, ok := sanitaryLog.TanqueLog().UsageLogs()
		if ok {
			a.UsageLogAccumulator(tanqueLogs, dayWindow, day, "tanque",house)
		}

}

func (a *AccumulatorDay) UsageLogAccumulator(usages []*logData.Usage,
		dayWindow map[uint8]*AccumulatorDay,day uint8, sanitaryType string,house *entities.House) {

	for i := 0; i < len(usages); i++ {
		hour, day := a.GetHourDay(usages[i],day,house) // hour aparentemente ta dando mais que 24
		liters := usages[i].WaterConsumption()
		dayWindow[day].accumulatorHour[hour].IndividualDeviceAddConsumption(sanitaryType,liters)

	}

}

func (a *AccumulatorDay) GetHourDay(usage *logData.Usage, day uint8,house *entities.House) (int, uint8) {
		
	var hour int
	startUsage := usage.StartUsage()

	if startUsage >= 0 && startUsage < 86400 {
		hour = a.getHour(startUsage)
		if hour >= 24 {
			log.Fatal("Caso 1")
		}
		if hour < 0 {
			log.Fatal("Caso 2")
		}
		return hour, day
	} else if startUsage >= 86400 {
		hour := a.getHour(startUsage) -24
		if hour >= 24 {
			log.Fatal("Caso 3")
		}
		if hour < 0 {
			log.Fatal("Caso 4")
		}
		return hour, day + 1
	} else {
		hour := a.getHour(startUsage) +24
		if hour >= 24 {
			log.Fatal("Caso 5")
		}
		if hour < 0 {
			log.Fatal("Caso 6")
		}
		return hour, day - 1
	}
		
}

// RoundFloat2 arredonda float64 para 2 casas decimais
func RoundFloat2(f float64) float64 {
	return math.Round(f*100) / 100
}

// RoundAccumulatorDayValues percorre todos os valores do dia e arredonda os litros
func (a *AccumulatorDay) RoundAccumulatorDayValues() {
	for _, hour := range a.accumulatorHour {
		for _, device := range hour.sanitaryDevice {
			if unit, ok := device.(*AccumulatorUnit); ok {
				unit.WaterConsumption = RoundFloat2(unit.WaterConsumption)
			}
		}
	}
}

func OrderedDeviceKeys() []string {
	return []string{
		"toilet",
		"shower",
		"wash_bassin",
		"wash_machine",
		"dish_washer",
		"tanque",
	}
}


func (a *AccumulatorDay) PrintHourlyWaterConsumption() {
	// Mapeia tipos de dispositivos e inicializa totais
	deviceTotals := make(map[string]float64)
	var totalPerHour [24]float64
	var grandTotal float64

	if len(a.accumulatorHour) != 24 {
		log.Println("Dados incompletos: esperados 24 acumuladores de hora.")
		return
	}

	fmt.Println("Consumo por dispositivo:")

	// Primeiro, descobrimos todos os tipos de dispositivos existentes
	// Pegando do primeiro acumulador (hora 0), pois todos têm as mesmas chaves
	deviceTypes := OrderedDeviceKeys()

	// Para cada tipo de dispositivo, printar hora a hora
	for _, device := range deviceTypes {
		fmt.Printf("\nDispositivo: %s\n", device)
		for hour := 0; hour < 24; hour++ {
			value := a.accumulatorHour[hour].IndividualDeviceTotal(device)
			fmt.Printf("Hora %02d: %.2f Litros\n", hour, value)

			// Acumula para total por hora e total por dispositivo
			totalPerHour[hour] += value
			deviceTotals[device] += value
			grandTotal += value
		}
	}

	// Total por hora
	fmt.Println("\nTotal por hora:")
	for hour := 0; hour < 24; hour++ {
		fmt.Printf("Hora %02d: %.2f Litros\n", hour, totalPerHour[hour])
	}

	// Total geral
	fmt.Printf("\nTotal completo do dia: %.2f Litros\n", grandTotal)
}