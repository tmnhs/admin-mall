<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="用户id:">
          <el-input v-model.number="formData.uid" clearable placeholder="请输入"/>
        </el-form-item>
        <el-form-item label="物品名称:">
          <el-input v-model="formData.name" clearable placeholder="请输入"/>
        </el-form-item>
        <el-form-item label="分类id:">
          <el-input v-model.number="formData.categoryId" clearable placeholder="请输入"/>
        </el-form-item>
        <el-form-item label="排序权值:">
          <el-input v-model.number="formData.sort" clearable placeholder="请输入"/>
        </el-form-item>
        <el-form-item label="预估价格:">
          <el-input-number v-model="formData.minPrice" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item label="物品相关描述:">
          <el-input v-model="formData.description" clearable placeholder="请输入"/>
        </el-form-item>
        <el-form-item label="交易方式:">
          <el-select v-model="formData.tradeWay" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in tradeWayOptions" :key="key" :label="item.label" :value="item.value"/>
          </el-select>
        </el-form-item>
        <el-form-item label="是否已购得:">
          <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="是"
                     inactive-text="否" clearable></el-switch>
        </el-form-item>
        <el-form-item label="是否上架:">
          <el-switch v-model="formData.publishStatus" active-color="#13ce66" inactive-color="#ff4949" active-text="是"
                     inactive-text="否" clearable></el-switch>
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'PurchaseItem'
}
</script>

<script setup>
import {
  createPurchaseItem,
  updatePurchaseItem,
  findPurchaseItem
} from '@/api/purchase'

// 自动获取字典
import {getDictFunc} from '@/utils/format'
import {useRoute, useRouter} from "vue-router"
import {ElMessage} from 'element-plus'
import {ref} from 'vue'

const route = useRoute()
const router = useRouter()
const type = ref('')
const tradeWayOptions = ref([])
const saleStatusOptions = ref([])
const formData = ref({
  uid: 0,
  name: '',
  categoryId: 0,
  sort: 0,
  minPrice: 0,
  description: '',
  tradeWay: undefined,
  status: false,
  publishStatus: false,
  createdTime: 0,
})

// 初始化方法
const init = async () => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findPurchaseItem({ID: route.query.id})
    if (res.code === 0) {
      formData.value = res.data.repurchaseItem
      type.value = 'update'
    }
  } else {
    type.value = 'create'
  }
  tradeWayOptions.value = await getDictFunc('tradeWay')
  saleStatusOptions.value = await getDictFunc('saleStatus')
}

init()
// 保存按钮
const save = async () => {
  let res
  switch (type.value) {
    case 'create':
      res = await createPurchaseItem(formData.value)
      break
    case 'update':
      res = await updatePurchaseItem(formData.value)
      break
    default:
      res = await createPurchaseItem(formData.value)
      break
  }
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '创建/更改成功'
    })
  }
}

// 返回按钮
const back = () => {
  router.go(-1)
}

</script>

<style>
</style>
