import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/manage/user/login',
    method: 'post',
    data
  })
}

export function getInfo() {
  return request({
    url: '/manage/user/info',
    method: 'get'
  })
}

export function logout() {
  return request({
    url: '/manage/user/logout',
    method: 'post'
  })
}
