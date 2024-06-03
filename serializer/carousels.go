package serializer

import "gin_mall_tmp/model"

type Carousel struct {
	Id        uint   `json:"id"`
	ImgPath   string `json:"img_path"`
	ProductId uint   `json:"product_id"`
	CreateAt  int64  `json:"create_at"`
}

// 單一
func BuildCarousel(item *model.Carousel) Carousel {
	return Carousel{
		Id:        item.ID,
		ImgPath:   item.ImgPath,
		ProductId: item.ProductId,
		CreateAt:  item.CreatedAt.Unix(),
	}
}

// 多個
func BuildCarousels(items []model.Carousel) (carousels []Carousel) {
	for _, item := range items {
		carousels = append(carousels, BuildCarousel(&item))
	}
	return carousels
}
