package mall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mall"
	mallReq "github.com/flipped-aurora/gin-vue-admin/server/model/mall/request"
	"time"
)

type PurchaseItemService struct {
}

// CreatePurchaseItem 创建PurchaseItem记录
func (purchaseItemService *PurchaseItemService) CreatePurchaseItem(purchaseItem mall.PurchaseItem) (err error) {
	purchaseItem.CreatedTime = time.Now().Unix() * 1000
	err = global.GVA_DBList[MallDataBase].Table(PurchaseItemTableName).Table(PurchaseItemTableName).Create(&purchaseItem).Error
	return err
}

// DeletePurchaseItem 删除PurchaseItem记录
func (purchaseItemService *PurchaseItemService) DeletePurchaseItem(purchaseItem mall.PurchaseItem) (err error) {
	err = global.GVA_DBList[MallDataBase].Table(PurchaseItemTableName).Delete(&purchaseItem).Error
	return err
}

// DeletePurchaseItemByIds 批量删除PurchaseItem记录
func (purchaseItemService *PurchaseItemService) DeletePurchaseItemByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DBList[MallDataBase].Table(PurchaseItemTableName).Delete(&[]mall.PurchaseItem{}, "id in ?", ids.Ids).Error
	return err
}

// UpdatePurchaseItem 更新PurchaseItem记录
func (purchaseItemService *PurchaseItemService) UpdatePurchaseItem(purchaseItem mall.PurchaseItem) (err error) {
	purchaseItem.UpdatedTime = time.Now().Unix() * 1000
	//err = global.GVA_DBList[MallDataBase].Table(PurchaseItemTableName).Where("uid = ?",purchaseItem.Uid).Updates(map[string]interface{}{
	//	"publish_status": purchaseItem.PublishStatus}).Error
	err = global.GVA_DBList[MallDataBase].Table(PurchaseItemTableName).Save(&purchaseItem).Error
	return err
}

// GetPurchaseItem 根据id获取PurchaseItem记录
func (purchaseItemService *PurchaseItemService) GetPurchaseItem(id int64) (err error, purchaseItem mall.PurchaseItem) {
	err = global.GVA_DBList[MallDataBase].Table(PurchaseItemTableName).Where("id = ?", id).First(&purchaseItem).Error
	return
}

// GetPurchaseItemInfoList 分页获取PurchaseItem记录
func (purchaseItemService *PurchaseItemService) GetPurchaseItemInfoList(info mallReq.PurchaseItemSearch) (err error, list []mall.PurchaseItem, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DBList[MallDataBase].Table(PurchaseItemTableName)
	var purchaseItems []mall.PurchaseItem
	var result []mall.PurchaseItem
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Uid != nil && *info.Uid > 0 {
		db = db.Where("uid = ?", info.Uid)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.CategoryId != nil && *info.CategoryId > 0 {
		db = db.Where("category_id = ?", info.CategoryId)
	}
	if info.TradeWay != nil {
		db = db.Where("trade_way = ?", info.TradeWay)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&purchaseItems).Error
	categoryService := new(CategoryService)
	userService := new(UserService)
	for _, purchase := range purchaseItems {
		_, category := categoryService.GetCategory(*purchase.CategoryId)
		if category != nil {
			purchase.CategoryName = category.Name
		}
		user, _ := userService.GetProfileById(*purchase.Uid)
		if user != nil {
			purchase.UserName = user.UserName
		}
		result = append(result, purchase)
	}
	return err, result, total
}
