package v1

import (
	"gin_mall_tmp/pkg/util"
	"gin_mall_tmp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProductImgs(c *gin.Context) {
	var productImgService service.ProductImgService

	if err := c.ShouldBind(&productImgService); err == nil {
		res := productImgService.GetProductImgs(c.Request.Context(), c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.Logrus.Infoln("get product img error", err)
	}
}
