# golang-restAPI-Mysql
Build simpel restAPI with golang programming language to access mysql. in this project we use `gin` driver to build controller restAPI 

## Install and SETUP GO
Download and configure your workspace with latest version of Go and correct environment path.
- [Last Version](https://golang.org/dl/)
- [Windows](http://www.wadewegner.com/2014/12/easy-go-programming-setup-for-windows/)
- [Linux](http://www.tecmint.com/install-go-in-linux/)


## Install driver
On your Editor open terminal to access your directory project and install driver
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
## Run API
``` 
go run Application.go
```
