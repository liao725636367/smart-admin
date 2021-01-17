package models

import (
	"github.com/astaxie/beego"
	"smartapp/helper"
)

//
type Agent_user struct {
	Id	int	`gorm:"primary_key" json:" - "` //
	Username	string	`json:"username"` //用户名
	Password	string	`json:"-"` //密码
	Nickname	string	`json:"nickname"` //管理员昵称
	Web_title	string	`json:"web_title"` //网站标题
	Keywords	string	`json:"keywords"` //网站关键字
	Describe	string	`json:"describe"` //网站描述
	Company_name	string	`json:"company_name"` //公司名称
	Company_address	string	`json:"company_address"` //公司地址
	Record_number	string	`json:"record_number"` //备案号
	Qr_code	string	`json:"qr_code"` //公司宣传二维码
	Email	string	`json:"email"` //邮箱
	Telephone	string	`json:"telephone"` //电话
	Qq	int	`json:"qq"` //qq号
	Logtimes	bigint	`json:"logtimes"` //登录次数 每登一次加一
	Is_admin	tinyint	`json:"is_admin"` //是否是管理员
	Balance	decimal	`json:"balance"` //网站
	Domain	string	`json:"domain"` //代理商绑定域名
	Webname	string	`json:"webname"` //网站名称
	Logo	string	`json:"logo"` //
	Service_code	string	`json:"service_code"` //客服二维码
	Pid	int	`json:"pid"` //上级代理id
	Bind_time	timestamp	`json:"bind_time"` //关联上级时间
	Auth_id	tinyint	`gorm:"primary_key" json:" - "` //代理商关联权限id（管理员无视权限）
	Province_id	int	`json:"province_id"` //省id
	City_id	int	`json:"city_id"` //市id
	Region_id	int	`json:"region_id"` //区id
	Name	string	`json:"name"` //代理商法人姓名
	Discount_level	int	`json:"discount_level"` //代理商享受折扣等级id
	App_id	string	`json:"app_id"` //开放平台配置appid
	App_secret	string	`json:"app_secret"` //开放平台配置appsecret
	Token	string	`json:"token"` //开放平台配置token
	Encoding_aes_key	string	`json:"encoding_aes_key"` //开放平台配置encoding_aes_key
	Config_domain	string	`json:"config_domain"` //给授权方小程序设置的白名单域名
	Template_id	string	`json:"template_id"` //绑定的template_id
	Recharge_appid	string	`json:"recharge_appid"` //充值用的appid
	Recharge_mchid	string	`json:"recharge_mchid"` //充值商户id
	Recharge_key	string	`json:"recharge_key"` //充值商户key
	Default_user_level	int	`json:"default_user_level"` //默认客户级别
	Default_user_days	int	`json:"default_user_days"` //默认客户天数
	Sms_id	int	`json:"sms_id"` //阿里云短信配置id agent_sms 0表示不开启短信功能
	Sms_price	float32	`json:"sms_price"` //短信价格客户购买价格设置
	Sms_balance	float32	`json:"sms_balance"` //短信余额查询短信余额
	Storage_type	enum	`json:"storage_type"` //数据存储类型
	Oss_id	int	`json:"oss_id"` //阿里云存储表id
	Register_check	tinyint	`json:"register_check"` //验证码功能（1开启，2关闭）
	Switch_pc	tinyint	`json:"switch_pc"` //是否开启pc页面（1开，2关）
	Copyright	string	`json:"copyright"` //小程序版权
	Package_num	int	`json:"package_num"` //剩余的可创建代理数额
	Qiniu_id	int	`json:"qiniu_id"` //配置的七牛id
	Salt	string	`json:"salt"` //密码加盐
}
func (c *Agent_user) CheckPwd(password string)bool{
	beego.Debug("密码 admin123 加密后 是",helper.EncodePwd("admin123",c.Salt))
	return helper.CheckPwd(c.Password,c.Salt,password)
}
