package globals

import (
	"models"
)

// Shower map
var showers = make(map[uint32]*models.SanitaryDevice)

// GetShower retrieves a shower profile by ID
func GetShower(id uint32) (*models.SanitaryDevice, bool) {
	mu.RLock()
	defer mu.RUnlock()
	shower, exists := showers[id]
	return shower, exists
}

// SetShower stores or updates a shower profile
func SetShower(id uint32, profile *models.SanitaryDevice) {
	mu.Lock()
	defer mu.Unlock()
	showers[id] = profile
}

// DeleteShower removes a shower profile by ID
func DeleteShower(id uint32) {
	mu.Lock()
	defer mu.Unlock()
	delete(showers, id)
}