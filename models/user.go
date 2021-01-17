package models

import (
)

type User struct {
	Id	int	`gorm:"primary_key" json:" - "` //
	Username	string	`json:"username"` //用户名
	Password	string	`json:"password"` //密码
	Email	string	`json:"email"` //邮箱
	Telephone	string	`json:"telephone"` //电话
	Qq	int	`json:"qq"` //
	Logtimes	bigint	`json:"logtimes"` //登录次数 每登一次加一
	Created_at	timestamp	`json:"created_at"` //注册时间
	Auth_id	int	`json:"auth_id"` //用户权限类型
	Remain_time	timestamp	`json:"remain_time"` //到期时间
	Province	int	`json:"province"` //省
	City	int	`json:"city"` //市
	Area	int	`json:"area"` //区
	Address	string	`json:"address"` //地址
	Aid	int	`json:"aid"` //所属代理id
	Last_login_ip	string	`json:"last_login_ip"` //最后登录ip
	User_type	tinyint	`json:"user_type"` //用户类型 1 商户 2 员工
	Pid	int	`json:"pid"` //上级用户id（员工就直接登录为上级，但是使用当前用户权限）
	Theme	string	`json:"theme"` //模板(预留模板选择)
	Remain_sms	int	`json:"remain_sms"` //剩余短信条数
	Balance	decimal	`json:"balance"` //账户余额
}

//func init(){
//	DB.RegisterModel(new(User))
//}
func (m *User) TableName() string{
	return "user"
}