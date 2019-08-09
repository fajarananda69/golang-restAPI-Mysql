package controllers

import (
	"net/http"
	"restAPI/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

// get all data in person
func (idb *InDB) GetAllData(c *gin.Context) {
	var (
		table  []structs.Tb_student
		result gin.H
	)

	// idb.DB.Find(&table) //get all with struct (table's name same with struct's name)
	idb.DB.Table("tb_students").Scan(&table) // get all from table's name, scan "&" => struct

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

// to get one data with {id}
func (idb *InDB) GetDataById(c *gin.Context) {
	var (
		table  structs.Tb_student
		result gin.H
	)
	// with param(/mysql/1)
	id := c.Param("id")

	// // with form-data
	// id := c.PostForm("id")

	// // with query(/mysql/?id=1)
	// id := c.Query("id")

	err := idb.DB.Where("id = ?", id).First(&table).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": table,
			"count":  1,
		}
	}
	c.JSON(http.StatusOK, result)
}

// insert database
func (idb *InDB) InsertDB(c *gin.Context) {
	var (
		table  structs.Tb_student
		result gin.H
	)

	id := c.Param("id")
	name := c.Param("name")
	age, _ := strconv.Atoi(c.Param("age"))
	grade, _ := strconv.Atoi(c.Param("grade"))

	table.Id = id
	table.Name = name
	table.Age = age
	table.Grade = grade

	idb.DB.Create(&table)
	result = gin.H{
		"result": table,
	}
	c.JSON(http.StatusOK, result)
}

// update database by {id}
func (idb *InDB) UpdateDB(c *gin.Context) {
	id := c.Query("id")
	name := c.PostForm("name")
	age, _ := strconv.Atoi(c.PostForm("age"))
	grade, _ := strconv.Atoi(c.PostForm("grade"))

	var (
		table    structs.Tb_student
		newtable structs.Tb_student
		result   gin.H
	)

	table.Id = id
	newtable.Id = id
	newtable.Name = name
	newtable.Age = age
	newtable.Grade = grade

	err := idb.DB.Where("id = ?", id).First(&table).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	} else {
		err = idb.DB.Model(&table).Where("id = ?", id).Updates(newtable).Error
		if err != nil {
			result = gin.H{
				"result": "update failed",
			}
		} else {
			result = gin.H{
				"status":  200,
				"message": "success",
				"result":  newtable,
			}
		}
	}

	c.JSON(http.StatusOK, result)
}

// delete data with {id}
func (idb *InDB) DeleteDB(c *gin.Context) {
	var (
		table  structs.Tb_student
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&table).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	} else {
		err = idb.DB.Where("id = ?", id).Delete(&table).Error
		if err != nil {
			result = gin.H{
				"result": "delete failed",
			}
		} else {
			result = gin.H{
				"result": "Data deleted successfully",
			}
		}
	}

	c.JSON(http.StatusOK, result)
}
