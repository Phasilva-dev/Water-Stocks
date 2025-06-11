package house

import (
	"errors"
	"log"
	"math/rand/v2"
	"simulation/internal/entities/house/ds/sanitarysystem"
	"simulation/internal/entities/house/profile/count"
	"simulation/internal/entities/house/profile/demographics"
	"simulation/internal/entities/house/profile/sanitarydevice"
	"simulation/internal/entities/resident"
)

var (
	ErrNilResidentCountProfile  = errors.New("resident count profile cannot be nil")
	ErrNilAgeProfile            = errors.New("age profile cannot be nil")
	ErrNilOccupationProfile     = errors.New("occupation profile cannot be nil")
	ErrNilSanitaryCountProfile  = errors.New("sanitary count profile cannot be nil")
	ErrIdInvalid = errors.New("0 is a invalid ID")
	ErrNilSanitaryDevices = errors.New("should be have a sanitarydevice")
	ErrNilResidentProfiles = errors.New("should be have almost one residentProfile")
)

type HouseProfile struct {
	houseClassID uint32
	
	numResidentsProfile *count.ResidentCount
	ageProfile *demographics.Age

	occupationProfile *demographics.Occupation
	numSanitarysDevice *count.SanitaryCount
	
	residentprofiles map[uint32]*resident.ResidentProfile

	sanitaryDevices map[string]sanitarydevice.SanitaryDevice

	// No futuro, seria bom ter um profile para decidir as chances de uma casa ter cada tipo de sanitaryDevice*/ 
	
}

func NewHouseProfile(
	houseClassID uint32,
	numResidentsProfile *count.ResidentCount,
	ageProfile *demographics.Age,
	occupationProfile *demographics.Occupation,
	numSanitarysDevice *count.SanitaryCount,
	residentProfiles map[uint32]*resident.ResidentProfile,
	sanitaryDevices map[string]sanitarydevice.SanitaryDevice,

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
		return nil, ErrIdInvalid
	}
	if len(sanitaryDevices) == 0 {
		return nil, ErrNilSanitaryDevices
	}
	if len(residentProfiles) == 0{
		return nil, ErrNilResidentProfiles
	}

	return &HouseProfile{
		houseClassID:        houseClassID,
		numResidentsProfile: numResidentsProfile,
		ageProfile:          ageProfile,
		occupationProfile:   occupationProfile,
		numSanitarysDevice:  numSanitarysDevice,
		residentprofiles: residentProfiles,
		sanitaryDevices: sanitaryDevices,
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

func (h *HouseProfile) SanitaryDevices() map[string]sanitarydevice.SanitaryDevice {
	return h.sanitaryDevices
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

func (h *HouseProfile) GenerateOccupation(age uint8, rng *rand.Rand) (uint32, error) {
	occupation, err := h.occupationProfile.Sample(age,rng)
	if err != nil {
		log.Printf("Erro para idade %d: %v", age, err)
		return 0, err
	}
	return occupation, nil
}

func (h *HouseProfile) GenerateNumberOfSanitaryDevices(rng *rand.Rand, numberOfResidents uint8) (uint8,error) {
	return h.numSanitarysDevice.GenerateData(rng,numberOfResidents)
	
}

func (h *HouseProfile) GenerateSanitaryHouse(amount uint8) (*sanitarysystem.SanitaryHouse,error) {
	return sanitarysystem.NewSanitaryHouse(h.sanitaryDevices,amount)
}
