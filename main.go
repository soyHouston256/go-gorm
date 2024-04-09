package main

import (
	"fmt"

	"github.com/soyhouston256/go-gorm/model"
	"github.com/soyhouston256/go-gorm/storage"
)

func main() {
	driver := storage.Mysql
	storage.New(driver)

	product := make([]model.Product, 0)

	storage.DB().Find(&product)

	for _, p := range product {
		fmt.Printf("ID: %d - %s\n", p.ID, p.Name)
	}
}
