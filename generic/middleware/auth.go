package middleware

import "github.com/gin-gonic/gin"

// 每个认证策略，都有一个其对应的接口，这本身一点也不难
type AuthStrategy interface {
	AuthFunc() gin.HandlerFunc
}

type AuthOperator struct {
	strategy AuthStrategy
}

func (operator *AuthOperator) SetStrategy(strategy AuthStrategy) {
	operator.strategy = strategy
}

func (operator *AuthOperator) AuthFunc() gin.HandlerFunc {
	return operator.strategy.AuthFunc()
}
