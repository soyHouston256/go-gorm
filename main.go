package main

import (
	"github.com/soyhouston256/go-gorm/model"
	"github.com/soyhouston256/go-gorm/storage"
)

func main() {
	driver := storage.Mysql
	storage.New(driver)

	storage.DB().AutoMigrate(&model.Product{}, &model.InvoiceHeader{}, &model.InvoiceItem{})
}
