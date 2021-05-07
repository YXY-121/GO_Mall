package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

 var db *gorm.DB
func SetUpDB()  *gorm.DB{
	var err error
	dsn := "root:123@tcp(127.0.0.1:3306)/mall?charset=utf8mb4&parseTime=True&loc=Local"
	db,err =gorm.Open("mysql",dsn)
	if err!=nil{
		panic(err)
	}
	fmt.Println("初始化完成")
	db.SingularTable(true)
	db.AutoMigrate(User{},Good{})
	return db
}

func main() {



//	flag1:=AddUser(data)
//	fmt.Println(flag1)
//	flag2,_:=DeleteUser(1)
//	fmt.Println(flag2)



}
