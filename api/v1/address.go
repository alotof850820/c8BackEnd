package v1

import (
	"gin_mall_tmp/pkg/util"
	"gin_mall_tmp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAddress(c *gin.Context) {
	var address service.AddressService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&address); err == nil {
		res := address.CreateAddress(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.Logrus.Infoln("create address error", err)
	}
}

func GetAddresses(c *gin.Context) {
	var address service.AddressService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&address); err == nil {
		res := address.GetAddresses(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.Logrus.Infoln("get addresses error", err)
	}
}

func GetAddress(c *gin.Context) {
	var address service.AddressService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&address); err == nil {
		res := address.GetAddress(c.Request.Context(), c.Param("id"), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.Logrus.Infoln("get address error", err)
	}
}

func UpdateAddress(c *gin.Context) {
	var address service.AddressService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&address); err == nil {
		res := address.UpdateAddress(c.Request.Context(), c.Param("id"), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.Logrus.Infoln("update address error", err)
	}
}

func DeleteAddress(c *gin.Context) {
	var address service.AddressService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&address); err == nil {
		res := address.DeleteAddress(c.Request.Context(), c.Param("id"), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.Logrus.Infoln("delete address error", err)
	}
}
