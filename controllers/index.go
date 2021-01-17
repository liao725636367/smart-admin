package controllers

type IndexController struct {
	BaseController
}

func (c *IndexController) Get() {
	c.TplName="index.html"
	c.Show()
}
