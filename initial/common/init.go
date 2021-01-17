package common

import (
	"github.com/astaxie/beego"
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
	"smartapp/initial"
)

func init()  {
	//initial.TestInit() //初始化test插件
	beego.BConfig.WebConfig.AutoRender = false //不自动渲染模板
	beego.BConfig.WebConfig.Session.SessionName ="session_id" //使用session_id
	beego.BConfig.WebConfig.Session.SessionEnableSidInHTTPHeader =true //使用http头设置session
	beego.BConfig.WebConfig.Session.SessionNameInHTTPHeader ="Session_id"
	beego.BConfig.WebConfig.Session.SessionAutoSetCookie = false
	beego.BConfig.CopyRequestBody = true //允许json请求
	InitSql()
	InitLog()
	InitCache()
	err := beego.AddFuncMap("sliceFirst", initial.SliceFirst)
	if err!=nil{
		panic("注册模板函数失败:"+err.Error())
	}
	//使用jsoniter 代替json
	extra.RegisterFuzzyDecoders()
	initial.Json = jsoniter.ConfigCompatibleWithStandardLibrary

}
