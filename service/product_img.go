package service

import (
	"context"
	"gin_mall_tmp/dao"
	"gin_mall_tmp/serializer"
	"strconv"
)

type ProductImgService struct{}

func (s *ProductImgService) GetProductImgs(ctx context.Context, pId string) serializer.Response {
	productImgDao := dao.NewProductImgDao(ctx)
	productId, _ := strconv.Atoi(pId)
	productImgs, _ := productImgDao.GetProductImgs(uint(productId))

	return serializer.BuildListResponse(
		serializer.BuildProductImgs(productImgs),
		uint(len(productImgs)),
	)
}
