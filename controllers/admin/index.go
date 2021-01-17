package admin

import (
	

)

type IndexController struct {
	BaseAdminController
}
// @router / [get]
func (c *IndexController)  Index()  {
	user:=c.User
	//var str string
	//str="这是首页"+user.Username+"-"+strconv.Itoa(user.Id)
	//c.Ctx.ResponseWriter.Write([]byte(str))
	c.Data["user"]=user
	c.Show()
	//c.TplName="admin/index.html"
}