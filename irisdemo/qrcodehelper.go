package main

import (
"github.com/kataras/iris"
"github.com/kataras/iris/context"
//"github.com/skip2/go-qrcode"
)

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./views", ".html"))

	app.Get("/", func(ctx context.Context) {
		ctx.ViewData("message", "Hello world!")
		ctx.View("hello.html")
	})

	app.Get("/user/{id:long}", func(ctx context.Context) {
		userID, _ := ctx.Params().GetInt64("id")
		ctx.Writef("User ID: %d", userID)
	})

	app.Get("/qrcode", func(ctx context.Context) {
		userID, _ := ctx.Params().GetInt64("id")
		data := ctx.FormValue("qrdata")
		ctx.Writef("User ID: %d", userID)
		ctx.Writef("data: %v", data)
	})

	app.Run(iris.Addr(":8080"))
}
