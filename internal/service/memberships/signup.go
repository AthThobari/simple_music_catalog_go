package memberships

import (
	"database/sql"
	"errors"

	"github.com/AthThobari/simple_music_catalog_go/internal/models/memberships"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignUp(request memberships.SignUpRequest) error {
	existingUser, err := s.repository.GetUser(request.Email, request.Username, 0)
	if err != nil || err != sql.ErrNoRows {
		log.Error().Err(err).Msg("error get user from database")
		return err
	}

	if existingUser != nil {
		return errors.New("email or username exists")
	}
	pass,err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("error hash password")
		return err
	}

	model := memberships.User{
		Email: request.Email,
		Username: request.Username,
		Password: string(pass),
		CreatedBy: request.Email,
		UpdatedBy: request.Email,
	}
	return s.repository.CreateUser(model)
}