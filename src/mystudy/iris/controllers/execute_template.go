package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/_examples/http_responsewriter/quicktemplate/templates"
)
// ExecuteTemplate将“tmpl”部分模板渲染给`context＃ResponseWriter`。
func ExecuteTemplate(ctx iris.Context, tmpl templates.Partial) {
	ctx.Gzip(true)
	ctx.ContentType("text/html")
	templates.WriteTemplate(ctx.ResponseWriter(), tmpl)
}