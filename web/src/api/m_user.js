import service from '@/utils/request'





// @Tags User
// @Summary 更新User状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.User true "更新User状态"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /user/updateUserStatus [put]
export const updateUserStatus = (data) => {
  return service({
    url: '/user/updateUserStatus',
    method: 'put',
    data
  })
}

// @Tags User
// @Summary 用id查询User
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.User true "用id查询User"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /user/findUser [get]
export const findUser = (params) => {
  return service({
    url: '/user/findUser',
    method: 'get',
    params
  })
}

// @Tags User
// @Summary 分页获取User列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取User列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [get]
export const getUserList = (params) => {
  return service({
    url: '/user/getUserList',
    method: 'get',
    params
  })
}


// @Tags UserLoginLog
// @Summary 分页获取UserLoginLog列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取UserLoginLog列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserLoginLogList [get]
export const getUserLoginLogList = (params) => {
  return service({
    url: '/user/getUserLoginLogList',
    method: 'get',
    params
  })
}