<template>
  <div id="package_level" class="components-container">
    <el-card class="box-card package-card">
      <div slot="header" class="card-header clearfix">
        <aside>套餐级别设置</aside>
        <!--        <el-tag type="primary">套餐级别设置</el-tag>-->

      </div>
      <el-table
        size="medium"
        border
        :data="levels"
        style="width: 100%">
        <el-table-column
          header-align="center"
          align="center"
          prop="id"
          label="级别"
          width="180">
        </el-table-column>
        <el-table-column
          label="套餐名称"
        >
          <template slot-scope="scope">
            <el-input v-model="scope.row.name" placeholder="请输入套餐名称"></el-input>
          </template>
        </el-table-column>
        <el-table-column
          header-align="center"
          align="center"
          label="操作"
          width="100"
        >
          <template slot-scope="scope">
            <el-button type="danger" icon="el-icon-delete" size="small" @click="delRow(scope.row.id)"></el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="m-t-10">

        <el-button size="medium" type="primary" icon="el-icon-s-tools" @click="updateAll()" >
          保存改动
        </el-button>
        <el-button size="medium" type="success" icon="el-icon-plus" @click="ShowAddForm()">
          新增套餐
        </el-button>
      </div>
    </el-card>

    <el-dialog @close="resetAddForm" title="添加套餐"  :visible.sync="showAddForm"  width="50%">
      <el-form ref="addFormRef" :model="addForm" :rules="addFormRules" label-width="100">
        <el-form-item label="行业名称"  prop="name" >
          <el-input v-model="addForm.name" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="行业套餐数"  >
          <el-input type="number" :disabled="true"  v-model="addForm.nums" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="showAddForm = false">取 消</el-button>
        <el-button type="primary" @click="addRow()">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { addLevel, delLevel, getLevels, updateLevels } from '../../../api/manage/package'

export default {
  name: 'Levels',
  data() {
    return {
      showAddForm: false,
      levels: [],
      page: 1,
      addForm: {
        id: 0,
        name: '',
        nums: 0
      },
      addFormRules: {
        name: [
          {
            required: true,
            message: '请输入行业名称',
            trigger: 'blur'
          },
          { min: 3, max: 10, message: '行业名称长度在 3 到 10 个字符', trigger: 'blur' }
        ]
      }
    }
  },
  mounted() {
    this.getLevelList()
  },
  methods: {
    resetAddForm() {
      this.$refs.addFormRef.resetFields()
    },
    getLevelList() {
      getLevels({ page: this.page }).then(res => {
        if (res.status === 'success') {
          this.levels = res.data.list
        }
      })
    },
    delRow(id) {
      this.$confirm('此操作将永久删除此套餐且不可恢复,是否继续？', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(res => {
        delLevel(id).then(res => {
          if (res.status === 'success') {
            this.$message.success('删除成功')
            this.getLevelList()
          }
        })
      }).catch(res => res)
    },
    addRow() {
      this.$refs.addFormRef.validate(valid => {
        if (!valid) return
        addLevel(this.addForm).then(res => {
          if (res.status === 'success') {
            this.$message.success('添加成功')
            this.showAddForm = false
            this.getLevelList()
          }
        })
      })
    },
    updateAll() {
      this.$confirm('此操作将保存当前数据,是否继续？', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(res => {
        updateLevels(this.levels).then(res => {
          if (res.status === 'success') {
            this.$message.success('更新成功')
            this.getLevelList()
          }
        })
      }).catch(res => res)
    },
    ShowAddForm() {
      this.showAddForm = true
    }
  }
}
</script>

<style lang="scss" scoped>
  .package-card {
    .card-header {
      display: flex;
      justify-content: space-between;
      align-items: center;
    }
  }
</style>
