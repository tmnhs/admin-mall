package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/mall"

type IdleItemGorm struct {
	mall.IdleItem
	mall.ImageString
}
