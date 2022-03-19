import service from '@/utils/request'

// @Tags PurchaseItem
// @Summary 创建PurchaseItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PurchaseItem true "创建PurchaseItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /purchaseItem/createPurchaseItem [post]
export const createPurchaseItem = (data) => {
  return service({
    url: '/purchaseItem/createPurchaseItem',
    method: 'post',
    data
  })
}

// @Tags PurchaseItem
// @Summary 删除PurchaseItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PurchaseItem true "删除PurchaseItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /purchaseItem/deletePurchaseItem [delete]
export const deletePurchaseItem = (data) => {
  return service({
    url: '/purchaseItem/deletePurchaseItem',
    method: 'delete',
    data
  })
}

// @Tags PurchaseItem
// @Summary 删除PurchaseItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PurchaseItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /purchaseItem/deletePurchaseItem [delete]
export const deletePurchaseItemByIds = (data) => {
  return service({
    url: '/purchaseItem/deletePurchaseItemByIds',
    method: 'delete',
    data
  })
}

// @Tags PurchaseItem
// @Summary 更新PurchaseItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PurchaseItem true "更新PurchaseItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /purchaseItem/updatePurchaseItem [put]
export const updatePurchaseItem = (data) => {
  return service({
    url: '/purchaseItem/updatePurchaseItem',
    method: 'put',
    data
  })
}

// @Tags PurchaseItem
// @Summary 用id查询PurchaseItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PurchaseItem true "用id查询PurchaseItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /purchaseItem/findPurchaseItem [get]
export const findPurchaseItem = (params) => {
  return service({
    url: '/purchaseItem/findPurchaseItem',
    method: 'get',
    params
  })
}

// @Tags PurchaseItem
// @Summary 分页获取PurchaseItem列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取PurchaseItem列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /purchaseItem/getPurchaseItemList [get]
export const getPurchaseItemList = (params) => {
  return service({
    url: '/purchaseItem/getPurchaseItemList',
    method: 'get',
    params
  })
}
