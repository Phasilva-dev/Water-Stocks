package count

import (
	"math/rand/v2"
	"testing"

	"simulation/internal/dists"
)

// ResidentCount Test

func TestNewResidentCount(t *testing.T) {
	dist, err := dists.CreateDistribution("normal", 5, 1)
	if err != nil {
		t.Fatalf("failed to create distribution: %v", err)
	}

	rc := NewResidentCount(dist)
	if rc == nil {
		t.Fatal("expected non-nil ResidentCount")
	}
	if rc.dist == nil {
		t.Error("expected non-nil distribution in ResidentCount")
	}
}

func TestGenerateData_RangeCheck(t *testing.T) {
	dist, _ := dists.CreateDistribution("normal", 5, 1)
	rc := NewResidentCount(dist)

	rng := rand.New(rand.NewPCG(42, 54))

	for i := 0; i < 100; i++ {
		value := rc.GenerateData(rng)
		if value < 0 || value > 255 {
			t.Errorf("expected value between 0 and 255, got %d", value)
		}
	}
}

func TestGenerateData_DeterministicBehavior(t *testing.T) {
	// Usando uma distribuição com desvio zero (valor fixo)
	dist, _ := dists.CreateDistribution("normal", 7, 0)
	rc := NewResidentCount(dist)

	rng := rand.New(rand.NewPCG(1, 1))

	value := rc.GenerateData(rng)

	expected := uint8(7)
	if value != expected {
		t.Errorf("expected %d, got %d", expected, value)
	}
}

func TestGenerateData_UpperBound(t *testing.T) {
	// Cria uma distribuição que gera um valor muito alto
	dist, _ := dists.CreateDistribution("normal", 1000, 0)
	rc := NewResidentCount(dist)

	rng := rand.New(rand.NewPCG(1, 1))

	value := rc.GenerateData(rng)

	if value != 255 {
		t.Errorf("expected capped value 255, got %d", value)
	}
}

func TestGenerateData_LowerBound(t *testing.T) {
	// Cria uma distribuição que gera um valor muito baixo (negativo)
	dist, _ := dists.CreateDistribution("normal", -1000, 0)
	rc := NewResidentCount(dist)

	rng := rand.New(rand.NewPCG(1, 1))

	value := rc.GenerateData(rng)

	if value != 1 {
		t.Errorf("expected capped value 1, got %d", value)
	}
}

//SanitaryCount Test


func TestNewSanitaryCount(t *testing.T) {
	s := NewSanitaryCount()
	if s == nil {
		t.Fatal("expected non-nil SanitaryCount")
	}
}

func TestGenerateData_ErrorWhenNoResidents(t *testing.T) {
	s := NewSanitaryCount()
	rng := rand.New(rand.NewPCG(1, 1))

	_, err := s.GenerateData(rng, 0)
	if err == nil {
		t.Error("expected error for numResidents <= 0, got nil")
	}

	// Compara a string da mensagem de erro diretamente.
	expectedErrMsg := "invalid sanitaryCount data: house without residents"
	if err.Error() != expectedErrMsg {
		t.Errorf("expected error message %q, got %q", expectedErrMsg, err.Error())
	}

}

func TestGenerateData_ValidRange(t *testing.T) {
	s := NewSanitaryCount()
	rng := rand.New(rand.NewPCG(42, 54))

	for residents := uint8(1); residents <= 20; residents++ {
		for i := 0; i < 100; i++ {
			value, err := s.GenerateData(rng, residents)
			if err != nil {
				t.Fatalf("unexpected error for residents=%d: %v", residents, err)
			}
			if value < 1 || value > 4 {
				t.Errorf("expected value between 1 and 4, got %d for residents=%d", value, residents)
			}
		}
	}
}

func TestGenerateData_Deterministic(t *testing.T) {
	s := NewSanitaryCount()
	rng := rand.New(rand.NewPCG(100, 1)) //Essa seed faz value ser 1

	value, err := s.GenerateData(rng, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Como usamos uma seed fixa, esse valor deve ser sempre o mesmo
	expected := uint8(1) // Ajuste se necessário conforme sua distribuição real

	if value != expected {
		t.Errorf("expected %d, got %d", expected, value)
	}
}