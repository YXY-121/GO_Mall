package main

import (
	"fmt"
	"gin_demo/db/model"
	"gin_demo/router"
	"io/ioutil"
	"net/http"
)

func sayHello(a http.ResponseWriter,r *http.Request)  {
	b,_:=ioutil.ReadFile("./hello.txt")
	_, _  =fmt.Fprintln(a,string(b))
}


func main() {
	model.SetUpDB()
	r:=router.APIRouter()
	r.Run()
	//good:=model.Good{1,"haha","s",2,1,1}
	//
	//good1:=model.GetGoodById(good.GoodId)
	//fmt.Println(good1)
	//user1,_:=model.GetUser(3)
	//fmt.Println(user1)
	//data:=make(map[string] interface{})
	//data["id"]=1
	//data["name"]="测试1"
	//data["dept"]=1

}
