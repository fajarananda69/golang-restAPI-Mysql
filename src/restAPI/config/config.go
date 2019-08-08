package config

import (
	"restAPI/structs"

	"github.com/jinzhu/gorm"
)

//DBInit create connection to database
func DBInint() *gorm.DB {
	db, err := gorm.Open("mysql", "admin:admin@tcp(127.0.0.1)/whatsapp")
	if err != nil {
		panic("failed to connect to database")
	}

	// auto create table where table doesn't exist
	db.AutoMigrate(structs.Tb_student{})
	return db
}
