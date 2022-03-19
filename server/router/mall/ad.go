package mall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AdRouter struct {
}

// InitAdRouter 初始化 Ad 路由信息
func (s *AdRouter) InitAdRouter(Router *gin.RouterGroup) {
	adRouter := Router.Group("ad").Use(middleware.OperationRecord())
	adRouterWithoutRecord := Router.Group("ad")
	var adApi = v1.ApiGroupApp.MallApiGroup.AdApi
	{
		adRouter.POST("createAd", adApi.CreateAd)             // 新建Ad
		adRouter.DELETE("deleteAd", adApi.DeleteAd)           // 删除Ad
		adRouter.DELETE("deleteAdByIds", adApi.DeleteAdByIds) // 批量删除Ad
		adRouter.PUT("updateAd", adApi.UpdateAd)              // 更新Ad
	}
	{
		adRouterWithoutRecord.GET("findAd", adApi.FindAd)       // 根据ID获取Ad
		adRouterWithoutRecord.GET("getAdList", adApi.GetAdList) // 获取Ad列表
	}
}
