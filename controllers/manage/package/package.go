package manage_package

import (
	"github.com/astaxie/beego"
	"smartapp/controllers/manage"
	"smartapp/helper"
	"smartapp/initial"
	"smartapp/models"
)

type PackageController struct {
	manage.BaseManagerController
}
//获取套餐等级列表
// @router  /package/levels [get]
func (c *PackageController) Levels(){
	db:=initial.DB
	//页码
	//第一次登陆
	var err error
	var pack []models.Agent_package_auth
	returnData := make(map[string]interface{})
	model := db.Model(&models.Agent_package_auth{}).Where("aid = ? ",c.Agent.Id)
	returnData["pageData"],err =c.Pagination(model,&pack,10)
	if err!=nil{
		c.Error("服务器出现错误:"+err.Error())
		return
	}
	returnData["list"]=pack
	//beego.Debug("菜单内容是",c.Data["menus"])
	//c.Show()
	c.Success("获取用户信息成功",returnData)
}

//添加行业套餐
// @router  /package/levels [post]
func (c *PackageController) AddLevel(){

	levelData :=models.Agent_package_auth{}
	data:=c.Ctx.Input.RequestBody
	err:= initial.Json.Unmarshal(data, &levelData)
	beego.Debug("数据是",levelData)
	if err!=nil{
		c.Error("出现错误"+err.Error())
		return
	}
	//判断套餐数量
	var count int
	c.DB.Model(models.Agent_package_auth{}).Count(&count)
	if count >= 10{
		c.Error("最多可创建10个套餐")
		return
	}
	levelData.Id=0
	levelData.Created_at = helper.TimeStr()
	levelData.Updated_at = helper.TimeStr()
	levelData.Aid = c.Agent.Id
	levelData.Type = "agent"
	err =c.DB.Create(&levelData).Error
	if err!=nil{
		c.Error("出现错误"+err.Error())
		return
	}
	if levelData.Id > 0{
		c.Success("添加成功",levelData)
	}else{
		c.Error("添加失败")
	}
}
//批量更新行业套餐
// @router  /package/levels  [put]
func (c *PackageController) UpdateLevels(){
	 var levels []models.Agent_package_auth
	 var levelsOld  []models.Agent_package_auth
	data:=c.Ctx.Input.RequestBody
	err:= initial.Json.Unmarshal(data, &levels)
	if err!=nil{
		c.Error("出现错误"+err.Error())
		return
	}
	if len(levels) > 0{
		//查询旧数据
		err = c.DB.Where("aid = ? ",c.Agent.Id).Find(&levelsOld).Error
		if !c.DB.RecordNotFound()&&err!=nil{
			c.Error("出现错误"+err.Error())
			return
		}
		for _,item:=range levels{
			for _,item1:=range levelsOld{
				if item.Id==item1.Id && (item.Name!=item1.Name||item.Nums!=item1.Nums){
					err  = c.DB.Model(&item1).Select("name","nums").Updates(item).Error
					if err!=nil{
						c.Error("出现错误"+err.Error())
						return
					}
				}
			}

		}
		c.Success("修改成功")
		return

	}

	//查询旧的套餐数据
	//if err!=nil{
	//	c.Error("删除失败"+err.Error())
	//}else{
	//	c.Success("删除成功",id)
	//}
}

//删除行业套餐
// @router  /package/levels/:id  [delete]
func (c *PackageController) DelLevels(){
	id,_:=c.GetInt(":id",0)
	err := c.DB.Where("aid= ? and id = ?",c.Agent.Id,id).Delete(models.Agent_package_auth{}).Error
	if err!=nil{
		c.Error("删除失败"+err.Error())
	}else{
		c.Success("删除成功",id)
	}
}


