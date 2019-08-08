package structs

type Tb_student struct {
	// gorm.Model
	Id    string
	Name  string
	Age   int
	Grade int
}

type Response struct {
	Status  int
	Message string
	Data    []Tb_student
}
