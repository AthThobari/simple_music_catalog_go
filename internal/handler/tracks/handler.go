package tracks

import (
	"context"

	"github.com/AthThobari/simple_music_catalog_go/internal/models/spotify"
	"github.com/gin-gonic/gin"
)


//go:generate mockgen -source=handler.go -destination=handler_mock_test.go -package=tracks
type service interface {
	Search(ctx context.Context, query string, pageSize, pageIndex int) (*spotify.SearchResponse, error)
}

type Handler struct {
	*gin.Engine
	service service
}

func NewHandler(api *gin.Engine, service service) *Handler {
	return &Handler{
		api,
		service,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("/tracks")
	route.GET("/search", h.Search)
}