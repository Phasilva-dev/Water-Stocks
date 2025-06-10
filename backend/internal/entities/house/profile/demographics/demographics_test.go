package demographics

import (
	"math"
	"math/rand/v2"
	"testing"
	//"reflect"

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















func createOccupationTestData() (
	[]misc.Tuple[uint32, float64],
	[]misc.Tuple[uint32, float64],
	[]misc.Tuple[uint32, float64],
	*misc.PercentSelector[uint32],
	*misc.PercentSelector[uint32],
	*misc.PercentSelector[uint32],
	error,
) {
	under18 := []misc.Tuple[uint32, float64]{
		*misc.NewTuple(uint32(1), 60.0),
		*misc.NewTuple(uint32(2), 40.0),
	}

	adult := []misc.Tuple[uint32, float64]{
		*misc.NewTuple(uint32(3), 50.0),
		*misc.NewTuple(uint32(4), 50.0),
	}

	over65 := []misc.Tuple[uint32, float64]{
		*misc.NewTuple(uint32(5), 100.0),
	}

	u18Sel, err := misc.NewPercentSelector(under18)
	if err != nil {
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

	return under18, adult, over65, u18Sel, adultSel, over65Sel, nil
}

func setupOccupationTestData(t *testing.T) (
	[]misc.Tuple[uint32, float64],
	[]misc.Tuple[uint32, float64],
	[]misc.Tuple[uint32, float64],
	*misc.PercentSelector[uint32],
	*misc.PercentSelector[uint32],
	*misc.PercentSelector[uint32],
) {
	t.Helper()
	u18Data, adultData, over65Data, u18Sel, adultSel, over65Sel, err := createOccupationTestData()
	if err != nil {
		t.Fatalf("Falha ao criar dados de teste: %v", err)
	}
	return u18Data, adultData, over65Data, u18Sel, adultSel, over65Sel
}

func TestNewOccupation(t *testing.T) {
	t.Run("DadosValidos", func(t *testing.T) {
		_, _, _, u18Sel, adultSel, over65Sel := setupOccupationTestData(t)

		u18Selector, err := NewAgeRangeSelector(0, 17, u18Sel)
		if err != nil {
			t.Fatalf("Erro criando AgeRangeSelector: %v", err)
		}
		adultSelector, err := NewAgeRangeSelector(18, 64, adultSel)
		if err != nil {
			t.Fatalf("Erro criando AgeRangeSelector: %v", err)
		}
		over65Selector, err := NewAgeRangeSelector(65, 120, over65Sel)
		if err != nil {
			t.Fatalf("Erro criando AgeRangeSelector: %v", err)
		}

		selectors := []*AgeRangeSelector{
			u18Selector,
			adultSelector,
			over65Selector,
		}

		occ, err := NewOccupation(selectors)
		if err != nil {
			t.Errorf("NewOccupation retornou erro inesperado: %v", err)
		}
		if occ == nil {
			t.Fatal("NewOccupation retornou nil com dados válidos")
		}
	})

	t.Run("SelectorsVazio", func(t *testing.T) {
		occ, err := NewOccupation([]*AgeRangeSelector{})
		if err == nil {
			t.Error("Era esperado erro para selectors vazio, mas não retornou erro")
		}
		if occ != nil {
			t.Error("Occupation deveria ser nil quando selectors são vazios")
		}
	})

	t.Run("SelectorComNil", func(t *testing.T) {
		selectors := []*AgeRangeSelector{nil}
		occ, err := NewOccupation(selectors)
		if err == nil {
			t.Error("Era esperado erro para selector nil, mas não retornou erro")
		}
		if occ != nil {
			t.Error("Occupation deveria ser nil quando selector é nil")
		}
	})

	t.Run("FaixaIdadeInvalida", func(t *testing.T) {
		_, _, _, u18Sel, _, _ := setupOccupationTestData(t)

		// Criando selector inválido com minAge > maxAge
		_, err := NewAgeRangeSelector(20, 10, u18Sel)
		if err == nil {
			t.Error("Era esperado erro para faixa etária inválida (minAge > maxAge), mas não retornou erro")
		}
	})
}

func TestOccupationSample(t *testing.T) {
	_, _, _, u18Sel, adultSel, over65Sel := setupOccupationTestData(t)

	u18Selector, _ := NewAgeRangeSelector(0, 17, u18Sel)
	adultSelector, _ := NewAgeRangeSelector(18, 64, adultSel)
	over65Selector, _ := NewAgeRangeSelector(65, 120, over65Sel)

	selectors := []*AgeRangeSelector{
		u18Selector,
		adultSelector,
		over65Selector,
	}

	occ, err := NewOccupation(selectors)
	if err != nil {
		t.Fatalf("Erro ao criar Occupation: %v", err)
	}

	src := rand.NewPCG(42, 0)
	rng := rand.New(src)

	t.Run("AmostragemUnder18", func(t *testing.T) {
		possible := map[uint32]bool{1: true, 2: true}

		for i := 0; i < 20; i++ {
			result, err := occ.Sample(10, rng)
			if err != nil {
				t.Errorf("Erro inesperado ao amostrar: %v", err)
			}
			if !possible[result] {
				t.Errorf("Valor inesperado %d para idade 10", result)
			}
		}
	})

	t.Run("AmostragemAdulto", func(t *testing.T) {
		possible := map[uint32]bool{3: true, 4: true}

		for i := 0; i < 20; i++ {
			result, err := occ.Sample(30, rng)
			if err != nil {
				t.Errorf("Erro inesperado ao amostrar: %v", err)
			}
			if !possible[result] {
				t.Errorf("Valor inesperado %d para idade 30", result)
			}
		}
	})

	t.Run("AmostragemOver65", func(t *testing.T) {
		expected := uint32(5)

		for i := 0; i < 10; i++ {
			result, err := occ.Sample(70, rng)
			if err != nil {
				t.Errorf("Erro inesperado ao amostrar: %v", err)
			}
			if result != expected {
				t.Errorf("Esperava %d para idade 70, mas obteve %d", expected, result)
			}
		}
	})

	t.Run("IdadeSemFaixa", func(t *testing.T) {
		_, err := occ.Sample(150, rng)
		if err == nil {
			t.Error("Era esperado erro para idade sem faixa, mas não retornou erro")
		}
	})
}

func TestOccupationSelectorsGetter(t *testing.T) {
	_, _, _, u18Sel, adultSel, over65Sel := setupOccupationTestData(t)

	u18Selector, _ := NewAgeRangeSelector(0, 17, u18Sel)
	adultSelector, _ := NewAgeRangeSelector(18, 64, adultSel)
	over65Selector, _ := NewAgeRangeSelector(65, 120, over65Sel)

	expectedSelectors := []*AgeRangeSelector{
		u18Selector,
		adultSelector,
		over65Selector,
	}

	occ, err := NewOccupation(expectedSelectors)
	if err != nil {
		t.Fatalf("Erro ao criar Occupation: %v", err)
	}

	actualSelectors := occ.Selectors()

	if len(actualSelectors) != len(expectedSelectors) {
		t.Fatalf("Número de selectors retornados (%d) difere do esperado (%d)", len(actualSelectors), len(expectedSelectors))
	}

	for i, expected := range expectedSelectors {
		actual := actualSelectors[i]
		if expected.MinAge() != actual.MinAge() || expected.MaxAge() != actual.MaxAge() || expected.Selector() != actual.Selector() {
			t.Errorf("Selector %d não corresponde. Esperado %+v, Obtido %+v", i, expected, actual)
		}
	}
}