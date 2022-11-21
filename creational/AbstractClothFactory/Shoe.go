package AbstractClothFactory

import (
	"TMPS/creational/ClothPrototypes"
)

type IShoe interface {
	SetSize(size int)
	SetHeight(height int)
	GetType() string
	GetSize() int
	GetHeight() int
	GetName() string
	Clone() ClothPrototypes.Cloneable
}

type Shoe struct {
	name   string
	stype  string
	size   int
	height int
}

func (s Shoe) GetName() string {
	return s.name
}

func (s Shoe) GetType() string {
	return s.stype
}

func (s Shoe) SetSize(size int) {
	s.size = size
}

func (s Shoe) GetSize() int {
	return s.size
}

func (s Shoe) SetHeight(height int) {
	s.height = height
}

func (s Shoe) GetHeight() int {
	return s.height
}

// Implements the Cloneable interface
func (s Shoe) Clone() ClothPrototypes.Cloneable {
	pe := Shoe{
		name:   s.name,
		stype:  s.stype,
		size:   s.size,
		height: s.height,
	}
	return pe
}
