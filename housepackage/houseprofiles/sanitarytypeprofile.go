package houseprofiles

import (
	"misc"
	"sanitarydevice"
	"housedata"
	"math/rand/v2"
)

type SanitaryTypeProfile struct {
	toilet *misc.PercentSelector[*sanitarydevice.SanitaryDevice]
	shower *misc.PercentSelector[*sanitarydevice.SanitaryDevice]
	washBassin *misc.PercentSelector[*sanitarydevice.SanitaryDevice]

	washMachine *misc.PercentSelector[*sanitarydevice.SanitaryDevice]
	dishWasher *misc.PercentSelector[*sanitarydevice.SanitaryDevice]
	tanque *misc.PercentSelector[*sanitarydevice.SanitaryDevice]
}

func (s *SanitaryTypeProfile) Toilet() *misc.PercentSelector[*sanitarydevice.SanitaryDevice] {
	return s.toilet
}

func (s *SanitaryTypeProfile) Shower() *misc.PercentSelector[*sanitarydevice.SanitaryDevice] {
	return s.shower
}

func (s *SanitaryTypeProfile) WashBassin() *misc.PercentSelector[*sanitarydevice.SanitaryDevice] {
	return s.washBassin
}

func (s *SanitaryTypeProfile) WashMachine() *misc.PercentSelector[*sanitarydevice.SanitaryDevice] {
	return s.washMachine
}

func (s *SanitaryTypeProfile) DishWasher() *misc.PercentSelector[*sanitarydevice.SanitaryDevice] {
	return s.dishWasher
}

func (s *SanitaryTypeProfile) Tanque() *misc.PercentSelector[*sanitarydevice.SanitaryDevice] {
	return s.tanque
}


func NewSanitaryTypeProfile(selectors map[string]*misc.PercentSelector[*sanitarydevice.SanitaryDevice]) *SanitaryTypeProfile {
	return &SanitaryTypeProfile{
		toilet:      selectors["toilet"],
		shower:      selectors["shower"],
		washBassin:  selectors["washbassin"],
		washMachine: selectors["washmachine"],
		dishWasher:  selectors["dishwasher"],
		tanque:      selectors["tanque"],
	}
}

func (s *SanitaryTypeProfile) generateOne(
	selector *misc.PercentSelector[*sanitarydevice.SanitaryDevice], rng *rand.Rand,
) (*sanitarydevice.SanitaryDevice, error) {
	if selector == nil {
		return nil, nil
	}

	device, err := selector.Sample(rng)
	if err != nil {
		return nil, err
	}

	return device, nil
}

func (s *SanitaryTypeProfile) GenerateData(rng *rand.Rand, amount uint8) (*housedata.SanitaryHouse, error) {

	devices := make(map[string]*sanitarydevice.SanitaryDevice)

	var err error
	if devices["toilet"], err = s.generateOne(s.toilet, rng); err != nil {
		return nil, err
	}

	if devices["shower"], err = s.generateOne(s.shower, rng); err != nil {
		return nil, err
	}

	if devices["washBassin"], err = s.generateOne(s.washBassin, rng); err != nil {
		return nil, err
	}

	if devices["washMachine"], err = s.generateOne(s.washMachine, rng); err != nil {
		return nil, err
	}

	if devices["dishWasher"], err = s.generateOne(s.dishWasher, rng); err != nil {
		return nil, err
	}

	if devices["tanque"], err = s.generateOne(s.tanque, rng); err != nil {
		return nil, err
	}

	return housedata.NewSanitaryHouse(devices, amount), nil
}

