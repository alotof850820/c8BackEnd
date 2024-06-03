package service

import (
	"context"
	"gin_mall_tmp/dao"
	"gin_mall_tmp/model"
	"gin_mall_tmp/pkg/e"
	"gin_mall_tmp/pkg/util"
	"gin_mall_tmp/serializer"
	"mime/multipart"
	"strconv"
	"sync"
)

type ProductService struct {
	Id            uint   `json:"id" form:"id"`
	Name          string `json:"name" form:"name"`
	CategoryId    uint   `json:"category_id" form:"category_id"`
	Title         string `json:"title" form:"title"`
	Info          string `json:"info" form:"info"`
	ImgPath       string `json:"img_path" form:"img_path"`
	Price         string `json:"price" form:"price"`
	DiscountPrice string `json:"discount_price" form:"discount_price"`
	OnSale        bool   `json:"on_sale" form:"on_sale"`
	Num           int    `json:"num" form:"num"`
	model.BasePage
}

func (service *ProductService) CreateProduct(ctx context.Context, uId uint, files []*multipart.FileHeader) serializer.Response {
	var boss *model.User
	var err error
	code := e.Success
	bossDao := dao.NewUserDao(ctx)
	boss, _ = bossDao.GetUserById(uId)

	// 以第一張作為封面圖片
	tmp, _ := files[0].Open()
	path, err := UploadProductToLocalStatic(tmp, uId, service.Name)
	if err != nil {
		code = e.ErrorProductImgUpload
		util.Logrus.Infoln("product img upload error", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 定義商品
	product := &model.Product{
		Name:          service.Name,
		CategoryId:    service.CategoryId,
		Title:         service.Title,
		Info:          service.Info,
		ImgPath:       path,
		Price:         service.Price,
		DiscountPrice: service.DiscountPrice,
		OnSale:        true,
		Num:           service.Num,
		BossId:        uId,
		BossName:      boss.UserName,
		BossAvatar:    boss.Avatar,
	}

	// 新增商品至資料庫
	productDao := dao.NewProductDao(ctx)
	err = productDao.CreateProduct(product)
	if err != nil {
		code = e.Error
		util.Logrus.Infoln("create product error", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 併發創建
	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	for index, file := range files {
		num := strconv.Itoa(index)
		productImgDao := dao.NewProductImgDaoByDB(productDao.DB)
		tmp, _ = file.Open()
		path, err = UploadProductToLocalStatic(tmp, uId, service.Name+num)
		if err != nil {
			code = e.ErrorProductImgUpload
			util.Logrus.Infoln("product img upload error", err)
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		productImg := &model.ProductImg{
			ProductId: product.ID,
			ImgPath:   path,
		}

		// 新增商品圖片至資料庫
		err = productImgDao.CreateProductImg(productImg)
		if err != nil {
			code = e.Error
			util.Logrus.Infoln("create product img error", err)
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		wg.Done()
	}
	wg.Wait()

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProduct(product),
	}
}

func (service *ProductService) GetProducts(ctx context.Context) serializer.Response {
	var products []*model.Product
	var err error
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15 // 默認每頁15筆
	}
	condition := make(map[string]interface{})
	// 找分類
	if service.CategoryId != 0 {
		condition["category_id"] = service.CategoryId
	}

	productDao := dao.NewProductDao(ctx)
	total, err := productDao.GetProductTotalByCondition(condition) //耗效能
	if err != nil {
		code = e.Error
		util.Logrus.Infoln("get product total error", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		productDao := dao.NewProductDaoByDb(productDao.DB)
		products, _ = productDao.GetProductsByCondition(condition, service.BasePage) //找每頁商品
		wg.Done()
	}()
	wg.Wait()
	return serializer.BuildListResponse(
		serializer.BuildProducts(products),
		uint(total),
	)
}

func (service *ProductService) SearchProducts(ctx context.Context) serializer.Response {

	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15 // 默認每頁15筆
	}
	productDao := dao.NewProductDao(ctx)
	products, total, err := productDao.SearchProducts(service.Info, service.BasePage)
	if err != nil {
		code = e.Error
		util.Logrus.Infoln("get product error", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(
		serializer.BuildProducts(products),
		uint(total),
	)
}

func (service *ProductService) GetProduct(ctx context.Context, id string) serializer.Response {
	var product *model.Product
	var err error
	code := e.Success
	pId, _ := strconv.Atoi(id)
	productDao := dao.NewProductDao(ctx)
	product, err = productDao.GetProductById(uint(pId))
	if err != nil {
		code = e.Error
		util.Logrus.Infoln("get product error", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProduct(product),
	}
}
