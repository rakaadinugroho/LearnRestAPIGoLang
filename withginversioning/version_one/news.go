package version_one

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ViewDetail(context *gin.Context) {
	context.JSON(http.StatusOK, "view detail version 1")
}
