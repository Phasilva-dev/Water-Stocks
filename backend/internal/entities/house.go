package entities

import (
	"fmt"
	"math/rand/v2"
	"simulation/internal/dists"
	"simulation/internal/entities/house"
	"simulation/internal/entities/house/ds/sanitarysystem"
	"simulation/internal/log"
)

type House struct {
	houseClassID uint32
	residents []*Resident
	sanitaryHouse *sanitarysystem.SanitaryHouse
	houseProfile *house.HouseProfile
	residentLogs []*log.Resident
	//householdLogs []*log.Usage
}

func (h *House) SanitaryHouse() *sanitarysystem.SanitaryHouse {
	return h.sanitaryHouse
}

func (h *House) Residents() []*Resident {
	return h.residents
}

func (h *House) ResidentLogs() []*log.Resident {
	return h.residentLogs
}

func (h *House) HouseClassID() uint32 {
	return h.houseClassID
}

func NewHouse(houseClassID uint32, houseProfile *house.HouseProfile) *House {
	return &House{
		houseClassID:  houseClassID,
		residents:     nil,
		sanitaryHouse: nil,
		houseProfile:  houseProfile,
		residentLogs: nil,
		//householdLogs: nil,
	}
}

func (h *House) GenerateResidents(rng *rand.Rand) error {
	num := h.houseProfile.GenerateNumbersOfResidents(rng)
	h.residents = make([]*Resident, num)

	for i := uint8(0); i < num; i++ {
		var age uint8

		if i == 0 {
			// Garantir que o primeiro residente tenha idade >= 18
			for {
				age = h.houseProfile.GenerateAgeofResidents(rng)

				if age >= 18 {
					break
				}
			}
		} else {
			age = h.houseProfile.GenerateAgeofResidents(rng)
		}


		

		occupation, err := h.houseProfile.GenerateOccupation(age, rng)
		if err != nil {
			return err
		}

		profile, err := h.houseProfile.ResidentProfile(occupation)
		if err != nil {
			return fmt.Errorf("resident profile with occupation ID %d not found", occupation)
		}

		resident, err := NewResident(age, occupation,profile, h)
		if err != nil {
			return err
		}
		h.residents[i] = resident
	}

	return nil
}

func (h *House) GenerateSanitaryDeviceOfHouse(rng *rand.Rand) error {
	numberOfResidents := uint8(len(h.residents))
	amountOfSanitarys, err := h.houseProfile.GenerateNumberOfSanitaryDevices(rng,numberOfResidents)
	if err != nil {
		return err
	}
	sanitHouse, err := h.houseProfile.GenerateSanitaryHouse(amountOfSanitarys)
	if err != nil {
		return err
	}
	h.sanitaryHouse = sanitHouse
	return nil
}

func (h *House) GenerateHouseData (rng *rand.Rand) error {

	err := h.GenerateResidents(rng)
	if err != nil {
		return err
	}

	err = h.GenerateSanitaryDeviceOfHouse(rng)
	if err != nil {
		return err
	}

	return nil
}

func (h *House) GenerateHouseholdLog (day uint8, rng *rand.Rand) error {
	var freqWashMachine, freqWashDishWasher, freqTanque uint8
	for i := 0; i < len(h.residents); i++ {
		resident := h.residents[i]
		freqWashMachine += resident.dayData.Frequency().WashMachine()
		freqWashDishWasher += resident.dayData.Frequency().DishWasher()
		freqTanque += resident.dayData.Frequency().Tanque()
	}
	dist, _ := dists.CreateDistribution("uniform", 0, 100)
	p := dist.Sample(rng)
	resident := h.residents[0]
	if p >= 66.66 {
		resident.GenerateWashMachineLogs(freqWashMachine, day, rng)
	}
	resident.GenerateDishWasherLogs(freqWashDishWasher, day, rng)
	resident.GenerateTanqueLogs(freqTanque, day, rng)

	return nil
	
}

func (h *House) GenerateLogs (day uint8,rng *rand.Rand) error {

	
	h.residentLogs = make([]*log.Resident,len(h.residents))
	for i := 0; i < len(h.residents); i++ {
		residentLog, err := h.residents[i].GenerateLogs(day, rng)
		if err != nil {
			return err
		}
		h.residentLogs[i] = residentLog
	}
	h.GenerateHouseholdLog(day,rng)
	return nil

}
