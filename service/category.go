package service

import (
	"context"
	"gin_mall_tmp/dao"
	"gin_mall_tmp/pkg/e"
	"gin_mall_tmp/serializer"
)

type CategoryService struct{}

func (s *CategoryService) GetCategories(ctx context.Context) serializer.Response {
	code := e.Success
	categoryDao := dao.NewCategoryDao(ctx)
	categories, err := categoryDao.GetCategories()
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(
		serializer.BuildCategories(categories),
		uint(len(categories)),
	)
}
