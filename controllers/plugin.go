package controllers

import (
	"github.com/astaxie/beego"
	"smartapp/helper"
	"strings"
)

type PluginController struct {
	BaseController
}

func (c *PluginController) Show(tplName... string){
	//c.TplName=tplName
	file := helper.GetControllerStackFile(1)
	file= (strings.SplitN(file,"/plugins/",2))[1]
	file=(strings.SplitN(file,".go",2))[0]
	files :=strings.SplitN(file,"/controller/",2)
	beego.Debug("当前url1",c.Ctx.Input.URL(),"当前文件:",file)
	_,act:=c.GetControllerAndAction()
	if c.TplName == "" {
		c.TplName = "plugins/"+strings.ToLower(files[0])+"/"+strings.ToLower(files[1]) + "/" + strings.ToLower(act) + "." + c.TplExt
		//c.TplName= "/plugins/test"+strings.TrimLeft(c.TplName,)
	}else  if !strings.Contains(c.TplName, "plugins/"){
		c.TplName = "plugins/"+strings.ToLower(files[0])+"/"+strings.TrimLeft(c.TplName,"/")
	}
	_ = c.Render()
	return
}
