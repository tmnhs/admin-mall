package mall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mall"
	mallReq "github.com/flipped-aurora/gin-vue-admin/server/model/mall/request"
	"time"
)

type ReportService struct {
}

// CreateReport 创建Report记录

func (reportService *ReportService) CreateReport(report mall.Report) (err error) {
	report.CreatedTime = time.Now().Unix() * 1000
	err = global.GVA_DBList[MallDataBase].Table(ReportTableName).Create(&report).Error
	return err
}

// DeleteReport 删除Report记录

func (reportService *ReportService) DeleteReport(report mall.Report) (err error) {
	err = global.GVA_DBList[MallDataBase].Table(ReportTableName).Delete(&report).Error
	return err
}

// DeleteReportByIds 批量删除Report记录

func (reportService *ReportService) DeleteReportByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DBList[MallDataBase].Table(ReportTableName).Delete(&[]mall.Report{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateReport 更新Report记录

func (reportService *ReportService) UpdateReport(report mall.Report) (err error) {
	err = global.GVA_DBList[MallDataBase].Table(ReportTableName).Save(&report).Error
	return err
}

// GetReport 根据id获取Report记录

func (reportService *ReportService) GetReport(id int64) (err error, report mall.Report) {
	err = global.GVA_DBList[MallDataBase].Table(ReportTableName).Where("id = ?", id).First(&report).Error
	return
}

// GetReportInfoList 分页获取Report记录

func (reportService *ReportService) GetReportInfoList(info mallReq.ReportSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DBList[MallDataBase].Table(ReportTableName)
	var reports []mall.Report
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ItemId != nil {
		db = db.Where("item_id = ?", info.ItemId)
	}
	if info.Type != nil {
		db = db.Where("type = ?", info.Type)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&reports).Error
	return err, reports, total
}
