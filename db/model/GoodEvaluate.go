package model
type GoodEvaluate struct {
	EvaluateId int `gorm:"evaluate_id"`
	OrderId int `gorm:"order_id"`
	UserId int `gorm:"user_id"`
	GoodId int `gorm:"good_id"`
	Level int `gorm:"level"`
	Image string `gorm:"image"`
	Word string `gorm:"word"`
}