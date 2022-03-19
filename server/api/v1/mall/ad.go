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

type AdApi struct {
}

var adService = service.ServiceGroupApp.MallServiceGroup.AdService

// CreateAd 创建Ad
// @Tags Ad
// @Summary 创建Ad
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mall.Ad true "创建Ad"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ad/createAd [post]
func (adApi *AdApi) CreateAd(c *gin.Context) {
	var ad mall.Ad
	_ = c.ShouldBindJSON(&ad)
	if err := adService.CreateAd(ad); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAd 删除Ad
// @Tags Ad
// @Summary 删除Ad
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mall.Ad true "删除Ad"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ad/deleteAd [delete]
func (adApi *AdApi) DeleteAd(c *gin.Context) {
	var ad mall.Ad
	_ = c.ShouldBindJSON(&ad)
	if err := adService.DeleteAd(ad); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAdByIds 批量删除Ad
// @Tags Ad
// @Summary 批量删除Ad
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Ad"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /ad/deleteAdByIds [delete]
func (adApi *AdApi) DeleteAdByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := adService.DeleteAdByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAd 更新Ad
// @Tags Ad
// @Summary 更新Ad
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mall.Ad true "更新Ad"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ad/updateAd [put]
func (adApi *AdApi) UpdateAd(c *gin.Context) {
	var ad mall.Ad
	_ = c.ShouldBindJSON(&ad)
	if err := adService.UpdateAd(ad); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAd 用id查询Ad
// @Tags Ad
// @Summary 用id查询Ad
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mall.Ad true "用id查询Ad"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ad/findAd [get]
func (adApi *AdApi) FindAd(c *gin.Context) {
	var ad mall.Ad
	_ = c.ShouldBindQuery(&ad)
	if err, read := adService.GetAd(ad.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"read": read}, c)
	}
}

// GetAdList 分页获取Ad列表
// @Tags Ad
// @Summary 分页获取Ad列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mallReq.AdSearch true "分页获取Ad列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ad/getAdList [get]
func (adApi *AdApi) GetAdList(c *gin.Context) {
	var pageInfo mallReq.AdSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := adService.GetAdInfoList(pageInfo); err != nil {
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
