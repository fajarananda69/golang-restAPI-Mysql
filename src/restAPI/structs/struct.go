package structs

import "github.com/jinzhu/gorm"

type Person struct {
	gorm.Model
	ID    string
	Name  string
	Age   int
	Grade int
}
