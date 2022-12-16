package complain_manager

import "fmt"

type Waiter struct {
	next ComplainManager
}

func (r *Waiter) Execute(p *Customer) {
	if p.Problem == ChangeOrder {
		fmt.Println("Okay, we will bring the right order in 15-30 minutes")
		return
	}
	r.next.Execute(p)
}

func (r *Waiter) SetNext(next ComplainManager) {
	r.next = next
}
