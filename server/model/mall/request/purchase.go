package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mall"
)

type PurchaseItemSearch struct {
	mall.PurchaseItem
	request.PageInfo
}
