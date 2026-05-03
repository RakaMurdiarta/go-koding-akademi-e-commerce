package delivery

import "github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/common"

type AuthResponse struct {
	AccessToken  string          `json:"acc_token"`
	RefreshToken string          `json:"refresh_token"`
	Role         common.UserRole `json:"role"`
}

type UserRegister struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	CreatedAt string `json:"createdAt"`
}

type RegisterResponse struct {
	User  UserRegister `json:"user"`
	Token string       `json:"token"`
}

type LoginResponse struct {
	User  UserRegister `json:"user"`
	Token string       `json:"token"`
}

type RefreshTokenResponse struct {
	NewAccessToken string `json:"new_access_token"`
}
