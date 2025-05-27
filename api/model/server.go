package model

import (
	"log"

	"github.com/FastDefence/go_learn/api/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	dsn := "golearn:password@tcp(db:3306)/golearn_db?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(dsn + "database can't connect")
	}
	DB.AutoMigrate(model.User{})
}
