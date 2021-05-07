package service

import (
	"fmt"
	"gin_demo/db/model"
)
type Good struct {
	GoodId int  `json:"good_id"`
	Name  string `json:"name"`
	Image string `json:"image"`
	Price float32 `json:"price"`
	Num int `json:"num"`
	DiscountId int `json:"discount_id"`
}
func IsExistGood(id int) bool {
	flag,_:= model.IsGoodExist(id)
	return flag
}
func  CreateGood(id int,form map[string]interface{}) error{
	flag,err:=model.IsGoodExist(id)
	if err!=nil {
		return err
	}
	fmt.Println(flag)
	if flag==false{
		fmt.Println(form)

		 _,err2:=model.AddGood(form)
		if err2!=nil {
			return err2
		}
	}
	return nil
}
func GetGood(id int)  *model.Good{
	//先从缓存取，取不到再去数据库
	return model.GetGoodById(id)
}
func GetGoodList() *[]model.Good {
	//先从缓存取，取不到再去数据库 拿全部的数据
	return model.GetGoodByList()
}
func UpdateGood(id int,data interface{})bool {
	flag,_:=model.IsGoodExist(id)
	if !flag {
		return model.UpdateGood(id,data)
	}
	return false

}
func DeleteGood(id int) bool {
	flag,_:=model.IsGoodExist(id)
	if flag {
		return model.DeleteGood(id)
	}
	return false
}

