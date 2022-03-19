package mall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	MallReq "github.com/flipped-aurora/gin-vue-admin/server/model/mall/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserApi struct {
}

var userService = service.ServiceGroupApp.MallServiceGroup.UserService

// UpdateUserStatus UpdateUserStatus 更新用户状态 [0->封禁、1->正常]
// @Tags User
// @Summary 更新用户状态 [0->封禁、1->正常]
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mall.User true "更新用户状态"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新"}"
// @Router /user/updateUserStatus [put]
func (userApi *UserApi) UpdateUserStatus(c *gin.Context) {
	var param MallReq.ReqUpdateUsersStatus
	_ = c.ShouldBindJSON(&param)
	if err := userService.UpdateUserStatus(param.ID, param.Status); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功！", c)
	}
}

// DeleteUser 删除User
// @Tags User
// @Summary 删除User
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mall.User true "删除User"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /user/deleteUser [delete]

// DeleteUserByIds 批量删除User
// @Tags User
// @Summary 批量删除User
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除User"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /user/deleteUserByIds [delete]

// UpdateUser 更新User
// @Tags User
// @Summary 更新User
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mall.User true "更新User"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /user/updateUser [put]

// FindUser 用id查询User
// @Tags User
// @Summary 用id查询User
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mall.User true "用id查询User"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /user/findUser [get]

// GetUserList 分页获取User列表
// @Tags User
// @Summary 分页获取User列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mallReq.UserSearch true "分页获取User列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [get]
func (userApi *UserApi) GetUserList(c *gin.Context) {
	var pageInfo MallReq.UserSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if pageInfo.PageSize == 0 {
		pageInfo.PageSize = 10
	}
	if pageInfo.Page == 0 {
		pageInfo.Page = 1
	}
	if err, list, total := userService.GetUserInfoList(pageInfo); err != nil {
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

// GetUserLoginLogList 分页获取UserLoginLog列表
// @Tags UserLoginLog
// @Summary 分页获取UserLoginLog列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mallReq.UserLoginLogSearch true "分页获取UserLoginLog列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserLoginLogList [get]
func (userApi *UserApi) GetUserLoginLogList(c *gin.Context) {
	var pageInfo MallReq.UserLoginLogSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if pageInfo.PageSize == 0 {
		pageInfo.PageSize = 10
	}
	if pageInfo.Page == 0 {
		pageInfo.Page = 1
	}
	if err, list, total := userService.GetUserLoginLogList(pageInfo); err != nil {
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
