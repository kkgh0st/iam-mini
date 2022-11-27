package policy

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"iam-mini/generic/middlewares"
)

func (p *PolicyController) Delete(c *gin.Context) {

	if err := p.srv.Policies().Delete(c, c.GetString(middlewares.UsernameKey), c.Param("name"),
		metav1.DeleteOptions{}); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
