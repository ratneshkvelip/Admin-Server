package controllers

import(
	"fmt"
	"server/services"
	"server/common/utils"
	"github.com/kataras/iris/v12"
	"encoding/json"
)


func Login(ctx iris.Context){
	var resp interface{}
	var statusCode int

	w,_:=ctx.ResponseWriter(), ctx.Request()

	w.Header().Set("Context-Type","application/json")

	err, userId,password:=utils.ReadLoginData(ctx)
	fmt.Println(err)
	statusCode,resp =services.Login(userId,password)
	

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)

}