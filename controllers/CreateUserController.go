package controllers

import(

	"server/services"	
	"server/common/utils"	
	"github.com/kataras/iris/v12"
	"encoding/json"
	"fmt"
)



func CreateUserAccount(ctx iris.Context)  {
	var resp interface{} 
	var body utils.Data 
	var statusCode int
	w,_:=ctx.ResponseWriter(),ctx.Request()
	w.Header().Set("Allow-Origin","*")
	w.Header().Set("Content-Type","application/json")

	err,body:= utils.ReadCreateAccountBody(ctx)

	fmt.Println(err)

	statusCode,resp=services.CreateUserAccount(ctx,body)

	fmt.Println(resp)

	w.WriteHeader(statusCode)

	
	
	json.NewEncoder(w).Encode(resp)
	

}
