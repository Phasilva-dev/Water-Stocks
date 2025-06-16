package log

import (
    //"fmt"
)

type Usage struct {
    startUsage int32
    endUsage   int32
    flowRate   float64
}

// NewUsageLog cria uma nova instância de UsageLog
func NewUsage(startUsage int32, endUsage int32, flowRate float64) *Usage {

    /*if startUsage < 0 || endUsage < 0 {
        fmt.Printf("[WARN] Uso com horários negativos: start=%d, end=%d\n", startUsage, endUsage)
    }
    if endUsage < startUsage {
        fmt.Printf("[WARN] endUsage (%d) menor que startUsage (%d)\n", endUsage, startUsage)
    }
    if flowRate <= 0 {
        fmt.Printf("[WARN] flowRate inválido: %.2f\n", flowRate)
    }*/

    return &Usage{
        startUsage: startUsage,
        endUsage:   endUsage,
        flowRate:   flowRate,
    }
}

// GetStartUsage retorna o tempo de início de uso
func (u *Usage) StartUsage() int32 {
    return u.startUsage
}

// GetEndUsage retorna o tempo de fim de uso
func (u *Usage) EndUsage() int32 {
    return u.endUsage
}

// GetFlowRate retorna a taxa de fluxo
func (u *Usage) FlowRate() float64 {
    return u.flowRate
}

