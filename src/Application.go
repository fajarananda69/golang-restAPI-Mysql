package main

import (
	"restAPI/config"
	"restAPI/controllers"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// connection
	db := config.DBInint()
	inDB := controllers.InDB{DB: db}

	router := gin.Default()

	//get data by id
	router.GET("/mysql/:id", inDB.GetDataById)
	//get all data
	router.GET("/mysql", inDB.GetAllData)
	//insert data
	router.POST("/mysql/:id/:name/:age/:grade", inDB.InsertDB)
	//update data
	router.PUT("/mysql/", inDB.UpdateDB)
	//delete data
	router.DELETE("/mysql/:id", inDB.DeleteDB)

	// set server port
	router.Run(":3000")
}
