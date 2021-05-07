package router

import (
	api1 "gin_demo/router/api"
	"github.com/gin-gonic/gin"
)

func APIRouter() *gin.Engine {
	r:=gin.Default()
	api:=r.Group("/api/v1")
	api.GET("/get/:id",api1.CreateUser)
	r.GET("/getUser/:id",api1.GetUser)

	r.POST("/createGood",api1.AddGood)
	r.GET("/getGood/:id",api1.GetGood)
	r.GET("/getGoods/",api1.GetGoodList)




	return r
}
