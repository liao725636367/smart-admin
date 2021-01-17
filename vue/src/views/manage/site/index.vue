<template>
  <div class="app-container">
      <admin-tip operation="password">
        <li>系统省市县区镇基础数据都在此处设置, 请谨慎操作。</li>
        <li>"所在层级"位该区域所在层级，如省层级为1，市层级为2，区县层级为3 乡镇层级为4</li>
      </admin-tip>
    <el-form class="adminForm" status-icon ref="form" :model="form" :rules="rules" label-width="160px">

      <el-form-item label="网站名称"  prop="name">
        <el-col :span="8">
        <el-input v-model="form.name" placeholder="网站名称" />
        </el-col>
      </el-form-item>
      <el-form-item label="网站标题" prop="title">
        <el-col :span="8">
        <el-input v-model="form.title" placeholder="网站标题" />
        </el-col>
      </el-form-item>
      <el-form-item label="网站LOGO" prop="weblogo">
        <el-col :span="8">
        <el-upload
          name="weblogo"
          class="upload-demo"
          action="https://jsonplaceholder.typicode.com/posts/"
          :on-preview="handlePreview"
          :on-remove="handleRemove"
          list-type="picture">
          <el-button size="small" type="primary">点击上传</el-button>
          <div slot="tip" class="el-upload__tip">只能上传jpg/png文件，且不超过500kb</div>
        </el-upload>
        </el-col>
      </el-form-item>
      <el-form-item label="网站域名" prop="domain">
        <el-col :span="8">
        <el-input v-model="form.domain" placeholder="网站域名" />
        <div class="input-tips">(无需填写"http://")</div>
        </el-col>
      </el-form-item>
      <el-form-item label="网站关键词">
        <el-col :span="8">
        <el-input v-model="form.key_word" placeholder="网站关键词" />
        </el-col>
      </el-form-item>
      <el-form-item label="网站描述">
        <el-col :span="8">
        <el-input v-model="form.describe" placeholder="网站描述" />
        </el-col>
      </el-form-item>
      <el-form-item label="公司名称">
        <el-col :span="8">
        <el-input v-model="form.describe" placeholder="公司名称" />
        </el-col>
      </el-form-item>
      <el-form-item label="公司地址">
        <el-col :span="8">
        <el-input v-model="form.describe" placeholder="公司地址" />
        </el-col>
      </el-form-item>
      <el-form-item label="备案号">
        <el-col :span="8">
        <el-input v-model="form.record_numbe" placeholder="备案号" />
        </el-col>
      </el-form-item>
      <el-form-item label="二维码" prop="qrcode">
        <el-col :span="8">
        <el-upload
          name="qrcode"
          class="upload-demo"
          action="https://jsonplaceholder.typicode.com/posts/"
          :on-preview="handlePreview"
          :on-remove="handleRemove"
          list-type="picture">
          <el-button size="small" type="primary">点击上传</el-button>
          <div slot="tip" class="el-upload__tip">只能上传jpg/png文件，且不超过500kb</div>
        </el-upload>
        </el-col>
      </el-form-item>
      <el-form-item label="联系电话">
        <el-col :span="8">
        <el-input v-model="form.iPhone" placeholder="联系电话" />
        </el-col>
      </el-form-item>
      <el-form-item label="客服QQ">
        <el-col :span="8">
        <el-input v-model="form.iPhone" placeholder="QQ" />
        </el-col>
      </el-form-item>
      <el-form-item label="新增用户默认级别">
        <el-select  clearable v-model="form.user_agencyuserlevel" placeholder="请选择">
          <el-option
            v-for="item in user_agencyuserlevel"
            :key="item.value"
            :label="item.label"
            :value="item.value">
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="新增用户默认天数">
        <el-col :span="8">
        <el-input v-model="form.bay_number" placeholder="新增用户默认天数" />
        </el-col>
      </el-form-item>
      <el-form-item label="注册用户默认级别">
        <el-col :span="8">
        <el-select  clearable v-model="form.user_defalt_level" placeholder="请选择">
          <el-option
            v-for="item in user_agencyuserlevel"
            :key="item.value"
            :label="item.label"
            :value="item.value">
          </el-option>
        </el-select>
        </el-col>
      </el-form-item>
      <el-form-item label="注册用户默认级别">
        <el-col :span="8">
        <el-input v-model="form.defalt_level" placeholder="注册用户默认级别" />
        </el-col>
      </el-form-item>
      <el-form-item label="注册用户是否审核">
         <el-radio-group v-model="form.user_register">
              <el-radio :label="0">否</el-radio>
              <el-radio :label="1">是</el-radio>
            </el-radio-group>
      </el-form-item>
      <el-form-item label="注册是否需要验证码">
         <el-radio-group v-model="form.ster_check">
              <el-radio :label="0">否</el-radio>
              <el-radio :label="1">是</el-radio>
          </el-radio-group>
      </el-form-item>
      <el-form-item label="是否开启pc模板">
         <el-radio-group v-model="form.switch_pc">
              <el-radio :label="0">否</el-radio>
              <el-radio :label="1">是</el-radio>
          </el-radio-group>
      </el-form-item>
      <el-form-item label="用户界面(新增用户)">
         <el-radio-group v-model="form.user_pc">
              <el-radio :label="1">皮肤1</el-radio>
              <el-radio :label="2">皮肤2</el-radio>
              <el-radio :label="2">皮肤3</el-radio>
          </el-radio-group>
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="onSubmit">提交</el-button>
        <el-button @click="onCancel">Cancel</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {
  data() {
    var checkDomain = (rule, value, callback) => {
      if (value == '') {
        callback(new Error('请输入域名'))
      }
      var reg = /(www\.)?[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+(:\d+)*(\/\w+\.\w+)*$/
      if (!reg.test(value)) {
        callback(new Error('请输入正确的域名'))
      }
      callback()
    }
    return {
      form: {
        name: 'rrrr',
        title: '',
        weblogo: '',
        domain: '',
        delivery: false,
        type: [],
        resource: '',
        desc: '',
        user_agencyuserlevel: '1',
        user_defalt_level: '1'
      },
      rules: {
        name: [
          { required: true, message: '请输入网站名称', trigger: 'blur' }
        ],
        title: [
          { required: true, message: '请输入网站标题', trigger: 'blur' }
        ],
        weblogo: [
          { required: true, message: '请上传网站LOGO', trigger: 'blur' }
        ],

        domain: [
          { required: true, message: '请上传网站LOGO', trigger: 'blur' }
        ],
        domain: [
          { required: true, validator: checkDomain, trigger: 'blur' }
        ]
      },
      user_agencyuserlevel: [
        {
          value: '0',
          label: '内部测试版'
        },
        {
          value: '1',
          label: '基础会员'
        },
        {
          value: '2',
          label: '黄金会员'
        },
        {
          value: '3',
          label: '钻石会员'
        },
        {
          value: '4',
          label: '超级贵宾'
        },
        {
          value: '5',
          label: '行业版本'
        },
        {
          value: '6',
          label: '行业版2'
        },
        {
          value: '7',
          label: '行业版3'
        }
      ]
    }
  },
  methods: {
    onSubmit() {
      this.$message('submit!')
    },
    handleRemove(file, fileList) {
      console.log(file, fileList)
    },
    handlePreview(file) {
      console.log(file)
    },
    onCancel() {
      this.$message({
        message: 'cancel!',
        type: 'warning'
      })
    }
  }
}
</script>

<style scoped>
.line{
  text-align: center;
}
.input-tips{color: #606266;font-size: 0.75rem;}
</style>
