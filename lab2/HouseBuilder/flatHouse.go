package HouseBuilder

type FlatBuilder struct {
	House
}

func newFlatBuilder() *FlatBuilder {
	return &FlatBuilder{}
}

func (b *FlatBuilder) setWindowType() {
	b.WindowType = "High Windows"
}

func (b *FlatBuilder) setDoorType() {
	b.DoorType = "Metal Door"
}

func (b *FlatBuilder) setNumFloor() {
	b.Floor = 1
}

func (b *FlatBuilder) getHouse() House {
	return House{
		DoorType:   b.DoorType,
		WindowType: b.WindowType,
		Floor:      b.Floor,
	}
}
