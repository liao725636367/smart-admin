package controllers

import (
	"github.com/astaxie/beego"
)

type TestControllers struct {
	beego.Controller
}
func init(){

}
func (c *TestControllers) Get(){
	c.Ctx.WriteString("hello laravels")
}
