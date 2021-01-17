package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"smartapp/controllers"
	"smartapp/initial"
	"smartapp/models"
)

type BaseAdminController struct {
	controllers.BaseController
	User models.User
}
func (this *BaseAdminController) Prepare(){
	this.BaseController.Prepare()
	//html后缀默认html
	beego.Info("当前页面:"+this.Ctx.Input.URL())

	userId := this.GetSession("admin-uid")
	logs.Error(userId)
	 if userId == nil{
	 	this.Redirect("/admin/login",302)
	 }else{
	 	db:= initial.DB
	 	userId := userId.(int)
	 	user :=models.User{Id:userId}
	 	err :=db.First(&user).Error
	 	if err!=nil{
			this.Redirect("/admin/login",302)
		}else{
			this.User = user
		}

	 }
}
//请求结束的操作
func (this *BaseAdminController) Finish(){
	//获取url

	//记录日志操作

}
