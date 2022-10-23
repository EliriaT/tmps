package Clothes

import (
	"github.com/EliriaT/tmps/lab2/AbstractClothFactory"
	"github.com/EliriaT/tmps/lab2/ClothPrototypes"
	"log"
)

var ClothPrototypesList = ClothPrototypes.NewPrototypeManager()

func MakeWinterClothList() {
	//singleton
	winterFactory, err := AbstractClothFactory.GetClothesFactory("winter")
	if err != nil {
		log.Fatal(err)
	}
	Ugg := winterFactory.MakeShoe("Ugg boots")
	Ugg.SetHeight(2)
	SimpleBoots := winterFactory.MakeShoe("simple boots")
	SimpleBoots.SetHeight(3)

	vNeck := winterFactory.MakeShirt("v-neck sweater")
	vNeck.SetGender("F")
	turtleNeck := winterFactory.MakeShirt("turtle neck sweater")
	turtleNeck.SetGender("F")

	militaryPants := winterFactory.MakePants("military pants")
	militaryPants.SetStyle("military")
	isolatedPants := winterFactory.MakePants("waterproof pants")
	isolatedPants.SetStyle("clasic")

	ClothPrototypesList.Set(ClothPrototypesList.Length()+1, Ugg)
	ClothPrototypesList.Set(ClothPrototypesList.Length()+1, SimpleBoots)
	ClothPrototypesList.Set(ClothPrototypesList.Length()+1, vNeck)
	ClothPrototypesList.Set(ClothPrototypesList.Length()+1, turtleNeck)
	ClothPrototypesList.Set(ClothPrototypesList.Length()+1, militaryPants)
	ClothPrototypesList.Set(ClothPrototypesList.Length()+1, isolatedPants)
}

func MakeSummerClothList() {
	//singleton
	summerFactory, err := AbstractClothFactory.GetClothesFactory("summer")
	if err != nil {
		log.Fatal(err)
	}
	Sandals := summerFactory.MakeShoe("Sandals")
	Sandals.SetHeight(7)
	Loafers := summerFactory.MakeShoe("Loafers")
	Loafers.SetHeight(1)

	blouse := summerFactory.MakeShirt("blouse")
	blouse.SetGender("F")
	vNeckShirt := summerFactory.MakeShirt("v-neck shirt")
	vNeckShirt.SetGender("M")

	shorts := summerFactory.MakePants("jeans shorts")
	shorts.SetStyle("casual")
	shorts.SetGender("M")
	skirt := summerFactory.MakePants("skirt")
	skirt.SetStyle("elegant")
	skirt.SetGender("F")

	ClothPrototypesList.Set(ClothPrototypesList.Length()+1, Sandals)
	ClothPrototypesList.Set(ClothPrototypesList.Length()+1, Loafers)
	ClothPrototypesList.Set(ClothPrototypesList.Length()+1, blouse)
	ClothPrototypesList.Set(ClothPrototypesList.Length()+1, vNeckShirt)
	ClothPrototypesList.Set(ClothPrototypesList.Length()+1, shorts)
	ClothPrototypesList.Set(ClothPrototypesList.Length()+1, skirt)
}
