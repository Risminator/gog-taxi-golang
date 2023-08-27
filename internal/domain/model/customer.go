package model

type Customer struct {
	CustomerId int    `json:"customerId" gorm:"primaryKey"`
	Phone      string `json:"phone"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
}

func CreateCustomer(id int, phone string, firstName string, lastName string) Customer {
	return Customer{id, phone, firstName, lastName}
}

func (c *Customer) SetCustomerId(id int) {
	c.CustomerId = id
}

func (c *Customer) SetPhone(p string) {
	c.Phone = p
}

func (c *Customer) SetFirstName(f string) {
	c.FirstName = f
}

func (c *Customer) SetLastName(l string) {
	c.LastName = l
}
