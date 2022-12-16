package complain_manager

import "fmt"

type Administrator struct {
	next ComplainManager
}

func (r *Administrator) Execute(p *Customer) {
	if p.Problem == RentTable {
		fmt.Println("Yes , we have available tables at this hour. You can rent.")
		return
	}
	r.next.Execute(p)
}

func (r *Administrator) SetNext(next ComplainManager) {
	r.next = next
}
