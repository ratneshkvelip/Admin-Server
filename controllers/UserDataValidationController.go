package controllers

import(

	"server/services"	
	"github.com/kataras/iris/v12"
	"encoding/json"
	"fmt"
)


func UserDataValidation(ctx iris.Context){
	var resp interface{} 

	w,_:=ctx.ResponseWriter(),ctx.Request()
	w.Header().Set("Allow-Origin","*")
	w.Header().Set("Content-Type","application/json")

	 UserName:=ctx.FormValue("user_name")
	cb:=ctx.FormValue("password")
	statusCode,resp:=services.UserDataValidation(UserName)

	fmt.Println(resp)
	
	fmt.Println(cb)

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
	

}