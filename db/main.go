package main

import (
	"fmt"
	"gin_demo/db/model"
	"io/ioutil"
	"net/http"
)

func sayHello(a http.ResponseWriter,r *http.Request)  {
	b,_:=ioutil.ReadFile("./hello.txt")
	_, _  =fmt.Fprintln(a,string(b))
}


func main() {
	model.SetUpDB()

	//good:=model.Good{1,"haha","s",2,1,1}



	//data:=make(map[string] interface{})
	//data["good_id"]=63
	//data["name"]="测试1"
	//data["image"]="1"
	//data["price"]=float32(1)
	//data["num"]=2
	//data["discount_id"]=1
	//service.CreateGood(63,data)
	data:=make(map[string] interface{})
	data["good_id"]=1
	data["user_id"]=1
	data["car_id"]=1
	data["num"]=3
	car:=model.GoodCar{GoodId: 1,UserId:54,CarId: 1,Num:122}
	err:=car.UpdateCarItemNum()
	if err!=nil {
		panic(err)
	}
	//model.AddGood(data)

//	删除model.DeleteGood(1)
//修改
//	data:=make(map[string]interface{})
//	data["name"]="haha"
//	data["num"]=22
//
//	model.UpdateGood(2,data)



}
