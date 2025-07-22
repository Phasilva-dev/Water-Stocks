package usagemock

import (
	"fmt"
	"simulation/internal/log"
)

func warningUsage(usage *log.Usage, device string, decision int, percentil float64, min, max int32) {
	if usage.StartUsage() < 3600 * (-min) {
		fmt.Printf("[WARNING] dia anterior na distribuição de uso do %s (p = %.4f), (decisão = %d), (tempo = %d)\n", device, percentil, decision, usage.StartUsage())
	} else if usage.StartUsage() >= 3600 * (24 + max){
		fmt.Printf("[WARNING] dia posterior na distribuição de uso do %s (p = %.4f), (decisão = %d), (tempo = %d)\n", device, percentil, decision, usage.StartUsage())
	}


}

// InverteHorarioCiclico inverte um horário mantendo a referência cíclica
// Retorna o horário invertido podendo ser negativo para indicar dia anterior
/*func inverteHorarioCiclico(horario int32) int32 {
	const totalSegundosDia = 86400

	if horario >= 86400 {
		return horario - 86400
	}
	
	// Se for positivo, inverte normalmente
	if horario >= 0 {
		invertido := totalSegundosDia - (horario % totalSegundosDia)
		if invertido == totalSegundosDia {
			return 0
		}
		return invertido
	}
	
	// Se for negativo, calcula o equivalente positivo, inverte e retorna negativo
	absHorario := -horario
	invertidoPositivo := totalSegundosDia - (absHorario % totalSegundosDia)
	
	// Ajusta para não retornar -86400 (que seria equivalente a 0)
	if invertidoPositivo == totalSegundosDia {
		return 0
	}
	
	return -invertidoPositivo
}*/
