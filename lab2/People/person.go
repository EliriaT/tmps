package People

import (
	"github.com/EliriaT/tmps/lab2/Clothes"
	"github.com/EliriaT/tmps/lab2/HouseBuilder"
	"log"
	"math/rand"
)

type Person struct {
	Name          string
	House         HouseBuilder.House
	BoughtClothes []string
}

func (p *Person) BuyClothes() {
	numItems := Clothes.ClothPrototypesList.Length()

	randomNr := rand.Intn(numItems) + 1

	for i := 0; i < randomNr; i++ {
		num := rand.Intn(numItems) + 1
		cloth, err := Clothes.ClothPrototypesList.Get(num)
		if err != nil {
			log.Fatalf("could not get this cloth %d", num)
		}
		//log.Println(cloth.GetName())
		p.BoughtClothes = append(p.BoughtClothes, cloth.GetName())
		log.Printf("Person %s bought %s. Total inventory: %s \n", p.Name, cloth.GetName(), p.BoughtClothes)
	}

}
