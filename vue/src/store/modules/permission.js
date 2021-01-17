
/**
 * Use meta.role to determine if the current user has permission
 * @param roles
 * @param route
 */
function hasPermission(roles, route) {
  if (route.meta && route.meta.roles instanceof Array && route.meta.roles.length > 0) {
    const bool = route.meta.roles.every(role => roles[role] !== undefined)
    console.log('权限匹配结果', route.meta.roles, bool)
    return bool
    // return route.meta.roles.every(role => roles[role] !== undefined)// 每个节点都符合才获取权限
    // return roles.every(role => route.meta.roles.includes(role))
  } else {
    return true
  }
}

/**
 * Filter asynchronous routing tables by recursion
 * @param routes asyncRoutes
 * @param roles
 */
export function filterAsyncRoutes(routes, roles) {
  const res = []

  routes.forEach(route => {
    const tmp = { ...route }

    if (hasPermission(roles, tmp)) {
      if (tmp.children) {
        console.log('这里没有报错', tmp.children, roles)
        tmp.children = filterAsyncRoutes(tmp.children, roles)
      }
      if (typeof tmp.component === 'function' || (typeof (tmp.children) === 'object' && tmp.children.length > 0)) { // 有子菜单才显示父菜单
        console.log('插入组件', tmp, typeof tmp.component)
        res.push(tmp)
      }
    } else {
      console.log('没有权限菜单', tmp)
    }
  })

  return res
}
export default filterAsyncRoutes
