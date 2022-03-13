package request

import (
	"github.com/tmnhs/admin-mall/server/model/autocode"
	"github.com/tmnhs/admin-mall/server/model/common/request"
)

type {{.StructName}}Search struct{
    autocode.{{.StructName}}
    request.PageInfo
}