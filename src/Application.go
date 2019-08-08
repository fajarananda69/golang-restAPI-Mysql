package main

import (
	"restAPI/config"
	"restAPI/controllers"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInint()
	inDB := controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/person", inDB.GetPersons)
	router.Run(":3000")
}
