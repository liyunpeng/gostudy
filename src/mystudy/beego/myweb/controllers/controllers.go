package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller  // 这里可以看做是其他语言中的继承
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	//TODO beego 模板文件加载
	c.TplName = "beego_index.tpl"
}
