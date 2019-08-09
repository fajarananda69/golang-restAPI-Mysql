# golang-restAPI-Mysql
Build simpel restAPI with golang programming language to access mysql

## install driver
```
go get -u "github.com/gin-gonic/gin"
```
```
go get -u "github.com/jinzhu/gorm"
```
```
go get -u "github.com/jinzhu/gorm/dialects/mysql"
```
```
go get -u "github.com/go-sql-driver/mysql"
```


## connection
```
func DBInint() *gorm.DB {
	db, err := gorm.Open("mysql", "root:pass@tcp(127.0.0.1)/DBname")
	if err != nil {
		panic("failed to connect to database")
	}

	// auto create table where table doesn't exist
	db.AutoMigrate(structs.TableName{})
	return db
}
```
