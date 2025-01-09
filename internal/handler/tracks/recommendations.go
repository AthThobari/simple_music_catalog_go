package tracks

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetRecommendation(c *gin.Context) {
	ctx := c.Request.Context()

	trackID := c.Query("trackID")
	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	userID := c.GetUint("userID")
	response, err := h.service.GetRecommendation(ctx, userID, limit, trackID)
	if err != nil {
		log.Printf("Error in service.Search: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Printf("Search successful: response=%v", response)
	c.JSON(http.StatusOK, response)
}
