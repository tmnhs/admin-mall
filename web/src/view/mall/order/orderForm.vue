<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="订单编码:">
          <el-input v-model="formData.orderSn" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="买家id:">
          <el-input v-model.number="formData.receiverId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="类型:">
          <el-select v-model="formData.type" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in orderTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="物品id:">
          <el-input v-model.number="formData.itemId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="总交易金额:">
          <el-input-number v-model="formData.totalAmount" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item label="交易地点:">
          <el-input v-model="formData.address" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="交易时间:">
          <el-input v-model.number="formData.tradeTime" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:">
          <el-select v-model="formData.status" placeholder="请选择" clearable>
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
  name: 'Order'
}
</script>

<script setup>
import {
  createOrder,
  updateOrder,
  findOrder
} from '@/api/order'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findOrder({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reorder
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    orderTypeOptions.value = await getDictFunc('orderType')
    orderStatusOptions.value = await getDictFunc('orderStatus')
}

init()
// 保存按钮
const save = async() => {
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
      }
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
