package memberships

import "gorm.io/gorm"

type (
	User struct {
		gorm.Model
		Email     string `gorm:"unique;not null"`
		Username  string `gorm:"unique;not null"`
		Password  string `gorm:"not null"`
		CreatedBy string `gorm:"not null"`
		UpdatedBy string `gorm:"not null"`
	}
)

type (
	SignUpRequest struct {
		Email string `json:"email"`
		Password string `json:"password"`
		Username string `json:"username"`
	}

	LoginRequest struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}
)

type (
	LoginResponse struct {
		AccessToken string `json:"accessToken"`
	}
)
