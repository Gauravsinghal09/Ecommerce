package Controllers

import (
	"fmt"
	"github.com/Gauravsinghal09/Ecommerce/Models/Customer"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCustomer(c *gin.Context) {
	var customer Customer.Customers
	c.BindJSON(&customer)
	err := Customer.CreateCustomer(&customer)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"customer": customer, "message": "Successfully registered"})
	}
}

func GetCustomerById(c *gin.Context) {
	CustomerId := c.Params.ByName("CustomerId")
	var customer Customer.Customers
	err := Customer.GetCustomerByID(&customer, CustomerId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, customer)
	}
}
