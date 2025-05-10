package entities

import (
	"fmt"
	"math/rand/v2"
	"simulation/internal/entities/house"
	"simulation/internal/entities/house/ds/sanitarysystem"
	"simulation/internal/entities/house/profile/sanitarydevice"
)

type House struct {
	houseClassID uint32
	residents []*Resident
	sanitaryHouse *sanitarysystem.SanitaryHouse
	houseProfile *house.HouseProfile
}

func (h *House) SanitaryHouse() *sanitarysystem.SanitaryHouse {
	return h.sanitaryHouse
}

func (h *House) Residents() []*Resident {
	return h.residents
}

func NewHouse(houseClassID uint32, houseProfile *house.HouseProfile) *House {
	return &House{
		houseClassID:  houseClassID,
		residents:     []*Resident{},
		sanitaryHouse: nil,
		houseProfile:  houseProfile,
	}
}

func (h *House) GenerateResidents(rng *rand.Rand) error {
	num := h.houseProfile.GenerateNumbersOfResidents(rng)
	h.residents = make([]*Resident, num)

	for i := 0; i < int(num); i++ {
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

		

		occupation := h.houseProfile.GenerateOccupation(age, rng)

		profile, err := h.houseProfile.ResidentProfile(occupation)
		if err != nil {
			return fmt.Errorf("resident profile with occupation ID %d not found", occupation)
		}

		h.residents[i] = NewResident(age, occupation,profile, h)
	}

	return nil
}

func (h *House) GenerateSanitaryDeviceOfHouse(rng *rand.Rand,devices map[string]sanitarydevice.SanitaryDevice) error {
	numberOfResidents := uint8(len(h.residents))
	amountOfSanitarys, err := h.houseProfile.GenerateNumberOfSanitaryDevices(rng,numberOfResidents)
	if err != nil {
		return err
	}
	sanitHouse, err := h.houseProfile.GenerateSanitaryHouse(devices, amountOfSanitarys)
	if err != nil {
		return err
	}
	h.sanitaryHouse = sanitHouse
	return nil
}