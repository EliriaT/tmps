package ClothPrototypes

import "fmt"

type Cloneable interface {
	Clone() Cloneable
	GetName() string
}

type PrototypeManager struct {
	Prototypes map[int]Cloneable
}

func (p *PrototypeManager) Get(name int) (Cloneable, error) {
	if _, ok := p.Prototypes[name]; ok == false {
		return nil, fmt.Errorf("Wrong  type passed")
	}
	return p.Prototypes[name].Clone(), nil
}

func (p *PrototypeManager) Set(name int, prototype Cloneable) {
	p.Prototypes[name] = prototype
}
func (p *PrototypeManager) Length() int {
	return len(p.Prototypes)
}
func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{
		Prototypes: make(map[int]Cloneable),
	}
}
