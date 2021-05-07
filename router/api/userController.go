package api1

import (
	"fmt"
	"gin_demo/db/model"
	"github.com/gin-gonic/gin"
	"strconv"
)




func GetUser(c *gin.Context) {
	//在这调用db
	var id int
	id,_=strconv.Atoi(c.Param("id"))
	if id!=0 {

		fmt.Println(id)
		//user,_:= model.GetUser(id)
		//c.JSON(400, map[string]interface{}{"user":user})

		return
	}

		c.JSON(400, map[string]string{"user":"找不到"})
	return
}
// @Summary Import Image
// @Produce  json
// @Param image formData file true "Image File"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags/import [post]
func CreateUser(context *gin.Context)  {

	user := model.User{}
	context.BindJSON(&user)
	fmt.Println("user",user)

	if !user.IsEmpty(){
		data:=make(map[string]interface{})
		data["id"]=user.Id
		data["name"]=user.Name
		data["dept"]=user.Dept
		err:= model.AddUser(data)
		if err!=nil{
			context.JSON(400, map[string]string{"user":"出错了"})

			return
		}
		context.JSON(200, map[string]interface{}{"user":user})


	}
	context.JSON(400, map[string]string{"user":"出错了"})
}
func main()  {

	model.SetUpDB() //一定要连接数据库啊！不然就会报错了= =

	//r:=APIRouter()
	//r.Run()
}