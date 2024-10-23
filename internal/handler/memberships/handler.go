package memberships

import (
	"github.com/AthThobari/simple_music_catalog_go/internal/models/memberships"
	"github.com/gin-gonic/gin"
)

type service interface {
	SignUp(request memberships.SignUpRequest) error
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
	route:=h.Group("/memberships")
	route.POST("sign_up", h.SignUp)
}
