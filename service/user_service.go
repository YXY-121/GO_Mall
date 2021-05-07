package service

import (
	"gin_demo/db/model"
)

var user model.User

func  name(id int)  {
		model.GetUser(id)
}
func main()  {
	name(2)
}