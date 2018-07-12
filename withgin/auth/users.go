package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type FormUser struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type RawUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(context *gin.Context) { // Test Post
	username:= context.PostForm("username")
	context.String(http.StatusOK, "hello lagi %s", username)
}

func RegisterForm(context *gin.Context) {
	var user FormUser
	data := context.ShouldBind(&user)
	if data != nil {
		context.String(http.StatusInternalServerError, "invalid gaes")
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"data" : data,
	})
}

func Profile(context *gin.Context) { // Test Get
	username := context.Query("username")
	context.String(http.StatusOK, "Hello %s ", username)
}
func Category(context *gin.Context) { // Test from Path
	username := context.Param("username")
	context.String(http.StatusOK, "Hello %s", username)
}