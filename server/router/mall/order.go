package mall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OrderRouter struct {
}

// InitorderReturnRouter 初始化 Order 路由信息
func (s *OrderRouter) InitOrderRouter(Router *gin.RouterGroup) {
	orderReturnRouter := Router.Group("order").Use(middleware.OperationRecord())
	orderReturnRouterWithoutRecord := Router.Group("order")
	var orderApi = v1.ApiGroupApp.MallApiGroup.OrderApi
	{
		orderReturnRouter.POST("createOrder", orderApi.CreateOrder)                         // 新建Order
		orderReturnRouter.DELETE("deleteOrder", orderApi.DeleteOrder)                       // 删除Order
		orderReturnRouter.DELETE("deleteOrderByIds", orderApi.DeleteOrderByIds)             // 批量删除Order
		orderReturnRouter.PUT("updateOrder", orderApi.UpdateOrder)                          // 更新Order
		orderReturnRouter.POST("createOrderReturn", orderApi.CreateOrderReturn)             // 新建OrderReturn
		orderReturnRouter.DELETE("deleteOrderReturn", orderApi.DeleteOrderReturn)           // 删除OrderReturn
		orderReturnRouter.DELETE("deleteOrderReturnByIds", orderApi.DeleteOrderReturnByIds) // 批量删除OrderReturn
		orderReturnRouter.PUT("updateOrderReturn", orderApi.UpdateOrderReturn)              // 更新OrderReturn
	}
	{
		orderReturnRouterWithoutRecord.GET("findOrder", orderApi.FindOrder)                   // 根据ID获取Order
		orderReturnRouterWithoutRecord.GET("getOrderList", orderApi.GetOrderList)             // 获取Order列表
		orderReturnRouterWithoutRecord.GET("findOrderReturn", orderApi.FindOrderReturn)       // 根据ID获取OrderReturn
		orderReturnRouterWithoutRecord.GET("getOrderReturnList", orderApi.GetOrderReturnList) // 获取OrderReturn列表
	}
}
