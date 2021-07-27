package Router

import (
	"github.com/Gauravsinghal09/Ecommerce/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	product := r.Group("/product-api")
	{
		product.GET("product", Controllers.GetProducts)
		product.POST("product", Controllers.CreateProduct)
		product.GET("product/:ProductId", Controllers.GetProductById)
		product.PUT("product/:ProductId", Controllers.UpdateProduct)
		product.DELETE("product/:ProductId", Controllers.DeleteProduct)
	}

	order := r.Group("/order-api")
	{
		order.POST("order", Controllers.CreateOrder)
		order.GET("order/:OrderId", Controllers.GetOrderById)
		order.GET("order", Controllers.GetOrders)
	}

	customer := r.Group("/customer-api")
	{
		customer.POST("customer", Controllers.CreateCustomer)
		customer.GET("customer/:CustomerId", Controllers.GetCustomerById)
		customer.GET("order/:CustomerId", Controllers.GetOrdersHistory)
	}

	return r
}
