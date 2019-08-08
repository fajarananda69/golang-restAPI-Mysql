package config

import (
	"restAPI/structs"

	"github.com/jinzhu/gorm"
)

//DBInit create connection to database
func DBInint() *gorm.DB {
	db, err := gorm.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/whatsapp")
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(structs.Person{})
	return db
}
