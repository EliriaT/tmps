package HouseBuilder

import "fmt"

type House struct {
	WindowType string
	DoorType   string
	Floor      int
}

func (h *House) PrintHouse() {
	fmt.Printf("\n House Door Type: %s\n", h.DoorType)
	fmt.Printf(" House Window Type: %s\n", h.WindowType)
	fmt.Printf(" House Num Floor: %d\n", h.Floor)
}
