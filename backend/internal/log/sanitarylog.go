package log


type Sanitary struct {
	sanitaryType     string
	sanitaryDeviceID uint32
	usageLogs        []*Usage
}

func NewSanitary(sanitaryType string, sanitaryDeviceID uint32) *Sanitary {
	return &Sanitary{
		sanitaryType:     sanitaryType,
		sanitaryDeviceID: sanitaryDeviceID,
		usageLogs:        make([]*Usage, 0), // Initialize empty slice
	}
}

// AddUsageLog adds a new usage log to the sanitary device
func (s *Sanitary) AddUsageLog(log *Usage) {
	s.usageLogs = append(s.usageLogs, log)
}

// Getters
func (s *Sanitary) SanitaryType() string {
	return s.sanitaryType
}

func (s *Sanitary) SanitaryDeviceID() uint32 {
	return s.sanitaryDeviceID
}

// GetUsageLogs returns all usage logs
func (s *Sanitary) UsageLogs() ([]*Usage, bool) {
	if len(s.usageLogs) == 0{
		return nil,false
	}
	return s.usageLogs,true
}

func (s *Sanitary) ClearUsageLogs() {
	s.usageLogs = make([]*Usage, 0)
}