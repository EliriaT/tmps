package AbstractClothFactory

type WinterShirt struct {
	WoolType string
	Shirt
}

func (s WinterShirt) setWoolType(woolType string) {
	s.WoolType = woolType
}
func (s WinterShirt) GetWoolType() string {
	return s.WoolType
}

type WinterShoe struct {
	Fur string
	Shoe
}

func (s WinterShoe) setFur(fur string) {
	s.Fur = fur
}
func (s WinterShoe) GetFur() string {
	return s.Fur
}

type WinterPants struct {
	Insulated bool
	Pants
}

func (s WinterPants) setInsulated(insulated bool) {
	s.Insulated = insulated
}
func (s WinterPants) GetInsulated() bool {
	return s.Insulated
}
