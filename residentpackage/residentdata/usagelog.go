package residentdata

type UsageLog struct {
    startUsage int32
    endUsage   int32
    flowRate   float64
}

// NewUsageLog cria uma nova instância de UsageLog
func NewUsageLog(startUsage int32, endUsage int32, flowRate float64) *UsageLog {
    return &UsageLog{
        startUsage: startUsage,
        endUsage:   endUsage,
        flowRate:   flowRate,
    }
}

// GetStartUsage retorna o tempo de início de uso
func (u *UsageLog) GetStartUsage() int32 {
    return u.startUsage
}

// GetEndUsage retorna o tempo de fim de uso
func (u *UsageLog) GetEndUsage() int32 {
    return u.endUsage
}

// GetFlowRate retorna a taxa de fluxo
func (u *UsageLog) GetFlowRate() float64 {
    return u.flowRate
}

