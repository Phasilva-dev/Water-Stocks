package houseprofiles

import (
	"errors"
	"math/rand/v2"
)

type SanitaryCountProfile struct { }

func NewSanitaryCountProfile() *SanitaryCountProfile {
	return &SanitaryCountProfile{
	}
}

func (s *SanitaryCountProfile) GenerateData(rng *rand.Rand, numResidents uint8) (uint8, error) {
	if numResidents <= 0 {
		return 0,errors.New("house without residents")
	}

	// Gera um número aleatório inteiro entre 0 e 99 (inclusive)
	percent := rng.IntN(100)

	// O switch agora é baseado no número de residentes para escolher a distribuição
	switch numResidents {
	case 1:
		if percent < 79 { // 0-78 (79% de chance)
			return 1,nil
		} else if percent < 95 { // 79-94 (16% de chance)
			return 2, nil
		} else if percent < 99 { // 95-98 (4% de chance)
			return 3, nil
		} else { // 99 (1% de chance)
			return 4, nil
		}
	case 2:
		if percent < 70 { // 0-69
			return 1, nil
		} else if percent < 92 { // 70-91
			return 2, nil
		} else if percent < 98 { // 92-97
			return 3, nil
		} else { // 98-99
			return 4, nil
		}
	case 3:
		if percent < 69 { // 0-68
			return 1, nil
		} else if percent < 91 { // 69-90
			return 2, nil
		} else if percent < 98 { // 91-97
			return 3, nil
		} else { // 98-99
			return 4, nil
		}
	case 4:
		if percent < 65 { // 0-64
			return 1, nil
		} else if percent < 89 { // 65-88
			return 2, nil
		} else if percent < 97 { // 89-96
			return 3, nil
		} else { // 97-99
			return 4, nil
		}
	case 5:
		if percent < 67 { // 0-66
			return 1, nil
		} else if percent < 90 { // 67-89
			return 2, nil
		} else if percent < 97 { // 90-96
			return 3, nil
		} else { // 97-99
			return 4, nil
		}
	case 6:
		if percent < 69 { // 0-68
			return 1, nil
		} else if percent < 91 { // 69-90
			return 2, nil
		} else if percent < 97 { // 91-96
			return 3, nil
		} else { // 97-99
			return 4, nil
		}
	case 7:
		if percent < 69 { // 0-68
			return 1, nil
		} else if percent < 92 { // 69-91
			return 2, nil
		} else if percent < 97 { // 92-96
			return 3, nil
		} else { // 97-99
			return 4, nil
		}
	case 8:
		if percent < 70 { // 0-69
			return 1, nil
		} else if percent < 91 { // 70-90
			return 2, nil
		} else if percent < 97 { // 91-96
			return 3, nil
		} else { // 97-99
			return 4, nil
		}
	case 9:
		if percent < 70 { // 0-69
			return 1, nil
		} else if percent < 91 { // 70-90
			return 2, nil
		} else if percent < 97 { // 91-96
			return 3, nil
		} else { // 97-99
			return 4, nil
		}
	case 10:
		if percent < 69 { // 0-68
			return 1, nil
		} else if percent < 90 { // 69-89
			return 2, nil
		} else if percent < 97 { // 90-96
			return 3, nil
		} else { // 97-99
			return 4, nil
		}
	case 11:
		if percent < 69 { // 0-68
			return 1, nil
		} else if percent < 90 { // 69-89
			return 2, nil
		} else if percent < 96 { // 90-95
			return 3, nil
		} else { // 96-99
			return 4, nil
		}
	case 12:
		if percent < 67 { // 0-66
			return 1, nil
		} else if percent < 88 { // 67-87
			return 2, nil
		} else if percent < 95 { // 88-94
			return 3, nil
		} else { // 95-99
			return 4, nil
		}
	case 13:
		if percent < 66 { // 0-65
			return 1, nil
		} else if percent < 87 { // 66-86
			return 2, nil
		} else if percent < 95 { // 87-94
			return 3, nil
		} else { // 95-99
			return 4, nil
		}
	default: // para numResidents >14 (ou inválido, se tratado acima)
		if percent < 63 { // 0-62
			return 1, nil
		} else if percent < 84 { // 63-83
			return 2, nil
		} else if percent < 92 { // 84-91
			return 3, nil
		} else { // 92-99
			return 4, nil
		}
	}
}
