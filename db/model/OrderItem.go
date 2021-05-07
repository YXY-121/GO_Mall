package model
type OrderItem struct {
	OrderId int `gorm:"order_id"`
	OrderItemId int `gorm:"order_item_id"`
	GoodId int   `gorm:"good_id"` //一个item对一种商品以及其价格 `
	Num int `gorm:"num"`
	Price float32 `gorm:"price"`

}