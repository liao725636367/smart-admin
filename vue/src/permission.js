import router from './router'
import store from './store'
import { Message } from 'element-ui'
import NProgress from 'nprogress' // progress bar
import 'nprogress/nprogress.css' // progress bar style
import { getToken } from '@/utils/auth' // get token from cookie
import getPageTitle from '@/utils/get-page-title'

NProgress.configure({ showSpinner: false }) // NProgress Configuration

const whiteList = ['/manage/login', '/user/login'] // no redirect whitelist

router.beforeEach(async(to, from, next) => {
  // start progress bar
  NProgress.start()

  // set page title
  document.title = getPageTitle(to.meta.title)
  console.log('当前跳转路径', to.path)
  // determine whether the user has logged in
  const hasToken = getToken()
  const splitPath = to.path.split('/')
  let module_name = ''
  if (splitPath[0] === 'manage' || splitPath[1] === 'manage') {
    module_name = 'manage'
  }
  if (splitPath[0] === 'user' || splitPath[1] === 'user') {
    module_name = 'user'
  }

  console.log('当前模块', module_name)
  if (hasToken) {
    console.log('有token跳转链接', to.path)
    if (to.path === '/' + module_name + '/login') {
      // if is logged in, redirect to the home page
      next({ path: '/' + module_name + '/' })
      NProgress.done()
    } else {
      const hasGetUserInfo = store.getters[`${module_name}_name`]
      if (hasGetUserInfo || !module_name) { // 不是用户模块不需要用户信息
        console.log('"进入next')
        // 进入同一个页面就刷新
        // if (to.path === from.path && module_name) {
        //   next(`/${module_name}/redirect/${to.path}`)
        // }
        next()
      } else {
        try {
          // get user info
          const { roles } = await store.dispatch(module_name + '/getInfo')

          // generate accessible routes map based on roles
          const accessRoutes = await store.dispatch(module_name + '/generateRoutes', roles)
          console.log('tmp.children', accessRoutes)

          // dynamically add accessible routes
          router.addRoutes(accessRoutes)

          // hack method to ensure that addRoutes is complete
          // set the replace: true, so the navigation will not leave a history record
          // next({ ...to, replace: true })
          console.log('添加页面后进入路由')
          next({ ...to, replace: true })
        } catch (error) {
          // remove token and go to login page to re-login
          await store.dispatch(module_name + '/resetToken')
          Message.error(error || 'Has Error')
          next(`/${module_name}/login?redirect=${to.path}`)
          NProgress.done()
        }
      }
    }
  } else {
    /* has no token*/
    console.log('跳转链接', to.path)
    if (whiteList.indexOf(to.path) !== -1) {
      // in the free login whitelist, go directly
      next()
    } else {
      // other pages that do not have permission to access are redirected to the login page.
      next(`/${module_name}/login?redirect=${to.path}`)
      NProgress.done()
    }
  }
})

router.afterEach(() => {
  // finish progress bar
  NProgress.done()
})
