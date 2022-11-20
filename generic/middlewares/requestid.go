package middlewares

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

const (
	// XRequestIDKey defines X-Request-ID key string.
	XRequestIDKey = "X-Request-ID"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		//  这个
		rid := c.GetHeader(XRequestIDKey)

		if rid == "" {
			rid = uuid.Must(uuid.NewV4(), nil).String()
			// Request头部
			// Request.Header
			c.Request.Header.Set(XRequestIDKey, rid)
			// 这个是其context
			// c.Context
			c.Set(XRequestIDKey, rid)
		}

		// 头部
		// Writer.ResponseWriter.HeaderHeader，这其实一点也
		c.Writer.Header().Set(XRequestIDKey, rid)
		c.Next()
	}
}
