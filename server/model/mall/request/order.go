package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mall"
)

type OrderSearch struct {
	mall.Order
	request.PageInfo
}
type OrderReturnSearch struct {
	mall.OrderReturn
	request.PageInfo
}
type OrderHistorySearch struct {
	mall.OrderHistory
	request.PageInfo
}
