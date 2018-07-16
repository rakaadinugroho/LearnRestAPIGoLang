package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restframework/withgin/model"
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

// from database
func ShowUser(ctx *gin.Context) {
	var user []model.User
	err := model.DB.Find(&user).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string] interface{}{
			"status" : http.StatusInternalServerError,
			"message" : "Gagal",
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string] interface{}{
		"status" : http.StatusOK,
		"data"	: user,
	})
}
// from database with alias
func ShowPosting(ctx *gin.Context) {
	var posting []model.PostItem
	err := model.DB.Find(&posting).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string] interface{}{
			"status" : http.StatusInternalServerError,
			"message" : "Gagal",
		})
		return
	}
	ctx.JSON(http.StatusOK, map[string] interface{}{
		"status" : http.StatusOK,
		"data"	: posting,
	})
}
// function with where
func ShowDetailUser(ctx *gin.Context) {
	id:= ctx.Query("username")
	var user model.User
	err := model.DB.Where("username = ?", id).First(&user).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string] interface{}{
			"status" : http.StatusInternalServerError,
			"message" : "Gagal",
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string] interface{}{
		"status" : http.StatusOK,
		"data"	: user,
	})
}
// function create
func CreateUser(ctx *gin.Context) {
	var user model.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusOK, map[string] interface{} {
			"status" : http.StatusOK,
			"messages" : "gagal",
		})
		return
	}
	err = user.ValidationUser()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string] interface{}{
			"message" : err.Error(),
			"status" : http.StatusInternalServerError,
		})
		return
	}
	err = model.DB.Create(&user).Error
	if err != nil {
		ctx.JSON(500, map[string] interface{} {
			"status" : http.StatusInternalServerError,
			"message" : "error",
		})
		return
	}
	ctx.JSON(http.StatusInternalServerError, map[string] interface{}{
		"status" : http.StatusOK,
		"message" : user,
	})
}
// Sample Migration
func MigrateTable(ctx *gin.Context)  {
	err :=model.DB.AutoMigrate(&model.Activities{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string] interface{}{
			"status" : http.StatusForbidden,
			"message" : "gagal dimigrasi",
		})
		return
	}

	ctx.JSON(http.StatusInternalServerError, map[string] interface{}{
		"status" : http.StatusOK,
		"message" : "berhasil migrasi",
	})

}
// Create Midleware
func CheckHeaderAuthorization(ctx *gin.Context) {
	authorization := ctx.GetHeader("Authorization")
	if authorization != "12345" { //sample with hardcode 12345
		ctx.JSON(http.StatusInternalServerError, map[string] interface{} {
			"message" : "Unautorized",
		})
		ctx.Abort()
	}
	ctx.Next()
}