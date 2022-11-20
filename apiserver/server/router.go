package server

import (
	"github.com/gin-gonic/gin"
	"iam-mini/apiserver/controller/v1/secret"
	"iam-mini/generic/middleware"
	"iam-mini/generic/middleware/auth"
)

func initRouter(g *gin.Engine) {
	installMiddleware(g)
	installController(g)

}

func installMiddleware(g *gin.Engine) {

}

// 虚假的Store实例，之后会换成mysql
type DummyStore struct {
}

func installController(g *gin.Engine) {
	jwtStrategy, _ := newJWTAuth().(auth.JWTStrategy)
	g.POST("/login", jwtStrategy.LoginHandler)
	g.POST("/logout", jwtStrategy.LogoutHandler)
	g.POST("/refresh", jwtStrategy.LogoutHandler)

	// 自动认证的，加上这个，意味着我们
	auto := newAutoAuth()

	// version_1
	storeIns := DummyStore{}
	v1 := g.Group("/v1")
	{
		v1.Use(auto.AuthFunc())
		secretv1 := v1.Group("/secrets", middleware.Publish())
		{

			secretController := secret.NewSecretController(storeIns)
			secretv1.POST("", secretController.Create)
		}

	}

}
