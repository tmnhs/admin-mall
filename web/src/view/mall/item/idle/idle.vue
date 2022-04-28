<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="用户id">
          <el-input v-model="searchInfo.uid" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="物品名称">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="分类id">
          <el-input v-model="searchInfo.categoryId" placeholder="搜索条件" />
        </el-form-item>
           <el-form-item label="交易方式">
          <el-select v-model="searchInfo.tradeWay" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in tradeWayOptions" :key="key" :label="item.label" :value="item.value" />
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
        <el-table-column align="left" label="物品id" prop="ID" width="120" />
        <el-table-column align="left" label="用户id" prop="uid" width="120" />
        <el-table-column align="left" label="用户名" prop="username" width="120" />
        <el-table-column align="left" label="物品名称" prop="name" width="120" />
        <el-table-column align="left" label="分类id" prop="categoryId" width="120" />
        <el-table-column align="left" label="分类名" prop="categoryName" width="120" />
        <el-table-column align="left" label="品牌名" prop="brand" width="120" />
          <el-table-column align="left" label="图片"  width="200" >
           <template #default="scope">
                <CustomPic  v-for="(item,key) in scope.row.albumImages" :key="key" :label="item.label"  pic-type="file" :pic-src="item" style="display: inline-block;" />
          </template>
        </el-table-column>
        <el-table-column align="left" label="购买渠道" prop="purchaseChannel" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.purchaseChannel,purchaseChannelOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="几成新" prop="quality" width="120" />
        <el-table-column align="left" label="剩余保质期" prop="shelfLife" width="120" />
        <el-table-column align="left" label="排序权值" prop="sort" width="120" />
        <el-table-column align="left" label="原价" prop="originPrice" width="120" />
        <el-table-column align="left" label="价格" prop="price" width="120" />
        <el-table-column align="left" label="物品相关描述" prop="description" width="120" />
        <el-table-column align="left" label="数量" prop="stock" width="120" />
        <el-table-column align="left" label="交易方式" prop="tradeWay" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.tradeWay,tradeWayOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="是否已销售" prop="saleStatus" width="120">
            <template #default="scope">
                      <el-switch v-model="scope.row.saleStatus" active-color="#13ce66" inactive-color="#ff4949"  clearable disabled></el-switch>
           </template>
        </el-table-column>
        <el-table-column align="left" label="是否上架" prop="publishStatus" width="120">
            <template #default="scope">
                      <el-switch v-model="scope.row.publishStatus" active-color="#13ce66" inactive-color="#ff4949"  clearable disabled></el-switch>
           </template>
        </el-table-column>
         <el-table-column align="left" label="创建时间" prop="createdTime" width="200" >
            <template #default="scope">{{ formatUnixDate(scope.row.createdTime) }}</template>
        </el-table-column>
        <el-table-column align="left" label="按钮组" width="120">
            <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateIdleItemFunc(scope.row)">变更</el-button>
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
        <!-- <el-form-item label="用户id:">
          <el-input v-model.number="formData.uid" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="物品名称:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="分类id:">
          <el-input v-model.number="formData.categoryId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="品牌名:">
          <el-input v-model="formData.brand" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="图标"  required>
            <el-upload
            :action="`${path}/fileUploadAndDownload/upload`"
            :before-upload="checkFile"
            :headers="{ 'x-token': userStore.token }"
            :on-error="uploadError"
            :on-success="uploadSuccess"
            :show-file-list="false"
            class="upload-btn"
          >
            <template v-if="formData.albumImages">
                       <CustomPic  v-for="(item,key) in formData.albumImages" :key="key"  pic-type="file" :pic-src="item" />
              </template><br>
            <el-button size="small" type="primary">点击上传</el-button>

          </el-upload>
         </el-form-item>

        <el-form-item label="购买渠道:">
          <el-select v-model="formData.purchaseChannel" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in purchaseChannelOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="几成新:">
          <el-input v-model.number="formData.quality" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="剩余保质期:">
          <el-input v-model="formData.shelfLife" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="排序权值:">
          <el-input v-model.number="formData.sort" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="原价:">
          <el-input-number v-model="formData.originPrice"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="价格:">
          <el-input-number v-model="formData.price"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="物品相关描述:">
          <el-input v-model="formData.description" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="数量:">
          <el-input v-model.number="formData.stock" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="交易方式:">
          <el-select v-model="formData.tradeWay" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in tradeWayOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item> -->
        <el-form-item label="是否已销售:">
          <el-switch v-model="formData.saleStatus" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="是否上架:">
          <el-switch v-model="formData.publishStatus" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
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
  name: 'IdleItem'
}
</script>

<script setup>
import {
  createIdleItem,
  deleteIdleItem,
  deleteIdleItemByIds,
  updateIdleItem,
  findIdleItem,
  getIdleItemList
} from '@/api/idle'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatUnixDate,formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import CustomPic from '@/components/customPic/index.vue'
import { useUserStore } from '@/pinia/modules/user'
const userStore = useUserStore()
const path = ref(import.meta.env.VITE_BASE_API)
const Images = ref([])
const checkFile = (file) => {
    return true
}

const uploadSuccess = (res) => {
  if (res.code === 0) {
    formData.value.albumImages.push(res.data.file.url)
    ElMessage({
      type: 'success',
      message: '上传成功'
    })
  } else {
    ElMessage({
      type: 'warning',
      message: res.msg
    })
  }
}
const uploadError = () => {
  ElMessage({
    type: 'error',
    message: '上传失败'
  })
}
// 自动化生成的字典（可能为空）以及字段
const tradeWayOptions = ref([])
const saleStatusOptions = ref([])
const purchaseChannelOptions = ref([])
const formData = ref({
        uid: 0,
        name: '',
        categoryId: 0,
        brand: '',
        albumImages:[],
        purchaseChannel: undefined,
        quality: 0,
        shelfLife: '',
        sort: 0,
        originPrice: 0,
        price: 0,
        description: '',
        stock: 0,
        tradeWay: undefined,
        saleStatus: false,
        publishStatus: false,
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
  if (searchInfo.value.saleStatus === ""){
      searchInfo.value.saleStatus=null
  }
  if (searchInfo.value.publishStatus === ""){
      searchInfo.value.publishStatus=null
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
  const table = await getIdleItemList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    console.log(tableData.value)
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    tradeWayOptions.value = await getDictFunc('tradeWay')
    saleStatusOptions.value = await getDictFunc('saleStatus')
    purchaseChannelOptions.value = await getDictFunc('purchaseChannel')
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
            deleteIdleItemFunc(row)
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
      const res = await deleteIdleItemByIds({ ids })
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
const updateIdleItemFunc = async(row) => {
    const res = await findIdleItem({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reidleItem
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteIdleItemFunc = async (row) => {
    const res = await deleteIdleItem({ ID: row.ID })
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
        uid: 0,
        name: '',
        categoryId: 0,
        brand: '',
        purchaseChannel: undefined,
        quality: 0,
        shelfLife: '',
        sort: 0,
        originPrice: 0,
        price: 0,
        description: '',
        stock: 0,
        tradeWay: undefined,
        saleStatus: false,
        publishStatus: false,
        }
}
// 弹窗确定
const enterDialog = async () => {
      let res
      switch (type.value) {
        case 'create':
          res = await createIdleItem(formData.value)
          break
        case 'update':
          res = await updateIdleItem(formData.value)
          break
        default:
          res = await createIdleItem(formData.value)
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
