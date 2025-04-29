package globals

import (
	"models"
)

var toilets = make(map[uint32]*models.SanitaryDevice)

// GetToilet retrieves a toilet profile by ID
func GetToilet(id uint32) (*models.SanitaryDevice, bool) {
	mu.RLock()
	defer mu.RUnlock()
	toilet, exists := toilets[id]
	return toilet, exists
}

// SetToilet stores or updates a toilet profile
func SetToilet(id uint32, profile *models.SanitaryDevice) {
	mu.Lock()
	defer mu.Unlock()
	toilets[id] = profile
}

// DeleteToilet removes a toilet profile by ID
func DeleteToilet(id uint32) {
	mu.Lock()
	defer mu.Unlock()
	delete(toilets, id)
}