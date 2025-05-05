package residentdata

import (

)

type DailyData struct {
	routine *Routine
	frequency *Frequency
}

func NewDailyData(routine *Routine, frequency *Frequency) *DailyData {
	return &DailyData{
		routine: routine,
		frequency: frequency,
	}
}

func (d *DailyData) Routine() *Routine {
	return d.routine
}

func (d *DailyData) Frequency() *Frequency {
	return d.frequency
}


func (d *DailyData) SetRoutine(r *Routine) {
	d.routine = r
}

func (d *DailyData) SetFrequency(f *Frequency) {
	d.frequency = f
}
