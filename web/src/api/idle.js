import service from '@/utils/request'

// @Tags IdleItem
// @Summary 创建IdleItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdleItem true "创建IdleItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /idleItem/createIdleItem [post]
export const createIdleItem = (data) => {
  return service({
    url: '/idleItem/createIdleItem',
    method: 'post',
    data
  })
}

// @Tags IdleItem
// @Summary 删除IdleItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdleItem true "删除IdleItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idleItem/deleteIdleItem [delete]
export const deleteIdleItem = (data) => {
  return service({
    url: '/idleItem/deleteIdleItem',
    method: 'delete',
    data
  })
}

// @Tags IdleItem
// @Summary 删除IdleItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除IdleItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idleItem/deleteIdleItem [delete]
export const deleteIdleItemByIds = (data) => {
  return service({
    url: '/idleItem/deleteIdleItemByIds',
    method: 'delete',
    data
  })
}

// @Tags IdleItem
// @Summary 更新IdleItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IdleItem true "更新IdleItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /idleItem/updateIdleItem [put]
export const updateIdleItem = (data) => {
  return service({
    url: '/idleItem/updateIdleItem',
    method: 'put',
    data
  })
}

// @Tags IdleItem
// @Summary 用id查询IdleItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.IdleItem true "用id查询IdleItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /idleItem/findIdleItem [get]
export const findIdleItem = (params) => {
  return service({
    url: '/idleItem/findIdleItem',
    method: 'get',
    params
  })
}

// @Tags IdleItem
// @Summary 分页获取IdleItem列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取IdleItem列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /idleItem/getIdleItemList [get]
export const getIdleItemList = (params) => {
  return service({
    url: '/idleItem/getIdleItemList',
    method: 'get',
    params
  })
}
