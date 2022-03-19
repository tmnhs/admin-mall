package mall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FeedbackRouter struct {
}

// InitFeedbackRouter 初始化 Feedback 路由信息
func (s *FeedbackRouter) InitFeedbackRouter(Router *gin.RouterGroup) {
	feedbackRouter := Router.Group("feedback").Use(middleware.OperationRecord())
	feedbackRouterWithoutRecord := Router.Group("feedback")
	var feedbackApi = v1.ApiGroupApp.MallApiGroup.FeedbackApi
	{
		feedbackRouter.POST("createFeedback", feedbackApi.CreateFeedback)             // 新建Feedback
		feedbackRouter.DELETE("deleteFeedback", feedbackApi.DeleteFeedback)           // 删除Feedback
		feedbackRouter.DELETE("deleteFeedbackByIds", feedbackApi.DeleteFeedbackByIds) // 批量删除Feedback
		feedbackRouter.PUT("updateFeedback", feedbackApi.UpdateFeedback)              // 更新Feedback
	}
	{
		feedbackRouterWithoutRecord.GET("findFeedback", feedbackApi.FindFeedback)       // 根据ID获取Feedback
		feedbackRouterWithoutRecord.GET("getFeedbackList", feedbackApi.GetFeedbackList) // 获取Feedback列表
	}
}
