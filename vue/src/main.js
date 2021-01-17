import Vue from 'vue'

import 'normalize.css/normalize.css' // A modern alternative to CSS resets

import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
// import locale from 'element-ui/lib/locale/lang/en' // lang i18n

import '@/styles/index.scss' // global css

import App from './App'
import store from './store'
import router from './router'

import '@/icons' // icon
import '@/permission' // permission control

import AdminTip from './components/AdminTip/index.vue' // admin-tip
Vue.component('AdminTip', AdminTip)

import { TweenMax, Power0, Power1, Power2, Power3, Power4, Bounce } from 'gsap'
Object.defineProperty(Vue.prototype, '$TweenMax', { value: TweenMax })
Object.defineProperty(Vue.prototype, '$Power0', { value: Power0 })
Object.defineProperty(Vue.prototype, '$Power1', { value: Power1 })
Object.defineProperty(Vue.prototype, '$Power2', { value: Power2 })
Object.defineProperty(Vue.prototype, '$Power3', { value: Power3 })
Object.defineProperty(Vue.prototype, '$Power4', { value: Power4 })
Object.defineProperty(Vue.prototype, '$Bounce', { value: Bounce })

/**
 * If you don't want to use mock-server
 * you want to use MockJs for mock api
 * you can execute: mockXHR()
 *
 * Currently MockJs will be used in the production environment,
 * please remove it before going online ! ! !
 */
// if (process.env.NODE_ENV === 'production') {
//   const { mockXHR } = require('../mock')
//   mockXHR()
// }

// set ElementUI lang to EN
// Vue.use(ElementUI, { locale })
// 如果想要中文版 element-ui，按如下方式声明
ElementUI.Dialog.props.closeOnClickModal.default = false // 修改 el-dialog 默认点击遮照为不关闭
Vue.use(ElementUI)

Vue.config.productionTip = false

new Vue({
  el: '#app',
  router,
  store,
  render: h => h(App)
})
