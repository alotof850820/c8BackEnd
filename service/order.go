package service

import (
	"context"
	"fmt"
	"gin_mall_tmp/dao"
	"gin_mall_tmp/model"
	"gin_mall_tmp/pkg/e"
	"gin_mall_tmp/serializer"
	"math/rand"
	"strconv"
	"time"
)

type OrderService struct {
	ProductId uint    `form:"product_id" json:"product_id"`
	Num       int     `form:"num" json:"num"`
	AddressId uint    `form:"address_id" json:"address_id"`
	Money     float64 `form:"money" json:"money"`
	BossId    uint    `form:"boss_id" json:"boss_id"`
	UserId    uint    `form:"user_id" json:"user_id"`
	OrderNum  int     `form:"order_num" json:"order_num"`
	Type      int     `form:"type" json:"type"`
	model.BasePage
}

func (c *OrderService) CreateOrder(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	OrderDao := dao.NewOrderDao(ctx)
	Order := &model.Order{
		UserId:    uid,
		ProductId: c.ProductId,
		BossId:    c.BossId,
		Num:       c.Num,
		Money:     c.Money,
		Type:      1, // 1:未支付 2:已支付 默認1
	}

	addressDao := dao.NewAddressDao(ctx)

	// 获取地址檢驗地址是否存在
	address, err := addressDao.GetAddressByAId(c.AddressId)
	Order.AddressId = address.ID
	// 訂單號 自動生成隨機號+唯一號
	number := fmt.Sprintf("%09v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000000))
	productNum := strconv.Itoa(int(c.ProductId))
	userNum := strconv.Itoa(int(c.UserId))
	// 訂單號拼接
	Order.OrderNum, _ = strconv.ParseUint(number+productNum+userNum, 10, 64)

	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	err = OrderDao.CreateOrder(Order)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (c *OrderService) GetOrder(ctx context.Context, uid uint, oId string) serializer.Response {
	OrderId, _ := strconv.Atoi(oId)
	code := e.Success
	OrderDao := dao.NewOrderDao(ctx)

	// 拿訂單
	order, err := OrderDao.GetOrderByOId(uint(OrderId), uid)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 拿地址
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressByAId(order.AddressId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 拿產品
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(order.ProductId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildOrder(order, product, address),
	}
}

func (c *OrderService) GetOrders(ctx context.Context, uid uint) serializer.Response {
	code := e.Success

	if c.PageSize == 0 {
		c.PageSize = 15 //默認15
	}

	// 查詢
	condition := make(map[string]interface{})
	if c.Type != 0 { // 0:全部 1:未支付 2:已支付
		condition["type"] = c.Type
	}
	condition["user_id"] = uid

	OrderDao := dao.NewOrderDao(ctx)
	Order, total, err := OrderDao.GetOrderesByCondition(condition, c.BasePage)

	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildOrders(ctx, Order), uint(total))
}

func (c *OrderService) DeleteOrder(ctx context.Context, cid string, uid uint) serializer.Response {
	OrderId, _ := strconv.Atoi(cid)
	code := e.Success
	OrderDao := dao.NewOrderDao(ctx)

	err := OrderDao.DeleteOrderByOId(uint(OrderId), uid)

	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
