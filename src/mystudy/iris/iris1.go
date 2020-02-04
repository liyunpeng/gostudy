package iris1

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"

	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)


func newApp1() *iris.Application {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

	mvc.New(app).Handle(new(ExampleController))
	return app
}

func IrisMain() {
	app := newApp1()

	app.Run(iris.Addr(":8080"))
}

type ExampleController struct{}

func (c *ExampleController) Get() mvc.Result {
	return mvc.Response{
		ContentType: "text/html",
		Text:        "<h1>Welcome</h1>",
	}
}

func (c *ExampleController) GetPing() string {
	return "pong"
}

func (c *ExampleController) GetTest1() string {
	return "test1"
}

func (c *ExampleController) GetHello() interface{} {
	return map[string]string{"message": "Hello Iris!"}
}

func (c *ExampleController) BeforeActivation(b mvc.BeforeActivation) {
	anyMiddlewareHere := func(ctx iris.Context) {
		ctx.Application().Logger().Warnf("Inside /custom_path")
		ctx.Next()
	}
	b.Handle(
		"GET",
		"/custom_path",
		"CustomHandlerWithoutFollowingTheNamingGuide", anyMiddlewareHere)

}

func (c *ExampleController) CustomHandlerWithoutFollowingTheNamingGuide() string {
	return "hello from the custom handler without following the naming guide"
}