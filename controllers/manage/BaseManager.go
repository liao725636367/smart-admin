package manage

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"smartapp/controllers"
	"smartapp/helper"
	"smartapp/initial"
	"smartapp/initial/menu"
	"smartapp/models"
	"strings"
	"time"
)

type BaseManagerController struct{
	controllers.BaseController
	Agent models.Agent_user


}
func (c *BaseManagerController) Prepare(){
	//html后缀默认html
	c.BaseController.Prepare()
	//beego.Info("当前页面:"+c.Ctx.Input.URL())
	agentId :=c.GetSession("agent_id")
	aid,ok := agentId.(int)
	if !ok{
		c.Error("请登录",nil,map[string]interface{}{"err_code":403})
		//c.Redirect(beego.URLFor("manage.LoginController.Index"),302)
		return
	}
	authInfo :=initial.Cache.Get("agent_user_info_"+string(aid)) //用户数据缓存
	if userInfo ,ok:= authInfo.(models.Agent_user);ok{
		c.Agent = userInfo
	}else{
		db:=c.DB
		if agentId ==nil{
			c.Error("请登录",nil,map[string]int{"status":403})
			//c.Redirect(beego.URLFor("manage.LoginController.Index"),302)
			return
		}

		c.Agent=models.Agent_user{}
		err:=db.Where("id = ? ",aid).First(&c.Agent).Error
		beego.Debug("RecordNotFound 的判定是",db.RecordNotFound(),"agent_id",aid)
		if db.RecordNotFound()||c.Agent.Id < 1{
			c.DelSession("agent_id")
			c.Error("请登录",nil,map[string]int{"status":403})
			return
			//c.Redirect(beego.URLFor("manage.LoginController.Index"),302)
			//return
		}else if err!=nil{
			c.Error("数据库错误:"+err.Error())
			return
		}
		_ = initial.Cache.Put("agent_user_info_"+string(aid), c.Agent, time.Second*60)
	}

		//获取授权列表
		c.checkAccess()
	c.Data["agent"]=c.Agent
	c.Layout="manage/layout/layout.html"
		return


}

//获取权限菜单并筛选用户权限
func (c *BaseManagerController) checkAccess(){
	agent :=c.Agent
	menus:=make(map[string]string)

	for _,item:=range menu.ManageMenus{
		menus[item["key"]]=item["value"]
	}

	isAuth :=false //默认未授权
	//查询授权
	beego.Debug("是否超级管理员",agent.Is_admin)
	if agent.Is_admin >0{
		c.Data["is_auth"]=true
		c.Data["menus"]=menus
		c.Data["auth_name"]="超级管理员"
		return
	}
	//如果不是超级管理员就删除超级管理员菜单因为继承原因 就不显示

	//获取当前控制器和方法
	ctrl,act:=c.GetControllerAndAction()
	currPath :=ctrl+"."+act
	beego.Debug("当前访问权限",currPath)
	authInfo :=initial.Cache.Get("agent_auth_info_"+string(agent.Id))
	if authInfo,ok := authInfo.(map[string]interface{});ok{ //如果是 指定类型才处理
		menus ,ok1 :=authInfo["menus"].(map[string]string)
		auth ,ok2 :=authInfo["auth"].(models.Agent_package_auth)
		if ok1&&ok2{ //判断数据类型是否菜单和 菜单map
			for _,item:=range menus{
				if item == currPath{
					isAuth=true
				}
			}
			c.Data["is_auth"]= isAuth
			c.Data["menus"]=menus
			c.Data["auth_name"]= auth.Name
			beego.Debug("菜单数据使用了缓存")
			return
		}
	}

	var pages []string
	auth:=models.Agent_package_auth{}
	db:=c.DB
	err:=db.Where("id=?",agent.Auth_id).First(&auth).Error
	beego.Debug("RecordNotFound 的判定是",db.RecordNotFound())
	if db.RecordNotFound()||auth.Id <1{
		 pages=[]string{}
	}else if err!=nil{
		c.Error("数据库错误1"+err.Error())
		return
	}
	pages=strings.Split(auth.Nopages,",")



	//临时hash 存储 提高性能
	tmpMaps:=helper.SliceUniqueMap(pages)
	beego.Debug("数据库菜单数据",tmpMaps)
	for index,value:=range menus{
		if _,ok:=tmpMaps[value];ok{
			delete(menus,index)
		}
		if value == currPath{
			isAuth=true
		}
	}
	// 增加逻辑 不在路径黑名单的链接都能够被访问有些隐藏菜单需要的接口比如获取自己信息这种
	if _,ok:=tmpMaps[currPath];!ok{
		isAuth=true
	}
	_ = initial.Cache.Put("agent_auth_info_"+string(agent.Id), map[string]interface{}{ "menus": menus}, time.Duration(cache.DefaultEvery))
	if !isAuth {
		c.Error("无权访问此页面")
		//c.Redirect(c.URLFor("manage.LoginController.Index"),302)
		return
	}

	beego.Debug("是否授权",isAuth,"菜单列表",menus)

	c.Data["is_auth"]= isAuth
	c.Data["menus"]=menus
	c.Data["auth_name"]= auth.Name



}
