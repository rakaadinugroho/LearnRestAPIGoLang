package main
// Sample GoLang with Gin
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restframework/withgin/auth"
)

func main() {
	route := gin.Default()
	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"username" : "rakaadinugroho",
		})
	})
	route.GET("/profile", auth.Profile)
	route.POST("/register", auth.Register)
	route.POST("/registerform", auth.RegisterForm)
	route.GET("/profile/:username", auth.Category)
	route.Run(":8089")
}