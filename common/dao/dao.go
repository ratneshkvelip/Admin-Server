package dao

import(
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func DbClient() (*sql.DB) {
	host:="localhost"
	port:="5432"
	user:="postgres"
	password:="Rkvelip@135"
	dbname:="postgres"

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	dbClient, err := sql.Open("postgres", psqlInfo)

	if err!=nil{
		fmt.Println(err)
	}
	
	err=dbClient.Ping()
	if err!=nil{
		fmt.Println(err)
	}

	return dbClient
	
}