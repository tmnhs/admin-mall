import service from '@/utils/request'

// @Tags Ad
// @Summary 创建Ad
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Ad true "创建Ad"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ad/createAd [post]
export const createAd = (data) => {
  return service({
    url: '/ad/createAd',
    method: 'post',
    data
  })
}

// @Tags Ad
// @Summary 删除Ad
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Ad true "删除Ad"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ad/deleteAd [delete]
export const deleteAd = (data) => {
  return service({
    url: '/ad/deleteAd',
    method: 'delete',
    data
  })
}

// @Tags Ad
// @Summary 删除Ad
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Ad"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ad/deleteAd [delete]
export const deleteAdByIds = (data) => {
  return service({
    url: '/ad/deleteAdByIds',
    method: 'delete',
    data
  })
}

// @Tags Ad
// @Summary 更新Ad
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Ad true "更新Ad"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ad/updateAd [put]
export const updateAd = (data) => {
  return service({
    url: '/ad/updateAd',
    method: 'put',
    data
  })
}

// @Tags Ad
// @Summary 用id查询Ad
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Ad true "用id查询Ad"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ad/findAd [get]
export const findAd = (params) => {
  return service({
    url: '/ad/findAd',
    method: 'get',
    params
  })
}

// @Tags Ad
// @Summary 分页获取Ad列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Ad列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ad/getAdList [get]
export const getAdList = (params) => {
  return service({
    url: '/ad/getAdList',
    method: 'get',
    params
  })
}
