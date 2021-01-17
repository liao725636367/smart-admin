package manage

import (
	"fmt"
	"github.com/astaxie/beego"
	"smartapp/controllers"
	"smartapp/initial"
	"smartapp/models"
	"time"
)

type LoginController struct {
	controllers.BaseController
}
// @router /user/login [get]
func (c * LoginController) Index() {
	c.Show()
}
//代理商登录 员工登录准备使用另一个方法处理 另一个表 agent_staff 来存储，实现每个代理可以随意设置员工账号
// @router /user/login [post]
func (c *LoginController) DoLogin(){
	username := c.GetS("username")
	password :=c.GetS("password")

	//找出用户
	db:=initial.DB
	agent:=models.Agent_user{}
	 err :=db.Where("username = ?",username).First(&agent).Error

	if db.RecordNotFound()||agent.Id < 1 {
		beego.Error("用户不存在",fmt.Sprintf("%v",agent.Id))
		c.Error("用户不存在")
	}else if err!=nil {

		c.Error("数据库出现错误:"+err.Error())
		return
	}
	if !agent.CheckPwd(password){
		c.Error("密码错误")
		return
	}
	//获取权限 判断管理员
	//agentData,_:=json.Marshal(agent)
	//	//c.SetSession("agent",agentData)
	c.SetSession("agent_id",agent.Id)

	_ = initial.Cache.Put("first_login", 1,time.Second*60)
	c.Success(" 登录成功",agent)
	 return
	//员工 不在此登录，不做处理


}
//@router /user/logout [post]
func (c *LoginController) LoginOut(){
	c.DestroySession()
	c.Success("退出登录成功")
}
