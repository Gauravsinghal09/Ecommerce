package Controllers

import (
	"fmt"
	"github.com/Gauravsinghal09/Ecommerce/Models/Product"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProducts(c *gin.Context) {
	var product []Product.Products
	err := Product.GetAllProducts(&product)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		fmt.Println(product)
		c.JSON(http.StatusOK, product)
	}
}

func CreateProduct(c *gin.Context) {
	var product Product.Products
	c.BindJSON(&product)
	err := Product.CreateProduct(&product)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"product": product, "message": "Successfully registered"})
	}
}

func GetProductById(c *gin.Context) {
	productId := c.Params.ByName("ProductId")
	var product Product.Products
	err := Product.GetProductByID(&product, productId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func UpdateProduct(c *gin.Context) {
	var product Product.Products
	productId := c.Params.ByName("ProductId")
	err := Product.GetProductByID(&product, productId)
	if err != nil {
		c.JSON(http.StatusNotFound, product)
	}
	c.BindJSON(&product)
	err = Product.UpdateProduct(&product, productId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func DeleteProduct(c *gin.Context) {
	var product Product.Products
	productId := c.Params.ByName("ProductId")
	err := Product.DeleteProduct(&product, productId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + productId: "is deleted"})
	}
}
