// --- START OF FILE app.go ---
package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"simulation/cmd/simulation" // Seu import original
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// RunSimulation chama a função de simulação e envia atualizações de progresso.
func (a *App) RunSimulation(size, day, toiletType, showerType int, filename string) (string, error) {
	// Mensagem inicial
	runtime.EventsEmit(a.ctx, "simulationStatus", "Iniciando simulação...")

	if filename == "" {
		err := fmt.Errorf("o nome do arquivo não pode estar vazio")
		log.Println(err)
		return "", err
	}

	const outputDir = "simulations_output"
	analysisCsvFilename := filepath.Join(outputDir, fmt.Sprintf("%s_analysis.csv", filename))

	// MODIFICAÇÃO: Definimos a função de callback aqui.
	// Esta função será executada a cada dia concluído pela simulação.
	progressCallback := func(currentDay, totalDays int) {
		message := fmt.Sprintf("Dia %d de %d simulado...", currentDay, totalDays)
		log.Println(message) // Log para o terminal do Go
		runtime.EventsEmit(a.ctx, "simulationStatus", message) // Envia para o frontend
	}

	// MODIFICAÇÃO: Passamos a função 'progressCallback' para a camada de simulação.
	err := simulation.RunSimulation(size, day, toiletType, showerType, filename, progressCallback)
	if err != nil {
		log.Printf("A simulação falhou: %v", err)
		return "", err // Retorna o erro para o frontend
	}

	// Mensagem final antes de salvar o arquivo
	runtime.EventsEmit(a.ctx, "simulationStatus", "Simulação concluída. Gerando arquivo CSV...")

	successMessage := fmt.Sprintf("Simulação '%s' concluída com sucesso! Arquivo salvo em: %s", filename, analysisCsvFilename)
	log.Println(successMessage)
	
	// Retorna a mensagem final de sucesso para o frontend
	return successMessage, nil
}



// SelectFile abre um diálogo para o usuário selecionar um arquivo CSV.
func (a *App) SelectFile() (string, error) {
	selection, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Selecione um arquivo de análise",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Arquivos CSV (*.csv)",
				Pattern:     "*.csv",
			},
		},
	})
	if err != nil {
		return "", err
	}
	return selection, nil
}

// =================================================================
// LÓGICA PARA ANÁLISE DE GRÁFICOS
// =================================================================

// ----- Estruturas de Dados para o CSV -----
type Metadata struct {
	TotalPopulation     int
	TotalSimulationDays int
}
type DeviceSummary struct {
	Device              string
	TotalLitersConsumed float64
	TotalUses           int
}
type PulseData struct {
	HorarioSegundos int
	Consumptions    map[string]float64
}
type ParsedCSV struct {
	Metadata      Metadata
	DeviceSummary []DeviceSummary
	PulseData     []PulseData
	PulseHeaders  []string // Nomes dos aparelhos + "total"
}

// ----- Estruturas de Dados para os Resultados da Análise -----
type DeviceSummaryResult struct {
	Devices []string  `json:"devices"`
	Liters  []float64 `json:"liters"`
}
type PlotlyTrace struct {
	X    []int     `json:"x"`
	Y    []float64 `json:"y"`
	Mode string    `json:"mode,omitempty"`
	Type string    `json:"type"`
	Name string    `json:"name"`
}
type FullAnalysisReport struct {
	ConsumptionPerPerson  float64             `json:"consumptionPerPerson"`
	PeakCoefficient       float64             `json:"peakCoefficient"`
	HourlyAnalysis        []PlotlyTrace       `json:"hourlyAnalysis"`
	DeviceSummaryAnalysis DeviceSummaryResult `json:"deviceSummaryAnalysis"`
	PulseAnalysis         []PlotlyTrace       `json:"pulseAnalysis"`
}

// GenerateAnalysisAndOpenBrowser é a função principal que o frontend chamará.
func (a *App) GenerateAnalysisAndOpenBrowser(csvFilePath string) (string, error) {
	log.Printf("Iniciando análise para o arquivo: %s", csvFilePath)

	parsedData, err := a.parseCustomCSV(csvFilePath)
	if err != nil {
		log.Printf("ERRO no parse: %v", err)
		return "", fmt.Errorf("falha ao parsear o CSV: %w", err)
	}

	if len(parsedData.PulseData) == 0 || len(parsedData.DeviceSummary) == 0 {
		log.Printf("ERRO: Dados essenciais não encontrados. PulseData: %d, DeviceSummary: %d", len(parsedData.PulseData), len(parsedData.DeviceSummary))
		return "", fmt.Errorf("os dados do CSV não foram lidos corretamente, verifique o formato do arquivo e os logs do terminal")
	}

	report := a.performAllAnalyses(parsedData)
	reportJSONBytes, err := json.Marshal(report)
	if err != nil {
		log.Printf("ERRO ao converter relatório para JSON: %v", err)
		return "", fmt.Errorf("falha ao converter relatório para JSON: %w", err)
	}

	templateData := struct {
		FileName     string
		AnalysisJSON template.JS
	}{
		FileName:     filepath.Base(csvFilePath),
		AnalysisJSON: template.JS(reportJSONBytes),
	}

	tmpl, err := template.ParseFS(chartTemplate, "chart_template.html")
	if err != nil {
		log.Printf("ERRO ao carregar template: %v", err)
		return "", fmt.Errorf("falha ao carregar template HTML: %w", err)
	}

	tempFile, err := os.CreateTemp("", "hydro-report-*.html")
	if err != nil {
		log.Printf("ERRO ao criar arquivo temporário: %v", err)
		return "", fmt.Errorf("falha ao criar arquivo temporário: %w", err)
	}
	defer tempFile.Close()

	if err := tmpl.Execute(tempFile, templateData); err != nil {
		log.Printf("ERRO ao executar template: %v", err)
		return "", fmt.Errorf("falha ao executar template: %w", err)
	}

	absPath, _ := filepath.Abs(tempFile.Name())
	runtime.BrowserOpenURL(a.ctx, "file://"+absPath)

	log.Println("Relatório gerado e aberto com sucesso!")
	return "Relatório de análise gerado com sucesso!", nil
}

func (a *App) performAllAnalyses(data *ParsedCSV) FullAnalysisReport {
	days := float64(data.Metadata.TotalSimulationDays)
	if days == 0 {
		log.Println("Aviso: TotalSimulationDays é 0, assumindo 1 para evitar divisão por zero.")
		days = 1
	}

	// Análise 2: Consumo total por aparelho
	var totalLitersConsumed float64
	deviceTotals := make(map[string]float64)
	deviceOrder := []string{}
	for _, d := range data.DeviceSummary {
		totalLitersConsumed += d.TotalLitersConsumed
		deviceTotals[d.Device] = d.TotalLitersConsumed
		deviceOrder = append(deviceOrder, d.Device)
	}
	sort.Strings(deviceOrder)

	deviceSummaryResult := DeviceSummaryResult{
		Devices: make([]string, len(deviceOrder)),
		Liters:  make([]float64, len(deviceOrder)),
	}
	for i, deviceName := range deviceOrder {
		deviceSummaryResult.Devices[i] = deviceName
		deviceSummaryResult.Liters[i] = deviceTotals[deviceName] // Mantém o total aqui, pois o gráfico se chama "Consumo Total"
	}

	// Consumo médio diário por pessoa
	var consumptionPerPerson float64
	if data.Metadata.TotalPopulation > 0 {
		consumptionPerPerson = totalLitersConsumed / float64(data.Metadata.TotalPopulation) / days
	}

	// Agregando dados por segundo e por hora em um único loop
	secondsInDay := 86400
	hourlyDeviceTotals := make(map[string][24]float64)
	pulseDeviceTotals := make(map[string][]float64)
	
	// Inicializa os slices para os dados de pulso
	for _, h := range data.PulseHeaders {
		pulseDeviceTotals[h] = make([]float64, secondsInDay)
	}

	for _, p := range data.PulseData {
		secondOfDay := p.HorarioSegundos % secondsInDay
		hourOfDay := secondOfDay / 3600

		for device, consumption := range p.Consumptions {
			// Agrega dados de pulso para o dia médio
			pulseDeviceTotals[device][secondOfDay] += consumption

			// Agrega dados horários (exceto a coluna "total")
			if device != "total" {
				if _, ok := hourlyDeviceTotals[device]; !ok {
					hourlyDeviceTotals[device] = [24]float64{}
				}
				currentTotals := hourlyDeviceTotals[device]
				currentTotals[hourOfDay] += consumption
				hourlyDeviceTotals[device] = currentTotals
			}
		}
	}

	// ANÁLISE 1: Média de Consumo Horário (Gráfico de Barras Empilhadas)
	var hourlyAnalysis []PlotlyTrace
	totalHourlyConsumption := make([]float64, 24)
	hours := make([]int, 24)
	for i := 0; i < 24; i++ {
		hours[i] = i
	}

	// Ordena os dispositivos para uma legenda consistente
	pulseDeviceHeaders := []string{}
	for h := range hourlyDeviceTotals {
		pulseDeviceHeaders = append(pulseDeviceHeaders, h)
	}
	sort.Strings(pulseDeviceHeaders)

	for _, deviceName := range pulseDeviceHeaders {
		consumptions := hourlyDeviceTotals[deviceName]
		avgConsumptions := make([]float64, 24)
		for i := 0; i < 24; i++ {
			avgConsumption := consumptions[i] / days
			avgConsumptions[i] = avgConsumption
			totalHourlyConsumption[i] += avgConsumption // Acumula para o K1
		}

		percentage := 0.0
		if totalLitersConsumed > 0 {
			percentage = (deviceTotals[deviceName] / totalLitersConsumed) * 100
		}

		trace := PlotlyTrace{
			X:    hours,
			Y:    avgConsumptions,
			Type: "bar",
			Name: fmt.Sprintf("%s (%.1f%%)", deviceName, percentage),
		}
		hourlyAnalysis = append(hourlyAnalysis, trace)
	}

	// ANÁLISE 3: Coeficiente de Pico Horário (K1)
	var peakCoefficient, maxHourlyConsumption, sumOfHourlyAverages float64
	for _, v := range totalHourlyConsumption {
		sumOfHourlyAverages += v
		if v > maxHourlyConsumption {
			maxHourlyConsumption = v
		}
	}
	avgHourlyConsumption := sumOfHourlyAverages / 24
	if avgHourlyConsumption > 0 {
		peakCoefficient = maxHourlyConsumption / avgHourlyConsumption
	}
	
	// ANÁLISE 4: Gráfico de Pulso Médio Diário
	var pulseAnalysis []PlotlyTrace
	timeAxis := make([]int, secondsInDay)
	for i := 0; i < secondsInDay; i++ {
		timeAxis[i] = i
	}

	// Ordena os cabeçalhos para uma legenda consistente
	sortedPulseHeaders := make([]string, 0, len(data.PulseHeaders))
	for _, h := range data.PulseHeaders {
		sortedPulseHeaders = append(sortedPulseHeaders, h)
	}
	sort.Strings(sortedPulseHeaders)
	
	for _, h := range sortedPulseHeaders {
		avgPulseY := make([]float64, secondsInDay)
		for i, totalVal := range pulseDeviceTotals[h] {
			avgPulseY[i] = totalVal / days
		}
		pulseAnalysis = append(pulseAnalysis, PlotlyTrace{
			X:    timeAxis,
			Y:    avgPulseY,
			Mode: "lines",
			Type: "scatter",
			Name: h,
		})
	}

	// Monta o relatório final
	return FullAnalysisReport{
		ConsumptionPerPerson:  consumptionPerPerson,
		PeakCoefficient:       peakCoefficient,
		HourlyAnalysis:        hourlyAnalysis,
		DeviceSummaryAnalysis: deviceSummaryResult,
		PulseAnalysis:         pulseAnalysis,
	}
}

// parseCustomCSV lê o arquivo CSV estruturado em seções.
func (a *App) parseCustomCSV(filePath string) (*ParsedCSV, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	parsed := &ParsedCSV{}
	var currentSection string
	isPulseDataHeader := false

	log.Println("--- INICIANDO PARSING ---")

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			currentSection = strings.Trim(line, "[]")
			log.Printf("MUDANDO PARA A SEÇÃO: [%s]", currentSection)

			if currentSection == "PULSE_DATA" {
				isPulseDataHeader = true
			}
			continue
		}

		parts := strings.Split(line, ",")
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}

		switch currentSection {
		case "METADATA":
			if len(parts) >= 2 {
				key := parts[0]
				value := parts[1]
				if key == "Total Population" {
					parsed.Metadata.TotalPopulation, _ = strconv.Atoi(value)
				}
				if key == "Total Simulation Days" {
					parsed.Metadata.TotalSimulationDays, _ = strconv.Atoi(value)
				}
			}
		case "DEVICE_SUMMARY":
			if strings.ToLower(parts[0]) == "device" { // Ignora cabeçalho
				continue
			}
			if len(parts) >= 3 {
				liters, _ := strconv.ParseFloat(parts[1], 64)
				uses, _ := strconv.Atoi(parts[2])
				parsed.DeviceSummary = append(parsed.DeviceSummary, DeviceSummary{Device: parts[0], TotalLitersConsumed: liters, TotalUses: uses})
			}
		case "PULSE_DATA":
			if isPulseDataHeader {
				parsed.PulseHeaders = parts[1:]
				log.Printf("  Cabeçalhos de Pulso lidos do CSV: %v", parsed.PulseHeaders)
				isPulseDataHeader = false
				continue
			}

			if parsed.PulseHeaders != nil && len(parts) == len(parsed.PulseHeaders)+1 {
				horario, err := strconv.Atoi(parts[0])
				if err != nil {
					log.Printf("AVISO: Pulando linha de dados de pulso mal formatada (horário): %s", line)
					continue
				}

				consumptions := make(map[string]float64)
				for i, headerName := range parsed.PulseHeaders {
					val, err := strconv.ParseFloat(parts[i+1], 64)
					if err != nil {
						log.Printf("AVISO: Pulando valor mal formatado para '%s' na linha: %s", headerName, line)
						val = 0.0
					}
					consumptions[headerName] = val
				}
				parsed.PulseData = append(parsed.PulseData, PulseData{HorarioSegundos: horario, Consumptions: consumptions})
			}
		}
	}

	log.Println("--- FIM DO PARSING ---")
	log.Printf("  Metadados Finais: %+v", parsed.Metadata)
	log.Printf("  Registros Finais Device Summary: %d", len(parsed.DeviceSummary))
	log.Printf("  Registros Finais Pulse Data: %d", len(parsed.PulseData))

	return parsed, scanner.Err()
}
// --- END OF FILE app.go ---