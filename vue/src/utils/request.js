import axios from 'axios'
import { Message } from 'element-ui'
// import store from '@/store'
// import router from '@/router'
import { getToken, setToken, removeToken } from '@/utils/auth'

// create an axios instance
const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_API, // url = base url + request url
  // withCredentials: true, // send cookies when cross-domain requests
  timeout: 5000, // request timeout
  headers: {
    'Content-Type': 'application/json; charset=utf-8'
  }
})

// request interceptor
service.interceptors.request.use(
  config => {
    // do something before request is sent
    const session_id = getToken()

    if (session_id) {
      console.log('session_id是', session_id)
      // let each request carry token
      // ['X-Token'] is a custom headers key
      // please modify it according to the actual situation
      config.headers['session_id'] = session_id
    }
    return config
  },
  error => {
    // do something with request error
    console.log(error) // for debug
    return Promise.reject(error)
  }
)

// response interceptor
service.interceptors.response.use(
  /**
   * If you want to get http information such as headers or status
   * Please return  response => response
  */

  /**
   * Determine the request status by custom code
   * Here is just an example
   * You can also judge the status by HTTP Status Code
   */
  response => {
    const res = response.data
    const header = response.headers

    if (header['session_id']) {
      setToken(header['session_id'])
    }
    console.log('响应体token', header['session_id'])
    // if the custom code is not 20000, it is judged as an error.
    if (res.code === 403 || res.status !== 'success') { // 未授权页面或者错误信息
      Message({
        message: res.msg || '请求出现错误',
        type: 'error',
        duration: 5 * 1000
      })
      return Promise.reject(new Error(res.msg || '请求出现错误'))
    } else if (res.code === 401) { // 未登录
      removeToken()
      Message({
        message: '请登录',
        type: 'error',
        duration: 5 * 1000,
        onClose: function() {
          // router.push('/login')
        }
      })
      return res
    } else {
      console.log('res is', res)
      return res
    }
  },
  error => {
    console.log('err' + error) // for debug
    Message({
      message: error.message,
      type: 'error',
      duration: 5 * 1000
    })
    return Promise.reject(error)
  }
)

export default service
