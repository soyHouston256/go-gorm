package main

import (
	"github.com/soyhouston256/go-gorm/model"
	"github.com/soyhouston256/go-gorm/storage"
)

func main() {
	driver := storage.Mysql
	storage.New(driver)

	product := model.Product{}
	product.ID = 1

	invoice := model.InvoiceHeader{
		Client: "Max Houston Ramirez",
		InvoiceItems: []model.InvoiceItem{
			{
				ProductID: 2,
				Quantity:  2,
				Price:     100,
			},
			{
				ProductID: 3,
				Quantity:  2,
				Price:     100,
			},
		},
	}
	storage.DB().Create(&invoice)
}
