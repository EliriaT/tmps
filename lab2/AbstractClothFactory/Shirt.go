package AbstractClothFactory

import (
	"github.com/EliriaT/tmps/lab2/ClothPrototypes"
)

type IShirt interface {
	SetSize(size int)
	GetType() string
	GetSize() int
	SetGender(gender string)
	GetGender() string
	SetGenderless(genderless bool)
	GetGenderless() bool
	GetName() string
	Clone() ClothPrototypes.Cloneable
}

type Shirt struct {
	name       string
	stype      string
	size       int
	gender     string
	genderless bool
}

func (s Shirt) GetName() string {
	return s.name
}

func (s Shirt) GetType() string {
	return s.stype
}

func (s Shirt) SetSize(size int) {
	s.size = size
}

func (s Shirt) GetSize() int {
	return s.size
}

func (s Shirt) SetGender(gender string) {
	s.gender = gender
}

func (s Shirt) GetGender() string {
	return s.gender
}

func (s Shirt) SetGenderless(genderless bool) {
	s.genderless = genderless
}

func (s Shirt) GetGenderless() bool {
	return s.genderless
}

// Implements the Cloneable interface
func (s Shirt) Clone() ClothPrototypes.Cloneable {
	pe := Shirt{
		name:       s.name,
		stype:      s.stype,
		size:       s.size,
		gender:     s.gender,
		genderless: s.genderless,
	}
	return pe
}
