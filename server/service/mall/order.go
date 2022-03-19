package mall

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mall"
	mallReq "github.com/flipped-aurora/gin-vue-admin/server/model/mall/request"
	"time"
)

type OrderService struct {
}

// CreateOrder 创建Order记录
func (orderService *OrderService) CreateOrder(order mall.Order) (err error) {
	order.CreatedTime = time.Now().Unix() * 1000
	err = global.GVA_DBList[MallDataBase].Table(OrderTableName).Create(&order).Error
	return err
}

// DeleteOrder 删除Order记录
func (orderService *OrderService) DeleteOrder(order mall.Order) (err error) {
	err = global.GVA_DBList[MallDataBase].Table(OrderTableName).Delete(&order).Error
	return err
}

// DeleteOrderByIds 批量删除Order记录
func (orderService *OrderService) DeleteOrderByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DBList[MallDataBase].Table(OrderTableName).Delete(&[]mall.Order{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateOrder 更新Order记录
func (orderService *OrderService) UpdateOrder(order mall.Order) (err error) {
	err = global.GVA_DBList[MallDataBase].Table(OrderTableName).Save(&order).Error
	return err
}

// GetOrder 根据id获取Order记录
func (orderService *OrderService) GetOrder(id int64) (err error, order mall.Order) {
	err = global.GVA_DBList[MallDataBase].Table(OrderTableName).Where("id = ?", id).First(&order).Error
	return
}

// GetOrderInfoList 分页获取Order记录
func (orderService *OrderService) GetOrderInfoList(info mallReq.OrderSearch) (err error, list []mall.Order, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DBList[MallDataBase].Table(OrderTableName)
	var orders []mall.Order
	var result []mall.Order
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ID > 0 {
		db = db.Where("id = ?", info.ID)
	}

	if len(info.OrderSn) > 0 {
		db = db.Where("order_sn = ?", info.OrderSn)
	}
	if info.ReceiverId != nil && *info.ReceiverId > 0 {
		db = db.Where("receiver_id = ?", info.ReceiverId)
	}
	if info.Type != nil {
		db = db.Where("type = ?", info.Type)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&orders).Error
	userService := new(UserService)
	for _, order := range orders {
		user, _ := userService.GetProfileById(*order.ReceiverId)
		if user != nil {
			order.ReceiverName = user.UserName
		}
		result = append(result, order)
	}
	return err, result, total
}

// CreateOrderReturn 创建OrderReturn记录
func (orderService *OrderService) CreateOrderReturn(orderReturn mall.OrderReturn) (err error) {
	orderReturn.CreatedTime = time.Now().Unix() * 1000
	err = global.GVA_DBList[MallDataBase].Table(OrderReturnTableName).Create(&orderReturn).Error
	return err
}

// DeleteOrderReturn 删除OrderReturn记录
func (orderService *OrderService) DeleteOrderReturn(orderReturn mall.OrderReturn) (err error) {
	err = global.GVA_DBList[MallDataBase].Table(OrderReturnTableName).Delete(&orderReturn).Error
	return err
}

// DeleteOrderReturnByIds 批量删除OrderReturn记录
func (orderService *OrderService) DeleteOrderReturnByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DBList[MallDataBase].Table(OrderReturnTableName).Delete(&[]mall.OrderReturn{}, "order_id in ?", ids.Ids).Error
	return err
}

// UpdateOrderReturn 更新OrderReturn记录
func (orderService *OrderService) UpdateOrderReturn(orderReturn mall.OrderReturn) (err error) {
	err = global.GVA_DBList[MallDataBase].Table(OrderReturnTableName).Save(&orderReturn).Error
	return err
}

// GetOrderReturn 根据id获取OrderReturn记录
func (orderService *OrderService) GetOrderReturn(id int64) (err error, orderReturn mall.OrderReturn) {
	err = global.GVA_DBList[MallDataBase].Table(OrderReturnTableName).Where("order_id = ?", id).First(&orderReturn).Error
	return
}

// GetOrderReturnInfoList 分页获取OrderReturn记录
func (orderService *OrderService) GetOrderReturnInfoList(info mallReq.OrderReturnSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DBList[MallDataBase].Table(OrderReturnTableName)
	var orderReturns []mall.OrderReturn
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.OrderId > 0 {
		db = db.Where("order_id = ?", info.OrderId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&orderReturns).Error
	return err, orderReturns, total
}

// CreateOrderHistory 创建OrderHistory记录
func (orderService *OrderService) CreateOrderHistory(orderHistory mall.OrderHistory) (err error) {
	orderHistory.CreatedTime = time.Now().Unix() * 1000
	err = global.GVA_DBList[MallDataBase].Table(OrderHistoryTableName).Create(&orderHistory).Error
	return err
}

// DeleteOrderHistory 删除OrderHistory记录
func (orderService *OrderService) DeleteOrderHistory(orderHistory mall.OrderHistory) (err error) {
	err = global.GVA_DBList[MallDataBase].Table(OrderHistoryTableName).Delete(&orderHistory).Error
	return err
}

// DeleteOrderHistoryByIds 批量删除OrderHistory记录

func (orderService *OrderService) DeleteOrderHistoryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DBList[MallDataBase].Table(OrderHistoryTableName).Delete(&[]mall.OrderHistory{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateOrderHistory 更新OrderHistory记录

func (orderService *OrderService) UpdateOrderHistory(orderHistory mall.OrderHistory) (err error) {
	err = global.GVA_DBList[MallDataBase].Table(OrderHistoryTableName).Save(&orderHistory).Error
	return err
}

// GetOrderHistory 根据id获取OrderHistory记录

func (orderService *OrderService) GetOrderHistory(id int64) (err error, orderHistory mall.OrderHistory) {
	err = global.GVA_DBList[MallDataBase].Table(OrderHistoryTableName).Where("id = ?", id).First(&orderHistory).Error
	return
}

// GetOrderHistoryInfoList 分页获取OrderHistory记录

func (orderService *OrderService) GetOrderHistoryInfoList(info mallReq.OrderHistorySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DBList[MallDataBase].Table(OrderHistoryTableName)
	var orderHistorys []mall.OrderHistory
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.OrderId != nil {
		db = db.Where("order_id = ?", info.OrderId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&orderHistorys).Error
	return err, orderHistorys, total
}
