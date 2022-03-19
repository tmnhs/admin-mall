package mall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type IdleItemRouter struct {
}

// InitIdleItemRouter 初始化 IdleItem 路由信息
func (s *IdleItemRouter) InitIdleItemRouter(Router *gin.RouterGroup) {
	idleItemRouter := Router.Group("idleItem").Use(middleware.OperationRecord())
	idleItemRouterWithoutRecord := Router.Group("idleItem")
	var idleItemApi = v1.ApiGroupApp.MallApiGroup.IdleItemApi
	{
		idleItemRouter.POST("createIdleItem", idleItemApi.CreateIdleItem)             // 新建IdleItem
		idleItemRouter.DELETE("deleteIdleItem", idleItemApi.DeleteIdleItem)           // 删除IdleItem
		idleItemRouter.DELETE("deleteIdleItemByIds", idleItemApi.DeleteIdleItemByIds) // 批量删除IdleItem
		idleItemRouter.PUT("updateIdleItem", idleItemApi.UpdateIdleItem)              // 更新IdleItem
	}
	{
		idleItemRouterWithoutRecord.GET("findIdleItem", idleItemApi.FindIdleItem)       // 根据ID获取IdleItem
		idleItemRouterWithoutRecord.GET("getIdleItemList", idleItemApi.GetIdleItemList) // 获取IdleItem列表
	}
}
