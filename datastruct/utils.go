package datastruct

// ProfileTuple representa tempos de saída e retorno
type ProfileTuple struct {
	entryTime  int32
	exitTime int32
}

// NewProfileTuple cria uma nova instância de ProfileTuple
func NewProfileTuple(entryTime, exitTime int32) ProfileTuple {
	return ProfileTuple{
		entryTime:  entryTime,
		exitTime: exitTime,
	}
}

func (p *ProfileTuple) EntryTime() int32 {
	return p.entryTime
}

func (p *ProfileTuple) ExitTime() int32 {
	return p.exitTime
}

