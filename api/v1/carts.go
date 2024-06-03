package v1

import (
	"gin_mall_tmp/pkg/util"
	"gin_mall_tmp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCart(c *gin.Context) {
	var Cart service.CartService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&Cart); err == nil {
		res := Cart.CreateCart(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.Logrus.Infoln("create Cart error", err)
	}
}

func GetCarts(c *gin.Context) {
	var Cart service.CartService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&Cart); err == nil {
		res := Cart.GetCarts(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.Logrus.Infoln("get Cart error", err)
	}
}

func DeleteCart(c *gin.Context) {
	var Cart service.CartService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&Cart); err == nil {
		res := Cart.DeleteCart(c.Request.Context(), c.Param("id"), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.Logrus.Infoln("delete Cart error", err)
	}
}

func UpdateCart(c *gin.Context) {
	var Cart service.CartService
	if err := c.ShouldBind(&Cart); err == nil {
		res := Cart.UpdateCart(c.Request.Context(), c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.Logrus.Infoln("update Cart error", err)
	}
}
