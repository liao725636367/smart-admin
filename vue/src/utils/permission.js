import store from '@/store'

/**
 * @param {Array} value
 * @returns {Boolean}
 * @example see @/views/permission/directive.vue
 */
export default function checkPermission(value) {
  console.log('value', value)
  if (value && value instanceof Array && value.length > 0) {
    const roles = store.getters && store.getters.roles

    const hasPermission = value.every(role => {
      return roles[role] !== undefined
      // return roles.includes(role)
    })

    if (!hasPermission) {
      return false
    }
    return true
  } else {
    console.error(`need roles! Like v-permission="['admin','editor']"`)
    return false
  }
}
