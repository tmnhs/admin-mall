package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mall"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mall/response"
)

type UserSearch struct {
	response.RspUser
	request.PageInfo
}
type ReqUpdateUsersStatus struct {
	ID     int64 `json:"ID" form:"ID"`
	Status int   `json:"status" form:"status"`
}
type UserLoginLogSearch struct {
	mall.UserLoginLog
	request.PageInfo
}
