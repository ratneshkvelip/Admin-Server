package services

import (
	"server/common/dao"
	"server/common/query"
)

func Login(userId string, password string) (int, interface{}) {
	var user_password string
	var statusCode int

	resp := rsp{}

	if userId == "" {
		statusCode = 200
		resp.Status = "Fail"
		resp.Payload.Message = "Invalid user name."
		return statusCode, resp
	}
	db := dao.DbClient()

	defer db.Close()

	loginQuery := query.GetPassword()

	err := db.QueryRow(loginQuery, userId).Scan(&user_password)

	if err != nil {
		statusCode = 200
		resp.Status = "Fail"
		resp.Payload.Message = "Invalid user name."
		return statusCode, resp
	}

	if password == user_password {
		statusCode = 200
		resp.Status = "OK"
		resp.Payload.Message = "Successfully logged in."
	} else {
		statusCode = 200
		resp.Status = "Fail"
		resp.Payload.Message = "Bosdike password galat hai. Sahi kar."
	}

	return statusCode, resp

}
