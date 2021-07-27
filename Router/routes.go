package Router

import (
	"github.com/Gauravsinghal09/Ecommerce/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	product := r.Group("/retailer-api")
	{
		product.GET("product", Controllers.GetProducts)
		product.POST("product", Controllers.CreateProduct)
		product.GET("product/:id", Controllers.GetProductById)
		product.PUT("product/:id", Controllers.UpdateProduct)
		product.DELETE("product/:id", Controllers.DeleteProduct)
	}

	//order := r.Group("/customer-api")
	//{
	//	order.POST("order", Controllers.CreateOrder)
	//	order.GET("order/:id", Controllers.GetOrderByID)
	//	order.GET("customer/:id", Controllers.GetOrdersHistory)
	//}

	return r
}
