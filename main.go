package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"server/controllers"
)

func main() {
	app := iris.Default()

	//set environment variable

	//declare service end point

	app.Get("/ping1", func(ctx iris.Context) {
		fmt.Println("Successfully called test api")
		ctx.HTML("Fuck you bitch Mr. Sarvesh K.")
	})

	app.Get("/ping2", func(ctx iris.Context) {
		fmt.Println("")
		ctx.HTML("Ratnesk ki ma ki chut")
	})

	app.Put("/ping3", func(ctx iris.Context) {
		fmt.Println("Successfully called test api")
		ctx.HTML("Server up")
	})

	app.Post("/ping4", func(ctx iris.Context) {
		fmt.Println("Successfully called test api")
		ctx.HTML("Server up")
	})

	app.Get("/ping5", func(ctx iris.Context) {
		fmt.Println("Successfully called test api")
		ctx.HTML("Server up")
	})

	app.Post("/create-account", func(ctx iris.Context) {
		controllers.CreateUserAccount(ctx)		
	})

	app.Get("/login", func(ctx iris.Context) {
		controllers.Login(ctx)		
	})

	app.Get("/user-name", func(ctx iris.Context) {
		controllers.Login(ctx)		
	})
	
	app.Get("/mobile-number", func(ctx iris.Context) {
		controllers.Login(ctx)		
	})

	app.Get("/emailid", func(ctx iris.Context) {
		controllers.Login(ctx)		
	})

	app.Post("/attempt", func(ctx iris.Context) {
		controllers.UserDataValidation(ctx)		
	})

	app.Post("/transaction", func(ctx iris.Context) {
		controllers.UserDataValidation(ctx)		
	})


	app.Get("/data1", func(ctx iris.Context) {
		controllers.Data1(ctx)		
	})

	app.Get("/data2", func(ctx iris.Context) {
		controllers.Data2(ctx)		
	})

	app.Run(iris.Addr(":9000"))

}
