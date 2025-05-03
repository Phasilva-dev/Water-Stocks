package globals

import (
	"interfaces"
)

var toilets = make(map[uint32]interfaces.SanitaryDevice)

// GetToilet retrieves a toilet profile by ID
func GetToilet(id uint32) (interfaces.SanitaryDevice, bool) {
	mu.RLock()
	defer mu.RUnlock()
	toilet, exists := toilets[id]
	return toilet, exists
}

// SetToilet stores or updates a toilet profile
func SetToilet(id uint32, profile interfaces.SanitaryDevice) {
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