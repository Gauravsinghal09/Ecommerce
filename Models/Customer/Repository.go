package Customer

import (
	"github.com/Gauravsinghal09/Ecommerce/Config"
)

func CreateCustomer(customer *Customers) (err error) {
	if err = Config.DB.Create(customer).Error; err != nil {
		return err
	}
	return nil
}

func GetCustomerByID(customer *Customers, CustomerId string) (err error) {
	if err = Config.DB.Where("id = ?", CustomerId).First(customer).Error; err != nil {
		return err
	}
	return nil
}
