package services

import (
	"fmt"
	"server/common/query"
	"server/common/dao"
)

func Validation(userId string,password string, token string){
	Query:=query.GetName()
	testdata:=""
	db:=dao.DbClient()

	defer db.Close()
	row:=db.QueryRow(Query)
	
	err:=row.Scan(&testdata)	

	if err!=nil{
		fmt.Println(err)
		
	}

	fmt.Println(testdata)

}