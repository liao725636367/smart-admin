package menu
//前端权限菜单节点
var ManageMenus = []map[string]string{
	{
		"key":"index.index.get",
		"value":"IndexController.Index",
	},
	{
		"key":"index.index.post",
		"value":"IndexController.Index1",
	},
	{
		"key":"common.get_role",//获取所有权限
		"value":"CommonController.GetRole",
	},
	{
		"key":"package.levels.get",//获取行业套餐列表
		"value":"PackageController.Levels",
	},
	{
		"key":"package.levels.post",//添加行业套餐
		"value":"PackageController.AddLevel",
	},
	{
		"key":"package.levels.put",//更新行业套餐
		"value":"PackageController.UpdateLevels",
	},
	{
		"key":"package.levels.delete",//删除行业套餐
		"value":"PackageController.DelLevels",
	},
	{
		"key":"package.roles.get",//获取权限节点-当前用户已有
		"value":"AuthController.GetRoles",
	},

}

//var ManageMenus MenusMap
//var ManageMenus initial.MenusMap = initial.MenusMap{
//	initial.Map{
//		"cn_name": "首页",
//		"en_name": "welcome",
//		"icon":     "icon icon-edit",
//		"sort":    1, //这里必须是数字才能排序
//		"level":2,//一级还是二级菜单
//		"sub": []map[string]string{
//			{
//				"name":     "首页",
//				"path": "IndexController.Index",
//				"icon":     "icon icon-index",
//			},
//		},
//	},
//	initial.Map{
//		"cn_name": "站点管理",
//		"en_name": "system",
//		"icon":     "icon icon-edit",
//		"sort":    2, //这里必须是数字才能排序
//		"level":	1,//一级还是二级菜单
//		"sub": []map[string]string{
//			{
//				"name":     "站点设置",
//				"path": "IndexController.Index",
//
//			},
//		},
//	},
//}


func init() {
	//sort.Sort(ManageMenus)

}
