<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="分类名:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="上级分类id:">
          <el-input v-model.number="formData.parentId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="分类级别:">
          <el-select v-model="formData.level" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in categoryLevelOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="物品总数:">
          <el-input v-model.number="formData.itemCount" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="显示状态:">
          <el-switch v-model="formData.showStatus" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="排序权值:">
          <el-input v-model.number="formData.sort" clearable placeholder="请输入" />
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
            <template v-if="formData.icon">
                     <CustomPic pic-type="file" :pic-src="formData.icon" />
              </template><br>
            <el-button size="small" type="primary">点击上传</el-button>

          </el-upload>
         </el-form-item>
        <el-form-item label="分类描述:">
          <el-input v-model="formData.description" clearable placeholder="请输入" />
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
  name: 'Category'
}
</script>

<script setup>
import {
  createCategory,
  updateCategory,
  findCategory
} from '@/api/category'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
import { useUserStore } from '@/pinia/modules/user'
const userStore = useUserStore()
const path = ref(import.meta.env.VITE_BASE_API)
const checkFile = (file) => {
    return true
}

const uploadSuccess = (res) => {
  if (res.code === 0) {
    formData.value.icon=res.data.file.url
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
const route = useRoute()
const router = useRouter()
const type = ref('')
const categoryLevelOptions = ref([])
const formData = ref({
        name: '',
        parentId: 0,
        level: undefined,
        itemCount: 0,
        showStatus: false,
        sort: 0,
        icon: '',
        description: '',
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCategory({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recategory
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    categoryLevelOptions.value = await getDictFunc('categoryLevel')
}

init()
// 保存按钮
const save = async() => {
      let res
      switch (type.value) {
        case 'create':
          res = await createCategory(formData.value)
          break
        case 'update':
          res = await updateCategory(formData.value)
          break
        default:
          res = await createCategory(formData.value)
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
