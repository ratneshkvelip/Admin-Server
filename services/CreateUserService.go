package services

import (
	"server/common/dao"
	"server/common/query"
	"server/common/utils"

	"github.com/kataras/iris/v12"
)

type payload struct {
	Message string `json:"message"`
}

type rsp struct {
	Status  string  `json:"status"`
	Payload payload `json:"payload"`
}

func CreateUserAccount(ctx iris.Context, userData utils.Data) (int, interface{}) {
	var statusCode int
	Query := query.CreateAccount()
	db := dao.DbClient()
	resp := rsp{}

	tx, _ := db.BeginTx(ctx, nil)
	defer db.Close()
	defer tx.Rollback()
	// row:=db.QueryRow(Query)

	// err:=row.Scan(&testdata)

	_, err := tx.Exec(Query, userData.UserID, userData.UserName, userData.EmailId, userData.Gender, userData.MobileNumber, userData.Country, "login_token", userData.Password, "A", "user", "user")
	if err != nil {
		statusCode = 400
		resp.Status = "OK"
		resp.Payload.Message = "Error in creating account."
		return statusCode, resp
	}
	statusCode = 200
	resp.Status = "Failed"
	resp.Payload.Message = "Successfully created account."

	tx.Commit()

	return statusCode, resp

}
