package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"smartapp/controllers"
	"smartapp/initial"
	"smartapp/models"

	//"github.com/astaxie/beego/orm"
	//"smartapp/models"
)

// LoginController operations for Login
type LoginController struct {
	controllers.BaseController
}

// @router /login [get]
func (c *LoginController) Index() {


		//c.TplName="admin/login.html"
		c.Show()

}

func (c *LoginController) Get() {

}
// @router /login [post]
func (c *LoginController) DoLogin(){
	username :=c.Input().Get("username")
	password :=c.Input().Get("password")
	//c.TplName="admin/login.html"
	//c.Show()
	//html,_:=c.RenderString()

	db:= initial.DB
	if username==""||password==""{
		logs.Notice("获取的用户名密码是:",username,"---",password)
		//c.Ctx.ResponseWriter.Write([]byte("<script>alert('用户名或密码为空')</script>"+html))
		c.Error("用户名密码不能为空",nil)
	}else{
		user:=models.User{Username:username,Password:password}
		err:=db.Where("username=? and password =? ",username,password).Find(&user).Error
		if user.Id == 0{
			c.Error("用户不存在",nil)
		}
		if err!=nil{
			 msg:=fmt.Sprintf("%v",err)
			c.Error(msg,nil)
			return
		}else{

			c.SetSession("admin-uid",user.Id)
			//c.Redirect("/admin",302)
			c.Success(" 登录成功",map[string]interface{}{"url":beego.URLFor("IndexController.Index")})

		}

	}
	return

}
