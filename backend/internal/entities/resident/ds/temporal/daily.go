// Package temporal define estruturas para armazenar e gerenciar dados simulados
// diários, tipicamente para um residente, incluindo sua rotina e frequência de uso de água.
package temporal

import (
	"simulation/internal/entities/resident/ds/behavioral"
)

// DailyData encapsula os dados comportamentais de um residente para um único dia.
// Isso inclui sua rotina diária de presença e a frequência de uso de diferentes
// aparelhos de água.
type DailyData struct {
	routine   *behavioral.Routine   // A rotina diária do residente, incluindo horários de eventos.
	frequency *behavioral.Frequency // A frequência diária de uso de aparelhos de água.
}

// NewDailyData cria e retorna uma nova instância de DailyData.
//
// routine: O perfil de rotina comportamental para o dia.
// frequency: O perfil de frequência de uso de água para o dia.
//
// Retorna um ponteiro para a estrutura DailyData preenchida com os dados fornecidos.
func NewDailyData(routine *behavioral.Routine, frequency *behavioral.Frequency) *DailyData {
	return &DailyData{
		routine:   routine,
		frequency: frequency,
	}
}

// Routine retorna o perfil de rotina comportamental para o dia.
// Pode retornar nil se não houver um perfil de rotina definido.
func (d *DailyData) Routine() *behavioral.Routine {
	return d.routine
}

// Frequency retorna o perfil de frequência de uso de água para o dia.
// Pode retornar nil se não houver um perfil de frequência definido.
func (d *DailyData) Frequency() *behavioral.Frequency {
	return d.frequency
}

// SetRoutine define o perfil de rotina comportamental para o dia.
//
// r: O novo perfil de rotina a ser definido. Pode ser nil para limpar o perfil existente.
func (d *DailyData) SetRoutine(r *behavioral.Routine) {
	d.routine = r
}

// SetFrequency define o perfil de frequência de uso de água para o dia.
//
// f: O novo perfil de frequência a ser definido. Pode ser nil para limpar o perfil existente.
func (d *DailyData) SetFrequency(f *behavioral.Frequency) {
	d.frequency = f
}

// ClearData redefine os perfis de rotina e frequência para nil.
// Isso libera as referências aos dados comportamentais, efetivamente 'limpando'
// os dados diários desta instância.
func (d *DailyData) ClearData() {
	d.SetFrequency(nil)
	d.SetRoutine(nil)
}