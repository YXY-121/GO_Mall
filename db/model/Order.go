package model

import _ "github.com/go-playground/locales/en_IL"

type Order struct {
	OrderId int `gorm:"order_id"`
	UserID int `gorm:"user_id"`
	IsPay int `gorm:"is_pay"` //0没有付款 1是付款
	PayWay int ` gorm:"pay_way"`
	Address string `gorm:"address"`
	TotalPrice float32 `gorm:"total_price"`
}

func GetOrderByOrderId(orderId dint) (*Order,error){
	var o Order
	//根据某个订单号去查看详细信息
	err:=db.Model(Order{}).Where("order_id=?",o.OrderId).Find(o).Error
	if err!=nil {
		return &o,err
	}
	return nil,err
}
func   GetOrderByUserId(userId int) (*[]Order,error){
	//根据某个订单号去查看详细信息
	var o *[]Order
	err:=db.Model(Order{}).Where("order_id=?",userId).Find(o).Error
	if err!=nil {
		return nil,err
	}
	return o,nil
}
func (o Order) CreateOrder() error{
	err:=db.Model(Order{}).Create(o).Error
	if err!=nil {
		return err
	}
	return nil
}
func (o Order) UpdateAddress() error{
	//修改订单地址
	err:=db.Model(Order{}).Where("order_id=?",o.OrderId).Update(o).Error
	if err!=nil {
		return err
	}
	return nil
}
