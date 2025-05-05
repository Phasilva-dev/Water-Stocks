package interfaces

type DailyData interface {
	Routine() Routine
	Frequency() Frequency
	SetRoutine(r Routine)
	SetFrequency(f Frequency)
}

type Frequency interface {
	FreqToilet() uint8
	FreqShower() uint8
	FreqWashBassin() uint8
	FreqWashMachine() uint8
	FreqDishWasher() uint8
	FreqTanque() uint8
}

type Routine interface {
	Times() []int32
	SleepTime() int32
	WakeupTime() int32
	ReturnHome() int32
	WorkTime() int32
}

type UsageLog interface {
	GetStartUsage() int32
	GetEndUsage() int32
	GetFlowRate() float64
}

type SanitaryLog interface {
	AddUsageLog(log UsageLog)
	GetSanitaryType() string
	GetSanitaryDeviceID() uint32
	GetUsageLogs() []UsageLog
	ClearUsageLogs()
}

// ResidentSanitaryLog interface representa os métodos disponíveis para a struct ResidentSanitaryLog.
type ResidentSanitaryLog interface {
	GetToiletLog() SanitaryLog
	GetShowerLog() SanitaryLog
	GetWashBassinLog() SanitaryLog
	GetWashMachineLog() SanitaryLog
	GetDishWasherLog() SanitaryLog
	GetTanqueLog() SanitaryLog
}

