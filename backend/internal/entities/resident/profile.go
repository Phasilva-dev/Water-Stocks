// Package resident define perfis para residentes, agregando seus hábitos semanais
// para a geração de dados comportamentais (rotina e frequência de uso).
package resident

import (
	"errors"
	"math/rand/v2"
	"simulation/internal/entities/resident/ds/behavioral"
	"simulation/internal/entities/resident/profile/habits"
)

// ResidentProfile representa o perfil completo de um residente, incluindo sua
// ID de ocupação e seu perfil de hábitos comportamentais semanais.
type ResidentProfile struct {
	OccupationID  uint32
	weeklyProfile *habits.ResidentWeeklyProfile
}

// NewResidentProfile cria uma nova instância de ResidentProfile.
//
// Parâmetros:
//   - profile: O perfil de hábitos semanais do residente. Não pode ser nil.
//   - id: A ID de ocupação do residente. Não pode ser zero.
//
// Retorna um ponteiro para o ResidentProfile recém-criado ou um erro se os
// parâmetros forem inválidos.
func NewResidentProfile(profile *habits.ResidentWeeklyProfile, id uint32) (*ResidentProfile, error) {
	if profile == nil {
		return nil, errors.New("invalid ResidentProfile: weekly profile cannot be nil")
	}

	return &ResidentProfile{
		OccupationID:  id,
		weeklyProfile: profile,
	}, nil
}

// GenerateFrequency gera e retorna dados de frequência de uso para um dia específico da semana.
// Delega a geração ao perfil semanal configurado.
//
// Parâmetros:
//   - day: O dia da semana (0 para o primeiro dia, até 6 para o sétimo dia).
//   - rng: O gerador de números aleatórios a ser usado.
//
// Retorna um ponteiro para a estrutura behavioral.Frequency com os dados gerados.
func (r *ResidentProfile) GenerateFrequency(day uint8, rng *rand.Rand) *behavioral.Frequency {
	// A responsabilidade de verificar se o 'day' é válido para o weeklyProfile
	// (e se o perfil para aquele dia existe) está dentro de weeklyProfile.GenerateFrequency.
	return r.weeklyProfile.GenerateFrequency(day, rng)
}

// GenerateRoutine gera e retorna dados de rotina para um dia específico da semana.
// Delega a geração ao perfil semanal configurado.
//
// Parâmetros:
//   - day: O dia da semana (0 para o primeiro dia, até 6 para o sétimo dia).
//   - rng: O gerador de números aleatórios a ser usado.
//
// Retorna um ponteiro para a estrutura behavioral.Routine com os dados gerados.
func (r *ResidentProfile) GenerateRoutine(day uint8, rng *rand.Rand) *behavioral.Routine {
	// A responsabilidade de verificar se o 'day' é válido para o weeklyProfile
	// (e se o perfil para aquele dia existe) está dentro de weeklyProfile.GenerateRoutine.
	return r.weeklyProfile.GenerateRoutine(day, rng)
}