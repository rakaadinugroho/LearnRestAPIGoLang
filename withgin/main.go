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
	route.GET("/showprofile", auth.ShowUser)
	route.GET("/showposting", auth.ShowPosting)
	route.GET("/detailprofile", auth.ShowDetailUser)
	//create user (from raw)
	route.POST("/createprofile", auth.CreateUser)
	// Migration
	route.GET("/migratedb", auth.MigrateTable)
	route.Run(":8089")
}