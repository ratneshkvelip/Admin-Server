package services


import(
	"fmt"
	"server/common/dao"
	"server/common/query"

)



func UserDataValidation(UserName string)(int, interface{}){
	var statusCode int
	Query := query.GetUserName()
	db := dao.DbClient()
	resp := rsp{}
	var user_name string
	defer db.Close()
	
	row:=db.QueryRow(Query,UserName)

	err:=row.Scan(&user_name)
	
	statusCode=200

	if err != nil {
		resp.Status ="OK"
		resp.Payload.Message = "Valid user name."
		return  statusCode,resp
	}
 	
	resp.Status = "Fail"
	resp.Payload.Message = "User Name already taken."

	fmt.Println(user_name)

	return  statusCode, resp

}