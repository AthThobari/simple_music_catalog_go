package memberships

import (
	"fmt"
	"net/http"

	"github.com/AthThobari/simple_music_catalog_go/internal/models/memberships"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	var req memberships.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	accessToken, err := h.service.Login(req)
	if err != nil {
		fmt.Println("test case ter execute")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, memberships.LoginResponse{
		AccessToken: accessToken,
	})
}