package complain_manager

import "fmt"

type Chef struct {
	next ComplainManager
}

func (r *Chef) Execute(p *Customer) {
	if p.Problem == HairInTheOrder {
		fmt.Println("We are extrimely sorry, please take this wine as an excuse.")
		return
	}
	r.next.Execute(p)
}

func (r *Chef) SetNext(next ComplainManager) {
	r.next = next
}
