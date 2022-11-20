package auth

import (
	"github.com/gin-gonic/gin"
	"iam-mini/generic/middleware"
	"net/http"
	"strings"
)

const authHeadCount = 2

type AutoStrategy struct {
	basic middleware.AuthStrategy
	jwt   middleware.AuthStrategy
}

var _ middleware.AuthStrategy = &AutoStrategy{}

func NewAutoStrategy(basic, jwt middleware.AuthStrategy) AutoStrategy {
	return AutoStrategy{
		basic: basic,
		jwt:   jwt,
	}
}

func (a AutoStrategy) AuthFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		operator := middleware.AuthOperator{}
		authHeader := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

		if len(authHeader) != authHeadCount {
			c.JSON(http.StatusBadRequest, "Authorization format is not satisfied")
			c.Abort() // 直接丢弃，不进行下一步操作
			return
		}

		switch authHeader[0] {
		case "Basic":
			operator.SetStrategy(a.basic)
		case "Bearer":
			operator.SetStrategy(a.jwt)
		default:
			c.JSON(http.StatusBadRequest, "Authorization format is not satisfied")
			c.Abort() // 直接丢弃，不进行下一步操作
			return
		}

		operator.AuthFunc()(c)

		c.Next()
	}
}
