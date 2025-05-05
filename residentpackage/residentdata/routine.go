// Package residentdata define estruturas de dados para armazenar informações
// relacionadas a residentes ou residências, como frequências de uso de água
// e sequências de eventos de rotina diária.
package residentdata

import () // Bloco de import mantido, pode ser ajustado por goimports.

// Routine armazena uma sequência ordenada de timestamps (`int32`) que
// representam os horários dos eventos chave na rotina diária de um residente
// (ex: acordar, sair, retornar, dormir).
//
// O acesso aos dados deve ser feito preferencialmente através dos métodos
// fornecidos (ex: WakeupTime(), SleepTime(), Times()).
type Routine struct {
	// times é um slice interno (não exportado) que contém os timestamps
	// dos eventos da rotina em ordem cronológica.
	times []int32
}

// NewRoutine cria e retorna uma nova instância de Routine.
//
// Recebe um slice de `int32` (`times`) que representa a sequência ordenada
// dos horários dos eventos da rotina.
//
// ATENÇÃO: Esta função armazena a referência direta ao slice `times` fornecido,
// *não* cria uma cópia. Modificações no slice `times` original *após* a
// chamada de NewRoutine afetarão a instância de Routine criada.
// Se a imutabilidade for necessária, passe uma cópia do slice para esta função.
func NewRoutine(times []int32) *Routine {
	return &Routine{
		times: times, // Armazena a referência ao slice original.
	}
}

// Times retorna o slice completo de timestamps (`int32`) que compõem a rotina.
//
// ATENÇÃO: Este método retorna uma referência direta ao slice interno `times`.
// Modificações no slice retornado afetarão diretamente o estado interno da
// instância de Routine. Se precisar dos dados sem risco de modificar o
// original, crie uma cópia do slice retornado.
func (r *Routine) Times() []int32 {
	return r.times
}

// SleepTime retorna o timestamp (`int32`) do último evento na rotina,
// que é assumido como o horário de dormir.
//
// Esta função pressupõe que o slice `times` não está vazio e que o último
// elemento sempre representa o horário de dormir. Causa pânico se `times` for vazio.
func (r *Routine) SleepTime() int32 {
	// Retorna o último elemento do slice.
	return r.times[len(r.times)-1]
}

// WakeupTime retorna o timestamp (`int32`) do primeiro evento na rotina,
// que é assumido como o horário de acordar.
//
// Esta função pressupõe que o slice `times` não está vazio.
// Causa pânico se `times` for vazio.
func (r *Routine) WakeupTime() int32 {
	// Retorna o primeiro elemento do slice.
	return r.times[0]
}

//Tudo abaixo desse comentario é basicamente um mock para seguir um modelo criado em MATLAB
//Caso queira se aventurar e criar seu proprio modelo, fique a vontade!
//Gostaria queisso tivesse sido um software de analise exploratoria de dados :( Por isso já fiz permitindo parametros.

// ReturnHome retorna o timestamp (`int32`) do terceiro evento na rotina
// (índice 2), que é assumido como o horário de retorno para casa.
//
// Esta função pressupõe que o slice `times` contém pelo menos 3 elementos
// e que o elemento no índice 2 representa o horário de retorno para casa.
// Causa pânico se `len(times) < 3`.
func (r *Routine) ReturnHome() int32 {
	// Retorna o terceiro elemento (índice 2).
	return r.times[2]
}

// WorkTime retorna o timestamp (`int32`) do segundo evento na rotina
// (índice 1), que é assumido como o horário de saída para o trabalho
// (ou início da atividade principal fora de casa).
//
// Esta função pressupõe que o slice `times` contém pelo menos 2 elementos
// e que o elemento no índice 1 representa o horário de saída.
// Causa pânico se `len(times) < 2`.
func (r *Routine) WorkTime() int32 {
	// Retorna o segundo elemento (índice 1).
	return r.times[1]
}



/* Código comentado mantido do original:
func (r *Routine) EntryHomeTime(index uint8) int32{
	return r.times[index]
}

func (r *Routine) ExitHomeTime(index uint8) int32{
	return r.times[index]
}*/