package memberships

import (
	"net/http"

	"github.com/AthThobari/simple_music_catalog_go/internal/models/memberships"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
	var req memberships.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.SignUp(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusCreated)
}
