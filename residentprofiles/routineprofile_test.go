package residentprofiles

import (
	"testing"
	"dists"
	"residentdata"
	//"image/color"
	//"fmt"
	"math/rand/v2"
	//"gonum.org/v1/plot"
	//"gonum.org/v1/plot/plotter"
	//"gonum.org/v1/plot/vg"
)

const (
	DAY_IN_SECONDS = 24 * 60 * 60
)

func TestNewRoutineProfileDist(t *testing.T) {
	// Cria uma distribuição válida para usar nos testes
	mockDist, _ := dists.NewNormalDist(20, 4)

	// Tabela de testes (table-driven tests)
	tests := []struct {
		name        string               // Nome descritivo do teste
		symbol      string               // Input: símbolo
		shift   int32                // Input: shift
		events      []dists.Distribution // Input: distribuições
		wantErr     bool                 // Se esperamos erro
		errMsg      string               // Mensagem de erro esperada
	}{
		{
			name:      "Caso válido",
			symbol:    "home",
			shift: 5,
			events:    []dists.Distribution{mockDist, mockDist},
			wantErr:   false,
		},
		{
			name:      "Número ímpar de eventos",
			symbol:    "work",
			shift: 5,
			events:    []dists.Distribution{mockDist},
			wantErr:   true,
			errMsg:    "number of elements in events must be even",
		},
		{
			name:      "shift negativa",
			symbol:    "school",
			shift: -1,
			events:    []dists.Distribution{mockDist, mockDist},
			wantErr:   true,
			errMsg:    "shift must be positive",
		},
		{
			name:      "Distribuição nil",
			symbol:    "park",
			shift: 5,
			events:    []dists.Distribution{mockDist, nil},
			wantErr:   true,
			errMsg:    "no distribution can be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Executa a função
			got, err := NewRoutineProfileDist(tt.shift, tt.events)

			// Verifica se o erro é o esperado
			if tt.wantErr {
				if err == nil {
					t.Fatal("Esperava erro, mas não ocorreu")
				}
				if err.Error() != tt.errMsg {
					t.Errorf("Mensagem de erro incorreta\nEsperada: %s\nObtida: %s", tt.errMsg, err.Error())
				}
				return
			}

			// Se não espera erro
			if err != nil {
				t.Fatalf("Erro inesperado: %v", err)
			}

			if got.shift != tt.shift {
				t.Errorf("shift incorreta\nEsperada: %d\nObtida: %d", tt.shift, got.shift)
			}

			if len(got.Events()) != len(tt.events) {
				t.Errorf("Número de eventos incorreto\nEsperado: %d\nObtido: %d", len(tt.events), len(got.Events()))
			}
		})
	}
}

func TestGenerateData(t *testing.T) {
	const (
		numRoutines = 1000   // Quantidade de rotinas a gerar
		symbol      = "home" // Símbolo esperado
		shift   = 60 * 15 // Gap mínimo entre entrada/saída
	)

	// 1. Cria uma distribuição determinística
	wakeUpDist, _ := dists.NewNormalDist(5*60*60, 0)
	leaveDist, _ := dists.NewNormalDist(7*60*60, 0)
	returnDist, _ := dists.NewNormalDist(18*60*60, 0)
	sleepDist, _ := dists.NewNormalDist(23*60*60, 0)

	// 2. Cria o perfil
	profile, err := NewRoutineProfileDist(
		shift,
		[]dists.Distribution{wakeUpDist, leaveDist, returnDist, sleepDist},
	)
	if err != nil {
		t.Fatalf("Falha ao criar perfil: %v", err)
	}

	// 3. Gera as rotinas
	src := rand.NewPCG(42,1)
	rng := rand.New(src) // Seed fixa para reprodutibilidade
	routines := make([]*residentdata.Routine, 0, numRoutines)

	for i := 0; i < numRoutines; i++ {
		routine := profile.GenerateData(rng)
		routines = append(routines, routine)
	}

	// 4. Verificações
	expectedTimes := []int32{5 * 60 * 60, 7 * 60 * 60, 18 * 60 * 60, 23 * 60 * 60}

	for i, routine := range routines {


		// Verifica quantidade de tempos
		if len(routine.Times()) != len(expectedTimes) {
			t.Fatalf("[Rotina %d] Número de tempos incorreto. Esperado: %d, Obtido: %d",
				i, len(expectedTimes), len(routine.Times()))
		}

		// Verifica valores e gap mínimo
		for j := 0; j < len(expectedTimes); j++ {
			if routine.Times()[j] != expectedTimes[j] {
				t.Errorf("[Rotina %d] Tempo na posição %d incorreto. Esperado: %d, Obtido: %d",
					i, j, expectedTimes[j], routine.Times()[j])
			}
		}

		// Verifica pares entrada/saída
		for k := 0; k < len(routine.Times()); k += 2 {
			// Verifica se há um par completo (entrada e saída)
			if k+1 >= len(routine.Times()) {
				t.Errorf("[Rotina %d] Número ímpar de tempos: %d", i, len(routine.Times()))
				break
			}

			entry := routine.Times()[k]    // Tempo de entrada
			exit := routine.Times()[k+1]  // Tempo de saída

			// Verifica se o gap mínimo foi respeitado
			if exit-entry < shift {
				t.Errorf("[Rotina %d] Gap inválido entre %d (entrada) e %d (saída). Esperado: mínimo %d",
					i, entry, exit, shift)
			}
		}
	}

	t.Logf("Geradas %d rotinas consistentes", numRoutines)
}

func TestGenerateDataBatch(t *testing.T) {
	// Configuração determinística
	const (
		numRoutines = 1000   // Quantidade de rotinas a gerar
		symbol      = "home" // Símbolo esperado
		shift   = 5      // Gap mínimo entre entrada/saída
	)

	// 1. Cria uma distribuição determinística (sempre retorna 10)
	mockDist, _ := dists.NewNormalDist(10, 0) // Média 10, desvio 0 = sempre 10

	// 2. Cria o perfil com 2 eventos (entrada e saída)
	profile, err := NewRoutineProfileDist(
		shift,
		[]dists.Distribution{mockDist, mockDist},
	)
	if err != nil {
		t.Fatalf("Falha ao criar perfil: %v", err)
	}

	// 3. Gera as rotinas
	src := rand.NewPCG(42,1)
	rng := rand.New(src) // Seed fixa para reprodutibilidade
	routines := make([]*residentdata.Routine, 0, numRoutines)

	for i := 0; i < numRoutines; i++ {
		routine := profile.GenerateData(rng)
		routines = append(routines, routine)
	}

	// 4. Verificações
	expectedTimes := []int32{10, 15} // Entrada: 10, Saída: 10 + shift = 15

	for i, routine := range routines {

		// Verifica quantidade de tempos
		if len(routine.Times()) != len(expectedTimes) {
			t.Fatalf("[Rotina %d] Número de tempos incorreto. Esperado: %d, Obtido: %d",
				i, len(expectedTimes), len(routine.Times()))
		}

		// Verifica valores e gap mínimo
		for j := 0; j < len(expectedTimes); j++ {
			if routine.Times()[j] != expectedTimes[j] {
				t.Errorf("[Rotina %d] Tempo na posição %d incorreto. Esperado: %d, Obtido: %d",
					i, j, expectedTimes[j], routine.Times()[j])
			}
		}

		// Verifica pares entrada/saída
		for k := 0; k < len(routine.Times()); k += 2 {
			if k+1 >= len(routine.Times()) {
				break
			}

			entry := routine.Times()[k]
			exit := routine.Times()[k+1]

			// O gap deve ser exatamente igual à shift
			if exit-entry != shift {
				t.Errorf("[Rotina %d] Gap inválido entre %d (entrada) e %d (saída). Esperado: %d",
					i, entry, exit, shift)
			}
		}
	}

	t.Logf("Geradas %d rotinas consistentes", numRoutines)
}

func TestGenerateDataReal(t *testing.T) {
	const (
		numRoutines = 100000   // Quantidade de rotinas a gerar
		symbol      = "clt"   // Símbolo esperado
		shift   = 60*15 // 15 minutos
	)

	// 1. Cria uma distribuição determinística
	wakeUpDist, _ := dists.NewNormalDist(5*60*60, 30*60)
	leaveDist, _ := dists.NewNormalDist(7*60*60, 30*60)
	returnDist, _ := dists.NewNormalDist(18*60*60, 30*60)
	sleepDist, _ := dists.NewNormalDist(22*60*60, 30*60)

	// 2. Cria o perfil
	profile, err := NewRoutineProfileDist(
		shift,
		[]dists.Distribution{wakeUpDist, leaveDist, returnDist, sleepDist},
	)
	if err != nil {
		t.Fatalf("Falha ao criar perfil: %v", err)
	}

	// 3. Gera as rotinas
	src := rand.NewPCG(42,1)
	rng := rand.New(src) // Seed fixa para reprodutibilidade
	routines := make([]*residentdata.Routine, 0, numRoutines)

	for i := 0; i < numRoutines; i++ {
		routine := profile.GenerateData(rng)
		routines = append(routines, routine)
	}

	// 4. Verificações
	var length int = 4

	for i, routine := range routines {


		// Verifica quantidade de tempos
		if len(routine.Times()) != length {
			t.Fatalf("[Rotina %d] Número de tempos incorreto. Esperado: %d, Obtido: %d",
				i, length, len(routine.Times()))
		}

		// Verifica pares entrada/saída
		for k := 0; k < len(routine.Times()); k += 2 {
			// Verifica se há um par completo (entrada e saída)
			if k+1 >= len(routine.Times()) {
				t.Errorf("[Rotina %d] Número ímpar de tempos: %d", i, len(routine.Times()))
				break
			}

			entry := routine.Times()[k]    // Tempo de entrada
			exit := routine.Times()[k+1]  // Tempo de saída

			// Verifica se o gap mínimo foi respeitado
			if exit-entry < shift {
				t.Errorf("[Rotina %d] Gap inválido entre %d (entrada) e %d (saída). Esperado: mínimo %d",
					i, entry, exit, shift)
			}
		}
	}
	println(routines[0].Times()[0])
	println(routines[0].Times()[1])
	println(routines[0].Times()[2])
	println(routines[0].Times()[3])
	
	/*
	// 5. Criação dos histogramas
	p := plot.New()
	p.Title.Text = "Distribuição de Atividades"
	p.X.Label.Text = "Hora do Dia"
	p.Y.Label.Text = "Frequência"
	p.Legend.Top = true
	p.Legend.Left = false

	// Bins de 1 hora (0-23h)
	bins := make([]float64, 24)
	for h := 0; h < 24; h++ {
			bins[h] = float64(h * 3600)
	}

	// Cores com transparência (RGBA)
	colors := []color.Color{
			color.RGBA{R: 0, G: 0, B: 255, A: 128},   // Azul - Acordar
			color.RGBA{R: 50, G: 205, B: 50, A: 128}, // Verde - Sair
			color.RGBA{R: 255, G: 0, B: 0, A: 128},   // Vermelho - Voltar
			color.RGBA{R: 128, G: 0, B: 128, A: 128}, // Roxo - Dormir
	}

	// Função para criar barras preenchidas
	createBars := func(times []float64, name string, clr color.Color) *plotter.BarChart {
			counts := make(plotter.Values, len(bins))
			
			// Contagem manual nos bins
			for _, t := range times {
					for i := 0; i < len(bins); i++ {
							if t >= bins[i] && t < bins[i]+3600 {
									counts[i]++
									break
							}
					}
			}
			
			// Cria barras preenchidas
			bars, err := plotter.NewBarChart(counts, 0.8)
			if err != nil {
					t.Fatalf("Erro ao criar barras: %v", err)
			}
			bars.Color = clr
			bars.LineStyle.Width = 0 // Remove bordas das barras
			return bars
	}

	// Extrai os tempos
	var wakeUp, leave, ret, sleep []float64
	for _, r := range routines {
			t := r.Times()
			wakeUp = append(wakeUp, float64(t[0]))
			leave = append(leave, float64(t[1]))
			ret = append(ret, float64(t[2]))
			sleep = append(sleep, float64(t[3]))
	}

	// Adiciona as séries
	wakeUpBars := createBars(wakeUp, "Acordar", colors[0])
	leaveBars := createBars(leave, "Sair", colors[1])
	retBars := createBars(ret, "Voltar", colors[2])
	sleepBars := createBars(sleep, "Dormir", colors[3])

	p.Add(wakeUpBars, leaveBars, retBars, sleepBars)

	// Adiciona legendas manualmente
	p.Legend.Add("Acordar", wakeUpBars)
	p.Legend.Add("Sair", leaveBars)
	p.Legend.Add("Voltar", retBars)
	p.Legend.Add("Dormir", sleepBars)

	// Configura eixo X para mostrar horas
	p.X.Tick.Marker = plot.TickerFunc(func(min, max float64) []plot.Tick {
			ticks := make([]plot.Tick, 24)
			for h := 0; h < 24; h++ {
					ticks[h] = plot.Tick{
							Value: float64(h * 3600),
							Label: fmt.Sprintf("%d", h),
					}
			}
			return ticks
	})

	// Salva o gráfico
	if err := p.Save(15*vg.Inch, 8*vg.Inch, "times_distribution.png"); err != nil {
			t.Fatal(err)
	}

	t.Logf("Geradas %d rotinas consistentes", numRoutines)*/
}