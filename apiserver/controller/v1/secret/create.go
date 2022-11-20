package secret

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *SecretController) Create(c *gin.Context) {
	var r Secret
	if err := c.ShouldBind(&r); err != nil {
		c.JSON(http.StatusBadRequest, "error")
	}

	c.JSON(http.StatusOK, "Good")
	return
}
