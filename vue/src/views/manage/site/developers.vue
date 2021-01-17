<template>
  <div class="app-container">
    <div slot="header" class="card-header clearfix">
      <aside>套餐级别设置</aside>
      <!--        <el-tag type="primary">套餐级别设置</el-tag>-->

    </div>
  <admin-tip operation="password">
    <li>配置后可使用扫码一键绑定小程序功能，不使用可不配置此项</li>
    <li>代理商空的情况下默认继承管理员</li>
  </admin-tip>
  <el-form class="adminForm" status-icon :rules="rules" ref="developers" :model="developers" label-width="160px">
    <el-form-item label="授权方式">
      <el-radio-group v-model="developers.empower">
        <el-radio :label="0">默认授权</el-radio>
        <el-radio :label="1">借用授权</el-radio>
      </el-radio-group>
    </el-form-item>
    <el-form-item label="AppID" prop="appid">
      <el-col :span="8">
        <el-input v-model="developers.appid" placeholder="请输入18位数的AppID" />
      </el-col>
    </el-form-item>
    <el-form-item label="AppSecret" prop="appsecret" >
      <el-col :span="8">
        <el-input v-model="developers.appsecret" placeholder="请输入32位的AppSecret" />
      </el-col>
    </el-form-item>
    <el-form-item label="token" prop="token">
      <el-col :span="8">
        <el-input v-model="developers.token" placeholder="请输入开放平台配置的TOKEN" />
      </el-col>
    </el-form-item>
    <el-form-item label="EncodingAESKey" prop="encodingaeskey">
      <el-col :span="8">
        <el-input v-model="developers.encodingaeskey" placeholder="请输入42的EncodingAESKey" />
      </el-col>
    </el-form-item>
    <el-form-item label="域名配置" prop="domain">
      <el-col :span="8">
        <el-input v-model="developers.domain" placeholder="小程序服务器域名" />
      </el-col>
    </el-form-item>
    <el-form-item label="TemplateId" >
      <el-col :span="8">
        <el-input v-model="developers.templateid" placeholder="此处为空自动读取模板库最新的模板ID" />
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
      developers: {
        empower: 0,
        appid: '',
        appsecret: '',
        token: '',
        encodingaeskey: '',
        domain: ''
      },
      rules: {
        appid: [
          { type: 'string', required: true, message: '请输入AppID' },
          { pattern: /^[0-9A-Za-z]{18}$/, message: '请输入正确的AppID' }
        ],
        appsecret: [
          { type: 'string', required: true, message: '请输入AppSecret' },
          { pattern: /^[0-9A-Za-z]{32}$/, message: '请输入正确的AppSecret' }
        ],
        token: [
          { type: 'string', required: true, message: '请输入Token' }
        ],
        encodingaeskey: [
          { type: 'string', required: true, message: '请输入EncodingAESKey' },
          { pattern: /^[0-9A-Za-z]{42}$/, message: '请输入正确的EncodingAESKey' }
        ],
        domain: [
          { type: 'string', required: true, message: '请输入域名' },
          { pattern: '/(www\\.)?[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+(:\\d+)*(\\/\\w+\\.\\w+)*$/', message: '请输入正确的域名' }
        ]
      }
    }
  }
}
</script>

<style scoped>

</style>
