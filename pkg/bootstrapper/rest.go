package bootstrapper

import (
	"net/http"

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
	_ = database.NewTransactionManager(s.DB)

	// Init Internal Route
	_ = s.initRoute()

	//Global Repository + Service

}

func (s *Server) initRoute() (v1 *echo.Group) {

	s.e.GET("/ping", ping)

	//Group Route API
	api := s.e.Group("/api")
	v1 = api.Group("/v1")
	v1.Group("")

	return v1

}

func ping(c *echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
