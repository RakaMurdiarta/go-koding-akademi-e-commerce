package delivery

import "github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/common"

type AuthResponse struct {
	AccessToken  string          `json:"acc_token"`
	RefreshToken string          `json:"refresh_token"`
	Role         common.UserRole `json:"role"`
}

type RefreshTokenResponse struct {
	NewAccessToken string `json:"new_access_token"`
}
