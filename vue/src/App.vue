<template>
  <div id="app">
    <router-view  />
  </div>
</template>

<script>
// https://panjiachen.github.io/vue-element-admin/#/icon/index 图标查看
export default {
  name: 'App',
  created() {
    // const s = document.createElement('link')
    // s.rel = 'stylesheet'
    // s.href = '//at.alicdn.com/t/font_1640213_xjy84pib4vs.css'
    // document.body.appendChild(s)
    const scriptTag = document.createElement('script')
    scriptTag.type = 'text/javascript'
    scriptTag.src = '//at.alicdn.com/t/font_1640213_u0lzvpja9od.js'
    document.body.appendChild(scriptTag)
  },
  computed: {
    key() {
      return this.$route.fullPath
    }
  }
}
</script>
<style lang="scss">
@import "assets/global";
</style>
