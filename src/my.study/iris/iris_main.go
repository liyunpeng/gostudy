package iris1

import (
	//"./controllers"

	"github.com/kataras/iris"
	"my.study/iris/controllers"
)

func newApp() *iris.Application {
	app := iris.New()
	app.Get("/", controllers.Index)
	app.Get("/{name}", controllers.Hello)
	return app
}

func Irismain1() {
	app := newApp()
	// http://localhost:8080
	// http://localhost:8080/yourname
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}