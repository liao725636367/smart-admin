package common

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"smartapp/initial"
	"time"
)

func InitCache(){
	initial.Cache = cache.NewMemoryCache() //使用内存缓存 默认60秒执行gc(就是回收内存)
	//err := initial.Cache.StartAndGC("")
	//if err!=nil{
	//	panic("连接缓存服务失败:"+err.Error())
	//}

	//调试模式缓存2秒过期
	if beego.AppConfig.String("runmode") == "dev" {
		initial.CacheTime = time.Second*6
	}else{
		initial.CacheTime = time.Second*60
	}
}