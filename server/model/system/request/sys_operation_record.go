package request

import (
	"github.com/tmnhs/admin-mall/server/model/common/request"
	"github.com/tmnhs/admin-mall/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
