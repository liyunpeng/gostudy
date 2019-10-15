package iris1

import (
	"github.com/kataras/iris"
	"time"
)

const refreshEvery = 10 * time.Second

func CacheClient() {
	app := iris.New()
	app.Use(iris.Cache304(refreshEvery))
	app.Logger().SetLevel("debug")
	// 等同于
	// app.Use(func(ctx iris.Context) {
	//     now := time.Now()
	//     if modified, err :=
	//     ctx.CheckIfModifiedSince(now.Add(-refresh)); !modified && err == nil {
	//         ctx.WriteNotModified()
	//         return
	//     }
	//     ctx.SetLastModified(now)
	//     ctx.Next()
	// })
	app.Get("/", greet)
	app.Run(iris.Addr(":8080"))
}

func greet(ctx iris.Context) {
	ctx.Header("X-Custom", "my  custom header")
	ctx.Writef("Hello World! %s", time.Now())
}
