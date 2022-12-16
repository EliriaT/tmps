package complain_manager

import "fmt"

type CEO struct {
	next ComplainManager
}

func (r *CEO) Execute(p *Customer) {
	if p.Problem == BuyBussiness {
		fmt.Println("Oh, you want to buy this business? Let's discuss this at 15:00 at a meal.")
		return
	}
	r.next.Execute(p)
}

func (r *CEO) SetNext(next ComplainManager) {
	r.next = next
}
