import Layout from '@/layout/user'
export const constRoutes = [
  {
    path: '/user/redirect',
    component: Layout,
    hidden: true,
    children: [
      {
        path: '/user/redirect/:path(.*)',
        component: () => import('@/views/redirect/index')
      }
    ]
  },
  {
    path: '/user/login',
    component: () => import('@/views/user/login/index'),
    hidden: true
  },
  {
    path: '/user/',
    component: Layout,
    redirect: '/user/dashboard',
    children: [{
      path: 'dashboard',
      name: 'userIndex',
      component: () => import('@/views/user/dashboard/index'),
      meta: { title: '首页', icon: 'dashboard', affix: true }
    }]
  }

]
export const asyncRoutes = [
  {
    path: '/user/example',
    component: Layout,
    alwaysShow: true, // will always show the root menu
    redirect: '/user/example/table',
    name: 'Example',
    meta: {
      title: '示例页面',
      icon: 'example'

    },

    children: [
      {
        path: 'table',
        name: 'Table',
        component: () => import('@/views/manage/table/index'),
        meta: {
          title: '表格',
          icon: 'table',
          roles: ['index.index.get']
        }
      }
    ]
  },
  {
    path: '/user/form',
    component: Layout,
    // alwaysShow: true, // will always show the root menu
    children: [
      {
        path: 'index',
        name: 'Form',
        component: () => import('@/views/manage/form/index'),
        meta: { title: '表单', icon: 'form' }
      }
    ]
  },
  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]
export default constRoutes
