package serializer

import (
	"gin_mall_tmp/conf"
	"gin_mall_tmp/model"
)

type ProductImg struct {
	ProductId uint   `json:"product_id"`
	ImgPath   string `json:"img_path"`
}

func BuildProductImg(item *model.ProductImg) ProductImg {
	return ProductImg{
		ProductId: item.ProductId,
		ImgPath:   conf.Host + conf.HttpPort + conf.ProductPath + item.ImgPath,
	}
}
func BuildProductImgs(items []*model.ProductImg) (productsImg []ProductImg) {
	for _, v := range items {
		productsImg = append(productsImg, BuildProductImg(v))
	}
	return
}
