package utils


import(
	"fmt"
	"github.com/kataras/iris/v12"
)



type Data struct{
	UserID string `json:"userId"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	EmailId string `json:"emailId"`
	Country string `json:"country"`
	MobileNumber string `json:"mobileNumber"`
	Age int `json:"age"`
	Gender string `json:"gender"`
	DOB string `json:"dateOfBirth"`
}

type LoginData struct{
	UserId string `json:"userId"`
	Password string `json:"password"`
}

func ReadCreateAccountBody(ctx iris.Context)(string,Data) {
	body :=Data{}
	err:=ctx.ReadBody(&body)
	fmt.Println(err)
	return "",body
}

func ReadLoginData(ctx iris.Context)(string,string,string ){
	userId:=ctx.FormValue("userId")
	password:=ctx.FormValue("password")
	
	return "",userId,password


}


