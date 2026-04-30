package bootstrapper

import (
	"net/http"

	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/modules/auth/provider"
	ImplService "github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/modules/auth/services/impl"
	userRepoImpl "github.com/RakaMurdiarta/go-koding-akademi-e-commerce/internal/modules/users/repository/impl"

	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/config"
	"github.com/RakaMurdiarta/go-koding-akademi-e-commerce/pkg/database"
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type Server struct {
	DB   *gorm.DB
	e    *echo.Echo
	conf *config.Config
}

func NewServer(e *echo.Echo, c *config.Config, db *gorm.DB) *Server {
	return &Server{e: e, conf: c, DB: db}
}

func (s *Server) InitAPI() {

	// Init Database Transaction
	txManager := database.NewTransactionManager(s.DB)

	// Init Internal Route
	privateRoute, publicRoute := s.initRoute()

	//Register Repository
	userRepo := userRepoImpl.NewUserRepository(txManager)

	//Register Services
	authService := ImplService.NewAuthService(userRepo, s.conf)

	//Register Provider Module
	provider.AuthProvide(privateRoute, publicRoute, s.conf, userRepo, authService)

}

func (s *Server) initRoute() (privateRoute *echo.Group, v1 *echo.Group) {

	s.e.GET("/ping", ping)

	//Group Route API
	api := s.e.Group("/api")
	v1 = api.Group("/v1")
	v1.Group("")
	privateRoute = v1.Group("")
	//Add Middleware here

	return privateRoute, v1

}

func ping(c *echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
