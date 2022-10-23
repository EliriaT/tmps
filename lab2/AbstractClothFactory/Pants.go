package AbstractClothFactory

import (
	"github.com/EliriaT/tmps/lab2/ClothPrototypes"
)

type IPants interface {
	SetSize(size int)
	GetType() string
	GetSize() int
	SetGender(gender string)
	GetGender() string
	SetStyle(style string)
	GetStyle() string
	GetName() string
	Clone() ClothPrototypes.Cloneable
}

type Pants struct {
	name   string
	stype  string
	size   int
	gender string
	style  string
}

func (s Pants) GetName() string {
	return s.name
}

func (s Pants) GetType() string {
	return s.stype
}

func (s Pants) SetSize(size int) {
	s.size = size
}

func (s Pants) GetSize() int {
	return s.size
}

func (s Pants) SetGender(gender string) {
	s.gender = gender
}

func (s Pants) GetGender() string {
	return s.gender
}

func (s Pants) SetStyle(style string) {
	s.style = style
}

func (s Pants) GetStyle() string {
	return s.style
}

// Implements the Cloneable interface
func (s Pants) Clone() ClothPrototypes.Cloneable {
	pe := Pants{
		name:   s.name,
		stype:  s.stype,
		size:   s.size,
		gender: s.gender,
		style:  s.style,
	}
	return pe
}
