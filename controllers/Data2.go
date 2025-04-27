package controllers

import(
	"server/services"

	"github.com/kataras/iris/v12"
	"encoding/json"
)


func Data2(ctx iris.Context){
	var resp interface{}
	var statusCode int

	w,_:=ctx.ResponseWriter(), ctx.Request()

	w.Header().Set("Context-Type","application/json")

	// err, userId,password:=utils.ReadLoginData(ctx)
	// fmt.Println(err)
	statusCode,resp =services.Data2()

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)

}