const state = {
  visitedViews: {
    user: [],
    manage: []
  },
  cachedViews: {
    user: [],
    manage: []
  }
}

const mutations = {
  ADD_VISITED_VIEW: (state, arg) => {
    const view = arg.view
    const type = arg.type
    console.log('添加标签参数', state, view, type, state.visitedViews[type])
    if (state.visitedViews[type].some(v => v.path === view.path)) return
    state.visitedViews[type].push(
      Object.assign({}, view, {
        title: view.meta.title || 'no-name'
      })
    )
  },
  ADD_CACHED_VIEW: (state, arg) => {
    const view = arg.view
    const type = arg.type
    if (state.cachedViews[type].includes(view.name)) return
    if (!view.meta.noCache) {
      state.cachedViews[type].push(view.name)
    }
  },

  DEL_VISITED_VIEW: (state, arg) => {
    const view = arg.view
    const type = arg.type
    for (const [i, v] of state.visitedViews[type].entries()) {
      if (v.path === view.path) {
        state.visitedViews[type].splice(i, 1)
        break
      }
    }
  },
  DEL_CACHED_VIEW: (state, arg) => {
    const view = arg.view
    const type = arg.type
    const index = state.cachedViews[type].indexOf(view.name)
    index > -1 && state.cachedViews[type].splice(index, 1)
  },

  DEL_OTHERS_VISITED_VIEWS: (state, arg) => {
    const view = arg.view
    const type = arg.type
    state.visitedViews[type] = state.visitedViews[type].filter(v => {
      return v.meta.affix || v.path === view.path
    })
  },
  DEL_OTHERS_CACHED_VIEWS: (state, arg) => {
    const view = arg.view
    const type = arg.type
    const index = state.cachedViews[type].indexOf(view.name)
    if (index > -1) {
      state.cachedViews = state.cachedViews[type].slice(index, index + 1)
    } else {
      // if index = -1, there is no cached tags
      state.cachedViews[type] = []
    }
  },

  DEL_ALL_VISITED_VIEWS: (state, type) => {
    // keep affix tags
    const affixTags = state.visitedViews[type].filter(tag => tag.meta.affix)
    state.visitedViews[type] = affixTags
  },
  DEL_ALL_CACHED_VIEWS: (state, type) => {
    state.cachedViews[type] = []
  },

  UPDATE_VISITED_VIEW: (state, arg) => {
    const view = arg.view
    const type = arg.type
    for (let v of state.visitedViews[type]) {
      if (v.path === view.path) {
        v = Object.assign(v, view)
        break
      }
    }
  }
}

const actions = {
  addView({ dispatch }, arg) {
    console.log(arguments)
    dispatch('addVisitedView', arg)
    dispatch('addCachedView', arg)
  },
  addVisitedView({ commit }, arg) {
    commit('ADD_VISITED_VIEW', arg)
  },
  addCachedView({ commit }, arg) {
    commit('ADD_CACHED_VIEW', arg)
  },

  delView({ dispatch, state }, arg) {
    return new Promise(resolve => {
      dispatch('delVisitedView', arg)
      dispatch('delCachedView', arg)
      resolve({
        visitedViews: [...state.visitedViews[arg.type]],
        cachedViews: [...state.cachedViews[arg.type]]
      })
    })
  },
  delVisitedView({ commit, state }, arg) {
    return new Promise(resolve => {
      commit('DEL_VISITED_VIEW', arg)
      resolve([...state.visitedViews[arg.type]])
    })
  },
  delCachedView({ commit, state }, arg) {
    return new Promise(resolve => {
      commit('DEL_CACHED_VIEW', arg)
      resolve([...state.cachedViews[arg.type]])
    })
  },

  delOthersViews({ dispatch, state }, arg) {
    return new Promise(resolve => {
      dispatch('delOthersVisitedViews', arg)
      dispatch('delOthersCachedViews', arg)
      resolve({
        visitedViews: [...state.visitedViews[arg.type]],
        cachedViews: [...state.cachedViews[arg.type]]
      })
    })
  },
  delOthersVisitedViews({ commit, state }, arg) {
    return new Promise(resolve => {
      commit('DEL_OTHERS_VISITED_VIEWS', arg)
      resolve([...state.visitedViews[arg.type]])
    })
  },
  delOthersCachedViews({ commit, state }, arg) {
    return new Promise(resolve => {
      commit('DEL_OTHERS_CACHED_VIEWS', arg)
      resolve([...state.cachedViews[arg.type]])
    })
  },

  delAllViews({ dispatch, state }, type) {
    return new Promise(resolve => {
      dispatch('delAllVisitedViews', type)
      dispatch('delAllCachedViews', type)
      resolve({
        visitedViews: [...state.visitedViews[type]],
        cachedViews: [...state.cachedViews[type]]
      })
    })
  },
  delAllVisitedViews({ commit, state }, type) {
    return new Promise(resolve => {
      commit('DEL_ALL_VISITED_VIEWS', type)
      resolve([...state.visitedViews[type]])
    })
  },
  delAllCachedViews({ commit, state }, type) {
    return new Promise(resolve => {
      commit('DEL_ALL_CACHED_VIEWS', type)
      resolve([...state.cachedViews[type]])
    })
  },

  updateVisitedView({ commit }, arg) {
    commit('UPDATE_VISITED_VIEW', arg)
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
