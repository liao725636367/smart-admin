<template>
   <div class="app-container">
      <admin-tip operation="password">
        <li>小程序接口域名设置仅可修改一次，确定后请勿来回切换</li>
      </admin-tip>
       <el-form class="adminForm" ref="wechatDomain" :model="wechatDomain" label-width="160px">
         <el-form-item label="小程序接口域名设置">
            <el-radio-group v-model="wechatDomain.choice">
                 <el-radio :label="0">用OEM域名（代理商需要配置https） </el-radio>
                 <el-radio :label="1">用平台主域名（代理商无需配置https） </el-radio>
             </el-radio-group>
         </el-form-item>
         <el-form-item>
             <el-button type="primary">提交</el-button>
               <el-button type="warning">重置</el-button>
               <el-button type="danger">返回</el-button>
           </el-form-item>
       </el-form>
   </div>
</template>

<script>
export default {
  data() {
    return {
      wechatDomain: {
        choice: ''
      }
    }
  }
}
</script>

<style scoped>

</style>
