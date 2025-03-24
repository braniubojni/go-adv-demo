package db

import (
	"fmt"
	"go/adv-demo/configs"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(conf *configs.Config) *Db {
	log.Println("Tring to connect to Postgres")
	db, err := gorm.Open(postgres.Open(conf.Db.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Postgres connected")
	return &Db{db}
}
