package mall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mall"
	mallReq "github.com/flipped-aurora/gin-vue-admin/server/model/mall/request"
)

type CategoryService struct {
}

// CreateCategory 创建Category记录
func (categoryService *CategoryService) CreateCategory(category mall.Category) (err error) {
	err = global.GVA_DBList[MallDataBase].Table(ItemCategoryTableName).Create(&category).Error
	return err
}

// DeleteCategory 删除Category记录
func (categoryService *CategoryService) DeleteCategory(category mall.Category) (err error) {
	err = global.GVA_DBList[MallDataBase].Table(ItemCategoryTableName).Delete(&category).Error
	return err
}

// DeleteCategoryByIds 批量删除Category记录

func (categoryService *CategoryService) DeleteCategoryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DBList[MallDataBase].Table(ItemCategoryTableName).Delete(&[]mall.Category{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCategory 更新Category记录
func (categoryService *CategoryService) UpdateCategory(category mall.Category) (err error) {
	err = global.GVA_DBList[MallDataBase].Table(ItemCategoryTableName).Save(&category).Error
	return err
}

// GetCategory 根据id获取Category记录
func (categoryService *CategoryService) GetCategory(id int64) (err error, category *mall.Category) {
	err = global.GVA_DBList[MallDataBase].Table(ItemCategoryTableName).Where("id = ?", id).First(&category).Error
	return
}

// GetCategoryInfoList 分页获取Category记录
func (categoryService *CategoryService) GetCategoryInfoList(info mallReq.CategorySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DBList[MallDataBase].Table(ItemCategoryTableName)
	var categorys []mall.Category
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ID > 0 {
		db = db.Where("id = ?", info.ID)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Level != nil {
		db = db.Where("level = ?", info.Level)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&categorys).Error
	return err, categorys, total
}
