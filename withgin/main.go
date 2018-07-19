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

	// http client
	route.GET("/grabuser", auth.GrabUser)

	//authentication
	route.POST("/auth", auth.AuthMiddleware.LoginHandler)
	dashboard := route.Group("/dashboard")
	dashboard.Use(auth.AuthMiddleware.MiddlewareFunc())
	{
		dashboard.GET("/hello", auth.SuccessAuthenticator)
		dashboard.GET("/refresh_token", auth.AuthMiddleware.RefreshHandler)
	}
	route.Run(":8089")
}