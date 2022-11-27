package middlewares

import "github.com/gin-gonic/gin"

const UsernameKey = "admin"

// Context is a middleware that injects common prefix fields to gin.Context.
func Context() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*				c.Set(log.KeyRequestID, c.GetString(XRequestIDKey))
						c.Set(log.KeyUsername, c.GetString(UsernameKey))*/
		c.Next()
	}
}
