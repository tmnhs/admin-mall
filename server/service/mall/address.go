package mall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mall"
	"go.uber.org/zap"
)

type AddressService struct {
}

func (a *AddressService) GetAddressById(id int64) (*mall.Address, error) {
	var address mall.Address
	err := global.GVA_DBList[MallDataBase].Table(AddressTableName).Where("id = ?", id).First(&address).Error
	if err != nil {
		global.GVA_LOG.Error("GetAddressById  err:%s", zap.Error(err))
		return nil, err
	}
	return &address, nil
}
func (a *AddressService) DeleteAddress(id int64) error {
	var address mall.Address
	err := global.GVA_DBList[MallDataBase].Table(AddressTableName).Where("id = ?", id).Delete(&address).Error
	if err != nil {
		global.GVA_LOG.Error("DeleteAddress  err:%s", zap.Error(err))
		return err
	}
	return nil
}
