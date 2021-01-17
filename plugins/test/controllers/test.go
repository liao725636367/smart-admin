package controllers

type TestController struct {
	BaseController
}
// @router /test/?:id  [get]
func (c *TestController) Index(){
	//c.TplName="index.html"
	//c.Success("成功2",nil)
	c.Show()
}
func (c *TestController) Get(){
	c.Success("成功",nil)
}