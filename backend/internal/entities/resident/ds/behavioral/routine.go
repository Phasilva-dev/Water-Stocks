// Package behavioral define estruturas para modelar a rotina diária de presença
// de residentes em casa, com base em eventos-chave ao longo do dia.
package behavioral

// Routine representa uma sequência de horários (timestamps em segundos desde a meia-noite)
// de eventos importantes na rotina de um residente, como acordar, sair, retornar e dormir.
// Os tempos devem estar em ordem cronológica.
type Routine struct {
	times []float64 // Timestamps dos eventos em segundos.
}

// NewRoutine cria e retorna uma nova instância de Routine com os horários fornecidos.
//
// times: Um slice de float64 representando os horários dos eventos. O slice é armazenado
//        por referência; portanto, modificações externas ao slice original afetarão esta rotina.
//        É responsabilidade do chamador garantir que os horários estejam em ordem cronológica.
//
// Retorna um ponteiro para a estrutura Routine.
func NewRoutine(times []float64) *Routine {
	return &Routine{times: times}
}

// Times retorna o slice completo de horários da rotina.
// O slice retornado é o mesmo armazenado internamente, ou seja, é uma referência.
// Modificações no slice retornado afetarão a rotina.
func (r *Routine) Times() []float64 {
	return r.times
}

func (r *Routine) EventTime(index int) float64 {
	if index >= 0 && index < len(r.times) {
		return r.times[index]
	}
	return 0

}

// WakeupTime retorna o horário do primeiro evento na rotina, tipicamente o horário de acordar.
//
// Pré-condição: O slice 'times' deve conter ao menos um elemento.
// Panics: Se o slice 'times' estiver vazio.
func (r *Routine) WakeupTime() float64 {
	return r.times[0]
}

// WorkTime retorna o horário do segundo evento na rotina, tipicamente a saída de casa para o trabalho.
//
// Pré-condição: O slice 'times' deve conter ao menos dois elementos.
// Panics: Se o slice 'times' tiver menos de dois elementos.
func (r *Routine) WorkTime() float64 {
	return r.times[1]
}

// ReturnHome retorna o horário do terceiro evento na rotina, tipicamente o retorno para casa.
//
// Pré-condição: O slice 'times' deve conter ao menos três elementos.
// Panics: Se o slice 'times' tiver menos de três elementos.
func (r *Routine) ReturnHome() float64 {
	return r.times[2]
}

// SleepTime retorna o horário do último evento na rotina, tipicamente o horário de dormir.
//
// Pré-condição: O slice 'times' deve conter ao menos um elemento.
// Panics: Se o slice 'times' estiver vazio.
func (r *Routine) SleepTime() float64 {
	return r.times[len(r.times)-1]
}

/*
As funções WakeupTime, WorkTime, ReturnHome e SleepTime foram projetadas
para seguir uma estrutura de modelo feita previamente em MATLAB. Elas assumem
uma ordem fixa e um número mínimo de eventos no slice `times` (e podem causar panic
se as pré-condições não forem atendidas).

Este código foi projetado com flexibilidade em mente para adaptações futuras
ou para a criação de novos modelos baseados em parâmetros dinâmicos. A ideia
inicial era criar um software de análise exploratória de dados,
então a estrutura está aberta para evolução!
*/

/*
Código legado comentado (exemplo de acesso genérico por índice):

func (r *Routine) EntryHomeTime(index uint8) float64 {
	return r.times[index]
}

func (r *Routine) ExitHomeTime(index uint8) float64 {
	return r.times[index]
}
*/