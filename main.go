package main

import (
	"github.com/soyhouston256/go-gorm/model"
	"github.com/soyhouston256/go-gorm/storage"
)

func main() {
	driver := storage.Mysql
	storage.New(driver)
	m := model.Product{
		Name:  "Php Course",
		Price: 100,
	}
	m2 := model.Product{
		Name:  "Python Course",
		Price: 100,
	}
	obs := "This is a golang course"
	m3 := model.Product{
		Name:        "Python Course",
		Price:       100,
		Observation: &obs,
	}

	storage.DB().Create(&m)
	storage.DB().Create(&m2)
	storage.DB().Create(&m3)

}
