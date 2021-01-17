package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["smartapp/plugins/test/controller:TestController"] = append(beego.GlobalControllerRouter["smartapp/plugins/test/controller:TestController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/test/?:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
