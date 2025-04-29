package globals

import (
	"models"
)

// WashBasin map
var washBasins = make(map[uint32]*models.SanitaryDevice)

// GetWashBasin retrieves a wash basin profile by ID
func GetWashBasin(id uint32) (*models.SanitaryDevice, bool) {
	mu.RLock()
	defer mu.RUnlock()
	washBasin, exists := washBasins[id]
	return washBasin, exists
}

// SetWashBasin stores or updates a wash basin profile
func SetWashBasin(id uint32, profile *models.SanitaryDevice) {
	mu.Lock()
	defer mu.Unlock()
	washBasins[id] = profile
}

// DeleteWashBasin removes a wash basin profile by ID
func DeleteWashBasin(id uint32) {
	mu.Lock()
	defer mu.Unlock()
	delete(washBasins, id)
}