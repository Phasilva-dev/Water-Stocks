
// --- START OF FILE app.go ---
package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log" // Importante para mensagens de depuração
	"os"
	"path/filepath"
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

// =================================================================
// SUAS FUNÇÕES ORIGINAIS (MANTIDAS)
// =================================================================

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
// NOVA LÓGICA PARA ANÁLISE DE GRÁFICOS
// =================================================================

// ----- Estruturas de Dados para o CSV -----
type Metadata struct {
	TotalPopulation     int
	TotalSimulationDays int
}
type DeviceSummary struct {
	Device string
	TotalLitersConsumed float64
	TotalUses int
}
type PulseData struct {
	HorarioSegundos int
	Consumptions map[string]float64
}
type ParsedCSV struct {
	Metadata      Metadata
	DeviceSummary []DeviceSummary
	PulseData     []PulseData
	PulseHeaders  []string // Nomes dos aparelhos + "total", sem "horario_segundos"
}

// ----- Estruturas de Dados para os Resultados da Análise -----
type HourlyResult struct {
	Hours       []int     `json:"hours"`
	Consumption []float64 `json:"consumption"`
}
type DeviceSummaryResult struct {
	Devices []string  `json:"devices"`
	Liters  []float64 `json:"liters"`
}
type PlotlyTrace struct {
	X    []int     `json:"x"`
	Y    []float64 `json:"y"`
	Mode string    `json:"mode"`
	Type string    `json:"type"`
	Name string    `json:"name"`
}
type FullAnalysisReport struct {
	ConsumptionPerPerson  float64             `json:"consumptionPerPerson"`
	PeakCoefficient       float64             `json:"peakCoefficient"`
	HourlyAnalysis        HourlyResult        `json:"hourlyAnalysis"`
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

	// Verificação de dados essenciais após o parsing
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

	// Prepara os dados do template com o tipo template.JS
	templateData := struct {
		FileName     string
		AnalysisJSON template.JS // <-- VOLTE PARA template.JS
	}{
		FileName:     filepath.Base(csvFilePath),
		AnalysisJSON: template.JS(reportJSONBytes), // <-- FAÇA O CAST PARA template.JS
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
	// Análise 1: Consumo Horário
	hourlyTotals := make([]float64, 24)
	hourlyResult := HourlyResult{Hours: make([]int, 24), Consumption: make([]float64, 24)}

	// Análise 2: Consumo por aparelho
	var totalLitersConsumed float64
	deviceSummaryResult := DeviceSummaryResult{Devices: make([]string, len(data.DeviceSummary)), Liters: make([]float64, len(data.DeviceSummary))}
	for i, d := range data.DeviceSummary {
		totalLitersConsumed += d.TotalLitersConsumed
		deviceSummaryResult.Devices[i] = d.Device
		deviceSummaryResult.Liters[i] = d.TotalLitersConsumed
	}

	// Análise 3 & 4 combinadas em um único loop
	var maxTotal, sumTotal float64
	var pulseAnalysis []PlotlyTrace

	pulseDataLimit := 7200
	if len(data.PulseData) < pulseDataLimit {
		pulseDataLimit = len(data.PulseData)
	}

	timeAxis := make([]int, pulseDataLimit)
	deviceData := make(map[string][]float64)
	for _, h := range data.PulseHeaders {
		deviceData[h] = make([]float64, pulseDataLimit)
	}

	for i, p := range data.PulseData {
		// Análise 1: Soma dos totais horários
		hour := (p.HorarioSegundos / 3600) % 24
		if total, ok := p.Consumptions["total"]; ok {
			hourlyTotals[hour] += total
		}
		
		// Análise 3: Encontra o pico e soma total
		if total, ok := p.Consumptions["total"]; ok {
			sumTotal += total
			if total > maxTotal {
				maxTotal = total
			}
		}

		// Análise 4: Prepara os dados do gráfico de pulso (limitado a 'pulseDataLimit')
		if i < pulseDataLimit {
			timeAxis[i] = p.HorarioSegundos
			for _, h := range data.PulseHeaders {
				if val, ok := p.Consumptions[h]; ok {
					deviceData[h][i] = val
				}
			}
		}
	}
	
	// Finaliza a Análise 1 (cálculo da média)
	for i := 0; i < 24; i++ {
		hourlyResult.Hours[i] = i
		if data.Metadata.TotalSimulationDays > 0 {
			hourlyResult.Consumption[i] = hourlyTotals[i] / float64(data.Metadata.TotalSimulationDays)
		}
	}

	// Finaliza a Análise 3 (cálculo do coeficiente de pico)
	var peakCoefficient float64
	if len(data.PulseData) > 0 {
		avgTotal := sumTotal / float64(len(data.PulseData))
		if avgTotal > 0 {
			peakCoefficient = maxTotal / avgTotal
		}
	}
	
	// Finaliza a Análise 4 (monta a estrutura para Plotly)
	for _, h := range data.PulseHeaders {
		pulseAnalysis = append(pulseAnalysis, PlotlyTrace{X: timeAxis, Y: deviceData[h], Mode: "lines", Type: "scatter", Name: h})
	}
	
	// Análise 2 (continuação)
	var consumptionPerPerson float64
	if data.Metadata.TotalPopulation > 0 && data.Metadata.TotalSimulationDays > 0 {
		consumptionPerPerson = totalLitersConsumed / float64(data.Metadata.TotalPopulation) / float64(data.Metadata.TotalSimulationDays)
	}

	// Monta o relatório final
	return FullAnalysisReport{
		ConsumptionPerPerson:  consumptionPerPerson,
		PeakCoefficient:       peakCoefficient,
		HourlyAnalysis:        hourlyResult,
		DeviceSummaryAnalysis: deviceSummaryResult,
		PulseAnalysis:         pulseAnalysis,
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
	isPulseDataHeader := false // Flag para identificar a linha de cabeçalho

	log.Println("--- INICIANDO PARSING (VERSÃO FINAL E ROBUSTA) ---")

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			currentSection = strings.Trim(line, "[]")
			log.Printf("MUDANDO PARA A SEÇÃO: [%s]", currentSection)
			
			if currentSection == "PULSE_DATA" {
				isPulseDataHeader = true // A próxima linha será o cabeçalho
			}
			continue // Pula para a próxima linha após identificar a seção
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
			if len(parts) >= 3 {
				liters, _ := strconv.ParseFloat(parts[1], 64)
				uses, _ := strconv.Atoi(parts[2])
				parsed.DeviceSummary = append(parsed.DeviceSummary, DeviceSummary{Device: parts[0], TotalLitersConsumed: liters, TotalUses: uses})
			}
		case "PULSE_DATA":
			if isPulseDataHeader {
				// A primeira linha são os cabeçalhos. Nós só precisamos dos cabeçalhos de consumo.
				// O primeiro é "horario_segundos", então pulamos ele.
				parsed.PulseHeaders = parts[1:]
				log.Printf("  Cabeçalhos de Pulso lidos do CSV: %v", parsed.PulseHeaders)
				isPulseDataHeader = false // Desativa a flag
				continue // Pula para a próxima linha que conterá dados
			}
			
			// Processamento normal das linhas de dados
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
						val = 0.0 // Define um valor padrão em caso de erro
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