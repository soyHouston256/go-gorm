package main

import (
	"fmt"

	"github.com/soyhouston256/go-gorm/model"
	"github.com/soyhouston256/go-gorm/storage"
)

func main() {
	driver := storage.Mysql
	storage.New(driver)

	product := model.Product{}
	product.ID = 1

	storage.DB().Delete(&product)

	fmt.Println(product)

}
