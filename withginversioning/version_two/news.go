package version_two

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func ViewDetail(context *gin.Context) {
	context.JSON(http.StatusOK, "view detail version 2")
}

