package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"log"
	"time"
)


func main() {
	keyMaps := make(map[int64]chan bool)

	app := iris.New()
	app.RegisterView(iris.HTML("./views", ".html"))

	app.Get("/", func(ctx context.Context) {
		ctx.ViewData("message", "Hello world!")
		ctx.View("hello.html")
	})

	app.Get("/user/{id:long}", func(ctx context.Context) {
		userID, _ := ctx.Params().GetInt64("id")
		log.Println("request received : ", userID)
		processTime := int(userID % 10) * 5
		//time.Sleep(time.Duration(processTime * int(time.Second)))
		waitKey := make(chan bool)
		keyMaps[userID] = waitKey
		_ = <- waitKey
		ctx.Writef("User ID: %d, process time : %d", userID, processTime)
	})

	app.Get("/wake/{id:long}", func(ctx context.Context) {
		userID, _ := ctx.Params().GetInt64("id")
		keyWaiter, ok := keyMaps[userID]
		if ok {
			keyWaiter <- true
		}
	})

	app.Get("/ping", func(ctx context.Context) {
		time.Sleep(time.Duration(5 * int(time.Second)))
		ctx.JSON(iris.Map{
			"message": "pong",
		})
	})

	app.Run(iris.Addr(":8880"))
}