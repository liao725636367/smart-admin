package plugins

import (
	"github.com/astaxie/beego"
	_ "smartapp/plugins/test/initial"
)
func init(){
	//initial.TestInit() //
	beego.Debug("初始化插件信息")
}