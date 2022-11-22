package secret

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"iam-mini/generic/middlewares"
)

func (s *SecretController) Delete(c *gin.Context) {
	opts := metav1.DeleteOptions{Unscoped: true}
	if err := s.srv.Secrets().Delete(c, c.GetString(middlewares.UsernameKey),
		c.Param("name"), opts); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, "delete successfully!")
}
