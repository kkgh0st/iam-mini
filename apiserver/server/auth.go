package server

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	ds "iam-mini/apiserver/datadef/v1"
	"iam-mini/generic/middlewares"
	"iam-mini/generic/middlewares/auth"
	"net/http"
	"time"
)

type loginInfo struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func newJWTAuth() middlewares.AuthStrategy {

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

func parseWithBody(c *gin.Context) (loginInfo, error) {
	var login loginInfo
	if err := c.ShouldBindJSON(&login); err != nil {
		return loginInfo{}, jwt.ErrFailedAuthentication
	}

	return login, nil
}

func authenticator() func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		var login loginInfo
		var err error
		// 这里直接解析相关头部
		if c.Request.Header.Get("Authorization") != "" {
			// 解析头
		} else {
			// 解析Body部分
			login, err = parseWithBody(c)
		}
		if err != nil {
			return "", jwt.ErrFailedAuthentication
		}

		/*
			这里调用Secure
		*/
		if login.Password == "admin" && login.Username == "admin" {
			// pass
		} else {
			return "", jwt.ErrFailedAuthentication
		}

		// 返回确切时间
		return &ds.User{
			Username:  login.Username,
			Password:  login.Password,
			LoginedAt: time.Now()}, nil
	}
}

func loginResponse() func(c *gin.Context, code int, token string, expire time.Time) {
	return nil
}

func newAutoAuth() middlewares.AuthStrategy {
	// newBasicAuth().(auth.BasicStrategy)，Basic这个策略我们暂时先不支持
	return auth.NewAutoStrategy(nil, newJWTAuth().(auth.JWTStrategy))
}

func refreshResponse() func(c *gin.Context, code int, token string, expire time.Time) {
	return nil
}
