import request from '@/utils/request'

// 获取套餐列表
export function getLevels(data) {
  return request({
    url: `/manage/package/levels`,
    method: 'GET',
    data
  })
}
// 添加套餐
export function addLevel(data) {
  return request({
    url: `/manage/package/levels`,
    method: 'POST',
    data
  })
}
// 更新多个套餐
export function updateLevels(data) {
  return request({
    url: `/manage/package/levels`,
    method: 'PUT',
    data
  })
}
// 删除套餐
export function delLevel(id) {
  return request({
    url: `/manage/package/levels/${id}`,
    method: 'DELETE'
  })
}
// 获取用户可设置权限列表
export function getRoles(package_id) {
  return request({
    url: `/manage/package/roles/${package_id}`,
    method: 'GET'
  })
}
// 设置套餐权限列表
export function SaveRoles(package_id, data) {
  return request({
    url: `/manage/package/roles/${package_id}`,
    method: 'PUT',
    data
  })
}

