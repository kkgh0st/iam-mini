package auth

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"iam-mini/generic/middleware"
	"iam-mini/generic/middleware/auth"
	"net/http"
	"time"
)

func NewJWTAuth() middleware.AuthStrategy {

	ginjwt, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            viper.GetString("jwt.Realm"),
		SigningAlgorithm: "HS256",
		Key:              []byte(viper.GetString("jwt.key")),
		Timeout:          viper.GetDuration("jwt.timeout"),
		MaxRefresh:       viper.GetDuration("jwt.max-refresh"),
		Authenticator:    authenticator(),
		LoginResponse:    loginResponse(),
		LogoutResponse: func(c *gin.Context, code int) {
			c.JSON(http.StatusOK, nil)
		},
		RefreshResponse: refreshResponse(),
		PayloadFunc:     payloadFunc(),
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return claims[jwt.IdentityKey]
		},
		IdentityKey:  "username",
		Authorizator: authorizator(),
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		SendCookie:    true,
		TimeFunc:      time.Now,
	})

	return auth.NewJWTStrategy(*ginjwt)
}

func authorizator() func(data interface{}, c *gin.Context) bool {
	return nil
}

func payloadFunc() func(data interface{}) jwt.MapClaims {
	return nil
}

func authenticator() func(c *gin.Context) (interface{}, error) {
	return nil
}

func loginResponse() func(c *gin.Context, code int, token string, expire time.Time) {
	return nil
}
func refreshResponse() func(c *gin.Context, code int, token string, expire time.Time) {
	return nil
}
