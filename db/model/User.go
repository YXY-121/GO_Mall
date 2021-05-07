package model

import (
	"reflect"
)

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Dept  int `json:"dept"`
}


func (u User)IsEmpty()  bool{
	return reflect.DeepEqual(u, User{})
	
}

func  ExistUserById(id int)(bool,error)  {
	var user User
	err:=db.Where("id=? ",id).First(&user).Error

	if err!=nil{
		return false,err
	}
	return true,err
}


func  GetUser(id int)(*User,error)  {
	var user User

	err:=db.Where("id =?", id).Find(&user).Error

	if err!=nil{
		return &user,err
	}

	return &user,err
}
//
//func GetUserList() ([]*User){
//	var users []*User
//	users,_=db.Find(&users).Error
//	return users
//}
func  DeleteUser(id int)(bool,error)  {
	err:= db.Delete(&User{},id).Error
	if err!=nil{
		return false,err
	}
	return true,err
}


func AddUser(data map[string]interface{})error  {
	user:= User{
		Id: data["id"].(int),
		Name:data["name"].(string),
		Dept: data["dept"].(int),
	}
	err:= db.Create(&user).Error
	if err!=nil{
		return err

	}
	return nil

}

func update(id int,data interface{}) (bool,error) {

		err:=db.Model(&User{}).Where("id=? ",id).Updates(data).Error
		if err!=nil{
			return false,err
		}
	return true, nil
}

