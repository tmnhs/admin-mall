package mall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PurchaseItemRouter struct {
}

// InitPurchaseItemRouter 初始化 PurchaseItem 路由信息
func (s *PurchaseItemRouter) InitPurchaseItemRouter(Router *gin.RouterGroup) {
	purchaseItemRouter := Router.Group("purchaseItem").Use(middleware.OperationRecord())
	purchaseItemRouterWithoutRecord := Router.Group("purchaseItem")
	var purchaseItemApi = v1.ApiGroupApp.MallApiGroup.PurchaseItemApi
	{
		purchaseItemRouter.POST("createPurchaseItem", purchaseItemApi.CreatePurchaseItem)             // 新建PurchaseItem
		purchaseItemRouter.DELETE("deletePurchaseItem", purchaseItemApi.DeletePurchaseItem)           // 删除PurchaseItem
		purchaseItemRouter.DELETE("deletePurchaseItemByIds", purchaseItemApi.DeletePurchaseItemByIds) // 批量删除PurchaseItem
		purchaseItemRouter.PUT("updatePurchaseItem", purchaseItemApi.UpdatePurchaseItem)              // 更新PurchaseItem
	}
	{
		purchaseItemRouterWithoutRecord.GET("findPurchaseItem", purchaseItemApi.FindPurchaseItem)       // 根据ID获取PurchaseItem
		purchaseItemRouterWithoutRecord.GET("getPurchaseItemList", purchaseItemApi.GetPurchaseItemList) // 获取PurchaseItem列表
	}
}
