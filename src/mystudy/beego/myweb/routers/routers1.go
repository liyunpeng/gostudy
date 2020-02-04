package routers

import (
	"github.com/astaxie/beego"
	"gostudy/src/mystudy/beego/myweb/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}