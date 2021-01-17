<template>
  <div class="package-auth  app-container">
    <div class="top-line">
      <span >用户级别</span>
      <el-select v-model="formData.package_id" @change="getAuthList()" size="small" placeholder="请选择">
        <el-option
          v-for="item in levels"
          :key="item.name"
          :label="item.name"
          :value="item.id">
        </el-option>

      </el-select>
      <el-button type="primary" size="small" @click="saveAuth()">保存</el-button>
    </div>
    <el-card class="box-card">
        <div class="panel panel-primary" v-for="(item) in filter_roles" :key="item.name">
          <div class="panel-header">
            <span>{{item.name}}</span>
          </div>
          <div class="panel-body">
            <ul >
              <li v-for="item1 in item.children" :key="item1.name">
                <el-checkbox label="1"  v-model="item1.checked"  >{{item1.name}}</el-checkbox>
              </li>

            </ul>
          </div>
        </div>
    </el-card>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { getLevels, getRoles, SaveRoles } from '../../../api/manage/package'
import { uniq } from '../../../utils/utils'
export default {
  name: 'PackageAuth',
  computed: {
    ...mapGetters([
      'manage_permission_routes'
    ])
  },
  data() {
    return {
      formData: {
        package_id: ''
      },
      roles: {},
      pack_roles: [],
      filter_roles: [],
      levels: []// 套餐列表
    }
  },
  async mounted() {
    await this.getLevelList()
    await this.getAuthList()
    console.log('manage_permission_routes', this.manage_permission_routes)
  },
  methods: {
    saveAuth() {
      const data = this.filter_roles
      const package_id = this.formData.package_id
      let nodes = []
      data.forEach(function(item, index) {
        if (!item.children) {
          return true
        }
        item.children.forEach(function(item1, index) {
          if (item1.checked) {
            nodes = nodes.concat(item1.nodes)
          }
        })
      })
      nodes = uniq(nodes)
      console.log(nodes)
      SaveRoles(package_id, { nodes: nodes }).then(res => {
        if (res.status === 'success') {
          this.$message.success('设置权限成功')
          // this.getLevelList()
        }
      })
    },
    async getLevelList() { // 获取套餐列表
      const { status, data } = await getLevels({ page: 1 })
      console.log(status, data)
      if (status === 'success') {
        this.levels = data.list
        console.log('levels 0', this.levels[0])
        this.formData.package_id = this.levels[0].id
      }
    },
    async getAuthList() { // 获取权限列表
      const package_id = this.formData.package_id
      console.log('package_id', package_id)
      const { status, data } = await getRoles(package_id)
      if (status !== 'success') return
      this.roles = data.menus
      this.pack_roles = data.pack_menus
      let node, is_auth, nodeName, children, has_menu
      const filter_roles = []
      this.manage_permission_routes.forEach(function(item, index) {
        // 一级菜单必定有权限
        if (!item.children) {
          return true
        }

        // const top_role = item
        // const name = top_role.meta.title
        children = []

        item.children.forEach(function(item1, index1) {
          if (item1.meta && item1.meta.roles) {
            nodeName = item1.meta.title
            is_auth = true
            has_menu = true
            for (const i in item1.meta.roles) {
              node = item1.meta.roles[i]
              if (!data.pack_menus[node]) {
                is_auth = false
              }
              if (!data.menus[node]) { // 没有菜单项
                has_menu = false
                break
              }
            }
            if (has_menu) {
              children.push({
                name: nodeName,
                checked: is_auth,
                nodes: item1.meta.roles
              })
            }
            console.log('roles', item1.meta.roles)
          }
        })
        if (children.length > 0) {
          filter_roles.push({
            name: item.meta.title,
            children: children
          })
        }
        console.log('item index', item, index)
      })
      this.filter_roles = filter_roles
      this.$message.success('刷新数据成功')
      console.log('过滤后权限列表', filter_roles)
    }
  }
}
</script>

<style lang="scss" scoped>
 .top-line{
   margin:10px 0;
   display: inline-flex;
   flex-direction: row; //横向排列
   justify-content: flex-start; //向左排列
   *{
     align-items: center;
     vertical-align:middle;
     margin-left: 5px;
   }
   span{
     line-height: 32px;
   }
 }
 .panel-primary{
   border-color: #337ab7;
 }
  .panel{
    box-sizing: border-box;
    background-color: #ffffff;
    border-radius: 4px;
    border:1px solid transparent;
    box-shadow: 0 1px 1px rgba(0,0,0,.05);
    .panel-header{
      color: #fff;
      background-color: #337ab7;
      border-color: #337ab7;

      padding: 10px 15px;
      align-items: center;
      border-bottom: 1px solid transparent;
      border-top-left-radius: 3px;
      border-top-right-radius: 3px;

    }
    .panel-body{
      padding:15px;
      ul{
        li{
          position: relative;
          display: inline-flex;
          justify-content: flex-start;
          flex-direction: row;
          padding-left: 20px;
          margin: 5px;
          font-weight: 400;
          vertical-align: middle;
          cursor: pointer;
        }
      }
    }
  }
</style>
