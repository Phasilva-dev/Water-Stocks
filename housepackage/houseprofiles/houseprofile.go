package houseprofiles

import (

	"math/rand/v2"
	//"housedata"

)

type HouseProfile struct {
	houseClassID uint32

	numResidentsProfile ResidentCountProfile
	ageProfile AgeProfile
	occupationProfile OccupationProfile

	numSanitarysDevice SanitaryCountProfile
	//sanitaryTypeProfile SanitaryTypeProfile //Preciso mudar a logica disso
	
}

func (h *HouseProfile) GenerateNumbersOfResidents(rng *rand.Rand) uint8 {
	return h.numResidentsProfile.GenerateData(rng)
}

func (h *HouseProfile) GenerateAgeofResidents(rng *rand.Rand) uint8 {
	return uint8(h.ageProfile.GenerateData(rng))
}

func (h *HouseProfile) GenerateOccupation(age uint8, rng *rand.Rand) uint32 {
	if age >= 18 && age < 65 {
		return h.occupationProfile.GenerateAdultSelector(rng)
	} else if age < 18 {
		return h.occupationProfile.GenerateUnderSelector(rng)
	}
	return h.occupationProfile.GenerateOverSelector(rng)
}

func (h *HouseProfile) GenerateNumberOfSanitaryDevices(rng *rand.Rand, numberOfResidents uint8) error {
	return h.numSanitarysDevice.GenerateData(rng,numberOfResidents)
	
}

func (h *HouseProfile) GetNumberOfSanitaryDevices() uint8 {
	return h.numSanitarysDevice.GetSanitaryCount()
}
func (h *HouseProfile) GenerateSanitaryHouse(devices map[string]uint32, amount uint8) (error) {
	return nil//housedata.NewSanitaryHouse(devices,amount),nil *housedata.SanitaryHouse, 
}
