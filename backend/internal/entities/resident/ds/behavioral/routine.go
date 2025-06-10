// Package behavioral define estruturas para modelar a rotina diária de presença
// de residentes em casa, com base em eventos-chave ao longo do dia.
package behavioral

import () // Pode ser ajustado automaticamente por goimports.

// Routine representa uma sequência de horários (timestamps) de eventos diários
// na rotina de um residente, como acordar, sair, retornar e dormir.
type Routine struct {
	times []float64 // Timestamps em ordem cronológica.
}

// NewRoutine cria uma nova instância de Routine com os horários fornecidos.
//
// Nota: o slice passado não é copiado — alterações externas ao slice original
// afetam diretamente a rotina. Se desejar imutabilidade, passe uma cópia.
func NewRoutine(times []float64) *Routine {
	return &Routine{times: times}
}

// Métodos de acesso à rotina.

// Times retorna o slice completo de horários da rotina.
// O slice retornado é o mesmo armazenado internamente.
func (r *Routine) Times() []float64 {
	return r.times
}

// WakeupTime retorna o primeiro horário da rotina (ex: acordar).
// Presume que o slice não está vazio.
func (r *Routine) WakeupTime() float64 {
	return r.times[0]
}

// WorkTime retorna o segundo horário da rotina (ex: saída para o trabalho).
// Presume ao menos dois eventos no slice.
func (r *Routine) WorkTime() float64 {
	return r.times[1]
}


// ReturnHome retorna o terceiro horário da rotina (ex: retorno para casa).
// Presume ao menos três eventos no slice.
func (r *Routine) ReturnHome() float64 {
	return r.times[2]
}

// SleepTime retorna o último horário da rotina (ex: dormir).
// Presume que o slice não está vazio.
func (r *Routine) SleepTime() float64 {
	return r.times[len(r.times)-1]
}

/*
Tanto ReturnHome quanto SleepTime foram feitos apenas para seguir um modelo feito previamente em MATLAB
Observação: As funções acima seguem uma estrutura baseada em um modelo MATLAB
anterior. Elas presumem uma ordem fixa de eventos no slice `times`.

Sinta-se à vontade para adaptar esse modelo ou criar um novo baseado
em parâmetros dinâmicos — este código foi projetado com essa flexibilidade em mente.

A ideia inicial era criar um software de análise exploratória de dados,
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
