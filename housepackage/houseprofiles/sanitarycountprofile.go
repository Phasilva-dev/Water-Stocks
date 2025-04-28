package houseprofiles

import (
	"errors"
	"math/rand/v2"
)

type SanitaryCountProfile struct {
	sanitaryNum uint8
}

func NewSanitaryCountProfile() *SanitaryCountProfile {
	return &SanitaryCountProfile{
		sanitaryNum: 0,
	}
}

func (s *SanitaryCountProfile) GenerateData(rng *rand.Rand, numResidents uint8) error {
	if numResidents <= 0 {
		return errors.New("house without residents")
	}

	// Gera um número aleatório inteiro entre 0 e 99 (inclusive)
	percent := rng.IntN(100)

	// O switch agora é baseado no número de residentes para escolher a distribuição
	switch numResidents {
	case 1:
		if percent < 79 { // 0-78 (79% de chance)
			s.sanitaryNum = 1
		} else if percent < 95 { // 79-94 (16% de chance)
			s.sanitaryNum = 2
		} else if percent < 99 { // 95-98 (4% de chance)
			s.sanitaryNum = 3
		} else { // 99 (1% de chance)
			s.sanitaryNum = 4
		}
	case 2:
		if percent < 70 { // 0-69
			s.sanitaryNum = 1
		} else if percent < 92 { // 70-91
			s.sanitaryNum = 2
		} else if percent < 98 { // 92-97
			s.sanitaryNum = 3
		} else { // 98-99
			s.sanitaryNum = 4
		}
	case 3:
		if percent < 69 { // 0-68
			s.sanitaryNum = 1
		} else if percent < 91 { // 69-90
			s.sanitaryNum = 2
		} else if percent < 98 { // 91-97
			s.sanitaryNum = 3
		} else { // 98-99
			s.sanitaryNum = 4
		}
	case 4:
		if percent < 65 { // 0-64
			s.sanitaryNum = 1
		} else if percent < 89 { // 65-88
			s.sanitaryNum = 2
		} else if percent < 97 { // 89-96
			s.sanitaryNum = 3
		} else { // 97-99
			s.sanitaryNum = 4
		}
	case 5:
		if percent < 67 { // 0-66
			s.sanitaryNum = 1
		} else if percent < 90 { // 67-89
			s.sanitaryNum = 2
		} else if percent < 97 { // 90-96
			s.sanitaryNum = 3
		} else { // 97-99
			s.sanitaryNum = 4
		}
	case 6:
		if percent < 69 { // 0-68
			s.sanitaryNum = 1
		} else if percent < 91 { // 69-90
			s.sanitaryNum = 2
		} else if percent < 97 { // 91-96
			s.sanitaryNum = 3
		} else { // 97-99
			s.sanitaryNum = 4
		}
	case 7:
		if percent < 69 { // 0-68
			s.sanitaryNum = 1
		} else if percent < 92 { // 69-91
			s.sanitaryNum = 2
		} else if percent < 97 { // 92-96
			s.sanitaryNum = 3
		} else { // 97-99
			s.sanitaryNum = 4
		}
	case 8:
		if percent < 70 { // 0-69
			s.sanitaryNum = 1
		} else if percent < 91 { // 70-90
			s.sanitaryNum = 2
		} else if percent < 97 { // 91-96
			s.sanitaryNum = 3
		} else { // 97-99
			s.sanitaryNum = 4
		}
	case 9:
		if percent < 70 { // 0-69
			s.sanitaryNum = 1
		} else if percent < 91 { // 70-90
			s.sanitaryNum = 2
		} else if percent < 97 { // 91-96
			s.sanitaryNum = 3
		} else { // 97-99
			s.sanitaryNum = 4
		}
	case 10:
		if percent < 69 { // 0-68
			s.sanitaryNum = 1
		} else if percent < 90 { // 69-89
			s.sanitaryNum = 2
		} else if percent < 97 { // 90-96
			s.sanitaryNum = 3
		} else { // 97-99
			s.sanitaryNum = 4
		}
	case 11:
		if percent < 69 { // 0-68
			s.sanitaryNum = 1
		} else if percent < 90 { // 69-89
			s.sanitaryNum = 2
		} else if percent < 96 { // 90-95
			s.sanitaryNum = 3
		} else { // 96-99
			s.sanitaryNum = 4
		}
	case 12:
		if percent < 67 { // 0-66
			s.sanitaryNum = 1
		} else if percent < 88 { // 67-87
			s.sanitaryNum = 2
		} else if percent < 95 { // 88-94
			s.sanitaryNum = 3
		} else { // 95-99
			s.sanitaryNum = 4
		}
	case 13:
		if percent < 66 { // 0-65
			s.sanitaryNum = 1
		} else if percent < 87 { // 66-86
			s.sanitaryNum = 2
		} else if percent < 95 { // 87-94
			s.sanitaryNum = 3
		} else { // 95-99
			s.sanitaryNum = 4
		}
	default: // para numResidents >= 14 (ou inválido, se tratado acima)
		if percent < 63 { // 0-62
			s.sanitaryNum = 1
		} else if percent < 84 { // 63-83
			s.sanitaryNum = 2
		} else if percent < 92 { // 84-91
			s.sanitaryNum = 3
		} else { // 92-99
			s.sanitaryNum = 4
		}
	}
	return nil
}

func (s *SanitaryCountProfile) GetSanitaryCount() uint8 {
	return s.sanitaryNum
}