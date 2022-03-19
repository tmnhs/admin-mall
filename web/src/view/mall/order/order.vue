<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
         <el-form-item label="订单id">
          <el-input v-model="searchInfo.ID" placeholder="搜索条件" />
        </el-form-item>
         <el-form-item label="订单编码">
          <el-input v-model="searchInfo.orderSn" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="买家id">
          <el-input v-model="searchInfo.receiverId" placeholder="搜索条件" />
        </el-form-item>
          <el-form-item label="类型">
          <el-select v-model="searchInfo.type" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in orderTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
           <el-form-item label="状态">
          <el-select v-model="searchInfo.status" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in orderStatusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button size="small" type="text" @click="deleteVisible = false">取消</el-button>
                <el-button size="small" type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="订单id" prop="ID" width="100" />
        <el-table-column align="left" label="订单编码" prop="orderSn" width="240" />
        <el-table-column align="left" label="买家id" prop="receiverId" width="120" />
        
        <el-table-column align="left" label="买家名" prop="receiverName" width="120" />
        
        <el-table-column align="left" label="类型" prop="type" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.type,orderTypeOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="物品id" prop="itemId" width="120" />
        <el-table-column align="left" label="总交易金额" prop="totalAmount" width="120" />
        <el-table-column align="left" label="交易地点" prop="address" width="120" />
         <el-table-column align="left" label="交易时间" prop="tradeTime" width="200" >
            <template #default="scope">{{ formatUnixDate(scope.row.tradeTime) }}</template>
        </el-table-column>
        <el-table-column align="left" label="状态" prop="status" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.status,orderStatusOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="订单备注" prop="note" width="120" />
        <el-table-column align="left" label="是否已确认" prop="confirmStatus" width="120">
                  <template #default="scope">
                      <el-switch v-model="scope.row.confirmStatus" active-color="#13ce66" inactive-color="#ff4949"  clearable disabled></el-switch>
                  </template>
        </el-table-column>
        <el-table-column align="left" label="是否已删除" prop="deleteStatus" width="120">
                  <template #default="scope">
                      <el-switch v-model="scope.row.deleteStatus" active-color="#13ce66" inactive-color="#ff4949"  clearable disabled></el-switch>
                  </template>
        </el-table-column>
        
        <el-table-column align="left" label="创建时间" prop="createdTime" width="200" >
            <template #default="scope">{{ formatUnixDate(scope.row.createdTime) }}</template>
        </el-table-column>
        <el-table-column align="left" label="按钮组" width="120">
            <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateOrderFunc(scope.row)">变更</el-button>
            <el-button type="text" icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="订单编号:">
          <el-input v-model="formData.orderSn" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="买家id:">
          <el-input v-model.number="formData.receiverId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="类型:">
          <el-select v-model="formData.type" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in orderTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="物品id:">
          <el-input v-model.number="formData.itemId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="总交易金额:">
          <el-input-number v-model="formData.totalAmount"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="交易地点:">
          <el-input v-model="formData.address" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="交易时间:">
          <el-input v-model.number="formData.tradeTime" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:">
          <el-select v-model="formData.status" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in orderStatusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="订单备注:">
          <el-input v-model="formData.note" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否已确认:">
          <el-switch v-model="formData.confirmStatus" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="是否已删除:">
          <el-switch v-model="formData.deleteStatus" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'Order'
}
</script>

<script setup>
import {
  createOrder,
  deleteOrder,
  deleteOrderByIds,
  updateOrder,
  findOrder,
  getOrderList
} from '@/api/order'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatUnixDate,formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const orderTypeOptions = ref([])
const orderStatusOptions = ref([])
const formData = ref({
        orderSn: '',
        receiverId: 0,
        type: undefined,
        itemId: 0,
        totalAmount: 0,
        address: '',
        tradeTime: 0,
        status: undefined,
        note: '',
        confirmStatus: false,
        deleteStatus: false,
        createdTime: 0,
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
  if (searchInfo.value.confirmStatus === ""){
      searchInfo.value.confirmStatus=null
  }
  if (searchInfo.value.deleteStatus === ""){
      searchInfo.value.deleteStatus=null
  }
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
  const table = await getOrderList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    orderTypeOptions.value = await getDictFunc('orderType')
    orderStatusOptions.value = await getDictFunc('orderStatus')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteOrderFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteOrderByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateOrderFunc = async(row) => {
    const res = await findOrder({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reorder
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteOrderFunc = async (row) => {
    const res = await deleteOrder({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        orderSn: '',
        receiverId: 0,
        type: undefined,
        itemId: 0,
        totalAmount: 0,
        address: '',
        tradeTime: 0,
        status: undefined,
        note: '',
        confirmStatus: false,
        deleteStatus: false,
        createdTime: 0,
        }
}
// 弹窗确定
const enterDialog = async () => {
      let res
      switch (type.value) {
        case 'create':
          res = await createOrder(formData.value)
          break
        case 'update':
          res = await updateOrder(formData.value)
          break
        default:
          res = await createOrder(formData.value)
          break
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
        closeDialog()
        getTableData()
      }
}
</script>

<style>
</style>
