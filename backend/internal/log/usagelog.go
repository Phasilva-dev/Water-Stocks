package log

import (
    "fmt"
)

type Usage struct {
    startUsage int32
    endUsage   int32
    flowRate   float64
}

// NewUsageLog cria uma nova instância de UsageLog
func NewUsage(startUsage int32, endUsage int32, flowRate float64) (*Usage, error) {

    /*if startUsage < 0 || endUsage < 0 {
        fmt.Printf("[WARN] Uso no dia anterior: start=%d, end=%d\n", startUsage, endUsage)
    }
    if startUsage > 86400 || endUsage > 86400 {
        fmt.Printf("[WARN] Uso no dia seguinte: start=%d, end=%d\n", startUsage, endUsage)
    }*/
    if endUsage < startUsage {
        return nil, fmt.Errorf("[ERROR] endUsage (%d) menor que startUsage (%d)\n", endUsage, startUsage)
    }
    if flowRate <= 0 {
        return nil, fmt.Errorf("[ERROR] flowRate inválido: %.2f\n", flowRate)
    }

    

    return &Usage{
        startUsage: startUsage,
        endUsage:   endUsage,
        flowRate:   flowRate,
    }, nil
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

