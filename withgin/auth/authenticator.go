package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/appleboy/gin-jwt"
	"net/http"
	"time"
	"restframework/withgin/model"
)

var AuthMiddleware *jwt.GinJWTMiddleware
func init() {
	AuthMiddleware = &jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte("secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(userId string, password string, c *gin.Context) (interface{}, bool) {
			if (userId == "admin" && password == "admin") || (userId == "test" && password == "test") {
				return &model.UserJWT{
					UserName:  userId,
					LastName:  "Bo-Yi",
					FirstName: "Wu",
				}, true
			}

			return nil, false
		},
		Authorizator: func(user interface{}, c *gin.Context) bool {
			if v, ok := user.(string); ok && v == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup: "header:Authorization",
		TokenHeadName: "Bearer",
		TimeFunc: time.Now,
	}
}

func SuccessAuthenticator(ctx *gin.Context) {
	claims := jwt.ExtractClaims(ctx)
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"userID" : claims["id"],
			"message" : "welcome with JWT",
		},
	)
}




