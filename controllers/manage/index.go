package manage

import (
	"smartapp/initial"
)

type IndexController struct {
	 BaseManagerController
}
// @router  / [get]
func (c *IndexController) Index(){
	//第一次登陆
	firstLogin :=initial.Cache.Get("first_login")
	c.Data["first_login"]=false
	if isFirst,ok:= firstLogin.(bool); isFirst &&ok{
		c.Data["first_login"]=true
	}
	//beego.Debug("菜单内容是",c.Data["menus"])
	//c.Show()
	c.Success("获取用户信息成功",c.Data["agent"])
}
// @router  / [post]
func (c *IndexController) Index1(){
	c.Success("获取用户信息不支持post")
}
// @router /order/list [get]
func (c *IndexController) SearchOrder(){
	content :=`{"status":"success","msg":"获取成功","data":{"total":20,"items":[{"order_no":"EDE4A7BC-d2Cc-fcdD-df21-Ae3b99Ea5f16","timestamp":843440507335,"username":"Joseph Williams","price":14641.8,"status":"pending"},{"order_no":"F07ffEc7-Ff2E-6Ee9-B4f5-C5b391eD8F13","timestamp":843440507335,"username":"Maria Allen","price":4391.3,"status":"success"},{"order_no":"7EC3eca6-Cf6C-49C3-eAC2-6BceDE4BD94d","timestamp":843440507335,"username":"Shirley Miller","price":5821.8,"status":"pending"},{"order_no":"Cd054c14-EdA1-7b37-3FaD-D58BdDeBFA6c","timestamp":843440507335,"username":"John Davis","price":13432,"status":"success"},{"order_no":"de7C2Bb9-79f6-Ee46-D1f8-Ec5d4B8F8AbB","timestamp":843440507335,"username":"Matthew Anderson","price":4637.17,"status":"success"},{"order_no":"e11CBbb7-1cfB-F50D-D7d2-11063132b594","timestamp":843440507335,"username":"Jason Rodriguez","price":9542,"status":"success"},{"order_no":"2430Ac48-fbCD-1ee8-fF09-F98863Ad0ab2","timestamp":843440507335,"username":"Joseph Martinez","price":2515.3,"status":"pending"},{"order_no":"bF4be0E9-Ef4F-0c89-CDAC-7f6d83F77FEa","timestamp":843440507335,"username":"Eric Thomas","price":7113.8,"status":"success"},{"order_no":"eab5F5Ac-B926-8F4e-Bb9F-5614e40477c5","timestamp":843440507335,"username":"Paul Moore","price":9910.1,"status":"success"},{"order_no":"ab1764c9-9153-3C37-e5eB-01Eb2FDcdb5B","timestamp":843440507335,"username":"Nancy Young","price":11181.5,"status":"success"},{"order_no":"a9FdeEc1-3F3c-7f31-ccE4-Ad6038dA7EbC","timestamp":843440507335,"username":"Susan Rodriguez","price":10025.86,"status":"success"},{"order_no":"EFF1fAC7-Ad38-afc3-eeF8-566e4ef9C92c","timestamp":843440507335,"username":"Donna Walker","price":10580.4,"status":"success"},{"order_no":"7803B41b-DAAb-8Ea2-143f-Da54A18BBFCf","timestamp":843440507335,"username":"Jason Jackson","price":10839.2,"status":"success"},{"order_no":"E44D662E-eF27-E35F-B71b-66f6bcd85d33","timestamp":843440507335,"username":"Amy Williams","price":14659.83,"status":"success"},{"order_no":"49C3cBEE-b4CA-5cd7-Bc41-Bb857696Ff1f","timestamp":843440507335,"username":"Joseph Robinson","price":9361.4,"status":"pending"},{"order_no":"F24E6dc9-F5db-Bd77-C5Bb-7aD5083d3cAa","timestamp":843440507335,"username":"Elizabeth Thomas","price":11317,"status":"success"},{"order_no":"Ef1D83e8-b9dD-67d8-2db8-540D7cAE9F2D","timestamp":843440507335,"username":"Mary Hernandez","price":11949,"status":"pending"},{"order_no":"D2D2356A-5c71-f1B2-9f21-Be8D268C3079","timestamp":843440507335,"username":"Jennifer Walker","price":9881.1,"status":"success"},{"order_no":"01CEE4fC-Be8f-FeF2-8BaC-5f5B2baA16F9","timestamp":843440507335,"username":"Gary Jones","price":9336,"status":"success"},{"order_no":"Ffd9BFFF-08Ff-8288-AF42-d8eee6Dde95F","timestamp":843440507335,"username":"Scott Harris","price":4669.7,"status":"pending"}]}}`
	c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	c.Ctx.ResponseWriter.Write([]byte(content))

}
