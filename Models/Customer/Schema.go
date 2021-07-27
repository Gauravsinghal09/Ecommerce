package Customer

import "gorm.io/gorm"

type Customers struct {
	gorm.Model
	FirstName string
	LastName  string
}

func (customer *Customers) TableName() string {
	return "customers"
}
