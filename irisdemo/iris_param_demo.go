package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main()  {
	app := iris.Default()

	// This handler will match /user/john but will not match neither /user/ or /user.
	app.Get("/user/{name}", func(ctx context.Context) {
		name := ctx.Params().Get("name")
		ctx.Writef("Hello %s", name)
	})

	// This handler will match /users/42
	// but will not match /users/-1 because uint should be bigger than zero
	// neither /users or /users/.
	app.Get("/users/{id:uint64}", func(ctx iris.Context) {
		id := ctx.Params().GetUint64Default("id", 0)
		ctx.Writef("User With ID: %d", id)
	})

	// However, this one will match /user/john/send and also /user/john/everything/else/here
	// but will not match /user/john neither /user/john/.
	app.Post("/user/{name:string}/{action:path}", func(ctx context.Context) {
		name := ctx.Params().Get("name")
		action := ctx.Params().Get("action")
		message := name + " is " + action
		ctx.WriteString(message)
	})

	app.Run(iris.Addr(":8080"))
}
