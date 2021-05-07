package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
)

type Good struct {
	GoodId int  `gorm:"good_id"`
	Name  string `gorm:"name"`
	Image string `gorm:"image"`
	Price float32 `gorm:"price"`
	Num int `gorm:"num"`
	DiscountId int `gorm:"discount_id"`
}

func IsEmpty(good *Good)  bool{
	return reflect.DeepEqual(good, Good{})

}
func IsGoodExist(id int) (bool,error){
	var good Good
	err:=db.Where("good_id=?",id).Find(&good).Error
	if err!=nil{
		if err==gorm.ErrRecordNotFound {
			return false,nil
		}
	}

	return true,err
}

//查看单个商品的所有信息
func  GetGoodById(id int)*Good  {

	var good Good
	fmt.Println(id)
	err:=db.Where("good_id=?",id).Find(&good).Error
	if err!=nil{
		return nil
	}
	return &good


}
//主页面上，查看所有商品的图片、价格、名字
func  GetGoodByList() *[]Good{
	goods:=new([]Good)
	err:=db.Find(&goods).Error
	if err!=nil{
		return nil
	}
	return goods
}
func  AddGood(data map[string]interface{}) (bool,error) {
	good1:=Good{
			GoodId: data["good_id"].(int),
			Name:data["name"].(string),
			Image:data["image"].(string),
			Price :data["price"].(float32),
			Num:data["num"].(int),
			DiscountId:data["discount_id"].(int),
	}
	fmt.Println("good",good1)
		err:=db.Create(&good1).Error
		if err!=nil{
			return false,err
		}
		return true,nil

}
func UpdateGood(id int,data interface{})bool  {
	err:=db.Model(&Good{}).Where("good_id=?",id).Updates(data).Error
	if err!=nil{
		return false
	}
	return true
}
func  DeleteGood(id int) bool {
	err:=db.Where("good_id=?",id).Delete(Good{}).Error
	if err!=nil {
		return false
	}
	return  true
}

