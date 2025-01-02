package tracks

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Search(c *gin.Context) {
	ctx := c.Request.Context()

	query := c.Query("query")
	pageSizeStr := c.Query("PageSize")
	pageIndexStr := c.Query("PageIndex")
	log.Printf("Received Search request: query=%s, PageSize=%s, PageIndex=%s", query, pageSizeStr, pageIndexStr)

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 10
		log.Printf("Invalid PageSize '%s', defaulting to %d", pageSizeStr, pageSize)
	}

	pageIndex, err := strconv.Atoi(pageIndexStr)
	if err != nil {
		pageIndex = 1
		log.Printf("Invalid PageIndex '%s', defaulting to %d", pageIndexStr, pageIndex)
	}

	response, err := h.service.Search(ctx, query, pageSize, pageIndex)
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
