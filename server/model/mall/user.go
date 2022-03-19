package mall

type Profile struct {
	UID       int64  `gorm:"column:uid"`
	UserName  string `json:"username" form:"username" gorm:"column:username;comment:;"`
	Gender    *int   `json:"gender" form:"gender" gorm:"column:gender;comment:;size:1;"`
	AddressID int64  `json:"address" form:"address" gorm:"column:address_id;comment:;size:1;"`
	IsOnline  int    `json:"isOnline" form:"isOnline" gorm:"column:is_online;comment:;size:1;"`
	Status    int    `json:"status" from:"status" gorm:"column:status"`
}

type UserStatisInfo struct {
	UID           int64 `gorm:"column:uid"`
	OrderCount    int   `json:"orderCount" form:"orderCount" gorm:"column:finish_order_count;comment:;"`
	FollowCount   int   `json:"followCount" form:"followCount" gorm:"column:follow_count;comment:;"`
	LoginCount    int   `json:"loginCount" form:"loginCount" gorm:"column:login_count;comment:;"`
	LastLoginTime int64 `json:"lastLoginTime" form:"lastLoginTime" gorm:"column:last_login_time;comment:;"`
}

type Address struct {
	ID     int64  `json:"ID" gorm:"column:id"`
	Region int    `json:"region" gorm:"column:region;comment:地址区域，1表示韵苑，2表示沁苑，3表示紫菘，4表示其他"`
	Other  string `json:"other" gorm:"column:other"`
	Unit   string `json:"unit" gorm:"column:unit"`
}

// UserLoginLog 结构
type UserLoginLog struct {
	ID        int64  `json:"ID" gorm:"column:id"`
	UID       int64  `json:"uid" gorm:"column:uid"`
	Username  string `json:"username" form:"username" gorm:"column:username;comment:;"`
	Ip        string `json:"ip" form:"ip" gorm:"column:ip;comment:;"`
	LoginType *int   `json:"loginType" form:"loginType" gorm:"column:login_type;comment:;size:1;"`
	City      string `json:"city" form:"city" gorm:"column:city;comment:;"`
	LoginTime int64  `json:"loginTime" form:"loginTime" gorm:"column:login_time;comment:;"`
}
