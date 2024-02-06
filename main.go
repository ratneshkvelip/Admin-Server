package main

import(
	"fmt"
	"github.com/kataras/iris/v12"
)

func main() {
	app:=iris.Default()

	app.Listen(":8000")

	fmt.Println("hello World")
	
}
