package Controllers

import (
	"github.com/Gauravsinghal09/Ecommerce/Models/Order"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOrders(c *gin.Context) {
	var order []Order.Orders
	err := Order.GetOrders(&order)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

func CreateOrder(c *gin.Context) {
	var order Order.Orders
	c.BindJSON(&order)
	err := Order.CreateOrder(&order)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

func GetOrderById(c *gin.Context) {
	OrderId := c.Params.ByName("OrderId")
	var order Order.Orders
	err := Order.GetOrderById(&order, OrderId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}
