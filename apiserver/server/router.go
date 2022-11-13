package server

import (
	"github.com/gin-gonic/gin"
	"iam-mini/generic/middleware/auth"
)

func initRouter(g *gin.Engine) {
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

}
