package mall

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mall"
	mallReq "github.com/flipped-aurora/gin-vue-admin/server/model/mall/request"
	"time"
)

type IdleItemService struct {
}

// CreateIdleItem 创建IdleItem记录

func (idleItemService *IdleItemService) CreateIdleItem(idleItem mallReq.ReqIdleItem) (err error) {
	imageByte, err := json.Marshal(idleItem.AlbumImages)
	if err != nil {
		return err
	}
	idleItem.CreatedTime = time.Now().Unix() * 1000
	idleItemGorm := &mall.IdleItemGorm{
		IdleItem: idleItem.IdleItem,
	}
	idleItemGorm.AlbumImages = string(imageByte)
	err = global.GVA_DBList[MallDataBase].Table(IdleItemTableName).Create(&idleItemGorm).Error
	return err
}

// DeleteIdleItem 删除IdleItem记录

func (idleItemService *IdleItemService) DeleteIdleItem(idleItem mall.IdleItem) (err error) {
	err = global.GVA_DBList[MallDataBase].Table(IdleItemTableName).Delete(&idleItem).Error
	return err
}

// DeleteIdleItemByIds 批量删除IdleItem记录
func (idleItemService *IdleItemService) DeleteIdleItemByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DBList[MallDataBase].Table(IdleItemTableName).Delete(&[]mall.IdleItem{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateIdleItem 更新IdleItem记录
func (idleItemService *IdleItemService) UpdateIdleItem(idleItem mallReq.ReqIdleItem) (err error) {
	idleItem.UpdatedTime = time.Now().Unix() * 1000
	imageByte, err := json.Marshal(idleItem.AlbumImages)
	if err != nil {
		return err
	}
	idleItemGorm := &mall.IdleItemGorm{
		IdleItem: idleItem.IdleItem,
	}
	idleItemGorm.AlbumImages = string(imageByte)
	err = global.GVA_DBList[MallDataBase].Table(IdleItemTableName).Save(&idleItemGorm).Error
	return err
}

// GetIdleItem 根据id获取IdleItem记录
func (idleItemService *IdleItemService) GetIdleItem(id int64) (error, *mallReq.ReqIdleItem) {
	var idleItem mall.IdleItemGorm
	err := global.GVA_DBList[MallDataBase].Table(IdleItemTableName).Where("id = ?", id).First(&idleItem).Error
	if err != nil {
		return err, nil
	}

	categoryService := new(CategoryService)
	userService := new(UserService)
	_, category := categoryService.GetCategory(*idleItem.CategoryId)
	if category != nil {
		idleItem.CategoryName = category.Name
	}
	user, _ := userService.GetProfileById(*idleItem.Uid)
	if user != nil {
		idleItem.UserName = user.UserName
	}
	idleRsp := mallReq.ReqIdleItem{
		IdleItem: idleItem.IdleItem,
	}
	_ = json.Unmarshal([]byte(idleItem.AlbumImages), &idleRsp.AlbumImages)

	return nil, &idleRsp
}

// GetIdleItemInfoList 分页获取IdleItem记录
func (idleItemService *IdleItemService) GetIdleItemInfoList(info mallReq.IdleItemSearch) (err error, list []mallReq.ReqIdleItem, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DBList[MallDataBase].Table(IdleItemTableName)
	var idleItems []mall.IdleItemGorm
	var result []mallReq.ReqIdleItem
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
	err = db.Limit(limit).Offset(offset).Find(&idleItems).Error
	categoryService := new(CategoryService)
	userService := new(UserService)
	for _, idle := range idleItems {
		_, category := categoryService.GetCategory(*idle.CategoryId)
		if category != nil {
			idle.CategoryName = category.Name
		}
		user, _ := userService.GetProfileById(*idle.Uid)
		if user != nil {
			idle.UserName = user.UserName
		}
		idleRsp := mallReq.ReqIdleItem{
			IdleItem: idle.IdleItem,
		}
		_ = json.Unmarshal([]byte(idle.AlbumImages), &idleRsp.AlbumImages)
		/*	if err != nil {
			fmt.Println("err:",err)
			return
		}*/
		result = append(result, idleRsp)
		fmt.Println(idleRsp)
	}
	return err, result, total
}
