package main

import (
	"github.com/astaxie/beego"
	_ "smartapp/initial/common"

	_ "smartapp/initial/plugins"
	_ "smartapp/routers"
)




func main() {
	beego.Run()
	

}



