package secret

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"iam-mini/generic/middlewares"
)

func (s *SecretController) Get(c *gin.Context) {
	secret, err := s.srv.Secrets().Get(c, c.GetString(middlewares.UsernameKey),
		c.Param("name"), metav1.GetOptions{})
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, secret)
}
