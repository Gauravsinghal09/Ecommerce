package main

import (
	"fmt"
	"github.com/Gauravsinghal09/Ecommerce/Config"
	"github.com/Gauravsinghal09/Ecommerce/Models/Customer"
	"github.com/Gauravsinghal09/Ecommerce/Models/Order"
	"github.com/Gauravsinghal09/Ecommerce/Models/Product"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	var er error
	Config.DB, er = gorm.Open(mysql.Open(Config.DbURL(Config.BuildDBConfig())), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if er != nil {
		fmt.Println("Status", er)
	}

	Config.DB.AutoMigrate(&Customer.Customers{})
	Config.DB.AutoMigrate(&Product.Products{})
	Config.DB.AutoMigrate(&Order.Orders{})
	//routes := Routes.SetupRoutes()
	//routes.Run()
}
