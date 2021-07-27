package Controllers

import (
	"fmt"
	"github.com/Gauravsinghal09/Ecommerce/Models/Order"
	"time"
)

type CustomerMap struct {
	Map map[int]time.Time
	//Queue []Order.Orders
}

var CMap CustomerMap

func OrderQueue(order *Order.Orders) {

	val, ok := CMap.Map[order.CustomerId]

	if !ok {
		EligibleOrder(order)
		CMap.Map[order.CustomerId] = order.UpdatedAt
	} else {
		if time.Time.Sub(time.Now(), val) > 2*time.Minute {
			EligibleOrder(order)
			CMap.Map[order.CustomerId] = order.UpdatedAt
		} else {
			fmt.Println("Sleeping for a minute")
			time.Sleep(time.Minute)
			OrderQueue(order)
		}
	}
}

func EligibleOrder(order *Order.Orders) error {
	if err := Order.UpdateOrder(order); err != nil {
		return err
	}
	return nil
}

/*


 */
