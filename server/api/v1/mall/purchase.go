package mall

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mall"
	mallReq "github.com/flipped-aurora/gin-vue-admin/server/model/mall/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PurchaseItemApi struct {
}

var purchaseItemService = service.ServiceGroupApp.MallServiceGroup.PurchaseItemService

// CreatePurchaseItem 创建PurchaseItem
// @Tags PurchaseItem
// @Summary 创建PurchaseItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mall.PurchaseItem true "创建PurchaseItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /purchaseItem/createPurchaseItem [post]
func (purchaseItemApi *PurchaseItemApi) CreatePurchaseItem(c *gin.Context) {
	var purchaseItem mall.PurchaseItem
	_ = c.ShouldBindJSON(&purchaseItem)
	if err := purchaseItemService.CreatePurchaseItem(purchaseItem); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePurchaseItem 删除PurchaseItem
// @Tags PurchaseItem
// @Summary 删除PurchaseItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mall.PurchaseItem true "删除PurchaseItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /purchaseItem/deletePurchaseItem [delete]
func (purchaseItemApi *PurchaseItemApi) DeletePurchaseItem(c *gin.Context) {
	var purchaseItem mall.PurchaseItem
	_ = c.ShouldBindJSON(&purchaseItem)
	if err := purchaseItemService.DeletePurchaseItem(purchaseItem); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePurchaseItemByIds 批量删除PurchaseItem
// @Tags PurchaseItem
// @Summary 批量删除PurchaseItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PurchaseItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /purchaseItem/deletePurchaseItemByIds [delete]
func (purchaseItemApi *PurchaseItemApi) DeletePurchaseItemByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := purchaseItemService.DeletePurchaseItemByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePurchaseItem 更新PurchaseItem
// @Tags PurchaseItem
// @Summary 更新PurchaseItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mall.PurchaseItem true "更新PurchaseItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /purchaseItem/updatePurchaseItem [put]
func (purchaseItemApi *PurchaseItemApi) UpdatePurchaseItem(c *gin.Context) {
	var purchaseItem mall.PurchaseItem
	_ = c.ShouldBindJSON(&purchaseItem)
	fmt.Println(purchaseItem)
	if err := purchaseItemService.UpdatePurchaseItem(purchaseItem); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPurchaseItem 用id查询PurchaseItem
// @Tags PurchaseItem
// @Summary 用id查询PurchaseItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mall.PurchaseItem true "用id查询PurchaseItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /purchaseItem/findPurchaseItem [get]
func (purchaseItemApi *PurchaseItemApi) FindPurchaseItem(c *gin.Context) {
	var purchaseItem mall.PurchaseItem
	_ = c.ShouldBindQuery(&purchaseItem)
	if err, repurchaseItem := purchaseItemService.GetPurchaseItem(purchaseItem.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repurchaseItem": repurchaseItem}, c)
	}
}

// GetPurchaseItemList 分页获取PurchaseItem列表
// @Tags PurchaseItem
// @Summary 分页获取PurchaseItem列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mallReq.PurchaseItemSearch true "分页获取PurchaseItem列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /purchaseItem/getPurchaseItemList [get]
func (purchaseItemApi *PurchaseItemApi) GetPurchaseItemList(c *gin.Context) {
	var pageInfo mallReq.PurchaseItemSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := purchaseItemService.GetPurchaseItemInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		fmt.Println(list)

		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
