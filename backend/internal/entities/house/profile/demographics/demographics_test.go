package demographics

import (
	"math"
	"math/rand/v2"
	"testing"
	"reflect"

	"simulation/internal/dists"
	"simulation/internal/misc"
)

// Testa a criação de Age e se ela retorna a distribuição corretamente
func TestNewAge(t *testing.T) {
	dist, _ := dists.CreateDistribution("normal", 30, 5)
	age := NewAge(dist)

	if age.AgeDist() != dist {
		t.Error("AgeDist does not return the correct distribution")
	}
}

// Testa se o valor gerado está dentro do intervalo correto (0-255)
func TestGenerateDataWithinBounds(t *testing.T) {
	dist, _ := dists.CreateDistribution("normal", 30, 5)
	age := NewAge(dist)

	rng := rand.New(rand.NewPCG(1, 1))

	for i := 0; i < 1000; i++ {
		value := age.GenerateData(rng)
		if value < 0 || value > 255 {
			t.Errorf("Generated value out of bounds: got %d", value)
		}
	}
}

// Testa se valores negativos são tratados corretamente (convertidos para positivos)
func TestGenerateDataHandlesNegative(t *testing.T) {
	// Distribuição com média negativa para forçar valores negativos
	dist, _ := dists.CreateDistribution("normal", -50, 0)
	age := NewAge(dist)

	rng := rand.New(rand.NewPCG(2, 2))

	value := age.GenerateData(rng)
	if value != 50 {
		t.Errorf("Expected 50 from abs(-50), got %d", value)
	}
}

// Testa se valores muito altos são limitados para 255
func TestGenerateDataClampsToMaxUint8(t *testing.T) {
	dist, _ := dists.CreateDistribution("normal", 500, 0)
	age := NewAge(dist)

	rng := rand.New(rand.NewPCG(3, 3))

	value := age.GenerateData(rng)
	if value != math.MaxUint8 {
		t.Errorf("Expected value to be clamped to 255, got %d", value)
	}
}















// Teste de Occupation

func createOccupationTestData() (
	[]misc.Tuple[uint32, float64],
	[]misc.Tuple[uint32, float64],
	[]misc.Tuple[uint32, float64],
	*misc.PercentSelector[uint32],
	*misc.PercentSelector[uint32],
	*misc.PercentSelector[uint32],
	error,
) {
	// Definindo os dados brutos para cada faixa etária
	under18 := []misc.Tuple[uint32, float64]{
		*misc.NewTuple(uint32(1), 60.0), // Ex: Estudante 1 (60%)
		*misc.NewTuple(uint32(2), 40.0), // Ex: Estudante 2 (40%)
	}

	adult := []misc.Tuple[uint32, float64]{
		*misc.NewTuple(uint32(3), 50.0), // Ex: Trabalhador (50%)
		*misc.NewTuple(uint32(4), 50.0), // Ex: Desempregado/Outro (50%)
	}

	over65 := []misc.Tuple[uint32, float64]{
		*misc.NewTuple(uint32(5), 100.0), // Ex: Aposentado (100%)
	}

	// Criando os PercentSelectors a partir dos dados brutos
	u18Sel, err := misc.NewPercentSelector(under18)
	if err != nil {
		// Retorna os dados brutos também, caso o chamador precise deles mesmo em caso de erro
		return under18, adult, over65, nil, nil, nil, err
	}

	adultSel, err := misc.NewPercentSelector(adult)
	if err != nil {
		return under18, adult, over65, u18Sel, nil, nil, err
	}

	over65Sel, err := misc.NewPercentSelector(over65)
	if err != nil {
		return under18, adult, over65, u18Sel, adultSel, nil, err
	}

	// Retorna os dados brutos E os seletores criados com sucesso
	return under18, adult, over65, u18Sel, adultSel, over65Sel, nil
}


// setupOccupationTestData é uma função auxiliar para obter os dados de teste
// e seletores, tratando qualquer erro fatalmente no teste.
func setupOccupationTestData(t *testing.T) (
	[]misc.Tuple[uint32, float64],
	[]misc.Tuple[uint32, float64],
	[]misc.Tuple[uint32, float64],
	*misc.PercentSelector[uint32],
	*misc.PercentSelector[uint32],
	*misc.PercentSelector[uint32],
) {
	t.Helper() // Marca esta função como um helper, para que as linhas de erro apontem para o teste chamador
	u18Data, adultData, over65Data, u18Sel, adultSel, over65Sel, err := createOccupationTestData()
	if err != nil {
		// Se a função auxiliar de teste falhar, o teste não pode prosseguir
		t.Fatalf("Falha ao criar dados de teste usando createOccupationTestData: %v", err)
	}
	return u18Data, adultData, over65Data, u18Sel, adultSel, over65Sel
}

func TestNewOccupation(t *testing.T) {
	t.Run("DadosValidos", func(t *testing.T) {
		// Arrange: Obter dados válidos da função auxiliar de teste
		_, _, _, u18Sel, adultSel, over65Sel := setupOccupationTestData(t)

		// Act: Chamar o construtor com os dados válidos
		occ, err := NewOccupation(u18Sel, adultSel, over65Sel)

		// Assert: Verificar se não houve erro e o objeto foi criado
		if err != nil {
			t.Errorf("NewOccupation com dados válidos retornou um erro inesperado: %v", err)
		}
		if occ == nil {
			t.Fatal("NewOccupation com dados válidos retornou nil")
		}
	})

	t.Run("DadosInvalidosPropagaErro", func(t *testing.T) {
		// Arrange: Criar dados que devem causar um erro no NewPercentSelector
		// Exemplo: Slice vazio geralmente é inválido para seletores baseados em percentual.
		//invalidData := []misc.Tuple[uint32, float64]{}

		// Usamos setupOccupationTestData para obter dados válidos para os outros grupos.
		_, _, _, u18Sel, adultSel, _ := setupOccupationTestData(t) // Ignora os seletores aqui

		// Act: Tentar criar Occupation com dados inválidos para um grupo
		occ, err := NewOccupation(u18Sel, adultSel, nil)

		// Assert: Verificar se um erro foi retornado e o objeto é nil
		if err == nil {
			t.Error("NewOccupation com dados inválidos era esperado retornar um erro, mas não retornou")
		}
		if occ != nil {
			t.Errorf("NewOccupation com dados inválidos era esperado retornar nil, mas retornou %v", occ)
		}
		// Opcional: Verificar se a mensagem de erro contém alguma indicação do problema
		// if !strings.Contains(err.Error(), "falha ao criar PercentSelector") { ... }
	})

	// Adicionar mais casos de teste para diferentes tipos de dados inválidos se necessário
	// (Ex: porcentagens que não somam 100%, dados com valores zero ou negativos, etc.)
}

func TestOccupationGetters(t *testing.T) {
	// Arrange: Configurar os dados de teste e criar um objeto Occupation
	_, _, _, expectedU18Sel, expectedAdultSel, expectedOver65Sel := setupOccupationTestData(t)

	occ, err := NewOccupation(expectedU18Sel, expectedAdultSel, expectedOver65Sel)
	if err != nil {
		t.Fatalf("Falha ao criar objeto Occupation para testes de getters: %v", err)
	}

	// Act & Assert: Chamar cada getter e comparar o seletor retornado
	t.Run("Under18Selector", func(t *testing.T) {
		actualSel := occ.Under18Selector()
		// Comparar ponteiros é uma forma eficaz de verificar se é o mesmo objeto seletor
		if actualSel != expectedU18Sel {
			t.Errorf("Under18Selector() retornou um ponteiro de seletor diferente. Esperado %p, Obtido %p", expectedU18Sel, actualSel)
			// Se a comparação de ponteiro não for suficiente, você precisaria comparar o *conteúdo*
			// dos seletores, o que pode ser complexo.
		}
	})

	t.Run("AdultSelector", func(t *testing.T) {
		actualSel := occ.AdultSelector()
		if actualSel != expectedAdultSel {
			t.Errorf("AdultSelector() retornou um ponteiro de seletor diferente. Esperado %p, Obtido %p", expectedAdultSel, actualSel)
		}
	})

	t.Run("Over65Selector", func(t *testing.T) {
		actualSel := occ.Over65Selector()
		if actualSel != expectedOver65Sel {
			t.Errorf("Over65Selector() retornou um ponteiro de seletor diferente. Esperado %p, Obtido %p", expectedOver65Sel, actualSel)
		}
	})
}

func TestOccupationSampling(t *testing.T) {
	// Arrange: Configurar os dados de teste e criar um objeto Occupation
	_, _, _, u18Sel, adultSel, over65Sel := setupOccupationTestData(t)
	occ, err := NewOccupation(u18Sel, adultSel, over65Sel)
	if err != nil {
		t.Fatalf("Falha ao criar objeto Occupation para testes de amostragem: %v", err)
	}

	src := rand.NewPCG(42, 0)
	rng := rand.New(src)

	// Testar o caso com 100% de probabilidade (Over65)
	t.Run("GenerateOver65Selector", func(t *testing.T) {
		// De acordo com createOccupationTestData, over65 sempre retorna 5
		expectedID := uint32(5)

		for i := 0; i < 10; i++ {
			sampledID := occ.GenerateOver65Selector(rng)
			if sampledID != expectedID {
				t.Errorf("GenerateOver65Selector() amostra %d: Esperado %d, Obtido %d", i, expectedID, sampledID)
			}

			if sampledID == 0 {
                 t.Errorf("GenerateOver65Selector() amostra %d retornou 0, indicando possível erro", i)
            }
		}
	})

	// Testar casos com probabilidades diferentes de 100% (Under18, Adult)
	// Não podemos testar uma única amostra, mas podemos testar se o valor retornado
	// está *dentro* dos possíveis resultados definidos nos dados de teste.
	// Testar a *distribuição* de probabilidade corretamente exigiria muitas amostras
	// e seria mais um teste para misc.PercentSelector do que para Occupation.
	t.Run("GenerateUnder18Selector", func(t *testing.T) {
		// Resultados possíveis: 1 (60%) ou 2 (40%) conforme createOccupationTestData
		possibleIDs := map[uint32]bool{1: true, 2: true}

		// Amostrar algumas vezes e verificar se o resultado é um dos esperados
		for i := 0; i < 20; i++ { // Aumentamos um pouco as amostras para ter mais confiança
			sampledID := occ.GenerateUnder18Selector(rng)
			if _, ok := possibleIDs[sampledID]; !ok {
				// reflect.ValueOf(possibleIDs).MapKeys() é uma forma de obter as chaves do map para a mensagem de erro
				t.Errorf("GenerateUnder18Selector() amostra %d: Obtido ID inesperado %d. Esperado um de %v", i, sampledID, reflect.ValueOf(possibleIDs).MapKeys())
			}
			if sampledID == 0 {
                t.Errorf("GenerateUnder18Selector() amostra %d retornou 0, indicando possível erro", i)
            }
		}
	})

	t.Run("GenerateAdultSelector", func(t *testing.T) {
		// Resultados possíveis: 3 (50%) ou 4 (50%) conforme createOccupationTestData
		possibleIDs := map[uint32]bool{3: true, 4: true}

		// Amostrar algumas vezes e verificar se o resultado é um dos esperados
		for i := 0; i < 20; i++ {
			sampledID := occ.GenerateAdultSelector(rng)
			if _, ok := possibleIDs[sampledID]; !ok {
				t.Errorf("GenerateAdultSelector() amostra %d: Obtido ID inesperado %d. Esperado um de %v", i, sampledID, reflect.ValueOf(possibleIDs).MapKeys())
			}
             if sampledID == 0 {
                t.Errorf("GenerateAdultSelector() amostra %d retornou 0, indicando possível erro", i)
            }
		}
	})
}

// É uma boa prática testar a própria função de setup de teste, mesmo que simples.
func TestCreateOccupationTestData(t *testing.T) {
	u18Data, adultData, over65Data, u18Sel, adultSel, over65Sel, err := createOccupationTestData()

	if err != nil {
		t.Fatalf("createOccupationTestData retornou erro: %v", err)
	}

	// Verificar se os slices de dados não são nulos/vazios (assumindo que não deveriam ser)
	if u18Data == nil || len(u18Data) == 0 {
		t.Error("Under18 data is nil or empty")
	}
	if adultData == nil || len(adultData) == 0 {
		t.Error("Adult data is nil or empty")
	}
	if over65Data == nil || len(over65Data) == 0 {
		t.Error("Over65 data is nil or empty")
	}

	// Verificar se os seletores foram criados (não nulos)
	if u18Sel == nil {
		t.Error("Under18 selector is nil")
	}
	if adultSel == nil {
		t.Error("Adult selector is nil")
	}
	if over65Sel == nil {
		t.Error("Over65 selector is nil")
	}

	// Opcional: Adicionar verificações mais detalhadas sobre o conteúdo esperado dos dados/seletores
}

