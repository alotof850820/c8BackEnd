package service

import (
	"context"
	"errors"
	"fmt"
	"gin_mall_tmp/dao"
	"gin_mall_tmp/model"
	"gin_mall_tmp/pkg/e"
	"gin_mall_tmp/pkg/util"
	"gin_mall_tmp/serializer"
	"strconv"
)

type PaymentService struct {
	OrderId   uint    `form:"order_id" json:"order_id"`
	Money     float64 `form:"money" json:"money"`
	OrderNo   string  `form:"order_no" json:"order_no"`
	ProductId uint    `form:"product_id" json:"product_id"`
	PayTime   string  `form:"pay_time" json:"pay_time"`
	Sign      string  `form:"sign" json:"sign"`
	BossId    uint    `form:"boss_id" json:"boss_id"`
	BossName  string  `form:"boss_name" json:"boss_name"`
	Num       int     `form:"num" json:"num"`
	Key       string  `form:"key" json:"key"`
}

func (s *PaymentService) OrderPay(ctx context.Context, uid uint) serializer.Response {
	util.Encrypt.SetKey(s.Key)
	code := e.Success

	// 拿訂單
	orderDao := dao.NewOrderDao(ctx)
	tx := orderDao.Begin() // 確保操作能成功 失敗會回朔
	order, err := orderDao.GetOrderByOId(s.OrderId, uid)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	money := order.Money
	num := order.Num
	money = money * float64(num)
	// 拿用戶準備扣錢
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uid)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 解密錢 減訂單 再加密保存
	moneyStr := util.Encrypt.AesDecoding(user.Money)
	moneyFloat, _ := strconv.ParseFloat(moneyStr, 64)

	// 金額不足
	if moneyFloat-money < 0.0 {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  errors.New("金額不足").Error(),
		}
	}

	// 額外加密
	finMoney := fmt.Sprintf("%f", moneyFloat-money) // 将各种类型的数据格式化为字符串。
	finMoney = util.Encrypt.AesEncoding(finMoney)
	user.Money = finMoney

	// 更新user
	userDao = dao.NewUserDaoByDb(userDao.DB) // 重新指定DB
	if err := userDao.UpdateUserById(uid, user); err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 商家加錢
	bossDao := dao.NewUserDao(ctx)
	boss, err := bossDao.GetUserById(s.BossId)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	bossMoneyStr := util.Encrypt.AesDecoding(boss.Money)
	bossMoneyFloat, _ := strconv.ParseFloat(bossMoneyStr, 64)
	bossMoneyFloat += money
	bossMoney := fmt.Sprintf("%f", bossMoneyFloat)
	boss.Money = util.Encrypt.AesEncoding(bossMoney)
	if err := bossDao.UpdateUserById(s.BossId, boss); err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 商品數目減
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(s.ProductId)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	product.Num -= num
	if err := productDao.UpdateProduct(s.ProductId, product); err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 訂單刪除
	if err := orderDao.DeleteOrderByOId(s.OrderId, uid); err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 自己商品+1 同商品? 用什麼判斷等
	userProduct := &model.Product{
		Name:          product.Name,
		CategoryId:    product.CategoryId,
		Title:         product.Title,
		Info:          product.Info,
		ImgPath:       product.ImgPath,
		Price:         product.Price,
		DiscountPrice: product.DiscountPrice,
		OnSale:        false,
		Num:           1,
		BossId:        uid,
		BossName:      user.UserName,
		BossAvatar:    user.Avatar,
	}

	if err := productDao.CreateProduct(userProduct); err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	tx.Commit()

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
