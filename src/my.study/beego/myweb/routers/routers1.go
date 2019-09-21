package routers

import (
	"github.com/astaxie/beego"
	"my.study/beego/myweb/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}