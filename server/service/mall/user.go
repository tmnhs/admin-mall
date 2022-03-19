package mall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mall"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mall/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mall/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
)

type UserService struct {
}

// GetUserInfoList 分页获取User记录
func (userService *UserService) GetUserInfoList(info request.UserSearch) (error, []response.RspUser, int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DBList[MallDataBase].Table(ProfileTableName)
	var profiles []mall.Profile
	var users []response.RspUser
	var total int64
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ID > 0 {
		db = db.Where("uid = ?", info.ID)
	}
	if info.UserName != "" {
		db = db.Where("username LIKE ?", "%"+info.UserName+"%")
	}
	if info.Gender != nil {
		db = db.Where("gender = ?", info.Gender)
	}
	err := db.Count(&total).Error
	if err != nil {
		return err, nil, 0
	}
	err = db.Limit(limit).Offset(offset).Find(&profiles).Error
	addressService := new(AddressService)
	for _, profile := range profiles {
		address, _ := addressService.GetAddressById(profile.AddressID)
		userStatis, _ := userService.GetUserStatis(profile.UID)
		//if userStatis==nil{
		//	total--
		//	continue
		//}
		user := response.RspUser{
			ID:            profile.UID,
			UserName:      profile.UserName,
			Gender:        profile.Gender,
			Address:       address,
			IsOnline:      profile.IsOnline,
			Status:        profile.Status,
			OrderCount:    userStatis.OrderCount,
			FollowCount:   userStatis.FollowCount,
			LoginCount:    userStatis.LoginCount,
			LastLoginTime: utils.Stamp2Time(userStatis.LastLoginTime),
		}
		users = append(users, user)
	}
	return err, users, total
}

func (userService *UserService) GetProfileById(uid int64) (*mall.Profile, error) {
	var profile mall.Profile
	err := global.GVA_DBList[MallDataBase].Table(ProfileTableName).Where("uid = ? ", uid).First(&profile).Error
	if err != nil && err != ErrRedisNotFound {
		global.GVA_LOG.Error("GetProfileById err: %s", zap.Error(err))
		return nil, err
	}
	return &profile, nil
}

func (userService *UserService) GetUserStatis(uid int64) (*mall.UserStatisInfo, error) {
	var userStatis mall.UserStatisInfo
	err := global.GVA_DBList[MallDataBase].Table(UserStatisticsInfoTableName).Where("uid = ? ", uid).First(&userStatis).Error
	if err != nil && err != ErrRedisNotFound {
		global.GVA_LOG.Error("GetProfileById err: %s", zap.Error(err))
		return nil, err
	}
	return &userStatis, nil
}

//UpdateUserStatus 更新用户状态 [0->封禁、1->正常]
func (userService *UserService) UpdateUserStatus(uid int64, status int) error {
	err := global.GVA_DBList[MallDataBase].Table(ProfileTableName).Where("uid = ? ", uid).Updates(map[string]interface{}{
		"status": status}).Error
	if err != nil {
		global.GVA_LOG.Error("UpdateUserStatus err: %s", zap.Error(err))
		return nil
	}
	return nil
}

//分页获取用户登录日志
// GetUserInfoList 分页获取User记录
func (userService *UserService) GetUserLoginLogList(info request.UserLoginLogSearch) (error, []mall.UserLoginLog, int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DBList[MallDataBase].Table(LoginLogTableName)
	var loginLogs []mall.UserLoginLog
	var total int64
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Username != "" {
		db = db.Where("username LIKE ?", "%"+info.Username+"%")
	}
	if info.City != "" {
		db = db.Where("city LIKE ?", "%"+info.City+"%")
	}
	if info.LoginType != nil {
		db = db.Where("login_type = ?", info.LoginType)
	}
	err := db.Count(&total).Error
	if err != nil {
		return err, nil, 0
	}
	err = db.Limit(limit).Offset(offset).Find(&loginLogs).Error
	return err, loginLogs, total
}
