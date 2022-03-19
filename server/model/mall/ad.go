package mall

// Ad 结构体
type Ad struct {
	ID             int64  `json:"ID" form:"ID"`
	Name           string `json:"name" form:"name" gorm:"column:name;comment:;"`
	Type           *int   `json:"type" form:"type" gorm:"column:type;comment:;"`
	Image          string `json:"image" form:"image" gorm:"column:image;comment:;"`
	Width          *int   `json:"width" form:"width" gorm:"column:width;comment:;"`
	Height         *int   `json:"height" form:"height" gorm:"column:height;comment:;"`
	Introduce      string `json:"introduce" form:"introduce" gorm:"column:introduce;comment:;"`
	Url            string `json:"url" form:"url" gorm:"column:url;comment:;"`
	Status         *bool  `json:"status" form:"status" gorm:"column:status;comment:;"`
	CreatedTime    int64  `json:"createdTime" form:"createdTime" gorm:"column:created_time;comment:;"`
	ExpirationTime int64  `json:"expirationTime" form:"expirationTime" gorm:"column:expiration_time;comment:;"`
}
