package guardar
/*
import (
	"fmt"
	"log"
	"simulation/internal/entities"
	"sort"
	logData "simulation/internal/log"
)




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
}*/