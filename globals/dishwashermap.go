package globals

import (
	"interfaces"
)

// DishWasher map
var dishWashers = make(map[uint32]interfaces.SanitaryDevice)

// GetDishWasher retrieves a dish washer profile by ID
func GetDishWasher(id uint32) (interfaces.SanitaryDevice, bool) {
	mu.RLock()
	defer mu.RUnlock()
	dishWasher, exists := dishWashers[id]
	return dishWasher, exists
}

// SetDishWasher stores or updates a dish washer profile
func SetDishWasher(id uint32, profile interfaces.SanitaryDevice) {
	mu.Lock()
	defer mu.Unlock()
	dishWashers[id] = profile
}

// DeleteDishWasher removes a dish washer profile by ID
func DeleteDishWasher(id uint32) {
	mu.Lock()
	defer mu.Unlock()
	delete(dishWashers, id)
}