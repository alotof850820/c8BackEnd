package v1

import (
	"gin_mall_tmp/pkg/util"
	"gin_mall_tmp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateFavorite(c *gin.Context) {
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	createFavoriteService := service.FavoriteService{}
	if err := c.ShouldBind(&createFavoriteService); err == nil {
		res := createFavoriteService.CreateFavorite(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.Logrus.Infoln("create Favorite error", err)
	}
}

func GetFavorite(c *gin.Context) {
	getFavoriteService := service.FavoriteService{}
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&getFavoriteService); err == nil {
		res := getFavoriteService.GetFavorite(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.Logrus.Infoln("get Favorite error", err)
	}
}

func DeleteFavorite(c *gin.Context) {
	deleteFavoriteService := service.FavoriteService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteFavoriteService); err == nil {
		res := deleteFavoriteService.DeleteFavorite(c.Request.Context(), c.Param("id"), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.Logrus.Infoln("delete Favorite error", err)
	}
}
