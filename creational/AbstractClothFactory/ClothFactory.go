package AbstractClothFactory

import (
	"fmt"
	"sync"
)

type IClothesFactory interface {
	//the abstract factory creator method
	MakeShoe(name string) IShoe
	MakeShirt(name string) IShirt
	MakePants(name string) IPants
}

var winterFactory *WinterClothes
var summerFactory *SummerClothes

var winterFactoryLock = &sync.Mutex{}
var summerFactoryLock = &sync.Mutex{}

func GetClothesFactory(season string) (IClothesFactory, error) {
	if season == "winter" {
		if winterFactory == nil {
			winterFactoryLock.Lock()
			defer winterFactoryLock.Unlock()
			if winterFactory == nil {
				fmt.Println("Creating single instance now.")
				winterFactory = &WinterClothes{}
			} else {
				fmt.Println("Single instance already created.")
			}
		} else {
			fmt.Println("Single instance already created.")
		}
		return winterFactory, nil
	}

	if season == "summer" {
		if summerFactory == nil {
			summerFactoryLock.Lock()
			defer summerFactoryLock.Unlock()
			if summerFactory == nil {
				fmt.Println("Creating single instance now.")
				summerFactory = &SummerClothes{}
			} else {
				fmt.Println("Single instance already created.")
			}
		} else {
			fmt.Println("Single instance already created.")
		}
		return summerFactory, nil
	}

	return nil, fmt.Errorf("Wrong season type passed")
}
