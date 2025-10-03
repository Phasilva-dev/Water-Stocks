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

// RunSimulation chama a função de simulação do seu pacote 'simulation'
func (a *App) RunSimulation(size, day, toiletType, showerType int, filename string) {
	simulation.RunSimulation(size, day, toiletType, showerType, filename)
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
	HourlyAnalysis        []PlotlyTrace       `json:"hourlyAnalysis"` // Modificado para stacked bar
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

// performAllAnalyses executa todos os cálculos e retorna a estrutura do relatório final.
func (a *App) performAllAnalyses(data *ParsedCSV) FullAnalysisReport {
	// Análise 2 (Consumo por aparelho e Consumo per capita)
	var totalLitersConsumed float64
	deviceTotals := make(map[string]float64)
	deviceOrder := []string{}
	for _, d := range data.DeviceSummary {
		totalLitersConsumed += d.TotalLitersConsumed
		deviceTotals[d.Device] = d.TotalLitersConsumed
		deviceOrder = append(deviceOrder, d.Device)
	}
	sort.Strings(deviceOrder) // Garante uma ordem consistente

	deviceSummaryResult := DeviceSummaryResult{
		Devices: make([]string, len(deviceOrder)),
		Liters:  make([]float64, len(deviceOrder)),
	}
	for i, deviceName := range deviceOrder {
		deviceSummaryResult.Devices[i] = deviceName
		deviceSummaryResult.Liters[i] = deviceTotals[deviceName]
	}

	var consumptionPerPerson float64
	if data.Metadata.TotalPopulation > 0 && data.Metadata.TotalSimulationDays > 0 {
		consumptionPerPerson = totalLitersConsumed / float64(data.Metadata.TotalPopulation) / float64(data.Metadata.TotalSimulationDays)
	}

	// Análise 1, 3 e 4 (Consumo Horário, Pico e Pulso) em um único loop
	hourlyDeviceTotals := make(map[string][24]float64)
	var maxTotal, sumTotal float64

	pulseDataLen := len(data.PulseData)
	timeAxis := make([]int, pulseDataLen)
	deviceData := make(map[string][]float64)

	// Inicializa os slices para cada aparelho
	for _, h := range data.PulseHeaders {
		deviceData[h] = make([]float64, pulseDataLen)
	}

	for i, p := range data.PulseData {
		hour := (p.HorarioSegundos / 3600) % 24
		
		// Soma para cada dispositivo na hora correta
		for device, consumption := range p.Consumptions {
			// Não incluir o total na análise por dispositivo
			if device != "total" {
				if _, ok := hourlyDeviceTotals[device]; !ok {
					hourlyDeviceTotals[device] = [24]float64{}
				}
				currentTotals := hourlyDeviceTotals[device]
				currentTotals[hour] += consumption
				hourlyDeviceTotals[device] = currentTotals
			}
		}

		// Análise 3: Encontra o pico e soma total
		if total, ok := p.Consumptions["total"]; ok {
			sumTotal += total
			if total > maxTotal {
				maxTotal = total
			}
		}

		// Análise 4: Prepara os dados do gráfico de pulso
		timeAxis[i] = p.HorarioSegundos
		for _, h := range data.PulseHeaders {
			if val, ok := p.Consumptions[h]; ok {
				deviceData[h][i] = val
			}
		}
	}

	// Finaliza Análise 1: Preparar dados para o gráfico de barras empilhadas
	var hourlyAnalysis []PlotlyTrace
	hours := make([]int, 24)
	for i := 0; i < 24; i++ {
		hours[i] = i
	}

	// Garante uma ordem consistente das legendas
	pulseDeviceHeaders := []string{}
	for _, h := range data.PulseHeaders {
		if h != "total" { // não inclua a linha total na pilha do gráfico de barras
			pulseDeviceHeaders = append(pulseDeviceHeaders, h)
		}
	}
	sort.Strings(pulseDeviceHeaders)

	for _, deviceName := range pulseDeviceHeaders {
		consumptions, ok := hourlyDeviceTotals[deviceName]
		if !ok {
			continue // Pula se o dispositivo não tiver dados
		}
		
		avgConsumptions := make([]float64, 24)
		if data.Metadata.TotalSimulationDays > 0 {
			for i := 0; i < 24; i++ {
				avgConsumptions[i] = consumptions[i] / float64(data.Metadata.TotalSimulationDays)
			}
		}

		// Calcula a porcentagem do consumo total para a legenda
		deviceTotalConsumption := deviceTotals[deviceName]
		percentage := 0.0
		if totalLitersConsumed > 0 {
			percentage = (deviceTotalConsumption / totalLitersConsumed) * 100
		}
		
		trace := PlotlyTrace{
			X:    hours,
			Y:    avgConsumptions,
			Type: "bar",
			Name: fmt.Sprintf("%s (%.1f%%)", deviceName, percentage),
		}
		hourlyAnalysis = append(hourlyAnalysis, trace)
	}

	// Finaliza Análise 3: Coeficiente de pico
	var peakCoefficient float64
	if pulseDataLen > 0 {
		avgTotal := sumTotal / float64(pulseDataLen)
		if avgTotal > 0 {
			peakCoefficient = maxTotal / avgTotal
		}
	}

	// Finaliza Análise 4: Gráfico de pulso
	var pulseAnalysis []PlotlyTrace
	sort.Strings(data.PulseHeaders) // Garante a ordem da legenda
	for _, h := range data.PulseHeaders {
		pulseAnalysis = append(pulseAnalysis, PlotlyTrace{X: timeAxis, Y: deviceData[h], Mode: "lines", Type: "scatter", Name: h})
	}

	// Monta o relatório final
	return FullAnalysisReport{
		ConsumptionPerPerson:  consumptionPerPerson,
		PeakCoefficient:       peakCoefficient,
		HourlyAnalysis:        hourlyAnalysis, // Nova estrutura de dados
		DeviceSummaryAnalysis: deviceSummaryResult,
		PulseAnalysis:         pulseAnalysis, // Agora com todos os dados
	}
}

// parseCustomCSV lê o arquivo CSV estruturado em seções.
func (a *App) parseCustomCSV(filePath string) (*ParsedCSV, error) {
	file, err := os.Open(filePath)
	if err != nil { return nil, err }
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
				if key == "Total Population" { parsed.Metadata.TotalPopulation, _ = strconv.Atoi(value) }
				if key == "Total Simulation Days" { parsed.Metadata.TotalSimulationDays, _ = strconv.Atoi(value) }
			}
		case "DEVICE_SUMMARY":
			// Ignora a linha de cabeçalho
			if strings.ToLower(parts[0]) == "device" {
				continue
			}
			if len(parts) >= 3 {
				liters, _ := strconv.ParseFloat(parts[1], 64)
				uses, _ := strconv.Atoi(parts[2])
				parsed.DeviceSummary = append(parsed.DeviceSummary, DeviceSummary{Device: parts[0], TotalLitersConsumed: liters, TotalUses: uses})
			}
		case "PULSE_DATA":
			if isPulseDataHeader {
				// A primeira linha são os cabeçalhos. Nós só precisamos dos cabeçalhos de consumo.
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