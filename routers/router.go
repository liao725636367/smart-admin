package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"smartapp/controllers"
	"smartapp/controllers/admin"
)

func init() {
	//错误页面处理
	//预请求处理
	beego.Options("/*", func(c *context.Context) {
		c.ResponseWriter.Write([]byte("ok"))
	})
	beego.Router("/",&controllers.IndexController{})
	beego.Router("/test",&controllers.TestControllers{})
	beego.ErrorController(&controllers.ErrorController{})
	ns:=beego.NewNamespace("/admin",
		beego.NSInclude(&admin.IndexController{}),
		beego.NSInclude(&admin.LoginController{}),
	)
	beego.AddNamespace(ns)



}
