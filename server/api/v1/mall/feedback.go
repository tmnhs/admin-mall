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

type FeedbackApi struct {
}

var feedbackService = service.ServiceGroupApp.MallServiceGroup.FeedbackService

// CreateFeedback 创建Feedback
// @Tags Feedback
// @Summary 创建Feedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mall.Feedback true "创建Feedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /feedback/createFeedback [post]
func (feedbackApi *FeedbackApi) CreateFeedback(c *gin.Context) {
	var feedback mall.Feedback
	_ = c.ShouldBindJSON(&feedback)
	if err := feedbackService.CreateFeedback(feedback); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFeedback 删除Feedback
// @Tags Feedback
// @Summary 删除Feedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mall.Feedback true "删除Feedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /feedback/deleteFeedback [delete]
func (feedbackApi *FeedbackApi) DeleteFeedback(c *gin.Context) {
	var feedback mall.Feedback
	_ = c.ShouldBindJSON(&feedback)
	if err := feedbackService.DeleteFeedback(feedback); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFeedbackByIds 批量删除Feedback
// @Tags Feedback
// @Summary 批量删除Feedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Feedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /feedback/deleteFeedbackByIds [delete]
func (feedbackApi *FeedbackApi) DeleteFeedbackByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := feedbackService.DeleteFeedbackByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFeedback 更新Feedback
// @Tags Feedback
// @Summary 更新Feedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mall.Feedback true "更新Feedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /feedback/updateFeedback [put]
func (feedbackApi *FeedbackApi) UpdateFeedback(c *gin.Context) {
	var feedback mall.Feedback
	_ = c.ShouldBindJSON(&feedback)
	if err := feedbackService.UpdateFeedback(feedback); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFeedback 用id查询Feedback
// @Tags Feedback
// @Summary 用id查询Feedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mall.Feedback true "用id查询Feedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /feedback/findFeedback [get]
func (feedbackApi *FeedbackApi) FindFeedback(c *gin.Context) {
	var feedback mall.Feedback
	_ = c.ShouldBindQuery(&feedback)
	if err, refeedback := feedbackService.GetFeedback(feedback.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refeedback": refeedback}, c)
	}
}

// GetFeedbackList 分页获取Feedback列表
// @Tags Feedback
// @Summary 分页获取Feedback列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mallReq.FeedbackSearch true "分页获取Feedback列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /feedback/getFeedbackList [get]
func (feedbackApi *FeedbackApi) GetFeedbackList(c *gin.Context) {
	var pageInfo mallReq.FeedbackSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := feedbackService.GetFeedbackInfoList(pageInfo); err != nil {
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
