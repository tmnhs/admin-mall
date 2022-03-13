package request

import (
	"github.com/tmnhs/admin-mall/server/model/common/request"
	"github.com/tmnhs/admin-mall/server/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
