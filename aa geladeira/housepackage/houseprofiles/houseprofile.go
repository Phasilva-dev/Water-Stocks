package houseprofiles

import (

	"math/rand/v2"
	"housedata"
	"errors"

)

var (
	ErrNilResidentCountProfile  = errors.New("resident count profile cannot be nil")
	ErrNilAgeProfile            = errors.New("age profile cannot be nil")
	ErrNilOccupationProfile     = errors.New("occupation profile cannot be nil")
	ErrNilSanitaryCountProfile  = errors.New("sanitary count profile cannot be nil")
)

type HouseProfile struct {
	houseClassID uint32

	numResidentsProfile *ResidentCountProfile
	ageProfile *AgeProfile
	occupationProfile *OccupationProfile

	numSanitarysDevice *SanitaryCountProfile
	//sanitaryTypeProfile *houseprofiles.SanitaryTypeProfile /*
	// No futuro, seria bom ter um profile para decidir as chances de uma casa ter cada tipo de sanitaryDevice*/ 
	
}

func NewHouseProfile(
	houseClassID uint32,
	numResidentsProfile *ResidentCountProfile,
	ageProfile *AgeProfile,
	occupationProfile *OccupationProfile,
	numSanitarysDevice *SanitaryCountProfile,
) (*HouseProfile, error) {
	if numResidentsProfile == nil {
		return nil, ErrNilResidentCountProfile
	}
	if ageProfile == nil {
		return nil, ErrNilAgeProfile
	}
	if occupationProfile == nil {
		return nil, ErrNilOccupationProfile
	}
	if numSanitarysDevice == nil {
		return nil, ErrNilSanitaryCountProfile
	}

	return &HouseProfile{
		houseClassID:        houseClassID,
		numResidentsProfile: numResidentsProfile,
		ageProfile:          ageProfile,
		occupationProfile:   occupationProfile,
		numSanitarysDevice:  numSanitarysDevice,
	}, nil
}

// Getters
func (h *HouseProfile) HouseClassID() uint32 {
	return h.houseClassID
}

func (h *HouseProfile) NumResidentsProfile() *ResidentCountProfile {
	return h.numResidentsProfile
}

func (h *HouseProfile) AgeProfile() *AgeProfile {
	return h.ageProfile
}

func (h *HouseProfile) OccupationProfile() *OccupationProfile {
	return h.occupationProfile
}

func (h *HouseProfile) NumSanitarysDevice() *SanitaryCountProfile {
	return h.numSanitarysDevice
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

func (h *HouseProfile) GenerateNumberOfSanitaryDevices(rng *rand.Rand, numberOfResidents uint8) (uint8,error) {
	return h.numSanitarysDevice.GenerateData(rng,numberOfResidents)
	
}

func (h *HouseProfile) GenerateSanitaryHouse(devices map[string]uint32, amount uint8) (*housedata.SanitaryHouse,error) {
	return housedata.NewSanitaryHouse(devices,amount)
}
