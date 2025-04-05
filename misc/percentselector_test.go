package misc

import (
	"testing"
	"time"
	"math/rand/v2"
)

// --- Teste de validação ---

func TestValidateValues(t *testing.T) {
	tests := []struct {
		name    string
		input   map[string]float64
		wantErr bool
	}{
		{"Valid", map[string]float64{"a": 30, "b": 70}, false},
		{"Exceeds 100%", map[string]float64{"a": 80, "b": 30}, true},
		{"Negative value", map[string]float64{"a": -10, "b": 40}, true},
	}

	for _, tc := range tests {
		err := validateValues(tc.input)
		if (err != nil) != tc.wantErr {
			t.Errorf("Test %q failed: expected error: %v, got: %v", tc.name, tc.wantErr, err)
		}
	}
}

// --- Teste de construção do PercentSelector ---

func TestNewPercentSelector(t *testing.T) {
	valid := map[string]float64{"x": 60, "y": 40}
	selector, err := NewPercentSelector(valid)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if selector == nil || len(selector.values) != 2 {
		t.Errorf("Expected selector with 2 values, got: %v", selector)
	}
}

// --- Teste da amostragem probabilística ---

func TestSampleDistribution(t *testing.T) {
	probs := map[string]float64{
		"a": 20.0,
		"b": 80.0,
	}
	selector, err := NewPercentSelector(probs)
	if err != nil {
		t.Fatalf("Failed to create selector: %v", err)
	}

	rng := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()),1))

	counts := map[string]int{"a": 0, "b": 0}
	runs := 100000

	for i := 0; i < runs; i++ {
		sample, err := selector.Sample(rng)
		if err != nil {
			t.Fatalf("Sample error: %v", err)
		}
		counts[sample]++
	}

	ratioA := float64(counts["a"]) / float64(runs)
	ratioB := float64(counts["b"]) / float64(runs)

	if ratioA < 0.15 || ratioA > 0.25 {
		t.Errorf("Sample ratio for 'a' out of expected range: got %.2f", ratioA)
	}
	if ratioB < 0.75 || ratioB > 0.85 {
		t.Errorf("Sample ratio for 'b' out of expected range: got %.2f", ratioB)
	}
}
