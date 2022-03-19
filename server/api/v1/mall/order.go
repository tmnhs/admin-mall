package mall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mall"
	mallReq "github.com/flipped-aurora/gin-vue-admin/server/model/mall/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type OrderApi struct {
}

var orderService = service.ServiceGroupApp.MallServiceGroup.OrderService

// CreateOrder 创建Order
// @Tags Order
// @Summary 创建Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mall.Order true "创建Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /order/createOrder [post]
func (orderApi *OrderApi) CreateOrder(c *gin.Context) {
	var order mall.Order
	_ = c.ShouldBindJSON(&order)
	if err := orderService.CreateOrder(order); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOrder 删除Order
// @Tags Order
// @Summary 删除Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mall.Order true "删除Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /order/deleteOrder [delete]
func (orderApi *OrderApi) DeleteOrder(c *gin.Context) {
	var order mall.Order
	_ = c.ShouldBindJSON(&order)
	if err := orderService.DeleteOrder(order); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOrderByIds 批量删除Order
// @Tags Order
// @Summary 批量删除Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /order/deleteOrderByIds [delete]
func (orderApi *OrderApi) DeleteOrderByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := orderService.DeleteOrderByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOrder 更新Order
// @Tags Order
// @Summary 更新Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mall.Order true "更新Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /order/updateOrder [put]
func (orderApi *OrderApi) UpdateOrder(c *gin.Context) {
	var order mall.Order
	_ = c.ShouldBindJSON(&order)
	if err := orderService.UpdateOrder(order); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOrder 用id查询Order
// @Tags Order
// @Summary 用id查询Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mall.Order true "用id查询Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /order/findOrder [get]
func (orderApi *OrderApi) FindOrder(c *gin.Context) {
	var order mall.Order
	_ = c.ShouldBindQuery(&order)
	if err, reorder := orderService.GetOrder(order.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reorder": reorder}, c)
	}
}

// GetOrderList 分页获取Order列表
// @Tags Order
// @Summary 分页获取Order列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mallReq.OrderSearch true "分页获取Order列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /order/getOrderList [get]
func (orderApi *OrderApi) GetOrderList(c *gin.Context) {
	var pageInfo mallReq.OrderSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := orderService.GetOrderInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// CreateOrderReturn 创建OrderReturn
// @Tags OrderReturn
// @Summary 创建OrderReturn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mall.OrderReturn true "创建OrderReturn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderReturn/createOrderReturn [post]
func (orderApi *OrderApi) CreateOrderReturn(c *gin.Context) {
	var orderReturn mall.OrderReturn
	_ = c.ShouldBindJSON(&orderReturn)
	if err := orderService.CreateOrderReturn(orderReturn); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOrderReturn 删除OrderReturn
// @Tags OrderReturn
// @Summary 删除OrderReturn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mall.OrderReturn true "删除OrderReturn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /orderReturn/deleteOrderReturn [delete]
func (orderApi *OrderApi) DeleteOrderReturn(c *gin.Context) {
	var orderReturn mall.OrderReturn
	_ = c.ShouldBindJSON(&orderReturn)
	if err := orderService.DeleteOrderReturn(orderReturn); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOrderReturnByIds 批量删除OrderReturn
// @Tags OrderReturn
// @Summary 批量删除OrderReturn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OrderReturn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /orderReturn/deleteOrderReturnByIds [delete]
func (orderApi *OrderApi) DeleteOrderReturnByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := orderService.DeleteOrderReturnByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOrderReturn 更新OrderReturn
// @Tags OrderReturn
// @Summary 更新OrderReturn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mall.OrderReturn true "更新OrderReturn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /orderReturn/updateOrderReturn [put]
func (orderApi *OrderApi) UpdateOrderReturn(c *gin.Context) {
	var orderReturn mall.OrderReturn
	_ = c.ShouldBindJSON(&orderReturn)
	if err := orderService.UpdateOrderReturn(orderReturn); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOrderReturn 用id查询OrderReturn
// @Tags OrderReturn
// @Summary 用id查询OrderReturn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mall.OrderReturn true "用id查询OrderReturn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /orderReturn/findOrderReturn [get]
func (orderApi *OrderApi) FindOrderReturn(c *gin.Context) {
	var orderReturn mall.OrderReturn
	_ = c.ShouldBindQuery(&orderReturn)
	if err, reorderReturn := orderService.GetOrderReturn(orderReturn.OrderId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reorderReturn": reorderReturn}, c)
	}
}

// GetOrderReturnList 分页获取OrderReturn列表
// @Tags OrderReturn
// @Summary 分页获取OrderReturn列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mallReq.OrderReturnSearch true "分页获取OrderReturn列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderReturn/getOrderReturnList [get]
func (orderApi *OrderApi) GetOrderReturnList(c *gin.Context) {
	var pageInfo mallReq.OrderReturnSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := orderService.GetOrderReturnInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// CreateOrderHistory 创建OrderHistory
// @Tags OrderHistory
// @Summary 创建OrderHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mall.OrderHistory true "创建OrderHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderHistory/createOrderHistory [post]
func (orderApi *OrderApi) CreateOrderHistory(c *gin.Context) {
	var orderHistory mall.OrderHistory
	_ = c.ShouldBindJSON(&orderHistory)
	if err := orderService.CreateOrderHistory(orderHistory); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOrderHistory 删除OrderHistory
// @Tags OrderHistory
// @Summary 删除OrderHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mall.OrderHistory true "删除OrderHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /orderHistory/deleteOrderHistory [delete]
func (orderApi *OrderApi) DeleteOrderHistory(c *gin.Context) {
	var orderHistory mall.OrderHistory
	_ = c.ShouldBindJSON(&orderHistory)
	if err := orderService.DeleteOrderHistory(orderHistory); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOrderHistoryByIds 批量删除OrderHistory
// @Tags OrderHistory
// @Summary 批量删除OrderHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OrderHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /orderHistory/deleteOrderHistoryByIds [delete]
func (orderApi *OrderApi) DeleteOrderHistoryByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := orderService.DeleteOrderHistoryByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOrderHistory 更新OrderHistory
// @Tags OrderHistory
// @Summary 更新OrderHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mall.OrderHistory true "更新OrderHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /orderHistory/updateOrderHistory [put]
func (orderApi *OrderApi) UpdateOrderHistory(c *gin.Context) {
	var orderHistory mall.OrderHistory
	_ = c.ShouldBindJSON(&orderHistory)
	if err := orderService.UpdateOrderHistory(orderHistory); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOrderHistory 用id查询OrderHistory
// @Tags OrderHistory
// @Summary 用id查询OrderHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mall.OrderHistory true "用id查询OrderHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /orderHistory/findOrderHistory [get]
func (orderApi *OrderApi) FindOrderHistory(c *gin.Context) {
	var orderHistory mall.OrderHistory
	_ = c.ShouldBindQuery(&orderHistory)
	if err, reorderHistory := orderService.GetOrderHistory(orderHistory.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reorderHistory": reorderHistory}, c)
	}
}

// GetOrderHistoryList 分页获取OrderHistory列表
// @Tags OrderHistory
// @Summary 分页获取OrderHistory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mallReq.OrderHistorySearch true "分页获取OrderHistory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderHistory/getOrderHistoryList [get]
func (orderApi *OrderApi) GetOrderHistoryList(c *gin.Context) {
	var pageInfo mallReq.OrderHistorySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := orderService.GetOrderHistoryInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
