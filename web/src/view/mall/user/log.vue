
<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="用户名">
          <el-input v-model="searchInfo.username" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="登录城市">
          <el-input v-model="searchInfo.city" placeholder="搜索条件" />
        </el-form-item>
           <el-form-item label="登录类型">
          <el-select v-model="searchInfo.loginType" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in loginTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
   
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column align="left" label="用户编号" prop="uid" width="120"/>
            <!-- <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template> -->
        <el-table-column align="left" label="用户名" prop="username" width="120" />
        <el-table-column align="left" label="ip地址" prop="ip" width="150" />
        <el-table-column align="left" label="登录类型" prop="loginType" width="150">
            <template #default="scope">
            {{ filterDict(scope.row.loginType,loginTypeOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="登录城市" prop="city" width="150" />
        <el-table-column align="left" label="登录时间" prop="loginTime" width="200" >
            <template #default="scope">{{ formatUnixDate(scope.row.loginTime) }}</template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
  
  </div>
</template>

<script>
export default {
  name: 'UserLoginLog'
}
</script>

<script setup>
import {
  getUserLoginLogList
} from '@/api/m_user'

// 全量引入格式化工具 请按需保留
import { getDictFunc,formatUnixDate, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const loginTypeOptions = ref([])
const formData = ref({
        username: '',
        ip: '',
        loginType: undefined,
        city: '',
        loginTime: new Date(),
        })

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getUserLoginLogList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    loginTypeOptions.value = await getDictFunc('loginType')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


</script>

<style>
</style>

