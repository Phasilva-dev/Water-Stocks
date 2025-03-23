package profiles

import (
	"dists"
)

// ProfileTupleDist representa distribuições de saída e retorno
type ProfileTupleDist struct {
	entryDist  dists.Distribution
	exitDist dists.Distribution
}

func (p *ProfileTupleDist) EntryDist() dists.Distribution {
	return p.entryDist
}

func (p *ProfileTupleDist) ExitDist() dists.Distribution {
	return p.exitDist
}

// NewProfileTupleDist cria uma nova instância de ProfileTupleDist
func (p *ProfileTupleDist) NewProfileTupleDist(entryDist, exitDist dists.Distribution) *ProfileTupleDist {
	return &ProfileTupleDist{
		entryDist:  entryDist,
		exitDist: exitDist,
	}
}