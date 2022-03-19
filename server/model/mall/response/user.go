package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/mall"
	"time"
)

type RspUser struct {
	ID            int64         `json:"ID" form:"ID"`
	UserName      string        `json:"username" form:"username" `
	Gender        *int          `json:"gender" form:"gender" `
	Address       *mall.Address `json:"address" form:"address" `
	IsOnline      int           `json:"isOnline" form:"isOnline" `
	Status        int           `json:"status" form:"status"`
	OrderCount    int           `json:"orderCount" form:"orderCount" `
	FollowCount   int           `json:"followCount" form:"followCount" `
	LoginCount    int           `json:"loginCount" form:"loginCount" `
	LastLoginTime time.Time     `json:"lastLoginTime" form:"lastLoginTime" `
}

// UserLoginLog 结构
type RspUserLoginLog struct {
	ID        int64     `json:"ID" gorm:"column:id"`
	UID       int64     `json:"uid" gorm:"column:uid"`
	Username  string    `json:"username" form:"username" gorm:"column:username;comment:;"`
	Ip        string    `json:"ip" form:"ip" gorm:"column:ip;comment:;"`
	LoginType *int      `json:"loginType" form:"loginType" gorm:"column:login_type;comment:;size:1;"`
	City      string    `json:"city" form:"city" gorm:"column:city;comment:;"`
	LoginTime time.Time `json:"loginTime" form:"loginTime" gorm:"column:login_time;comment:;"`
}
