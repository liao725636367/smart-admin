package models
type Agent_package_auth struct {
	Id	int	`gorm:"primary_key" json:"id"` //
	Type	enum	`json:"type"` //权限类别 staff 员工 agent 代理
	Aid	int	`json:"aid"` //各项类别: 当前权限设置所属代理
	Nopages	string	`json:"nopages"` //未授权页面列表 LoginController.index方式存储,逗号隔开
	Name	string	` json:"name" valid:"Required"` //行业套餐名称
	Nums	int	`json:"nums" ` //行业套餐可创建代理数量
	Created_at	timestamp	`json:"created_at"` //创建时间
	Updated_at	timestamp	`json:"updated_at"` //更新时间
}
