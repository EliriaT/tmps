package HouseBuilder

type VillaBuilder struct {
	House
}

func newVillaBuilder() *VillaBuilder {
	return &VillaBuilder{}
}

func (b *VillaBuilder) setWindowType() {
	b.WindowType = "Wooden Window"
}

func (b *VillaBuilder) setDoorType() {
	b.DoorType = "Wooden Door"
}

func (b *VillaBuilder) setNumFloor() {
	b.Floor = 3
}

func (b *VillaBuilder) getHouse() House {
	return House{
		DoorType:   b.DoorType,
		WindowType: b.WindowType,
		Floor:      b.Floor,
	}
}
