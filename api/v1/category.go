package v1

import (
	"gin_mall_tmp/pkg/util"
	"gin_mall_tmp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	var getCategoriesService = service.CategoryService{}
	if err := c.ShouldBind(&getCategoriesService); err == nil {
		res := getCategoriesService.GetCategories(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.Logrus.Infoln("get categories error", err)
	}
}
