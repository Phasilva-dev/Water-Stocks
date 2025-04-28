package house

import (
	"housedata"
	"resident"
)

type House struct {
	houseClassID uint32
	residents []resident.Resident
	sanitaryHouse *housedata.SanitaryHouse
	houseProfile *HouseProfile
}

func NewHouse(houseClassID uint32, houseProfile *HouseProfile) *House {
	return &House{
		houseClassID:  houseClassID,
		residents:     []resident.Resident{},
		sanitaryHouse: nil,
		houseProfile:  houseProfile,
	}
}

func 