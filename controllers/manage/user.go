package manage

type UserController struct {
	BaseManagerController
}
//@router /user/info [get]
func (c *UserController) GetInfo(){
	c.Success("获取用户信息成功",map[string]interface{}{"userInfo":c.Data["agent"],"roles":c.Data["menus"]})
}
