package datastruct

import (
	"unique"
	
)

type Routine struct {
	symbol unique.Handle[string]
	times []int32
}

func NewRoutine(symbol string, times []int32) *Routine {
	return &Routine{
		symbol: unique.Make(symbol), // Cria um handle único para o símbolo
		times:  times,               // Inicializa o slice de ProfileTuple
	}
}

func (r *Routine) SleepTime() int32 {
	return r.times[len(r.times)-1]
}

func (r *Routine) WakeupTime() int32 {
	return r.times[0]
}


