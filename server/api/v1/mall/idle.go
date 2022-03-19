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

type IdleItemApi struct {
}

var idleItemService = service.ServiceGroupApp.MallServiceGroup.IdleItemService

// CreateIdleItem 创建IdleItem
// @Tags IdleItem
// @Summary 创建IdleItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mall.IdleItem true "创建IdleItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /idleItem/createIdleItem [post]
func (idleItemApi *IdleItemApi) CreateIdleItem(c *gin.Context) {
	var idleItem mallReq.ReqIdleItem
	_ = c.ShouldBindJSON(&idleItem)
	if err := idleItemService.CreateIdleItem(idleItem); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIdleItem 删除IdleItem
// @Tags IdleItem
// @Summary 删除IdleItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mall.IdleItem true "删除IdleItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /idleItem/deleteIdleItem [delete]
func (idleItemApi *IdleItemApi) DeleteIdleItem(c *gin.Context) {
	var idleItem mall.IdleItem
	_ = c.ShouldBindJSON(&idleItem)
	if err := idleItemService.DeleteIdleItem(idleItem); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIdleItemByIds 批量删除IdleItem
// @Tags IdleItem
// @Summary 批量删除IdleItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除IdleItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /idleItem/deleteIdleItemByIds [delete]
func (idleItemApi *IdleItemApi) DeleteIdleItemByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := idleItemService.DeleteIdleItemByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIdleItem 更新IdleItem
// @Tags IdleItem
// @Summary 更新IdleItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mall.IdleItem true "更新IdleItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /idleItem/updateIdleItem [put]
func (idleItemApi *IdleItemApi) UpdateIdleItem(c *gin.Context) {
	var idleItem mallReq.ReqIdleItem
	_ = c.ShouldBindJSON(&idleItem)
	if err := idleItemService.UpdateIdleItem(idleItem); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIdleItem 用id查询IdleItem
// @Tags IdleItem
// @Summary 用id查询IdleItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mall.IdleItem true "用id查询IdleItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /idleItem/findIdleItem [get]
func (idleItemApi *IdleItemApi) FindIdleItem(c *gin.Context) {
	var idleItem mall.IdleItem
	_ = c.ShouldBindQuery(&idleItem)
	if err, reidleItem := idleItemService.GetIdleItem(idleItem.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reidleItem": reidleItem}, c)
	}
}

// GetIdleItemList 分页获取IdleItem列表
// @Tags IdleItem
// @Summary 分页获取IdleItem列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mallReq.IdleItemSearch true "分页获取IdleItem列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /idleItem/getIdleItemList [get]
func (idleItemApi *IdleItemApi) GetIdleItemList(c *gin.Context) {
	var pageInfo mallReq.IdleItemSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := idleItemService.GetIdleItemInfoList(pageInfo); err != nil {
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
