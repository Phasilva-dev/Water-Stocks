package house

import (

	"math/rand/v2"
	"simulation/internal/entities/house/ds/sanitarysystem"
	"simulation/internal/entities/house/profile/count"
	"simulation/internal/entities/house/profile/demographics"
	"simulation/internal/entities/house/profile/sanitarydevice"
	"simulation/internal/entities/resident"
	"errors"

)

var (
	ErrNilResidentCountProfile  = errors.New("resident count profile cannot be nil")
	ErrNilAgeProfile            = errors.New("age profile cannot be nil")
	ErrNilOccupationProfile     = errors.New("occupation profile cannot be nil")
	ErrNilSanitaryCountProfile  = errors.New("sanitary count profile cannot be nil")
	IdInvalid = errors.New("0 is a invalid ID")
)

type HouseProfile struct {
	houseClassID uint32
	
	numResidentsProfile *count.ResidentCount
	ageProfile *demographics.Age

	occupationProfile *demographics.Occupation
	numSanitarysDevice *count.SanitaryCount
	
	residentprofiles map[uint32]*resident.ResidentProfile

	// No futuro, seria bom ter um profile para decidir as chances de uma casa ter cada tipo de sanitaryDevice*/ 
	
}

func NewHouseProfile(
	houseClassID uint32,
	numResidentsProfile *count.ResidentCount,
	ageProfile *demographics.Age,
	occupationProfile *demographics.Occupation,
	numSanitarysDevice *count.SanitaryCount,
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
	if houseClassID == 0 {
		return nil, IdInvalid
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

func (h *HouseProfile) NumResidentsProfile() *count.ResidentCount {
	return h.numResidentsProfile
}

func (h *HouseProfile) AgeProfile() *demographics.Age {
	return h.ageProfile
}

func (h *HouseProfile) OccupationProfile() *demographics.Occupation {
	return h.occupationProfile
}

func (h *HouseProfile) NumSanitarysDevice() *count.SanitaryCount {
	return h.numSanitarysDevice
}

func (h *HouseProfile) ResidentProfile(ID uint32) (*resident.ResidentProfile, error) {
	p := h.residentprofiles[ID]
	if p != nil {
		return p,nil
	}
	return nil, errors.New("missing resident profile in house profile")
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
		return h.occupationProfile.GenerateUnder18Selector(rng)
	}
	return h.occupationProfile.GenerateOver65Selector(rng)
}

func (h *HouseProfile) GenerateNumberOfSanitaryDevices(rng *rand.Rand, numberOfResidents uint8) (uint8,error) {
	return h.numSanitarysDevice.GenerateData(rng,numberOfResidents)
	
}

func (h *HouseProfile) GenerateSanitaryHouse(devices map[string]sanitarydevice.SanitaryDevice, amount uint8) (*sanitarysystem.SanitaryHouse,error) {
	return sanitarysystem.NewSanitaryHouse(devices,amount)
}
