package storage

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

type Driver string

const (
	Mysql    Driver = "mysql"
	Postgres Driver = "postgres"
)

func New(d Driver) {
	switch d {
	case Mysql:
		newMysqlDB()
	case Postgres:
		newPostgresDB()
	default:
		newMysqlDB()
	}
}

func newPostgresDB() {
	once.Do(func() {
		var err error
		dsn := "host=localhost user=user password=admin dbname=golang port=5432 sslmode=disable TimeZone=Asia/Shanghai"

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatalf("cant open db connection: %v", err)
			panic(err)
		}
		fmt.Println("connected to db")
	})
}

func newMysqlDB() {
	once.Do(func() {
		var err error
		dsn := "root:root@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"

		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatalf("cant open db connection: %v", err)
			panic(err)
		}
		fmt.Println("connected to mysql db")
	})
}

func DB() *gorm.DB {
	return db
}
