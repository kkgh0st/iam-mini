package server

import (
	"github.com/gin-gonic/gin"
	"iam-mini/apiserver/controller/v1/secret"
	"iam-mini/apiserver/store/mysql"
	"iam-mini/generic/middlewares"
	"iam-mini/generic/middlewares/auth"
)

func InitRouter(g *gin.Engine) {
	installMiddleware(g)
	installController(g)

}

func installMiddleware(g *gin.Engine) {

}

func installController(g *gin.Engine) {
	jwtStrategy, _ := newJWTAuth().(auth.JWTStrategy)
	g.POST("/login", jwtStrategy.LoginHandler)
	g.POST("/logout", jwtStrategy.LogoutHandler)
	g.POST("/refresh", jwtStrategy.LogoutHandler)

	// 自动认证的，加上这个，意味着我们
	auto := newAutoAuth()

	// version_1
	storeIns, err := mysql.GetMySQLFactoryOr(nil)
	if err != nil {
		return
	}
	v1 := g.Group("/v1")
	{
		v1.Use(auto.AuthFunc())
		secretv1 := v1.Group("/secrets", middlewares.Publish())
		{

			secretController := secret.NewSecretController(storeIns)
			secretv1.POST("", secretController.Create)
		}

	}

}
