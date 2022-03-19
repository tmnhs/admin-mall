package mall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mall"
	mallReq "github.com/flipped-aurora/gin-vue-admin/server/model/mall/request"
	"time"
)

type FeedbackService struct {
}

// CreateFeedback 创建Feedback记录

func (feedbackService *FeedbackService) CreateFeedback(feedback mall.Feedback) (err error) {
	feedback.CreatedTime = time.Now().Unix() * 1000
	err = global.GVA_DBList[MallDataBase].Table(FeedBackTableName).Create(&feedback).Error
	return err
}

// DeleteFeedback 删除Feedback记录

func (feedbackService *FeedbackService) DeleteFeedback(feedback mall.Feedback) (err error) {
	err = global.GVA_DBList[MallDataBase].Table(FeedBackTableName).Delete(&feedback).Error
	return err
}

// DeleteFeedbackByIds 批量删除Feedback记录

func (feedbackService *FeedbackService) DeleteFeedbackByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DBList[MallDataBase].Table(FeedBackTableName).Delete(&[]mall.Feedback{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateFeedback 更新Feedback记录

func (feedbackService *FeedbackService) UpdateFeedback(feedback mall.Feedback) (err error) {
	err = global.GVA_DBList[MallDataBase].Table(FeedBackTableName).Save(&feedback).Error
	return err
}

// GetFeedback 根据id获取Feedback记录

func (feedbackService *FeedbackService) GetFeedback(id int64) (err error, feedback mall.Feedback) {
	err = global.GVA_DBList[MallDataBase].Table(FeedBackTableName).Where("id = ?", id).First(&feedback).Error
	return
}

// GetFeedbackInfoList 分页获取Feedback记录

func (feedbackService *FeedbackService) GetFeedbackInfoList(info mallReq.FeedbackSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DBList[MallDataBase].Table(FeedBackTableName)
	var feedbacks []mall.Feedback
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&feedbacks).Error
	return err, feedbacks, total
}
