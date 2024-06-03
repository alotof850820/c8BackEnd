package v1

import (
	"gin_mall_tmp/pkg/util"
	"gin_mall_tmp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatOrder(c *gin.Context) {
	var Order service.OrderService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&Order); err == nil {
		res := Order.CreateOrder(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.Logrus.Infoln("create Order error", err)
	}
}

func GetOrders(c *gin.Context) {
	var Order service.OrderService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&Order); err == nil {
		res := Order.GetOrders(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.Logrus.Infoln("get Order error", err)
	}
}

func GetOrder(c *gin.Context) {
	var Order service.OrderService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&Order); err == nil {
		res := Order.GetOrder(c.Request.Context(), claims.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.Logrus.Infoln("update Order error", err)
	}
}

func DeletOrder(c *gin.Context) {
	var Order service.OrderService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&Order); err == nil {
		res := Order.DeleteOrder(c.Request.Context(), c.Param("id"), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.Logrus.Infoln("delete Order error", err)
	}
}
