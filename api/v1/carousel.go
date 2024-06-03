package v1

import (
	"gin_mall_tmp/pkg/util"
	"gin_mall_tmp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListCarousels(c *gin.Context) {
	var listCarousels service.CarouselsService

	if err := c.ShouldBind(&listCarousels); err == nil {
		res := listCarousels.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.Logrus.Infoln("list carousels error", err)
	}
}
