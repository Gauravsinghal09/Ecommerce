package Product

import (
	"fmt"
	"github.com/Gauravsinghal09/Ecommerce/Config"
	"gorm.io/gorm/clause"
)

func GetAllProducts(product *[]Products) (err error) {

	if err = Config.DB.Preload(clause.Associations).Find(product).Error; err != nil {
		return err
	}
	fmt.Println(product)
	return nil
}

func CreateProduct(product *Products) (err error) {
	if err = Config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func GetProductByID(product *Products, id string) (err error) {
	if err = Config.DB.Preload(clause.Associations).Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProduct(product *Products, id string) (err error) {
	Config.DB.Save(product)
	return nil
}

func DeleteProduct(product *Products, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(product)
	return nil
}
