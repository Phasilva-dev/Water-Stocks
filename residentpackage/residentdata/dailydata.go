package residentdata

import (

)

type DailyData struct {
	routine *Routine
	frequency *Frequency
	usage *Usage  
}

func NewDailyData(routine *Routine, frequency *Frequency, usage *Usage) *DailyData {
	return &DailyData{
		routine: routine,
		frequency: frequency,
		usage: usage,
	}
}

func (d *DailyData) Routine() *Routine {
	return d.routine
}

func (d *DailyData) Frequency() *Frequency {
	return d.frequency
}

func (d *DailyData) Usage() *Usage {
	return d.usage
}

func (d *DailyData) SetRoutine(r *Routine) {
	d.routine = r
}

func (d *DailyData) SetFrequency(f *Frequency) {
	d.frequency = f
}

func (d *DailyData) SetUsage(u *Usage) {
	d.usage = u
}