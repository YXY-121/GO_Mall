package api1

import (
	"fmt"
	"gin_demo/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type GoodForm struct {
	GoodId int  `json:"good_id"`
	Name  string `json:"name"`
	Image string `json:"image"`
	Price float32 `json:"price"`
	Num int `json:"num"`
	DiscountId int `json:"discount_id"`
}

func GetGood(c *gin.Context) {
	//在这调用db
	var id int
	id,_=strconv.Atoi(c.Param("id"))

	if	id!=0 &&service.IsExistGood(id){
		good:=service.GetGood(id)
		c.JSON(200,good)
		return
	}
	c.JSON(400, map[string]string{"user":"找不到"})
	return
}
func GetGoodList(c *gin.Context)  {

	c.JSON(200, service.GetGoodList())
	return
}

func AddGood(c *gin.Context)  {
	data:= GoodForm{}
	c.BindJSON(&data)
	fmt.Println(data)
	if !service.IsExistGood(data.GoodId){
		err:=service.CreateGood(data.GoodId,convert(data))
		if err!=nil {
			c.JSON(400, "错误了")
			return
		}
	}
	c.JSON(200, service.GetGoodList())
	return
	}


func convert(form GoodForm)map[string] interface{} {
	good:=map[string]interface{}{
		"good_id":    form.GoodId,
		"name":       form.Name,
		"image":      form.Image,
		"price":      form.Price,
		"num":       form.Num,
		"discount_id": form.DiscountId,
	}
	return good
}