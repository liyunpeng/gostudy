package controllers

import (
	//"../templates"
	"github.com/kataras/iris"
	"github.com/kataras/iris/_examples/http_responsewriter/quicktemplate/templates"
)
//索引使用已编译的../templates/index.qtpl.go文件渲染我们的../templates/index.qtpl文件。
func Index(ctx iris.Context) {
	tmpl := &templates.Index{}
	//渲染模板
	ExecuteTemplate(ctx, tmpl)
}