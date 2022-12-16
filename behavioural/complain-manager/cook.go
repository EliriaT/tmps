package complain_manager

import "fmt"

type Cook struct {
	next ComplainManager
}

func (r *Cook) Execute(p *Customer) {
	if p.Problem == WrongPreparedOrder {
		fmt.Println("Okay, sorry, i will prepared the order according to the recipy")
		return
	}
	r.next.Execute(p)
}

func (r *Cook) SetNext(next ComplainManager) {
	r.next = next
}
