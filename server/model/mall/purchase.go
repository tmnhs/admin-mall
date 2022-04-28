package mall

type PurchaseItem struct {
	ID            int64    `json:"ID" form:"ID"`
	Uid           *int64   `json:"uid" form:"uid" gorm:"column:uid;comment:;"`
	Name          string   `json:"name" form:"name" gorm:"column:name;comment:;"`
	UserName      string   `json:"username" form:"username" gorm:"-" `
	CategoryName  string   `json:"categoryName" form:"categoryName" gorm:"-"`
	CategoryId    *int64   `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:;"`
	Sort          *int     `json:"sort" form:"sort" gorm:"column:sort;comment:;"`
	MaxPrice      *float64 `json:"maxPrice" form:"maxPrice" gorm:"column:max_price;comment:;"`
	MinPrice      *float64 `json:"minPrice" form:"minPrice" gorm:"column:min_price;comment:;"`
	Description   string   `json:"description" form:"description" gorm:"column:description;comment:;"`
	TradeWay      *int     `json:"tradeWay" form:"tradeWay" gorm:"column:trade_way;comment:;"`
	Status        bool     `json:"status" form:"status" gorm:"column:status;comment:已购状态，1表示已购，0表示未购得;"`
	PublishStatus bool     `json:"publishStatus" form:"publishStatus" gorm:"column:publish_status;comment:上架状态，1表示上架，0表示下架;"`
	CreatedTime   int64    `json:"createdTime" form:"createdTime" gorm:"column:created_time;comment:;"`
	UpdatedTime   int64    `json:"updatedTime" form:"updatedTime" gorm:"column:updated_time;comment:;"`
}
