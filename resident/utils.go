package resident

// ProfileTuple representa tempos de saída e retorno
type ProfileTuple struct {
	leaveTime  int32
	returnTime int32
}

// NewProfileTuple cria uma nova instância de ProfileTuple
func NewProfileTuple(leaveTime, returnTime int32) *ProfileTuple {
	return &ProfileTuple{
		leaveTime:  leaveTime,
		returnTime: returnTime,
	}
}

func (p *ProfileTuple) LeaveTime() int32 {
	return p.leaveTime
}

func (p *ProfileTuple) ReturnTime() int32 {
	return p.returnTime
}