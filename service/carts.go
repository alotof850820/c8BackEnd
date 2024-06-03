package service

import (
	"context"
	"gin_mall_tmp/dao"
	"gin_mall_tmp/model"
	"gin_mall_tmp/pkg/e"
	"gin_mall_tmp/serializer"
	"strconv"
)

type CartService struct {
	Id        uint `form:"id" json:"id"`
	BossId    uint `form:"boss_id" json:"boss_id"`
	ProductId uint `form:"product_id" json:"product_id"`
	Num       int  `form:"num" json:"num"`
}

func (c *CartService) CreateCart(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	// 判斷有無此商品
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(c.ProductId)

	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	cartDao := dao.NewCartDao(ctx)
	cart := &model.Cart{
		UserId:    uid,
		BossId:    c.BossId,
		ProductId: c.ProductId,
		Num:       uint(c.Num),
	}

	err = cartDao.CreateCart(cart)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	userDao := dao.NewUserDao(ctx)
	boss, err := userDao.GetUserById(c.BossId)

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
		Data:   serializer.BuildCart(cart, product, boss),
	}
}

func (c *CartService) GetCarts(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	CartDao := dao.NewCartDao(ctx)
	Cart, err := CartDao.GetCartesByUserId(uid)

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
		Data:   serializer.BuildCarts(ctx, Cart),
	}
}

func (c *CartService) DeleteCart(ctx context.Context, cid string, uid uint) serializer.Response {
	cartId, _ := strconv.Atoi(cid)
	code := e.Success
	CartDao := dao.NewCartDao(ctx)

	err := CartDao.DeleteCartByCId(uint(cartId), uid)

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

func (c *CartService) UpdateCart(ctx context.Context, cid string) serializer.Response {
	code := e.Success
	cartDao := dao.NewCartDao(ctx)
	cartId, _ := strconv.Atoi(cid)

	err := cartDao.UpdateCartNumByCId(uint(cartId), c.Num)

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
