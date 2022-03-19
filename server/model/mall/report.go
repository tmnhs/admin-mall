package mall

// Report 结构体
// 如果含有time.Time 请自行import time包
type Report struct {
	ID          int64 `json:"ID" form:"ID"`
	Uid         *int  `json:"uid" form:"uid" gorm:"column:uid;comment:;"`
	ItemId      *int  `json:"itemId" form:"itemId" gorm:"column:item_id;comment:;"`
	Type        *int  `json:"type" form:"type" gorm:"column:type;comment:;"`
	Status      *bool `json:"status" form:"status" gorm:"column:status;comment:;"`
	CreatedTime int64 `json:"createdTime" form:"createdTime" gorm:"column:created_time;comment:;"`
}
