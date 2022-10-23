package main

import (
	"github.com/EliriaT/tmps/lab2/Clothes"
	"github.com/EliriaT/tmps/lab2/HouseBuilder"
	"github.com/EliriaT/tmps/lab2/People"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	Clothes.MakeSummerClothList()
	Clothes.MakeWinterClothList()

	villaBuilder, err := HouseBuilder.GetBuilder("villa")
	if err != nil {
		log.Fatal(err)
	}
	flatBuilder, err := HouseBuilder.GetBuilder("flat")
	if err != nil {
		log.Fatal(err)
	}

	director := HouseBuilder.NewDirector(villaBuilder)

	people := make([]People.Person, 3, 3)
	people[0].Name = "Alex"
	people[1].Name = "Alina"
	people[2].Name = "Victor"
	people[0].House = director.BuildHouse()
	people[1].House = director.BuildHouse()
	// the third person will live in a flat
	director.SetBuilder(flatBuilder)
	people[2].House = director.BuildHouse()

	for _, person := range people {
		person.House.PrintHouse()
	}
	for i := range people {
		go people[i].BuyClothes()
	}
	time.Sleep(time.Millisecond * 1000)
}
