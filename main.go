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

	storage.DB().First(&product, 1)

	fmt.Println(product)

}
