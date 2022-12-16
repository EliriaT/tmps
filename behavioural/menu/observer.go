package menu

import "fmt"

// it is the client
type Observer interface {
	update([]string)
	getID() string
}

type Customer struct {
	Email string
}

func (c *Customer) update(menu []string) {
	fmt.Printf("Sending email to customer %s with new menu %s\n", c.Email, menu)
}

func (c *Customer) getID() string {
	return c.Email
}
