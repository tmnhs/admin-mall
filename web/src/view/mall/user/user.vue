<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
          <el-form-item label="编号">
          <el-input v-model="searchInfo.ID" placeholder="请输入用户编号" />
        </el-form-item>
        <el-form-item label="用户名">
          <el-input v-model="searchInfo.username" placeholder="请输入用户名" />
        </el-form-item>
        <el-form-item label="性别">
          <el-select v-model="searchInfo.gender" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in genderOptions" :key="key" :label="item.label" :value="item.value" />
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
        <!-- <el-table-column type="selection" width="55" /> -->
        <el-table-column align="left" label="编号"  width="100">
            <template #default="scope">{{ (scope.row.ID) }}</template>
        </el-table-column>
        <el-table-column align="left" label="用户名" prop="username" width="120" />
        <el-table-column align="left" label="性别" prop="gender" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.gender,genderOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="地址" prop="address.region" width="120">
            <template #default="scope" >
              <template v-if="scope.row.address.region==4">
              {{ scope.row.address.other }}
              </template>
               <template  v-else >
                 {{ filterDict(scope.row.address.region,regionOptions) }}{{ scope.row.address.unit }}
            </template>
            </template>
          
        </el-table-column>
        <el-table-column align="left" label="是否在线" prop="isOnline" width="120" >
             <template #default="scope">
            {{ filterDict(scope.row.isOnline,onlineOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="交易完成次数" prop="orderCount" width="120" />
        <el-table-column align="left" label="好友数量" prop="followCount" width="120" />
        <el-table-column align="left" label="登录次数" prop="loginCount" width="120" />
        <el-table-column align="left" label="最近登录时间" prop="lastLoginTime" width="180" >
            <template #default="scope">{{ formatDate(scope.row.lastLoginTime) }}</template>
        </el-table-column>
        <el-table-column align="left" label="状态" prop="status" width="120">
             <template #default="scope">
            {{ filterDict(scope.row.status,userStatusOptions) }}
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateUserStatusFunc(scope.row)">变更</el-button>
            </template>
            
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
  name: 'User'
}
</script>

<script setup>
import {
  updateUserStatus,
  getUserList
} from '@/api/m_user'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const genderOptions = ref([])
const regionOptions = ref([])
const onlineOptions= ref([])
const userStatusOptions=ref([])


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
  const table = await getUserList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    genderOptions.value = await getDictFunc('gender')
    regionOptions.value = await getDictFunc('region')
    onlineOptions.value=await getDictFunc('Online')
    userStatusOptions.value=await getDictFunc('userStatus')
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 更新用户状态
const updateUserStatusFunc = async(row) => {
    const res = await updateUserStatus({ ID: row.ID, status:row.status==0?1:0})
     ElMessageBox.confirm('确定要更改吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
                if (res.code === 0) {
                  ElMessage({
                    type: 'success',
                    message: '更改成功'
                  })
                  getTableData()
                }
        })

}




</script>

<style>
</style>
