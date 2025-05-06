package log

type Usage struct {
    startUsage int32
    endUsage   int32
    flowRate   float64
}

// NewUsageLog cria uma nova instância de UsageLog
func NewUsage(startUsage int32, endUsage int32, flowRate float64) *Usage {
    return &Usage{
        startUsage: startUsage,
        endUsage:   endUsage,
        flowRate:   flowRate,
    }
}

// GetStartUsage retorna o tempo de início de uso
func (u *Usage) GetStartUsage() int32 {
    return u.startUsage
}

// GetEndUsage retorna o tempo de fim de uso
func (u *Usage) GetEndUsage() int32 {
    return u.endUsage
}

// GetFlowRate retorna a taxa de fluxo
func (u *Usage) GetFlowRate() float64 {
    return u.flowRate
}

