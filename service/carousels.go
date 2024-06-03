package service

import (
	"context"
	"gin_mall_tmp/dao"
	"gin_mall_tmp/pkg/e"
	"gin_mall_tmp/pkg/util"
	"gin_mall_tmp/serializer"
)

type CarouselsService struct{}

func (service *CarouselsService) List(ctx context.Context) serializer.Response {
	carouselsDao := dao.NewCarouselsDao(ctx)
	code := e.Success

	carousels, err := carouselsDao.GetListCarousels()
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
		serializer.BuildCarousels(carousels),
		uint(len(carousels)),
	)
}
