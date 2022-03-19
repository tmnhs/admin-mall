import service from '@/utils/request'

// @Tags OrderHistory
// @Summary 创建OrderHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderHistory true "创建OrderHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderHistory/createOrderHistory [post]
export const createOrderHistory = (data) => {
  return service({
    url: '/orderHistory/createOrderHistory',
    method: 'post',
    data
  })
}

// @Tags OrderHistory
// @Summary 删除OrderHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderHistory true "删除OrderHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /orderHistory/deleteOrderHistory [delete]
export const deleteOrderHistory = (data) => {
  return service({
    url: '/orderHistory/deleteOrderHistory',
    method: 'delete',
    data
  })
}

// @Tags OrderHistory
// @Summary 删除OrderHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OrderHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /orderHistory/deleteOrderHistory [delete]
export const deleteOrderHistoryByIds = (data) => {
  return service({
    url: '/orderHistory/deleteOrderHistoryByIds',
    method: 'delete',
    data
  })
}

// @Tags OrderHistory
// @Summary 更新OrderHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderHistory true "更新OrderHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /orderHistory/updateOrderHistory [put]
export const updateOrderHistory = (data) => {
  return service({
    url: '/orderHistory/updateOrderHistory',
    method: 'put',
    data
  })
}

// @Tags OrderHistory
// @Summary 用id查询OrderHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OrderHistory true "用id查询OrderHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /orderHistory/findOrderHistory [get]
export const findOrderHistory = (params) => {
  return service({
    url: '/orderHistory/findOrderHistory',
    method: 'get',
    params
  })
}

// @Tags OrderHistory
// @Summary 分页获取OrderHistory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取OrderHistory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderHistory/getOrderHistoryList [get]
export const getOrderHistoryList = (params) => {
  return service({
    url: '/orderHistory/getOrderHistoryList',
    method: 'get',
    params
  })
}
