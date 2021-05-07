package model

import "github.com/jinzhu/gorm"

type  GoodCar struct {
	CarId int `gorm:"car_id"` //carID好像可有可无
	UserId int `gorm:"user_id"` //主要靠这个来辨别
	GoodId int `gorm:"good_id"` //每一项的
	Num int `gorm:"num"`

}

func GetCar(userid int)  (*[]GoodCar,error){
	//找到所有的
	var car =[]GoodCar{}
	err:=db.Model(GoodCar{}).Where("user_id=?",userid).Find(&car).Error
	if err!=nil {
		return nil,err
	}
	return &car,nil
}
func (g GoodCar) AddItemToCar()error  {
	//car:=GoodCar{
	//	CarId:data["car_id"].(int),
	//	UserId :data["user_id"].(int), //主要靠这个来辨别
	//	GoodId :data["good_id"].(int), //每一项的
	//	Num : data["num"].(int),
	//
	//}
	err:=db.Model(GoodCar{}).Create(&g).Error
	if err!=nil {
		return err
	}
	return nil

}
func (g GoodCar) UpdateCarItemNum() error {
		err:=db.Model(GoodCar{}).Where("good_id=? and user_id=?",g.GoodId,g.UserId).Updates(g).Error
	if err!=nil {
		return err
	}
	return nil

}
func (g GoodCar) DeleteCarItem( )  error{
		err:=db.Model(GoodCar{}).Where("good_id=? and user_id=?",g.GoodId,g.UserId).Delete(GoodCar{}).Error
		if err!=nil{
			return err
		}
		return nil
}
func (g GoodCar)IsExistCarItem() bool {
	err:=db.Model(GoodCar{}).Where("good_id=? and user_id=?",g.GoodId,g.UserId).Error
	if err != nil {
		if err==gorm.ErrRecordNotFound {
			return false
		}
	}
	return true


}