// generate_samples.go
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"path/filepath" // NOVO: Pacote para manipular caminhos de arquivos de forma segura
	"strconv"
	"strings"

	"simulation/internal/dists" // Mantenha o seu caminho de importação
)

/*
// usageInfo (sem alterações)
var usageInfo = map[string]string{
	"exponential": "rate",
	"normal":      "mean stdDev",
	"uniform":     "min max",
	"weibull":     "shape scale",
	"gamma":       "shape scale",
	"lognormal":   "mu sigma",
	"triangle":    "min mode max",
	"loglogistic": "shape scale",
	"poisson":     "lambda",
}*/

// printUsage (atualizado com novas instruções)
func printUsage() {
	fmt.Println("Usage: go run generate_samples.go [-o filename] <distribution_type> [params...]")
	fmt.Println("\nSe apenas um nome de arquivo for fornecido com -o, ele será salvo na pasta padrão.")
	fmt.Println("\nOptions:")
	flag.PrintDefaults()
	fmt.Println("\nExample (salva 'weibull_run1.csv' na pasta padrão):")
	fmt.Println("  go run generate_samples.go -o weibull_run1 weibull 1.5871 0.05691")
	fmt.Println("\nExample (salva em um local específico):")
	fmt.Println(`  go run generate_samples.go -o C:\temp\outro_teste.csv weibull 1.5871 0.05691`)
}

const numSamples = 100000

// NOVO: Definição da pasta de salvamento padrão.
// Use duas barras invertidas `\\` em strings no Go para representar uma única `\`.
const defaultSavePath = `C:\Users\Pedro\Desktop\codigin`

func main() {
	// --- 1. Definir e Ler Flags e Argumentos ---

	var outputFilename string
	flag.StringVar(&outputFilename, "o", "samples.csv", "Nome do arquivo de saída (sem caminho para usar a pasta padrão)")
	flag.Usage = printUsage
	flag.Parse()

	// ======================================================================
	// --- LÓGICA PRINCIPAL ATUALIZADA ---
	// ======================================================================

	// Passo A: Garante que o nome do arquivo termina com .csv
	if !strings.HasSuffix(strings.ToLower(outputFilename), ".csv") {
		outputFilename += ".csv"
	}

	// Passo B: Verifica se o usuário forneceu um caminho completo ou apenas um nome de arquivo.
	// Se for apenas um nome de arquivo, nós o juntamos com a pasta padrão.
	isFullPath := strings.Contains(outputFilename, `\`) || strings.Contains(outputFilename, "/")
	finalPath := outputFilename
	if !isFullPath {
		// Garante que a pasta padrão exista. Se não, a cria.
		if err := os.MkdirAll(defaultSavePath, os.ModePerm); err != nil {
			log.Fatalf("❌ Failed to create default directory '%s': %v", defaultSavePath, err)
		}
		finalPath = filepath.Join(defaultSavePath, outputFilename)
	}

	// O resto do código usará a variável `finalPath`.
	// ======================================================================

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Error: Missing distribution type.")
		printUsage()
		return
	}

	distType := strings.ToLower(args[0])
	paramStrings := args[1:]

	var distParams []float64
	for _, paramStr := range paramStrings {
		p, err := strconv.ParseFloat(paramStr, 64)
		if err != nil {
			log.Fatalf("❌ Error: Invalid number provided for a parameter: '%s'. Please provide only numbers.", paramStr)
		}
		distParams = append(distParams, p)
	}

	dist, err := dists.CreateDistribution(distType, distParams...)
	if err != nil {
		log.Fatalf("❌ Error creating distribution: %v", err)
	}
	fmt.Printf("✅ Generating %d samples for: %s\n", numSamples, dist)

	rng := rand.New(rand.NewPCG(10, 20))

	// Usa `finalPath` para criar o arquivo no local correto.
	file, err := os.Create(finalPath)
	if err != nil {
		log.Fatalf("❌ Failed to create CSV file '%s': %v", finalPath, err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headerParts := []string{distType}
	headerParts = append(headerParts, paramStrings...)
	header := strings.Join(headerParts, ";")
	if err := writer.Write([]string{header}); err != nil {
		log.Fatalf("❌ Error writing smart header to CSV: %v", err)
	}

	for i := 0; i < numSamples; i++ {
		sample := dist.Sample(rng)
		sampleStr := strconv.FormatFloat(sample, 'f', -1, 64)
		if err := writer.Write([]string{sampleStr}); err != nil {
			log.Fatalf("❌ Error writing row %d to CSV: %v", i+1, err)
		}
	}

	fmt.Printf("✅ Successfully generated '%s' with %d samples and a smart header.\n", finalPath, numSamples)
}