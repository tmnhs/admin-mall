package request

import (
	"github.com/tmnhs/admin-mall/server/model/common/request"
	"github.com/tmnhs/admin-mall/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
