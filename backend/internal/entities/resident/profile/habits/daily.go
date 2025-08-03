package habits

import (
	"math/rand/v2"
	"simulation/internal/entities/resident/ds/behavioral"
	"simulation/internal/entities/resident/profile/frequency"
	"simulation/internal/entities/resident/profile/routine"
)

// ResidentDayProfile encapsula os perfis de rotina diária e de frequência de uso
// de aparelhos para um residente.
// Ele serve como um agregador para gerar os dados de comportamento diários.
type ResidentDayProfile struct {
	routineProfile      *routine.DayProfile
	frequencyProfileDay *frequency.ResidentDeviceProfiles
}

// NewResidentDayProfile cria uma nova instância de ResidentDayProfile.
//
// routine: O perfil de rotina diária do residente.
// frequency: O perfil de frequência de uso de aparelhos do residente.
//
// Retorna um ponteiro para o ResidentDayProfile recém-criado.
func NewResidentDayProfile(routine *routine.DayProfile, frequency *frequency.ResidentDeviceProfiles) *ResidentDayProfile {
	return &ResidentDayProfile{
		routineProfile:      routine,
		frequencyProfileDay: frequency,
	}
}

// RoutineProfile retorna o perfil de rotina diária associado a este ResidentDayProfile.
func (r *ResidentDayProfile) RoutineProfile() *routine.DayProfile {
	return r.routineProfile
}

// FrequencyProfileDay retorna o perfil de frequência de uso de aparelhos
// associado a este ResidentDayProfile.
func (r *ResidentDayProfile) FrequencyProfileDay() *frequency.ResidentDeviceProfiles {
	return r.frequencyProfileDay
}

// GenerateRoutine gera e retorna uma nova rotina diária para o residente,
// baseada no routine.RoutineProfile configurado neste perfil do dia.
//
// rng: O gerador de números aleatórios a ser usado para a geração dos dados.
//
// Retorna um ponteiro para a estrutura behavioral.Routine gerada.
func (r *ResidentDayProfile) GenerateRoutine(rng *rand.Rand) (*behavioral.Routine, error) {
	return r.routineProfile.GenerateData(rng)
}

// GenerateFrequency gera e retorna novos dados de frequência de uso de aparelhos
// para o residente, baseados no frequency.FrequencyProfileDay configurado neste perfil do dia.
//
// rng: O gerador de números aleatórios a ser usado para a geração dos dados.
//
// Retorna um ponteiro para a estrutura behavioral.Frequency gerada.
func (r *ResidentDayProfile) GenerateFrequency(rng *rand.Rand) (*behavioral.Frequency, error) {
	return r.frequencyProfileDay.GenerateData(rng)
}