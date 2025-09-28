package analysis

import "fmt"

// Supondo que os tipos são os mesmos usados anteriormente
func PrintUsagesOverview(dayWindow map[uint8]*AccumulatorDay, sanitaryTypes []string) {
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