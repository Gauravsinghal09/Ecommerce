package Order

import (
	"github.com/Gauravsinghal09/Ecommerce/Models/Customer"
	"github.com/Gauravsinghal09/Ecommerce/Models/Product"
	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	CustomerId int
	ProductId  int
	Customer   Customer.Customers
	Product    Product.Products
	Status     string
}

func (order *Orders) TableName() string {
	return "orders"
}
