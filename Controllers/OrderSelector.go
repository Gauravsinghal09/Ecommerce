package Controllers

import (
	"fmt"
	"github.com/Gauravsinghal09/Ecommerce/Models/Order"
	"sync"
	"time"
)

type CustomerMap struct {
	Map sync.Map
	//Queue []Order.Orders
}

var CMap CustomerMap

func OrderQueue(order *Order.Orders) {

	time.Sleep(time.Second * 15)
	val, ok := CMap.Map.Load(order.CustomerId)
	if !ok {
		EligibleOrder(order)
		CMap.Map.Store(order.CustomerId, order.UpdatedAt)
	} else {
		lastUpdatedTime, _ := val.(time.Time)
		if time.Time.Sub(time.Now(), lastUpdatedTime) > 2*time.Minute {
			EligibleOrder(order)
			CMap.Map.Store(order.CustomerId, order.UpdatedAt)
		} else {
			fmt.Println("Sleeping for a minute")
			time.Sleep(time.Minute)
			OrderQueue(order)
		}
	}
}

func EligibleOrder(order *Order.Orders) error {
	time.Sleep(time.Second * 15)
	if err := Order.UpdateOrder(order); err != nil {
		return err
	}
	return nil
}

/*


 */
