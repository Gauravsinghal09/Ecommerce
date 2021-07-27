package Order

import (
	"github.com/Gauravsinghal09/Ecommerce/Config"
	"gorm.io/gorm"
)

func GetOrders(order *[]Orders) (err error) {
	if err = Config.DB.Find(order).Error; err != nil {
		return err
	}
	return nil
}

func CreateOrder(order *Orders) (err error) {
	err = Config.DB.Transaction(func(tx *gorm.DB) error {
		if err = Config.DB.Create(order).Error; err != nil {
			return err
		}
		if err = Config.DB.Model(order).Update("Status", "Order placed").Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func GetOrderById(order *Orders, OrderId string) (err error) {
	if err = Config.DB.Where("id = ?", OrderId).First(order).Error; err != nil {
		return err
	}
	return nil
}

func GetOrdersHistory(order *[]Orders, CustomerId string) (err error) {
	if err = Config.DB.Where("customer_id = ?", CustomerId).Find(order).Error; err != nil {
		return err
	}
	return nil
}
