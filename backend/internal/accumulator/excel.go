package accumulator

import (
	"fmt"
	"log"

	"simulation/internal/configs"

	"github.com/xuri/excelize/v2"
)

func (a *AccumulatorDay) ExportToExcel(filename string) error {
	f := excelize.NewFile()

	// Pegamos os dispositivos do primeiro acumulador (hora 0)
	deviceTypes := configs.OrderedDeviceKeys()
	// Cabeçalho
	f.SetCellValue("Sheet1", "A1", "Hour")
	for i, device := range deviceTypes {
		col := string(rune('B' + i)) // B, C, D...
		f.SetCellValue("Sheet1", fmt.Sprintf("%s1", col), device)
	}
	f.SetCellValue("Sheet1", fmt.Sprintf("%s1", string(rune('B'+len(deviceTypes)))), "Total")

	// Preencher linhas 2 a 25 (horas 0 a 23)
	for hour := 0; hour < 24; hour++ {
		row := hour + 2
		// Coluna A = hora
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), hour)

		// Valores por dispositivo
		var totalHour float64
		for i, device := range deviceTypes {
			col := string(rune('B' + i))
			value := a.accumulatorHour[hour].IndividualDeviceTotal(device)
			totalHour += value
			f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", col, row), value)
		}
		// Total da hora na última coluna
		totalCol := string(rune('B' + len(deviceTypes)))
		f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", totalCol, row), totalHour)
	}

	// Linha 26 = totais por dispositivo + total geral
	f.SetCellValue("Sheet1", "A26", "Total")

	var grandTotal float64
	for i, device := range deviceTypes {
		col := string(rune('B' + i))
		var deviceTotal float64
		for hour := 0; hour < 24; hour++ {
			deviceTotal += a.accumulatorHour[hour].IndividualDeviceTotal(device)
		}
		grandTotal += deviceTotal
		f.SetCellValue("Sheet1", fmt.Sprintf("%s26", col), deviceTotal)
	}
	// Total geral na última coluna
	totalCol := string(rune('B' + len(deviceTypes)))
	f.SetCellValue("Sheet1", fmt.Sprintf("%s26", totalCol), grandTotal)

	// Salvar arquivo
	if err := f.SaveAs(filename); err != nil {
		log.Printf("Erro ao salvar arquivo Excel: %v", err)
		return err
	}

	log.Printf("Arquivo Excel salvo com sucesso em: %s", filename)
	return nil
}