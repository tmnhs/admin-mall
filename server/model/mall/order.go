package mall

// Order 结构体
// 如果含有time.Time 请自行import time包
type Order struct {
	ID            int64    `json:"ID" form:"ID"`
	OrderSn       string   `json:"orderSn" form:"orderSn" gorm:"column:order_sn;comment:;"`
	ReceiverId    *int64   `json:"receiverId" form:"receiverId" gorm:"column:receiver_id;comment:;"`
	ReceiverName  string   `json:"receiverName" form:"receiverName" gorm:"-"`
	Type          *int     `json:"type" form:"type" gorm:"column:type;comment:;"`
	ItemId        *int     `json:"itemId" form:"itemId" gorm:"column:item_id;comment:;"`
	TotalAmount   *float64 `json:"totalAmount" form:"totalAmount" gorm:"column:total_amount;comment:;"`
	Address       string   `json:"address" form:"address" gorm:"column:address;comment:;"`
	TradeTime     *int     `json:"tradeTime" form:"tradeTime" gorm:"column:trade_time;comment:;"`
	Status        *int     `json:"status" form:"status" gorm:"column:status;comment:;"`
	Note          string   `json:"note" form:"note" gorm:"column:note;comment:;"`
	ConfirmStatus *bool    `json:"confirmStatus" form:"confirmStatus" gorm:"column:confirm_status;comment:;"`
	DeleteStatus  *bool    `json:"deleteStatus" form:"deleteStatus" gorm:"column:delete_status;comment:;"`
	CreatedTime   int64    `json:"createdTime" form:"createdTime" gorm:"column:created_time;comment:;"`
}

// OrderReturn 结构体
type OrderReturn struct {
	ID          int64  `json:"ID" form:"ID"`
	OrderId     int64  `json:"orderId" form:"orderId" gorm:"column:order_id;comment:;"`
	Reason      string `json:"reason" form:"reason" gorm:"column:reason;comment:;"`
	Sort        *int   `json:"sort" form:"sort" gorm:"column:sort;comment:;"`
	Status      *bool  `json:"status" form:"status" gorm:"column:status;comment:;"`
	CreatedTime int64  `json:"createdTime" form:"createdTime" gorm:"column:created_time;comment:;"`
}

// OrderHistory 结构体
type OrderHistory struct {
	ID          int64  `json:"ID" form:"ID"`
	OrderId     *int   `json:"orderId" form:"orderId" gorm:"column:order_id;comment:;"`
	Type        *int   `json:"type" form:"type" gorm:"column:type;comment:;"`
	Uid         *int   `json:"uid" form:"uid" gorm:"column:uid;comment:;"`
	OrderStatus *int   `json:"orderStatus" form:"orderStatus" gorm:"column:order_status;comment:;"`
	Note        string `json:"note" form:"note" gorm:"column:note;comment:;"`
	CreatedTime int64  `json:"createdTime" form:"createdTime" gorm:"column:created_time;comment:;"`
}
