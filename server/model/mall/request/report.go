package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mall"
)

type ReportSearch struct {
	mall.Report
	request.PageInfo
}
