package misc

import (
    "math/rand/v2"
    "testing"
    "time"
)

// --- Teste de validação ---
func TestValidateValues(t *testing.T) {
    tests := []struct {
        name    string
        input   []Tuple[string, float64]
        wantErr bool
    }{
        {
            name: "Valid",
            input: []Tuple[string, float64]{
                {key: "a", value: 30},
                {key: "b", value: 70},
            },
            wantErr: false,
        },
        {
            name: "Exceeds 100%",
            input: []Tuple[string, float64]{
                {key: "a", value: 80},
                {key: "b", value: 30},
            },
            wantErr: true,
        },
        {
            name: "Negative value",
            input: []Tuple[string, float64]{
                {key: "a", value: -10},
                {key: "b", value: 40},
            },
            wantErr: true,
        },
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            err := validateValues(tc.input)
            if (err != nil) != tc.wantErr {
                t.Errorf("expected error: %v, got: %v", tc.wantErr, err)
            }
        })
    }
}

// --- Teste de construção do PercentSelector ---
func TestNewPercentSelector(t *testing.T) {
    valid := []Tuple[string, float64]{
        {key: "x", value: 60},
        {key: "y", value: 40},
    }
    selector, err := NewPercentSelector(valid)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    if selector == nil || len(selector.values) != 2 {
        t.Errorf("Expected selector with 2 values, got: %v", selector)
    }

    // Verificar se os valores cumulativos estão corretos
    expectedCumulative := []float64{60, 100}
    for i, tuple := range selector.values {
        if tuple.Value() != expectedCumulative[i] {
            t.Errorf("Expected cumulative value %.2f at index %d, got %.2f", expectedCumulative[i], i, tuple.Value())
        }
    }
}

// --- Teste da amostragem probabilística ---
func TestSampleDistribution(t *testing.T) {
    probs := []Tuple[string, float64]{
        {key: "a", value: 20.0},
        {key: "b", value: 80.0},
    }
    selector, err := NewPercentSelector(probs)
    if err != nil {
        t.Fatalf("Failed to create selector: %v", err)
    }

    rng := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), 1))

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

    // Tolerância de 5% para flutuações estatísticas
    if ratioA < 0.15 || ratioA > 0.25 {
        t.Errorf("Sample ratio for 'a' out of expected range: got %.2f", ratioA)
    }
    if ratioB < 0.75 || ratioB > 0.85 {
        t.Errorf("Sample ratio for 'b' out of expected range: got %.2f", ratioB)
    }
}