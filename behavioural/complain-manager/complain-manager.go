package complain_manager

type ComplainManager interface {
	Execute(*Customer)
	SetNext(ComplainManager)
}

type Complaint int

const (
	ChangeOrder Complaint = iota
	WrongPreparedOrder
	HairInTheOrder
	RentTable
	BuyBussiness
)

//const (
//	Waiter Complaint = iota
//	Cook
//	Chef
//	Administrator
//	CEO
//)
