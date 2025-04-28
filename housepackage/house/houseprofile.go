package house

import (
	"houseprofiles"
	"math/rand/v2"
	"housedata"
)

type HouseProfile struct {
	houseClassID uint32
	numResidentsProfile *houseprofiles.ResidentCountProfile
	ageProfile *houseprofiles.AgeProfile
	numSanitarysDevice *houseprofiles.SanitaryCountProfile
	sanitaryTypeProfile *houseprofiles.SanitaryTypeProfile
	
	
}

func (h *HouseProfile) GenerateNumbersOfResidents(rng *rand.Rand) uint8 {
	return h.numResidentsProfile.GenerateData(rng)
}

func (h *HouseProfile) GenerateAgeofResidents(rng *rand.Rand) uint8 {
	return uint8(h.ageProfile.GenerateData(rng))
}

func (h *HouseProfile) GenerateNumberOfSanitaryDevices(rng *rand.Rand, numberOfResidents uint8) error {
	return h.numSanitarysDevice.GenerateData(rng,numberOfResidents)
	
}

func (h *HouseProfile) GetNumberOfSanitaryDevices() uint8 {
	return h.numSanitarysDevice.GetSanitaryCount()
}
func (h *HouseProfile) GenerateSanitaryHouse(rng *rand.Rand, amount uint8) (*housedata.SanitaryHouse, error) {
	return h.sanitaryTypeProfile.GenerateData(rng,amount)
}
