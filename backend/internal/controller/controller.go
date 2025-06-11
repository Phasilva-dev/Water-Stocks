package controller

import (
	//"fmt"
	"fmt"
	"log"

	"simulation/internal/entities"
	"simulation/internal/entities/house"


	"math/rand/v2"
	"time"
)



func setHouses(profile *house.HouseProfile, houses []*entities.House, size int, rng *rand.Rand) {

	for i := 0; i < size; i++ {

		houses[i] = entities.NewHouse(1,profile)
		if err := houses[i].GenerateHouseData(rng); err != nil {
			log.Fatalf("Erro ao criar a casa %d : %v", i, err)
		} 

	}
	fmt.Println("Casas criadas ")

}


func RunSimulation(size, day, toiletType, showerType int) {

	
	rng := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), 0))

	profile := defaultHouseProfile(toiletType, showerType)
	houses := make([]*entities.House, size)

	setHouses(profile, houses, size, rng)

	for i := uint8(0); i < uint8(day); i++ {
		for j := 0; j < size; j++ {
			if err := houses[j].GenerateLogs(i, rng); err != nil {
			log.Fatalf("Erro ao gerar logs da casa %d no dia %d: %v", j, i, err)
		}
			//fmt.Println(j)
		}
	}

	fmt.Println("Passou")
	/*residents := houses[0].ResidentLogs()
	if len(residents) == 0 || residents[0] == nil {
    	log.Fatal("Erro: ResidentLogs vazio ou com valor nulo")
	}
	PrintLogLines(residents[0].ToLogLines())*/
	
}