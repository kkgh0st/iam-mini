package middleware

import "github.com/gin-gonic/gin"

func Publish() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		/*		if c.Writer.Status() != http.StatusOK {
					log.L(c).Debugf("request failed with http status code `%d`, ignore publish message", c.Writer.Status())

					return
				}

				var resource string

				pathSplit := strings.Split(c.Request.URL.Path, "/")
				if len(pathSplit) > 2 {
					resource = pathSplit[2]
				}

				method := c.Request.Method

				switch resource {
				case "policies":
					notify(c, method, load.NoticePolicyChanged)
				case "secrets":
					notify(c, method, load.NoticeSecretChanged)
				default:
				}*/
	}
}
