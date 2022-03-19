package mall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OrderHistoryRouter struct {
}

// InitOrderHistoryRouter 初始化 OrderHistory 路由信息
func (s *OrderHistoryRouter) InitOrderHistoryRouter(Router *gin.RouterGroup) {
	orderHistoryRouter := Router.Group("orderHistory").Use(middleware.OperationRecord())
	orderHistoryRouterWithoutRecord := Router.Group("orderHistory")
	var orderHistoryApi = v1.ApiGroupApp.MallApiGroup.OrderApi
	{
		orderHistoryRouter.POST("createOrderHistory", orderHistoryApi.CreateOrderHistory)             // 新建OrderHistory
		orderHistoryRouter.DELETE("deleteOrderHistory", orderHistoryApi.DeleteOrderHistory)           // 删除OrderHistory
		orderHistoryRouter.DELETE("deleteOrderHistoryByIds", orderHistoryApi.DeleteOrderHistoryByIds) // 批量删除OrderHistory
		orderHistoryRouter.PUT("updateOrderHistory", orderHistoryApi.UpdateOrderHistory)              // 更新OrderHistory
	}
	{
		orderHistoryRouterWithoutRecord.GET("findOrderHistory", orderHistoryApi.FindOrderHistory)       // 根据ID获取OrderHistory
		orderHistoryRouterWithoutRecord.GET("getOrderHistoryList", orderHistoryApi.GetOrderHistoryList) // 获取OrderHistory列表
	}
}
