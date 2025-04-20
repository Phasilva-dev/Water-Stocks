package residentdata

import (
)

type Routine struct {
	times []int32
}

func NewRoutine(times []int32) *Routine {
	return &Routine{
		times:  times,               // Inicializa o slice de ProfileTuple
	}
}

func (r *Routine) Times() []int32 {
	return r.times
}

func (r *Routine) SleepTime() int32 {
	return r.times[len(r.times)-1]
}

func (r *Routine) WakeupTime() int32 {
	return r.times[0]
}

func (r *Routine) EntryHomeTime(index uint8) int32{
	return r.times[index]
}

func (r *Routine) ExitHomeTime(index uint8) int32{
	return r.times[index]
}