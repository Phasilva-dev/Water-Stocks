package temporal

import (
	"simulation/internal/entities/resident/ds/behavioral"

)

type DailyData struct {
	routine *behavioral.Routine
	frequency *behavioral.Frequency
}

func NewDailyData(routine *behavioral.Routine, frequency *behavioral.Frequency) *DailyData {
	return &DailyData{
		routine: routine,
		frequency: frequency,
	}
}

func (d *DailyData) Routine() *behavioral.Routine {
	return d.routine
}

func (d *DailyData) Frequency() *behavioral.Frequency {
	return d.frequency
}


func (d *DailyData) SetRoutine(r *behavioral.Routine) {
	d.routine = r
}

func (d *DailyData) SetFrequency(f *behavioral.Frequency) {
	d.frequency = f
}

func (d *DailyData) ClearData() {
	d.SetFrequency(nil)
	d.SetRoutine(nil)
}
