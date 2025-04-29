package globals

import (
	"models"
)

// Global storage for resident profiles
var (
	residents = make(map[uint32]*models.ResidentProfile)
)

// GetResident retrieves a resident profile by ID
func GetResident(id uint32) (*models.ResidentProfile, bool) {
	mu.RLock()
	defer mu.RUnlock()
	resident, exists := residents[id]
	return resident, exists
}

// SetResident stores or updates a resident profile
func SetResident(id uint32, profile *models.ResidentProfile) {
	mu.Lock()
	defer mu.Unlock()
	residents[id] = profile
}

// DeleteResident removes a resident profile by ID
func DeleteResident(id uint32) {
	mu.Lock()
	defer mu.Unlock()
	delete(residents, id)
}