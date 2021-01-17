package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	"smartapp/helper"
	"smartapp/initial"
	"strconv"
	"strings"
)

type BaseController struct {
	beego.Controller
	DB *gorm.DB
}

func (c *BaseController) Prepare() {
	c.TplExt = "html" //默认html后缀
	//ctl, act := c.GetControllerAndAction()
	//beego.Error("控制器和方法", ctl, act)
	//beego.Debug("路径", beego.AppPath)
	c.Data["title"] = "智能小程序"
	//针对options返回信息
	beego.Debug("请求方法",c.Ctx.Request.Method)
	c.DB =initial.DB
	//if c.Ctx.Request.Method =="options"{
	//
	//}
	//c.Ctx.Output.Header("Access-Control-Expose-Headers", "*")
	//c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	//c.Ctx.Output.Header("Access-Control-Allow-Methods", "*")
	//c.Ctx.Output.Header("Access-Control-Allow-Headers", "*")

}

//返回正确ajax json 信息 预留方便后续统一修改
func (c *BaseController) Success(msg string, data... interface{})  {
	jsonData := make(map[string]interface{})
	jsonData["status"] = "success"
	jsonData["msg"] = msg
	jsonData["data"] = nil
	if len(data) >= 1 {
		jsonData["data"] = data[0]

	}
	//只取第三个参数为额外的 返回json数据字段 比如
	if len(data) > 1 {
		if exts, ok := data[1].(map[string]interface{}); ok {
			for index, item := range exts {
				jsonData[index] = item
			}
		}

	}
	c.Data["json"]=jsonData
	c.ServeJSON()

}

//返回错误ajax json 信息  预留方便后续统一修改
func (c *BaseController) Error(msg string, data... interface{}) {
	//isAjaxMethod:=this.IsAjax()
	jsonData := make(map[string]interface{})
	jsonData["status"] = "error"
	jsonData["msg"] = msg
	jsonData["data"] = nil
	if len(data) >= 1 {
		jsonData["data"] = data[0]

	}
	//只取第三个参数为额外的 返回json数据字段 比如
	if len(data) > 1 {
		if exts, ok := data[1].(map[string]interface{}); ok {
			for index, item := range exts {
				jsonData[index] = item
			}
		}

	}
	c.Data["json"]=jsonData
	c.ServeJSON()


}

//自动定位模板目录
//使用统一函数方便记录
func (c *BaseController) ParseTplName(tplName ...string) {
	beego.Debug("base-当前url", c.Ctx.Input.URL())
	//获取文件路径以ccontrollers为分界线自动找到模板

	if len(tplName) > 0 && tplName[0] != "" {
		c.TplName = tplName[0]
		//beego.Error("当前模板路径:", c.TplName)
	} else if c.TplName == "" {
		file := helper.GetControllerStackFile(1)
		//beego.Error("parseTplName 模板文件路径" , helper.FileWithLineNum())
		files := strings.Split(file, "/controllers/")
		file = files[len(files)-1] //避免客户目录带了controllers 不可控因素
		file = (strings.SplitN(file, ".go", 2))[0]
		//beego.Debug("模板路径文件", file)
		_, act := c.GetControllerAndAction()
		//beego.Debug("当前url", c.Ctx.Input.URL())
		c.TplName = strings.ToLower(file) + "/" + strings.ToLower(act) + "." + c.TplExt

	}
	//beego.Error("当前模板:",c.TplName)
}
func (c *BaseController) Show(tplName ...string) {
	//_, file, line, ok:=runtime.Caller(1)
	//beego.Error("show 模板文件路径" , helper.FileWithLineNum(),"当前文件信息", file, line, ok)
	c.ParseTplName(tplName...)
	_ = c.Render()

}
func (c *BaseController) Finish(){
}

func (this *BaseController) RequestBody() []byte {
	return this.Ctx.Input.RequestBody
}
func (this *BaseController) decodeRawRequestBodyJson() map[string]interface{} {
	var mm interface{}
	requestBody := make(map[string]interface{})
	_ = json.Unmarshal(this.RequestBody(), &mm)
	if mm != nil {
		var m1 map[string]interface{}
		m1 = mm.(map[string]interface{})
		for k, v := range m1 {
			requestBody[k] = v
		}
	}
	return requestBody
}
func (this *BaseController) JsonData() map[string]interface{} {
	return this.decodeRawRequestBodyJson()
}

func (c *BaseController) GetValue(key string, def ...interface{}) interface{} {
	var value interface{}
	//先尝试获取表单数据

	value  = c.Ctx.Input.Query(key)
	if value == nil || value == ""{
		data:=c.JsonData()
		value =data[key]
		if value == nil {
			if len(def) > 0{
				return def[0]
			}
		}else{
			return value
		}
	}else{
		beego.Debug(fmt.Sprintf("获取的请求值 %s:%v",key,value))
		return value
	}
	return nil


}
//获取 string
func (c *BaseController) GetS(key string, def ...string ) string {
	value := fmt.Sprintf("%v",c.GetValue(key))
	if value!="nil"{
		return value
	}else{
		if len(def) >0{
			return def[0]
		}else{
			return ""
		}
	}
}
//获取int类型数据
func (c *BaseController) GetI(key string, def ...int ) (int,error) {
	strv := fmt.Sprintf("%v",c.GetValue(key))
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return strconv.Atoi(strv)
}
//获取布尔值
// GetBool returns input value as bool or the default value while it's present and input is blank.
func (c *BaseController) GetB(key string, def ...bool) (bool, error) {
	strv := fmt.Sprintf("%v",c.GetValue(key))
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return strconv.ParseBool(strv)
}
//获取float值
// GetFloat returns input value as float64 or the default value while it's present and input is blank.
func (c *BaseController) GetF(key string, def ...float64) (float64, error) {
	strv := fmt.Sprintf("%v",c.GetValue(key))
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return strconv.ParseFloat(strv, 64)
}
//数据统一分页处理
func (c *BaseController) Pagination(db *gorm.DB,result interface{},pageSize int)(pageData map[string]interface{},err error){
	page,_:=c.GetI("page",1)
	var count int
	db.Count(&count)
	totalPage:= count/pageSize
	if page < 1{
		page =1
	}else if  page >totalPage{
		page=totalPage
	}
	offset := (page -1)*pageSize
	err =db.Limit(pageSize).Offset(offset).Find(result).Error
	if !db.RecordNotFound()&&err!=nil{
		return nil,err
	}
	pageData =make(map[string]interface{})
	pageData["totalPage"]=totalPage
	pageData["total"]=count
	pageData["pageSize"]=pageSize
	pageData["page"]=page
	return pageData,nil

}


