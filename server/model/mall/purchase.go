package mall

// PurchaseItem 结构体
type PurchaseItem struct {
	ID           int64    `json:"ID" form:"ID"`
	Uid          *int64   `json:"uid" form:"uid" gorm:"column:uid;comment:;"`
	UserName     string   `json:"username" form:"username" gorm:"-" `
	CategoryName string   `json:"categoryName" form:"categoryName" gorm:"-"`
	Name         string   `json:"name" form:"name" gorm:"column:name;comment:;"`
	CategoryId   *int64   `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:;"`
	Sort         *int     `json:"sort" form:"sort" gorm:"column:sort;comment:;"`
	Price        *float64 `json:"price" form:"price" gorm:"column:price;comment:;"`
	Description  string   `json:"description" form:"description" gorm:"column:description;comment:;"`
	TradeWay     *int     `json:"tradeWay" form:"tradeWay" gorm:"column:trade_way;comment:;"`
	Status       *bool    `json:"status" form:"status" gorm:"column:status;comment:;"`
	DeleteStatus *bool    `json:"deleteStatus" form:"deleteStatus" gorm:"column:delete_status;comment:;"`
	CreatedTime  int64    `json:"createdTime" form:"createdTime" gorm:"column:created_time;comment:;"`
	UpdatedTime  int64    `json:"updatedTime" form:"updatedTime" gorm:"column:updated_time;comment:;"`
}
