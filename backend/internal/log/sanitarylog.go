package log


type Sanitary struct {
	sanitaryType     string
	sanitaryDeviceID uint32
	usageLogs        []*Usage
}

func NewSanitary(sanitaryType string, sanitaryDeviceID uint32, usages []*Usage) *Sanitary {
	return &Sanitary{
		sanitaryType:     sanitaryType,
		sanitaryDeviceID: sanitaryDeviceID,
		usageLogs:        usages, // Initialize empty slice
	}
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