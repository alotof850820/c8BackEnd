package service

import (
	"context"
	"gin_mall_tmp/dao"
	"gin_mall_tmp/model"
	"gin_mall_tmp/pkg/e"
	"gin_mall_tmp/pkg/util"
	"gin_mall_tmp/serializer"
	"strconv"
)

type FavoriteService struct {
	ProductId  uint `form:"product_id" json:"product_id"`
	BossId     uint `form:"boss_id" json:"boss_id"`
	FavoriteId uint `form:"favorite_id" json:"favorite_id"`
	model.BasePage
}

func (f *FavoriteService) CreateFavorite(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	favoriteDao := dao.NewFavoriteDao(ctx)
	// 判斷是否已經收藏
	exist, _ := favoriteDao.FavoriteExistOrNot(f.ProductId, uId)
	if exist {

		code = e.ErrorFavoriteExist
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}

	}

	// 拿user信息
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		util.Logrus.Infoln(err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 拿boss信息
	bossDao := dao.NewUserDao(ctx)
	boss, err := bossDao.GetUserById(f.BossId)
	if err != nil {
		util.Logrus.Infoln(err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 拿product信息
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(f.ProductId)
	if err != nil {
		util.Logrus.Infoln(err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 新增收藏
	favorite := &model.Favorite{
		UserId:    uId,
		ProductId: f.ProductId,
		BossId:    f.BossId,
		User:      *user,
		Product:   *product,
		Boss:      *boss,
	}

	err = favoriteDao.CreateFavorite(favorite)
	if err != nil {
		util.Logrus.Infoln(err)
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

func (f *FavoriteService) GetFavorite(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	favoriteDao := dao.NewFavoriteDao(ctx)
	favorite, err := favoriteDao.GetFavorite(uId)
	if err != nil {
		util.Logrus.Infoln(err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(
		serializer.BuildFavorites(ctx, favorite),
		uint(len(favorite)),
	)
}

func (f *FavoriteService) DeleteFavorite(ctx context.Context, fId string, uId uint) serializer.Response {
	code := e.Success
	favoriteDao := dao.NewFavoriteDao(ctx)
	favoriteId, _ := strconv.Atoi(fId)
	err := favoriteDao.DeleteFavorite(uId, uint(favoriteId))
	if err != nil {
		util.Logrus.Infoln(err)
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
