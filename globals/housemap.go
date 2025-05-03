package globals

import (
	"interfaces"
	"sync"
)

// Global storage for house profiles
var (
	houses = make(map[uint32]interfaces.HouseProfile)
	// mu ensures thread-safe access to the maps
	mu sync.RWMutex
)

// GetHouse retrieves a house profile by ID
func GetHouse(id uint32) (interfaces.HouseProfile, bool) {
	mu.RLock()
	defer mu.RUnlock()
	house, exists := houses[id]
	return house, exists
}

// SetHouse stores or updates a house profile
func SetHouse(id uint32, profile interfaces.HouseProfile) {
	mu.Lock()
	defer mu.Unlock()
	houses[id] = profile
}

// DeleteHouse removes a house profile by ID
func DeleteHouse(id uint32) {
	mu.Lock()
	defer mu.Unlock()
	delete(houses, id)
}