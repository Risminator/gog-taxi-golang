// PROBLEMS HERE

package model

// Phone string???
type Customer struct {
	CustomerId int    `json:"customer_id"`
	Phone      string `json:"phone"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
}

// Phone string?
func CreateCustomer(id int, phone string, firstName string, lastName string) Customer {
	return Customer{id, phone, firstName, lastName}
}

// Phone string?
func (c *Customer) SetPhone(p string) {
	c.Phone = p
}

func (c *Customer) SetFirstName(f string) {
	c.FirstName = f
}

func (c *Customer) SetLastName(l string) {
	c.LastName = l
}
