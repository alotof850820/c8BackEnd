package serializer

import (
	"context"
	"gin_mall_tmp/conf"
	"gin_mall_tmp/dao"
	"gin_mall_tmp/model"
)

type Favorite struct {
	UserId        uint   `json:"user_id"`
	ProductId     uint   `json:"product_id"`
	CreatedAt     int64  `json:"created_at"`
	Name          string `json:"name"`
	CategoryId    uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	OnSale        bool   `json:"on_sale"`
	Num           int    `json:"num"`
	BossId        uint   `json:"boss_id"`
}

func BuildFavorite(favorite *model.Favorite, product *model.Product, boss *model.User) Favorite {
	return Favorite{
		UserId:        favorite.UserId,
		ProductId:     favorite.ProductId,
		CreatedAt:     favorite.CreatedAt.Unix(),
		Name:          product.Name,
		CategoryId:    product.CategoryId,
		Title:         product.Title,
		Info:          product.Info,
		ImgPath:       conf.Host + conf.HttpPort + conf.ProductPath + product.ImgPath,
		Price:         product.Price,
		DiscountPrice: product.DiscountPrice,
		OnSale:        product.OnSale,
		Num:           product.Num,
		BossId:        boss.ID,
	}
}

func BuildFavorites(ctx context.Context, items []*model.Favorite) (favorites []Favorite) {
	productDao := dao.NewProductDao(ctx)
	// 查boss信息
	bossDao := dao.NewUserDao(ctx)
	for _, item := range items {
		product, err := productDao.GetProductById(item.ProductId)
		if err != nil {
			continue
		}
		boss, err := bossDao.GetUserById(item.UserId)
		if err != nil {
			continue
		}
		favorite := BuildFavorite(item, product, boss)
		favorites = append(favorites, favorite)
	}
	return favorites
}
