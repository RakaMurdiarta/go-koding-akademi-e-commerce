package provider

import (
	auth_handler "github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/modules/auth/handler"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/modules/auth/services"
	ImplService "github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/modules/auth/services/impl"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/modules/users/repository"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/config"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
)

func AuthProvide(
	privateRoute *echo.Group,
	publicRoute *echo.Group,
	conf *config.Config,
	userRepo repository.UserRepository,
	authService services.AuthService,

) {

	v := validator.New()

	google := ImplService.NewGoogleProvider(userRepo, conf)
	github := ImplService.NewGithubProvider(userRepo, conf)

	authHandler := auth_handler.NewAuthHandler(authService, v, google, github, conf)

	auth := publicRoute.Group("/auth")
	auth.POST("/signup", authHandler.SignUpLocal)
	auth.POST("/login", authHandler.Login)
	auth.POST("/refresh", authHandler.RefreshToken)

	oauth := publicRoute.Group("/oauth")
	oauth.GET("/google/login", authHandler.LoginGoogle)
	oauth.GET("/google/cb", authHandler.GoogleCallback)

	oauth.GET("/github/login", authHandler.LoginGithub)
	oauth.GET("/github/cb", authHandler.GithubCallback)

}
