package globals

import (
	"models"
)

// Tanque map
var tanques = make(map[uint32]*models.SanitaryDevice)

// GetTanque retrieves a tanque profile by ID
func GetTanque(id uint32) (*models.SanitaryDevice, bool) {
	mu.RLock()
	defer mu.RUnlock()
	tanque, exists := tanques[id]
	return tanque, exists
}

// SetTanque stores or updates a tanque profile
func SetTanque(id uint32, profile *models.SanitaryDevice) {
	mu.Lock()
	defer mu.Unlock()
	tanques[id] = profile
}

// DeleteTanque removes a tanque profile by ID
func DeleteTanque(id uint32) {
	mu.Lock()
	defer mu.Unlock()
	delete(tanques, id)
}