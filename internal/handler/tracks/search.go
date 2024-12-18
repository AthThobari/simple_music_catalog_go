package tracks

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Search(c *gin.Context) {
	ctx := c.Request.Context()

	query := c.Query("query")
	pageSizeStr := c.Query("PageSize")
	pageIndexStr := c.Query("PageIndex")

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 10
	}

	pageIndex, err := strconv.Atoi(pageIndexStr)
	if err != nil {
		pageIndex = 1
	}

	response, err := h.service.Search(ctx, query, pageSize, pageIndex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}