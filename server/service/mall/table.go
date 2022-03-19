package mall

import (
	"github.com/go-redis/redis/v8"
)

const (
	MallDataBase                = "hust_mall"
	UserTableName               = "m_user"
	ProfileTableName            = "m_profile"
	UserStatisticsInfoTableName = "m_user_statistics_info"
	LoginLogTableName           = "m_user_login_log"
	AddressTableName            = "m_address"
	ItemCategoryTableName       = "m_item_category"
	IdleItemTableName           = "m_idle_item"
	PurchaseItemTableName       = "m_purchase_item"
	OrderTableName              = "m_order"
	OrderReturnTableName        = "m_order_return"
	AdTableName                 = "m_ad"
	ReportTableName             = "m_report"
	FeedBackTableName           = "m_feedback"
	OrderHistoryTableName       = "m_order_operate_history"
)

const ErrRedisNotFound = redis.Nil
