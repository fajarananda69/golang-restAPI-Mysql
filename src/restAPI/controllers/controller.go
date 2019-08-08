package controllers

import (
	"net/http"
	"restAPI/structs"

	"github.com/gin-gonic/gin"
)

// // to get one data with {id}
// func (idb *InDB) GetPerson(c *gin.Context) {
// 	var (
// 		person structs.GroupList
// 		result gin.H
// 	)
// 	id := c.Param("id")
// 	err := idb.DB.Where("id = ?", id).First(&person).Error
// 	if err != nil {
// 		result = gin.H{
// 			"result": err.Error(),
// 			"count":  0,
// 		}
// 	} else {
// 		result = gin.H{
// 			"result": person,
// 			"count":  1,
// 		}
// 	}
// 	c.JSON(http.StatusOK, result)
// }

// get all data in person
func (idb *InDB) GetPersons(c *gin.Context) {
	var (
		table []structs.Tb_student
		// jojo   []GroupList
		result gin.H
	)

	// idb.DB.Find(&table)
	idb.DB.Table("tb_student").Scan(&table)
	if len(table) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"status":  200,
			"message": "success",
			"data":    table,
			"count":   len(table),
		}
	}

	c.JSON(http.StatusOK, result)
}
