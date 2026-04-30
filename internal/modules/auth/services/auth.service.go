package services

import (
	"context"

	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/models"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/modules/auth/delivery"
)

type AuthService interface {
	RegisterLocal(ctx context.Context, req *delivery.SingUpRequest) error
	LoginLocal(ctx context.Context, req *delivery.LoginRequest) (*delivery.AuthResponse, error)
	HandleOAuthCallback(ctx context.Context, data *delivery.OAuthUserRequest) (*models.User, error)
	RefreshToken(ctx context.Context, refreshToken string) (*delivery.RefreshTokenResponse, error)
}
