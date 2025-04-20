package resident


import(
	"residentdata"
)

type Resident struct {
	age uint8
	personID uint32
	dayData residentdata.DailyData
}