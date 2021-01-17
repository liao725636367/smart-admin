package manage

import "smartapp/initial/menu"

type CommonController struct{
	BaseManagerController
}

//获取所有权限
//@router /get_role [post]
func  (c *CommonController) GetRole(){
	menus:=make(map[string]string)
	for _,item :=range menu.ManageMenus{
		menus[item["key"]]=item["value"]
	}
	c.Success("获取成功",c.Data["menus"])
	return
}
