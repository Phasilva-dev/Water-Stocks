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