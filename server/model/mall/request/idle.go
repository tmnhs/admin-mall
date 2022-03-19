package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mall"
)

type IdleItemSearch struct {
	mall.IdleItem
	request.PageInfo
	mall.ImageArray
}
type ReqIdleItem struct {
	mall.IdleItem
	mall.ImageArray
}
