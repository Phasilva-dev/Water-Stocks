package count

import (
	"errors"
	"math/rand/v2"
)

// SanitaryCount é uma estrutura que fornece um método para gerar a contagem de usos sanitários.
// Atualmente, não possui campos, mas está projetada para permitir a adição futura de modelos
// ou parâmetros que influenciem a lógica de geração de dados.
type SanitaryCount struct{}

// NewSanitaryCount cria e retorna uma nova instância de SanitaryCount.
// Este construtor permite a futura extensão da estrutura com parâmetros de modelo.
func NewSanitaryCount() *SanitaryCount {
	return &SanitaryCount{}
}

// GenerateData gera uma contagem de usos sanitários baseada no número de residentes na casa.
// A lógica interna usa um modelo probabilístico discreto que varia de acordo com `numResidents`.
//
// rng: O gerador de números aleatórios a ser usado para a amostragem das probabilidades.
// numResidents: O número de residentes na casa (deve ser > 0).
//
// Retorna a contagem de usos sanitários (uint8) e um erro se `numResidents` for 0 ou menor.
// A contagem de usos varia entre 1 e 4, conforme as probabilidades definidas.
func (s *SanitaryCount) GenerateData(rng *rand.Rand, numResidents uint8) (uint8, error) {
	if numResidents <= 0 {
		return 0, errors.New("invalid sanitaryCount data: house without residents") // Retorna erro se não houver residentes.
	}

	// Gera um número inteiro pseudo-aleatório entre 0 e 99 (inclusive),
	// representando um percentil para a distribuição de probabilidades.
	percent := rng.IntN(100)

	// O switch define diferentes distribuições de probabilidade para a contagem de usos
	// sanitários (1, 2, 3 ou 4 usos), baseadas no número de residentes.
	switch numResidents {
	case 1:
		if percent < 79 { // 79% de chance (0-78)
			return 1, nil
		} else if percent < 95 { // 16% de chance (79-94)
			return 2, nil
		} else if percent < 99 { // 4% de chance (95-98)
			return 3, nil
		} else { // 1% de chance (99)
			return 4, nil
		}
	case 2:
		if percent < 70 { // 70% de chance (0-69)
			return 1, nil
		} else if percent < 92 { // 22% de chance (70-91)
			return 2, nil
		} else if percent < 98 { // 6% de chance (92-97)
			return 3, nil
		} else { // 2% de chance (98-99)
			return 4, nil
		}
	case 3:
		if percent < 69 { // 69% de chance (0-68)
			return 1, nil
		} else if percent < 91 { // 22% de chance (69-90)
			return 2, nil
		} else if percent < 98 { // 7% de chance (91-97)
			return 3, nil
		} else { // 2% de chance (98-99)
			return 4, nil
		}
	case 4:
		if percent < 65 { // 65% de chance (0-64)
			return 1, nil
		} else if percent < 89 { // 24% de chance (65-88)
			return 2, nil
		} else if percent < 97 { // 8% de chance (89-96)
			return 3, nil
		} else { // 3% de chance (97-99)
			return 4, nil
		}
	case 5:
		if percent < 67 { // 67% de chance (0-66)
			return 1, nil
		} else if percent < 90 { // 23% de chance (67-89)
			return 2, nil
		} else if percent < 97 { // 7% de chance (90-96)
			return 3, nil
		} else { // 3% de chance (97-99)
			return 4, nil
		}
	case 6:
		if percent < 69 { // 69% de chance (0-68)
			return 1, nil
		} else if percent < 91 { // 22% de chance (69-90)
			return 2, nil
		} else if percent < 97 { // 6% de chance (91-96)
			return 3, nil
		} else { // 3% de chance (97-99)
			return 4, nil
		}
	case 7:
		if percent < 69 { // 69% de chance (0-68)
			return 1, nil
		} else if percent < 92 { // 23% de chance (69-91)
			return 2, nil
		} else if percent < 97 { // 5% de chance (92-96)
			return 3, nil
		} else { // 3% de chance (97-99)
			return 4, nil
		}
	case 8:
		if percent < 70 { // 70% de chance (0-69)
			return 1, nil
		} else if percent < 91 { // 21% de chance (70-90)
			return 2, nil
		} else if percent < 97 { // 6% de chance (91-96)
			return 3, nil
		} else { // 3% de chance (97-99)
			return 4, nil
		}
	case 9:
		if percent < 70 { // 70% de chance (0-69)
			return 1, nil
		} else if percent < 91 { // 21% de chance (70-90)
			return 2, nil
		} else if percent < 97 { // 6% de chance (91-96)
			return 3, nil
		} else { // 3% de chance (97-99)
			return 4, nil
		}
	case 10:
		if percent < 69 { // 69% de chance (0-68)
			return 1, nil
		} else if percent < 90 { // 21% de chance (69-89)
			return 2, nil
		} else if percent < 97 { // 7% de chance (90-96)
			return 3, nil
		} else { // 3% de chance (97-99)
			return 4, nil
		}
	case 11:
		if percent < 69 { // 69% de chance (0-68)
			return 1, nil
		} else if percent < 90 { // 21% de chance (69-89)
			return 2, nil
		} else if percent < 96 { // 6% de chance (90-95)
			return 3, nil
		} else { // 4% de chance (96-99)
			return 4, nil
		}
	case 12:
		if percent < 67 { // 67% de chance (0-66)
			return 1, nil
		} else if percent < 88 { // 21% de chance (67-87)
			return 2, nil
		} else if percent < 95 { // 7% de chance (88-94)
			return 3, nil
		} else { // 5% de chance (95-99)
			return 4, nil
		}
	case 13:
		if percent < 66 { // 66% de chance (0-65)
			return 1, nil
		} else if percent < 87 { // 21% de chance (66-86)
			return 2, nil
		} else if percent < 95 { // 8% de chance (87-94)
			return 3, nil
		} else { // 5% de chance (95-99)
			return 4, nil
		}
	default: // Caso `numResidents` seja maior que 13.
		if percent < 63 { // 63% de chance (0-62)
			return 1, nil
		} else if percent < 84 { // 21% de chance (63-83)
			return 2, nil
		} else if percent < 92 { // 8% de chance (84-91)
			return 3, nil
		} else { // 8% de chance (92-99)
			return 4, nil
		}
	}
}