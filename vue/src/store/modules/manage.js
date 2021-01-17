import { login, logout, getInfo } from '@/api/manage/manage'
import { getToken, removeToken } from '@/utils/auth'
import { constantRoutes, resetRouter } from '@/router'

import { filterAsyncRoutes } from '@/store/modules/permission'
import { constRoutes, asyncRoutes } from '@/router/modules/manage'

const getDefaultState1 = () => {
  return {
    token: getToken(),
    name: '',
    avatar: '',
    roles_manage: [],
    roles_user: [],
    routes: [],
    addRoutes: []
  }
}

const state = getDefaultState1()

const mutations = {
  RESET_STATE: (state) => {
    Object.assign(state, getDefaultState1())
  },
  SET_TOKEN: (state, token) => {
    state.token = token
  },
  SET_NAME: (state, name) => {
    state.name = name
  },
  SET_AVATAR: (state, avatar) => {
    state.avatar = avatar
  },

  SET_ROLES: (state, roles) => {
    state.roles_user = roles
  },
  SET_ROUTES: (state, routes) => {
    state.addRoutes = routes
    state.routes = constantRoutes.concat(constRoutes, routes)
  }
}

const actions = {
  // user login
  login({ commit }, userInfo) {
    const { username, password } = userInfo
    return new Promise((resolve, reject) => {
      login({ username: username.trim(), password: password }).then(response => {
        // const { data } = response
        console.log('返回数据', response)
        // commit('SET_TOKEN', data.token)
        // setToken(data.token)
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },

  // get user info
  getInfo({ commit, state }) {
    return new Promise((resolve, reject) => {
      getInfo().then(response => {
        const { data } = response
        console.log('登录数据', data)
        if (!data) {
          reject('Verification failed, please Login again.')
        }

        const { userInfo, roles } = data
        console.log('name avatar', userInfo, roles)
        const { nickname, logo } = userInfo
        commit('SET_NAME', nickname)
        commit('SET_AVATAR', logo)
        commit('SET_ROLES', roles)
        resolve(data)
      }).catch(error => {
        reject(error)
      })
    })
  },
  generateRoutes({ commit }, roles) {
    return new Promise(resolve => {
      let accessedRoutes
      if (!roles || roles.length === 0) {
        accessedRoutes = asyncRoutes || []
      } else {
        accessedRoutes = filterAsyncRoutes(asyncRoutes, roles)
      }
      console.log('添加菜单', accessedRoutes)
      commit('SET_ROUTES', accessedRoutes)
      resolve(accessedRoutes)
    })
  },
  // user logout
  logout({ commit, state }) {
    return new Promise((resolve, reject) => {
      logout(state.token).then(() => {
        removeToken() // must remove  token  first
        resetRouter()
        commit('RESET_STATE')
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },

  // remove token
  resetToken({ commit }) {
    return new Promise(resolve => {
      removeToken() // must remove  token  first
      commit('RESET_STATE')
      resolve()
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}

