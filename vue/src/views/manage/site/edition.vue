<template>

<div class="app-container">
  <admin-tip>
   <li> 请输入更新的小程序版权</li>
  </admin-tip>
  <el-form CLASS="adminForm" status-icon ref="form" :model="form" label-width="160px" :rules="rules">
     <el-form-item label="小程序版权" prop="editionNum">
       <el-col :span="8">
       <el-input v-model="form.editionNum" placeholder="小程序版权" />
       </el-col>
     </el-form-item>
    <el-form-item>
      <el-button type="primary" :loading="true">提交</el-button>
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
      form: {
        editionNum: ''
      },
      rules: {
        editionNum: [
          { required: true, message: '请输入小程序版权', trigger: 'blur' }
        ]
      }
    }
  }
}
</script>

<style scoped>

</style>
