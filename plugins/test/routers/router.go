package routers

import (
	"github.com/astaxie/beego"
	"smartapp/plugins/test/controllers"
)

func init(){
	beego.Debug("插件test路由初始化")
	beego.Router("/plugin/test/index", &controllers.TestController{},"*:Index")
	ns1:=beego.NewNamespace("/plugin",
		beego.NSInclude(&controllers.TestController{}),
	)
	beego.AddNamespace(ns1)
}