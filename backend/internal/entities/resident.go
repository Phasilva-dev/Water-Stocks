package entities

import (
	"math/rand/v2"
	"simulation/internal/entities/resident/ds/temporal"
	"simulation/internal/entities/resident"
	"simulation/internal/log"
	"simulation/internal/usagemock"

)

type Resident struct {
	age uint8
	occupationID uint32 //Ocupação, exemplo, estudante
	dayData *temporal.DailyData
	residentProfile *resident.ResidentProfile
	house *House
}

func NewResident(age uint8, occupation uint32, profile *resident.ResidentProfile, house *House) *Resident {
	return &Resident{
		age: age,
		occupationID: occupation,
		dayData: temporal.NewDailyData(nil,nil),
		residentProfile: profile,
		house: house,
	}
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
	r.GenerateDailyData(day,rng)
	frequency := r.dayData.Frequency()
	routine := r.dayData.Routine()

	usageToilet := make([]*log.Usage, frequency.Toilet())
	usageShower := make([]*log.Usage, frequency.Shower())
	usageWashBassin := make([]*log.Usage, frequency.WashBassin())
	usageWashMachine := make([]*log.Usage, frequency.WashMachine())
	usageDishWasher := make([]*log.Usage, frequency.DishWasher())
	usageTanque := make([]*log.Usage, frequency.Tanque())

	sanitaryHouse := r.house.SanitaryHouse()

	for i := 0; i < len(usageToilet); i++ {
		usage, err := usagemock.GenerateToiletUsage(routine, sanitaryHouse.Toilet().Device(), rng)
		if err != nil {
			return nil, err
		}
		usageToilet[i] = usage
	}

	for i := 0; i < len(usageShower); i++ {
		usage, err := usagemock.GenerateShowerUsage(routine, sanitaryHouse.Shower().Device(), rng)
		if err != nil {
			return nil, err
		}
		usageShower[i] = usage
	}

	for i := 0; i < len(usageWashBassin); i++ {
		usage, err := usagemock.GenerateWashBassinUsage(routine, sanitaryHouse.WashBassin().Device(), rng)
		if err != nil {
			return nil, err
		}
		usageWashBassin[i] = usage
	}

	for i := 0; i < len(usageWashMachine); i++ {
		usage, err := usagemock.GenerateWashMachineUsage(routine, sanitaryHouse.WashMachine().Device(), rng)
		if err != nil {
			return nil, err
		}
		usageWashMachine[i] = usage
	}

	for i := 0; i < len(usageDishWasher); i++ {
		usage, err := usagemock.GenerateDishWasherUsage(routine, sanitaryHouse.DishWasher().Device(), rng)
		if err != nil {
			return nil, err
		}
		usageDishWasher[i] = usage
	}

	for i := 0; i < len(usageTanque); i++ {
		usage, err := usagemock.GenerateTanqueUsage(routine, sanitaryHouse.Tanque().Device(), rng)
		if err != nil {
			return nil, err
		}
		usageTanque[i] = usage
	}
	toiletLog := log.NewSanitary("toilet", sanitaryHouse.Toilet().Device().SanitaryDeviceID(),usageToilet)
	showerLog := log.NewSanitary("shower", sanitaryHouse.Shower().Device().SanitaryDeviceID(),usageShower)
	washBassinLog := log.NewSanitary("wash_bassin", sanitaryHouse.WashBassin().Device().SanitaryDeviceID(),usageWashBassin)
	washMachineLog := log.NewSanitary("wash_machine", sanitaryHouse.WashMachine().Device().SanitaryDeviceID(),usageWashMachine)
	dishWasherLog := log.NewSanitary("dish_washer", sanitaryHouse.DishWasher().Device().SanitaryDeviceID(),usageDishWasher)
	tanqueLog := log.NewSanitary("tanque", sanitaryHouse.Tanque().Device().SanitaryDeviceID(),usageTanque)

	residentSanitarylog := log.NewResidentSanitary(toiletLog,showerLog,washBassinLog,washMachineLog,dishWasherLog,tanqueLog)

	residentLog := log.NewResident(day+1, r.house.houseClassID,r.occupationID,r.age,residentSanitarylog)

	r.dayData.ClearData()
	return residentLog,nil
}
