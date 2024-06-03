package routes

import (
	api "gin_mall_tmp/api/v1"
	"gin_mall_tmp/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 初始化路由
func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())
	// 加載靜態檔案
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "欢迎来到mall后台管理系统",
			})
		})
		// 用戶操作
		v1.POST("/user/register", api.UserRegister)
		v1.POST("/user/login", api.UserLogin)

		// 輪波圖操作
		v1.POST("/carousels", api.ListCarousels)

		// 商品操作
		v1.GET("getProducts", api.GetProducts)
		v1.GET("getProduct/:id", api.GetProduct)
		v1.GET("getProductImgs/:id", api.GetProductImgs)
		v1.GET("getCategories", api.GetCategories)

		// 需要登入保護
		authApi := v1.Group("/")
		authApi.Use(middleware.JWT())
		{
			// 用戶操作
			authApi.PUT("user", api.UserUpdate)
			authApi.POST("avatar", api.UserUploadAvatar)
			authApi.POST("user/send-email", api.UserSendEmail)
			authApi.POST("user/valid-email", api.UserValidEmail)

			// 顯示金額
			authApi.POST("money", api.ShowMoney)

			// 商品操作
			authApi.POST("createProduct", api.CreateProduct)
			authApi.POST("searchProducts", api.SearchProducts)

			// 收藏夾操作
			authApi.GET("getFavorite", api.GetFavorite)
			authApi.POST("createFavorite", api.CreateFavorite)
			authApi.DELETE("deleteFavorite/:id", api.DeleteFavorite)

			// 地址操作
			authApi.POST("createAddress", api.CreateAddress)
			authApi.GET("getAddresses", api.GetAddresses)
			authApi.GET("getAddress/:id", api.GetAddress)
			authApi.PUT("updateAddress/:id", api.UpdateAddress)
			authApi.DELETE("deleteAddress/:id", api.DeleteAddress)

			// 購物車操作
			authApi.POST("createCart", api.CreateCart)
			authApi.GET("getCarts", api.GetCarts)
			authApi.PUT("updateCart/:id", api.UpdateCart)
			authApi.DELETE("deleteCart/:id", api.DeleteCart)

			// 訂單操作
			authApi.POST("creatOrder", api.CreatOrder)
			authApi.GET("getOrders", api.GetOrders)
			authApi.GET("getOrder/:id", api.GetOrder)
			authApi.DELETE("deletOrder/:id", api.DeletOrder)

			// 支付操作
			authApi.POST("orderPay", api.OrderPay)

		}
	}

	return r
}
