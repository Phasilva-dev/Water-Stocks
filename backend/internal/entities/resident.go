package entities

import (
	"errors"
	"math/rand/v2"
	"simulation/internal/entities/resident"
	"simulation/internal/entities/resident/ds/temporal"
	"simulation/internal/log"
	"simulation/internal/usagemock"

	"fmt"
)

type Resident struct {
	age uint8
	occupationID uint32 //Ocupação, exemplo, estudante
	dayData *temporal.DailyData
	residentProfile *resident.ResidentProfile
	house *House
}

func NewResident(age uint8, occupation uint32, profile *resident.ResidentProfile, house *House) (*Resident, error) {
	if occupation == 0 {
		return nil, errors.New("ocupação não pode ser zero")
	}
	if profile == nil {
		return nil, errors.New("perfil do residente é nil")
	}
	if house == nil {
		return nil, errors.New("casa é nil")
	}
	return &Resident{
		age: age,
		occupationID: occupation,
		dayData: temporal.NewDailyData(nil,nil),
		residentProfile: profile,
		house: house,
	}, nil
}

func (r *Resident) Age() uint8 {
	return r.age
}
func (r *Resident) OccupationID() uint32 {
	return r.occupationID
}

func (r *Resident) DayData() *temporal.DailyData {
	return r.dayData
}

func (r *Resident) ResidentProfile() *resident.ResidentProfile {
	return r.residentProfile
}

func (r *Resident) House() *House {
	return r.house
}


func (r *Resident) GenerateFrequency(day uint8, rng *rand.Rand) {
	r.dayData.SetFrequency(r.residentProfile.GenerateFrequency(day,rng))
}

func (r *Resident) GenerateRoutine(day uint8, rng *rand.Rand) {
	r.dayData.SetRoutine(r.residentProfile.GenerateRoutine(day,rng))
}

func (r *Resident) GenerateDailyData(day uint8, rng *rand.Rand) {
	r.GenerateRoutine(day,rng)
	r.GenerateFrequency(day,rng)
}

func (r *Resident) GenerateLogs(day uint8,rng *rand.Rand) (*log.Resident,error) {

	// ✅ Checagem 1: Resident.House não pode ser nil
	if r.house == nil {
		return nil, errors.New("resident.house is nil")
	}

	// ✅ Checagem 2: SanitaryHouse da House não pode ser nil
	sanitaryHouse := r.house.SanitaryHouse()
	if sanitaryHouse == nil {
		return nil, errors.New("resident.house.SanitaryHouse() returned nil")
	}

	// ✅ Checagem 3: Cada dispositivo também precisa ser não-nulo (opcional, mas ideal)
	if sanitaryHouse.Toilet() == nil {
		return nil, errors.New("sanitaryHouse.Toilet() is nil")
	}
	if sanitaryHouse.Shower() == nil {
		return nil, errors.New("sanitaryHouse.Shower() is nil")
	}
	if sanitaryHouse.WashBassin() == nil {
		return nil, errors.New("sanitaryHouse.WashBassin() is nil")
	}
	if sanitaryHouse.WashMachine() == nil {
		return nil, errors.New("sanitaryHouse.WashMachine() is nil")
	}
	if sanitaryHouse.DishWasher() == nil {
		return nil, errors.New("sanitaryHouse.DishWasher() is nil")
	}
	if sanitaryHouse.Tanque() == nil {
		return nil, errors.New("sanitaryHouse.Tanque() is nil")
	}

	r.GenerateDailyData(day,rng)
	frequency := r.dayData.Frequency()
	routine := r.dayData.Routine()

	usageToilet := make([]*log.Usage, frequency.Toilet())
	usageShower := make([]*log.Usage, frequency.Shower())
	usageWashBassin := make([]*log.Usage, frequency.WashBassin())
	usageWashMachine := make([]*log.Usage, 0)
	usageDishWasher := make([]*log.Usage, 0)
	usageTanque := make([]*log.Usage, 0)

	for i := 0; i < len(usageToilet); i++ {
		usage, err := usagemock.GenerateToiletUsage(routine, sanitaryHouse.Toilet().Device(), rng)
		if err != nil {
			return nil, fmt.Errorf("erro ao gerar logs para occupationID=%d: %w", r.occupationID, err)
		}
		usageToilet[i] = usage
	}

	for i := 0; i < len(usageShower); i++ {
		usage, err := usagemock.GenerateShowerUsage(routine, sanitaryHouse.Shower().Device(), rng)
		if err != nil {
			return nil, fmt.Errorf("erro ao gerar logs para occupationID=%d: %w", r.occupationID, err)
		}
		usageShower[i] = usage
	}

	for i := 0; i < len(usageWashBassin); i++ {
		usage, err := usagemock.GenerateWashBassinUsage(routine, sanitaryHouse.WashBassin().Device(), rng)
		if err != nil {
			return nil, fmt.Errorf("erro ao gerar logs para occupationID=%d: %w", r.occupationID, err)
		}
		usageWashBassin[i] = usage
	}

	/*
	for i := 0; i < len(usageWashMachine); i++ {
		usage, err := usagemock.GenerateWashMachineUsage(routine, sanitaryHouse.WashMachine().Device(), rng)
		if err != nil {
			return nil, fmt.Errorf("erro ao gerar logs para occupationID=%d: %w", r.occupationID, err)
		}
		usageWashMachine[i] = usage
	}

	for i := 0; i < len(usageDishWasher); i++ {
		usage, err := usagemock.GenerateDishWasherUsage(routine, sanitaryHouse.DishWasher().Device(), rng)
		if err != nil {
			return nil, fmt.Errorf("erro ao gerar logs para occupationID=%d: %w", r.occupationID, err)
		}
		usageDishWasher[i] = usage
	}

	for i := 0; i < len(usageTanque); i++ {
		usage, err := usagemock.GenerateTanqueUsage(routine, sanitaryHouse.Tanque().Device(), rng)
		if err != nil {
			return nil, fmt.Errorf("erro ao gerar logs para occupationID=%d: %w", r.occupationID, err)
		}
		usageTanque[i] = usage
	}*/
	
	toiletLog := log.NewSanitary("toilet", sanitaryHouse.Toilet().Device().SanitaryDeviceID(),usageToilet)
	showerLog := log.NewSanitary("shower", sanitaryHouse.Shower().Device().SanitaryDeviceID(),usageShower)
	washBassinLog := log.NewSanitary("wash_bassin", sanitaryHouse.WashBassin().Device().SanitaryDeviceID(),usageWashBassin)
	washMachineLog := log.NewSanitary("wash_machine", sanitaryHouse.WashMachine().Device().SanitaryDeviceID(),usageWashMachine)
	dishWasherLog := log.NewSanitary("dish_washer", sanitaryHouse.DishWasher().Device().SanitaryDeviceID(),usageDishWasher)
	tanqueLog := log.NewSanitary("tanque", sanitaryHouse.Tanque().Device().SanitaryDeviceID(),usageTanque)

	residentSanitarylog := log.NewResidentSanitary(toiletLog,showerLog,washBassinLog,washMachineLog,dishWasherLog,tanqueLog)

	residentLog := log.NewResident(day, r.house.houseClassID,r.occupationID,r.age,residentSanitarylog)

	//r.dayData.ClearData()
	return residentLog,nil
}

func (r *Resident) GenerateWashMachineLogs(freqWashMachine uint8,
	day uint8,rng *rand.Rand) (error) {

	if r.house == nil {
		return  errors.New("resident.house is nil") //nil,
	}

	// ✅ Checagem 2: SanitaryHouse da House não pode ser nil
	sanitaryHouse := r.house.SanitaryHouse()
	if sanitaryHouse == nil {
		return  errors.New("resident.house.SanitaryHouse() returned nil") //nil,
	}
	
	routine := r.dayData.Routine()

	usageWashMachine := make([]*log.Usage, freqWashMachine)

	for i := 0; i < len(usageWashMachine); i++ {
		usage, err := usagemock.GenerateWashMachineUsage(routine, sanitaryHouse.WashMachine().Device(), rng)
		if err != nil {
			return fmt.Errorf("erro ao gerar logs para occupationID=%d: %w", r.occupationID, err) //nil,
		}
		usageWashMachine[i] = usage
	}

	washMachineLog := log.NewSanitary("wash_machine", sanitaryHouse.WashMachine().Device().SanitaryDeviceID(),usageWashMachine)

	r.house.residentLogs[0].SanitaryLogs().SetWashMachineLog(washMachineLog)
	return nil

}

func (r *Resident) GenerateDishWasherLogs(freqWashDishWasher uint8,
	day uint8,rng *rand.Rand) (error) {

	if r.house == nil {
		return  errors.New("resident.house is nil") //nil,
	}

	// ✅ Checagem 2: SanitaryHouse da House não pode ser nil
	sanitaryHouse := r.house.SanitaryHouse()
	if sanitaryHouse == nil {
		return  errors.New("resident.house.SanitaryHouse() returned nil") //nil,
	}
	
	routine := r.dayData.Routine()

	usageDishWasher := make([]*log.Usage, freqWashDishWasher)


	for i := 0; i < len(usageDishWasher); i++ {
		usage, err := usagemock.GenerateDishWasherUsage(routine, sanitaryHouse.DishWasher().Device(), rng)
		if err != nil {
			return fmt.Errorf("erro ao gerar logs para occupationID=%d: %w", r.occupationID, err) //nil, 
		}
		usageDishWasher[i] = usage
	}

	dishWasherLog := log.NewSanitary("dish_washer", sanitaryHouse.DishWasher().Device().SanitaryDeviceID(),usageDishWasher)

	r.house.residentLogs[0].SanitaryLogs().SetDishWasherLog(dishWasherLog)

	return nil

}

func (r *Resident) GenerateTanqueLogs(freqTanque uint8,
	day uint8,rng *rand.Rand) (error) { //*log.Resident


	if r.house == nil {
		return  errors.New("resident.house is nil") //nil,
	}

	// ✅ Checagem 2: SanitaryHouse da House não pode ser nil
	sanitaryHouse := r.house.SanitaryHouse()
	if sanitaryHouse == nil {
		return  errors.New("resident.house.SanitaryHouse() returned nil") //nil,
	}
	
	routine := r.dayData.Routine()

	usageTanque := make([]*log.Usage, freqTanque)

	for i := 0; i < len(usageTanque); i++ {
		usage, err := usagemock.GenerateTanqueUsage(routine, sanitaryHouse.Tanque().Device(), rng)
		if err != nil {
			return fmt.Errorf("erro ao gerar logs para occupationID=%d: %w", r.occupationID, err) //nil, 
		}
		usageTanque[i] = usage
	}

	tanqueLog := log.NewSanitary("tanque", sanitaryHouse.Tanque().Device().SanitaryDeviceID(),usageTanque)

	r.house.residentLogs[0].SanitaryLogs().SetTanqueLog(tanqueLog)

	return nil

}
