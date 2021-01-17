import Layout from '@/layout/manage'
export const constRoutes = [
  {
    path: '/manage/redirect',
    component: Layout,
    hidden: true,
    children: [
      {
        path: '/manage/redirect/:path(.*)',
        component: () => import('@/views/redirect/index')
      }
    ]
  },
  {
    path: '/manage/login',
    component: () => import('@/views/manage/login/index'),
    hidden: true
  },
  {
    path: '/manage/',
    component: Layout,
    redirect: '/manage/dashboard',
    children: [{
      path: 'dashboard',
      name: 'manageIndex',
      component: () => import('@/views/manage/dashboard/index'),
      meta: { title: '首页', icon: 'dashboard', affix: true }
    }]
  }

]
export const asyncRoutes = [
  {
    path: 'manage/site',
    component: Layout,
    alwaysShow: true, // will always show the root menu
    meta: {
      title: '站点管理',
      icon: 'example'

    },
    redirect: '/manage/site/index',
    children: [
      {
        path: '/manage/site/index',
        name: 'Site',
        component: () => import('@/views/manage/site/index'),
        meta: { title: '站点设置' }
      },
      {
        path: '/manage/site/wechat-domain',
        name: 'SiteWechat',
        component: () => import('@/views/manage/site/wechat-domain'),
        meta: { title: '小程序接口域名' }
      },
      {
        path: '/manage/site/edition',
        name: 'WechatEdition',
        component: () => import('@/views/manage/site/edition'),
        meta: { title: '小程序版权设置' }
      },
      {
        path: '/manage/site/developers',
        name: 'WechatDevelopers',
        component: () => import('@/views/manage/site/developers'),
        meta: { title: '微信开放平台设置' }
      },
      {
        path: '/manage/site/baidu-developers',
        name: 'BaiduDevelopers',
        component: () => import('@/views/manage/site/baidu-developers'),
        meta: { title: '百度第三平台设置' }
      }
    ]
  },
  // 套餐管理
  {
    path: '/manage/package/example',
    component: Layout,
    alwaysShow: true, // will always show the root menu
    redirect: 'level',
    name: 'package',
    meta: {
      title: '套餐管理',
      icon: 'component'
    },
    children: [
      {
        path: 'level',
        name: 'level',
        component: () => import('@/views/manage/package/levels'),
        meta: {
          title: '套餐级别设置',
          // icon: 'dengji',
          roles: ['package.levels.get', 'package.levels.post', 'package.levels.put', 'package.levels.delete']
        }
      },
      {
        path: 'level_price',
        name: 'level_price',
        component: () => import('@/views/manage/package/level_price'),
        meta: {
          title: '套餐数量设置',
          // icon: 'dengji',
          roles: ['package.levels.get', 'package.levels.put', 'package.levels.delete']
        }
      },
      {
        path: 'package_auth',
        name: 'package_auth',
        component: () => import('@/views/manage/package/package_auth'),
        meta: {
          title: '用户权限设置',
          // icon: 'dengji',
          roles: ['package.roles.get']
        }
      }
    ]
  },
  // 套餐管理
  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]
export default constRoutes
