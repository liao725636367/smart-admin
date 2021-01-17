package common

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

//const (
//	LevelEmergency = iota // 紧急级别
//	LevelAlert                   // 报警级别
//	LevelCritical                // 严重错误级别
//	LevelError                   // 错误级别
//	LevelWarning              // 警告级别
//	LevelNotice                 // 注意级别
//	LevelInformational       // 报告级别
//	LevelDebug                 // 除错级别
//	)
func InitLog() {
	_ = beego.BeeLogger.SetLogger(logs.AdapterMultiFile, `
	{
	"filename": "logs/framework/app.log",
	"level": 7,
	"maxlines": 0,
	"maxsize": 0,
	"daily": true,
	"maxdays": 10,
	"color": true,
	"separate": [
		"emergency",
		"alert",
		"critical",
		"error",
		"warning",
		"notice",
		"info",
		"debug"
	]
}`)
	if beego.AppConfig.String("runmode") == "dev" {
		beego.BeeLogger.EnableFuncCallDepth(true) //输出文件名和行号
	}else{
		beego.BeeLogger.Async()
		beego.BeeLogger.SetLogFuncCallDepth(3)    //
	}
	beego.BeeLogger.SetLevel(beego.LevelDebug)



	//logs.Async(1e3)
	//logs.Notice()
	//beego.SetLogger("file", `{"filename":"logs/test.log"}`)
}
