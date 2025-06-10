package entities

import (
	"fmt"
	"math/rand/v2"
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

		h.residents[i] = NewResident(age, occupation,profile, h)
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
		return nil
	}

	err = h.GenerateSanitaryDeviceOfHouse(rng)
	if err != nil {
		return err
	}

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
	return nil

}
