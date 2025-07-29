package log

import (
	"fmt"
)

type Resident struct {
	day                  uint8
	houseClassID         uint32
	residentOccupationID uint32
	age                  uint8
	sanitaryLogs         *ResidentSanitary
}

// NewResident creates a new initialized Resident instance
func NewResident(
	day uint8,
	houseClassID uint32,
	residentOccupationID uint32,
	age uint8,
	log *ResidentSanitary,
) *Resident {
	return &Resident{
		day:                  day,
		houseClassID:         houseClassID,
		residentOccupationID: residentOccupationID,
		age:                  age,
		sanitaryLogs:         log,
	}
}

// Getters
func (r *Resident) Day() uint8 {
	return r.day
}

func (r *Resident) HouseClassID() uint32 {
	return r.houseClassID
}

func (r *Resident) ResidentOccupationID() uint32 {
	return r.residentOccupationID
}

func (r *Resident) Age() uint8 {
	return r.age
}

func (r *Resident) SanitaryLogs() *ResidentSanitary {
	return r.sanitaryLogs
}





func secondsToHHMMSS(seconds int32) string {
    h := seconds / 3600
    m := (seconds % 3600) / 60
    s := seconds % 60
    return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

func (r *Resident) ToLogLines() []string {
    lines := []string{}
    sanitaryLogs := []*Sanitary{
        r.sanitaryLogs.ToiletLog(),
        r.sanitaryLogs.ShowerLog(),
        r.sanitaryLogs.WashBassinLog(),
        r.sanitaryLogs.WashMachineLog(),
        r.sanitaryLogs.DishWasherLog(),
        r.sanitaryLogs.TanqueLog(),
    }

    for _, sanitary := range sanitaryLogs {
        usages, ok := sanitary.UsageLogs()
        if !ok {
            continue // sem uso, pula
        }

        for _, usage := range usages {
            line := fmt.Sprintf("%d | %d | %d | %d | %s | %d | %s | %s | %.2f",
                r.Day(),
                r.HouseClassID(),
                r.ResidentOccupationID(),
                r.Age(),
                sanitary.SanitaryType(),
                sanitary.SanitaryDeviceID(),
                secondsToHHMMSS(usage.StartUsage()),
                secondsToHHMMSS(usage.EndUsage()),
                usage.FlowRate(),
            )
            lines = append(lines, line)
        }
    }

    return lines
}

