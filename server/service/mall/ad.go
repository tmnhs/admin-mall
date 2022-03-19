package mall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mall"
	mallReq "github.com/flipped-aurora/gin-vue-admin/server/model/mall/request"
	"time"
)

type AdService struct {
}

// CreateAd 创建Ad记录

func (adService *AdService) CreateAd(ad mall.Ad) (err error) {
	ad.CreatedTime = time.Now().Unix() * 1000
	err = global.GVA_DBList[MallDataBase].Table(AdTableName).Create(&ad).Error
	return err
}

// DeleteAd 删除Ad记录

func (adService *AdService) DeleteAd(ad mall.Ad) (err error) {
	err = global.GVA_DBList[MallDataBase].Table(AdTableName).Delete(&ad).Error
	return err
}

// DeleteAdByIds 批量删除Ad记录

func (adService *AdService) DeleteAdByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DBList[MallDataBase].Table(AdTableName).Delete(&[]mall.Ad{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAd 更新Ad记录

func (adService *AdService) UpdateAd(ad mall.Ad) (err error) {
	err = global.GVA_DBList[MallDataBase].Table(AdTableName).Save(&ad).Error
	return err
}

// GetAd 根据id获取Ad记录

func (adService *AdService) GetAd(id int64) (err error, ad mall.Ad) {
	err = global.GVA_DBList[MallDataBase].Table(AdTableName).Where("id = ?", id).First(&ad).Error
	return
}

// GetAdInfoList 分页获取Ad记录

func (adService *AdService) GetAdInfoList(info mallReq.AdSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DBList[MallDataBase].Table(AdTableName)
	var ads []mall.Ad
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.Type != nil {
		db = db.Where("type = ?", info.Type)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&ads).Error
	return err, ads, total
}
