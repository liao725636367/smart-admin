const getters = {
  sidebar: state => state.app.sidebar,
  device: state => state.app.device,
  token: state => state.user.token,
  manage_avatar: state => state.manage.avatar,
  user_avatar: state => state.user.avatar,
  user_name: state => state.user.name,
  manage_name: state => state.manage.name,
  visitedViews: state => state.tagsView.visitedViews,
  cachedViews: state => state.tagsView.cachedViews,
  manage_permission_routes: state => state.manage.routes,
  user_permission_routes: state => state.user.routes
}
export default getters
