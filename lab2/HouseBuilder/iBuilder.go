package HouseBuilder

import "fmt"

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

func GetBuilder(builderType string) (IBuilder, error) {
	if builderType == "flat" {
		return newFlatBuilder(), nil
	}

	if builderType == "villa" {
		return newVillaBuilder(), nil
	}
	return nil, fmt.Errorf("Wrong house type passed")
}
