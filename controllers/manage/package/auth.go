package manage_package

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"smartapp/controllers/manage"
	"smartapp/helper"
	"smartapp/initial/menu"
	"smartapp/models"
	"strings"
)

type AuthController struct {
	manage.BaseManagerController
}
//获取行业权限
//@router /package/roles/:id [get]
func (c *AuthController) GetRoles(){
	menus :=c.Data["menus"]
	 id:=c.GetS(":id")
	 packData :=models.Agent_package_auth{}

	menus1,ok:=menus.(map[string]string)
	 if !ok{
	 	 c.Error("没有权限")
		 return
	 }
	 err :=c.DB.Where("id=?",id).First(&packData).Error
	 if err != nil{
	 	 c.Error("获取失败")
	 }
	 noPages :=strings.Split(packData.Nopages,",")
	 //pageMenus:=menus //权限页面以当前用户权限为标准
		menus2 :=make(map[string]string)
		for index,value:=range menus1{
			menus2[index]=value
			loop:
			for _,noPage:=range noPages{
				if noPage == value{
					delete(menus1,index)
					break loop
				}
			}
		}

	c.Success("获取成功",map[string]interface{}{"pack_menus":menus1,"menus":menus2})
}
//设置行业权限 根据当前用户权限继承后 再加上设置后没有的节点进行扩充
//@router /package/roles/:id [put]
func (c *AuthController) SavePackageAuth(){
	var dataStruct = struct { //结构体接收数据
		Nodes []string
	}{}
	packageId,_:=c.GetInt(":id")
	data := c.Ctx.Input.RequestBody
	err:=json.Unmarshal(data,&dataStruct)
	if err!=nil{
		c.Error("数据格式错误")
		return
	}
	nodes:=dataStruct.Nodes

	var packData models.Agent_package_auth
	var baseNoPages []string

	if c.Agent.Auth_id >0 { //先找到当前用户没有权限的页面
		err =c.DB.Where("id=?",c.Agent.Auth_id).First(&packData).Error
		if err != nil{
			c.Error("获取用户权限失败")
			return
		}
		baseNoPages = strings.Split(packData.Nopages,",")

	}else{
		baseNoPages = []string{}
	}
	baseNoPagesMap := helper.SliceUniqueMap(baseNoPages) //这里是控制器方法
	nodesMap := helper.SliceUniqueMap(nodes) //这里是节点别名
	var SetNoPages []string
	for _,item:=range menu.ManageMenus{
		_,ok:=baseNoPagesMap[item["value"]]
		_,ok1:=nodesMap[item["key"]]
		if ok||!ok1{//用户已经没有权限，或者设置了没有权限的，视为没有权限节点
			SetNoPages = append(SetNoPages,item["value"] )
		}
	}
	beego.Debug("设置的没权限节点",SetNoPages)

	c.DB.Model(models.Agent_package_auth{}).Where("id=?",packageId).Update("Nopages",strings.Join(SetNoPages,","))
	c.Success("设置成功")
}
