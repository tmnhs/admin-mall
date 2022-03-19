package mall

// IdleItem 结构体
type IdleItem struct {
	ID              int64    `json:"ID" form:"ID"`
	Uid             *int64   `json:"uid" form:"uid" gorm:"column:uid;comment:;"`
	UserName        string   `json:"username" form:"username" gorm:"-" `
	CategoryName    string   `json:"categoryName" form:"categoryName" gorm:"-"`
	Name            string   `json:"name" form:"name" gorm:"column:name;comment:;"`
	CategoryId      *int64   `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:;"`
	Brand           string   `json:"brand" form:"brand" gorm:"column:brand;comment:;"`
	PurchaseChannel *int     `json:"purchaseChannel" form:"purchaseChannel" gorm:"column:purchase_channel;comment:;"`
	Quality         *int     `json:"quality" form:"quality" gorm:"column:quality;comment:;"`
	ShelfLife       string   `json:"shelfLife" form:"shelfLife" gorm:"column:shelf_life;comment:;"`
	Sort            *int     `json:"sort" form:"sort" gorm:"column:sort;comment:;"`
	OriginPrice     *float64 `json:"originPrice" form:"originPrice" gorm:"column:origin_price;comment:;"`
	Price           *float64 `json:"price" form:"price" gorm:"column:price;comment:;"`
	Description     string   `json:"description" form:"description" gorm:"column:description;comment:;"`
	Stock           *int     `json:"stock" form:"stock" gorm:"column:stock;comment:;"`
	TradeWay        *int     `json:"tradeWay" form:"tradeWay" gorm:"column:trade_way;comment:;"`
	SaleStatus      *bool    `json:"saleStatus" form:"saleStatus" gorm:"column:sale_status;comment:;"`
	PublishStatus   *bool    `json:"publishStatus" form:"publishStatus" gorm:"column:publish_status;comment:;"`
	CreatedTime     int64    `json:"createdTime" form:"createdTime" gorm:"column:created_time;comment:;"`
	UpdatedTime     int64    `json:"updatedTime" form:"updatedTime" gorm:"column:updated_time;comment:;"`
}

type ImageArray struct {
	AlbumImages []string `json:"albumImages" form:"albumImages"`
}
type ImageString struct {
	AlbumImages string `json:"albumImages" form:"albumImages" gorm:"column:album_images"`
}

type IdleItemGorm struct {
	IdleItem
	ImageString
}
