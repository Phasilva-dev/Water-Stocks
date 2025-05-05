package interfaces

type SanitaryHouse interface {
	Toilet() SanitaryDeviceInstance
	Shower() SanitaryDeviceInstance
	WashBassin() SanitaryDeviceInstance
	WashMachine() SanitaryDeviceInstance
	DishWasher() SanitaryDeviceInstance
	Tanque() SanitaryDeviceInstance
}