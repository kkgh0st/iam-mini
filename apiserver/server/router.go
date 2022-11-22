package server

import (
	"github.com/gin-gonic/gin"
	"iam-mini/apiserver/controller/v1/secret"
	"iam-mini/apiserver/controller/v1/user"
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

		userv1 := v1.Group("/users")
		{
			userController := user.NewUserController(storeIns)
			userv1.POST("", userController.Create)
			userv1.DELETE(":name", userController.Delete)
			userv1.DELETE("", userController.DeleteCollection)
			userv1.PUT(":name/change-password", userController.ChangePassword)
			userv1.PUT(":name", userController.Update)
			userv1.GET(":name", userController.Get)
			userv1.GET("", userController.List)
		}

		secretv1 := v1.Group("/secrets", middlewares.Publish())
		{

			secretController := secret.NewSecretController(storeIns)
			secretv1.POST("", secretController.Create)
			// 这里将其索引作为参数，添加到context中
			secretv1.DELETE(":name", secretController.Delete)
			secretv1.PUT(":name", secretController.Update)
			secretv1.GET(":name", secretController.Get)
			secretv1.GET("", secretController.List)
		}

	}

}
