package model

type Customer struct {
	CustomerId int    `json:"customer_id" gorm:"primaryKey"`
	Phone      string `json:"phone"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
}

func CreateCustomer(id int, phone string, firstName string, lastName string) Customer {
	return Customer{id, phone, firstName, lastName}
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
