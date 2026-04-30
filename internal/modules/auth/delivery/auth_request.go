package delivery

import (
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/models"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/common"
)

type SingUpRequest struct {
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"min=8"`
}

func (s *SingUpRequest) ToDomain(hashPass *string, role common.UserRole, provider common.Provider) *models.User {
	return &models.User{
		FullName: s.FullName,
		Email:    s.Email,
		Password: hashPass,
		Role:     role,
		Provider: provider,
	}
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type OAuthUserRequest struct {
	Email      string
	FullName   string
	Provider   common.Provider
	ProviderID string
	AvatarURL  string
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
