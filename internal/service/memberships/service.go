package memberships

import (
	"github.com/AthThobari/simple_music_catalog_go/internal/configs"
	"github.com/AthThobari/simple_music_catalog_go/internal/models/memberships"
)

type repository interface {
	CreateUser(model memberships.User) error
	GetUser(email, username string, id uint) (*memberships.User, error)
}

type service struct {
	cfg        *configs.Config
	repository repository
}

func NewService(cfg *configs.Config, repository repository) *service {
	return &service{
		cfg: cfg,
		repository: repository,
	}
}
