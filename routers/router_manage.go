package routers

import (
	"github.com/astaxie/beego"
	"smartapp/controllers/manage"
	manage_package "smartapp/controllers/manage/package"
)

func init() {

	ns1:=beego.NewNamespace("/manage",
		beego.NSInclude(&manage.IndexController{}),
		beego.NSInclude(&manage.LoginController{}),
		beego.NSInclude(&manage.CommonController{}),
		beego.NSInclude(&manage.UserController{}),
		beego.NSInclude(&manage_package.PackageController{}),
		beego.NSInclude(&manage_package.AuthController{}),
	)
	beego.AddNamespace(ns1)


}
