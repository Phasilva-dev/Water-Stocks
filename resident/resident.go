package resident

import()

type resident struct {
	ocupacao uint32
	idade uint8
	perfil_rotina //Deve buscar de um perfil global com ID de ocupação + Idade + Classe
	perfil_frequencia
	perfil_hora_uso
	rotina rotina
	frequencia frequencia
	hora_uso hora_uso
}