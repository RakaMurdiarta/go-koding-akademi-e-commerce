package handler

import (
	"errors"
	"net/http"

	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/modules/auth/delivery"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/modules/auth/services"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/common"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/common/customs"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/common/response"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/config"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/shared"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
)

type authHandler struct {
	as     services.AuthService
	v      *validator.Validate
	google services.OauthService
	github services.OauthService
	conf   *config.Config
}

func NewAuthHandler(s services.AuthService, v *validator.Validate, google services.OauthService, github services.OauthService, conf *config.Config) *authHandler {
	return &authHandler{
		as:     s,
		v:      v,
		google: google,
		conf:   conf,
		github: github,
	}
}

func (a *authHandler) SignUpLocal(c *echo.Context) error {
	req := new(delivery.SingUpRequest)

	//bind
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponseError("Validation Failed", customs.HandleBindError(err)...))
	}

	//validate
	if err := a.v.StructCtx(c.Request().Context(), req); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponseError("Validation Failed", *customs.NewErrorValue("validation ", err.Error())))

	}

	if err := a.as.RegisterLocal(c.Request().Context(), req); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponseError(err.Error(), *customs.NewErrorValue("bussines_logic", err.Error())))

	}

	return c.JSON(http.StatusCreated, response.NewResponseSuccess(req.Email, "Account Created"))
}

func (a *authHandler) Login(c *echo.Context) error {

	req := new(delivery.LoginRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponseError("Validation Failed", customs.HandleBindError(err)...))
	}

	//validate
	if err := a.v.StructCtx(c.Request().Context(), req); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponseError("Validation Failed", *customs.NewErrorValue("validation ", err.Error())))

	}

	res, err := a.as.LoginLocal(c.Request().Context(), req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponseError(err.Error(), *customs.NewErrorValue("bussines_logic", err.Error())))
	}

	return c.JSON(http.StatusOK, response.NewResponseSuccess(res, "Login Successfully"))

}

func (a *authHandler) LoginGoogle(c *echo.Context) error {
	url := a.google.LoginRedirectUrl()
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func (a *authHandler) LoginGithub(c *echo.Context) error {
	url := a.github.LoginRedirectUrl()
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func (a *authHandler) GoogleCallback(c *echo.Context) error {
	code := c.QueryParam("code")
	providerResponse, err := a.google.Callback(c.Request().Context(), code)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponseError(err.Error(), *customs.NewErrorValue("bussines_logic", err.Error())))
	}

	user, err := a.as.HandleOAuthCallback(c.Request().Context(), &delivery.OAuthUserRequest{
		AvatarURL:  providerResponse.Picture,
		Email:      providerResponse.Email,
		Provider:   common.ProviderGoogle,
		ProviderID: providerResponse.Sub,
		FullName:   providerResponse.Name,
	})

	token, err := shared.GenerateToken(uint(user.ID), user.Role, a.conf.JwtSecretKey, 24)
	if err != nil {
		return errors.New("failed to generate token")
	}

	refreshToken, err := shared.GenerateToken(uint(user.ID), user.Role, a.conf.JwtSecretKey, 168)
	if err != nil {
		return errors.New("failed to generate refresh token")
	}

	resp := &delivery.AuthResponse{
		AccessToken:  token,
		RefreshToken: refreshToken,
		Role:         user.Role,
	}

	return c.JSON(http.StatusOK, response.NewResponseSuccess(resp, "Login Successfully"))
}

func (a *authHandler) GithubCallback(c *echo.Context) error {
	code := c.QueryParam("code")
	providerResponse, err := a.github.Callback(c.Request().Context(), code)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponseError(err.Error(), *customs.NewErrorValue("bussines_logic", err.Error())))
	}

	user, err := a.as.HandleOAuthCallback(c.Request().Context(), &delivery.OAuthUserRequest{
		AvatarURL:  providerResponse.Picture,
		Email:      providerResponse.Email,
		Provider:   common.ProviderGithub,
		ProviderID: providerResponse.Sub,
		FullName:   providerResponse.Name,
	})

	token, err := shared.GenerateToken(uint(user.ID), user.Role, a.conf.JwtSecretKey, 24)
	if err != nil {
		return errors.New("failed to generate token")
	}

	refreshToken, err := shared.GenerateToken(uint(user.ID), user.Role, a.conf.JwtSecretKey, 168)
	if err != nil {
		return errors.New("failed to generate refresh token")
	}

	resp := &delivery.AuthResponse{
		AccessToken:  token,
		RefreshToken: refreshToken,
		Role:         user.Role,
	}

	return c.JSON(http.StatusOK, response.NewResponseSuccess(resp, "Login Successfully"))
}

func (a *authHandler) RefreshToken(c *echo.Context) error {
	req := new(delivery.RefreshTokenRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponseError("Validation Failed", customs.HandleBindError(err)...))
	}

	if err := a.v.StructCtx(c.Request().Context(), req); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponseError("Validation Failed", *customs.NewErrorValue("validation ", err.Error())))

	}

	res, err := a.as.RefreshToken(c.Request().Context(), req.RefreshToken)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponseError(err.Error(), *customs.NewErrorValue("business_logic", err.Error())))
	}

	return c.JSON(http.StatusOK, response.NewResponseSuccess(res, "Refresh Token Fetched"))
}
