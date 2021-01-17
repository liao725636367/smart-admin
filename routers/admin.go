package routers

import (
	"github.com/astaxie/beego"
	"smartapp/controllers"
)

func init(){



	model :=beego.AppConfig.String("runmode")
	if model == "dev"{
		beego.Router("/test",&controllers.TestControllers{})
	}
}
