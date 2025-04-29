package globals

import (
	"sanitarydevice"
)

// WashMachine map
var washMachines = make(map[uint32]*sanitarydevice.SanitaryDevice)

// GetWashMachine retrieves a wash machine profile by ID
func GetWashMachine(id uint32) (*sanitarydevice.SanitaryDevice, bool) {
	mu.RLock()
	defer mu.RUnlock()
	washMachine, exists := washMachines[id]
	return washMachine, exists
}

// SetWashMachine stores or updates a wash machine profile
func SetWashMachine(id uint32, profile *sanitarydevice.SanitaryDevice) {
	mu.Lock()
	defer mu.Unlock()
	washMachines[id] = profile
}

// DeleteWashMachine removes a wash machine profile by ID
func DeleteWashMachine(id uint32) {
	mu.Lock()
	defer mu.Unlock()
	delete(washMachines, id)
}