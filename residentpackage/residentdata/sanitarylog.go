package residentdata


type SanitaryLog struct {
	sanitaryType     string
	sanitaryDeviceID uint16
	usageLogs        []*UsageLog
}

func NewSanitaryLog(sanitaryType string, sanitaryDeviceID uint16) *SanitaryLog {
	return &SanitaryLog{
		sanitaryType:     sanitaryType,
		sanitaryDeviceID: sanitaryDeviceID,
		usageLogs:        make([]*UsageLog, 0), // Initialize empty slice
	}
}

// AddUsageLog adds a new usage log to the sanitary device
func (s *SanitaryLog) AddUsageLog(log *UsageLog) {
	s.usageLogs = append(s.usageLogs, log)
}

// Getters
func (s *SanitaryLog) GetSanitaryType() string {
	return s.sanitaryType
}

func (s *SanitaryLog) GetSanitaryDeviceID() uint16 {
	return s.sanitaryDeviceID
}

// GetUsageLogs returns all usage logs
func (s *SanitaryLog) GetUsageLogs() []*UsageLog {
	return s.usageLogs
}

func (s *SanitaryLog) ClearUsageLogs() {
	s.usageLogs = make([]*UsageLog, 0)
}