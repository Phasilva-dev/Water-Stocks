package usagemock

import(
	//"simulation/internal/entities"
	//"simulation/internal/log"
)

// InverteHorarioCiclico inverte um horário mantendo a referência cíclica
// Retorna o horário invertido podendo ser negativo para indicar dia anterior
func inverteHorarioCiclico(horario int32) int32 {
	const totalSegundosDia = 86400
	
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
}
/*
func avaliableTime(log *log.ResidentSanitary) {
	toiletLog, isValid := log.ToiletLog().UsageLogs()
	toiletLog[].
	if isValid {
		for i in range len(toiletLog)
	}
	showerLog, isValid := log.ShowerLog().UsageLogs()
	washBassin, isValid := log.WashBassinLog().UsageLogs()
	washMachine, isValid := log.WashMachineLog().UsageLogs()
	dishWasher, isValid := log.DishWasherLog().UsageLogs()
	tanque, isValid := log.TanqueLog().UsageLogs()


}*/
/*
func checkCollision(start1, end1, start2, end2 int32) bool {
	return start1 < end2 && start2 < end1
}*/