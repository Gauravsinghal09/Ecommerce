package main

import (
	"fmt"
	"github.com/Gauravsinghal09/Ecommerce/Config"
	"github.com/Gauravsinghal09/Ecommerce/Controllers"
	"github.com/Gauravsinghal09/Ecommerce/Models/Customer"
	"github.com/Gauravsinghal09/Ecommerce/Models/Order"
	"github.com/Gauravsinghal09/Ecommerce/Models/Product"
	"github.com/Gauravsinghal09/Ecommerce/Router"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitialiseDB() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(Config.DbURL(Config.BuildDBConfig())), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}

func AutoMigrateDB() {
	Config.DB.AutoMigrate(&Customer.Customers{})
	Config.DB.AutoMigrate(&Product.Products{})
	Config.DB.AutoMigrate(&Order.Orders{})
}

func main() {
	var er error
	Config.DB, er = InitialiseDB()

	if er != nil {
		fmt.Println("Status", er)
	}

	AutoMigrateDB()
	Controllers.CurrentProduct = Controllers.CurrProduct{}

	routes := Router.SetupRoutes()
	routes.Run()
}
