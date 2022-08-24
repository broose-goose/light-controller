package server

import (
	"github.com/gin-gonic/gin"
	"light-controller/internal/types"
)

type Server struct {
	config Config
	bus    Bus
	router *gin.Engine
}

func NewServer(cfg Config, bus Bus) (*Server, error) {

	switch cfg.GetGinMode() {
	case types.GinDebug:
		gin.SetMode(gin.DebugMode)
		break
	case types.GinTest:
		gin.SetMode(gin.TestMode)
		break
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	srv := &Server{
		bus:    bus,
		router: router,
		config: cfg,
	}

	public := router.Group("/api")
	instantiatePublicRoutes(srv, public)

	private := router.Group("/api")
	instantiatePrivateRoutes(srv, private)

	return srv, nil
}

func instantiatePrivateRoutes(s *Server, rg *gin.RouterGroup) {

}
