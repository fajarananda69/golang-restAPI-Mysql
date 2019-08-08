package main

import (
	"restAPI/config"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInint()
	inDB := controllers.inDB{DB: db}

	router := gin.Default
}
