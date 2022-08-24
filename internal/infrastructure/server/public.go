package server

import (
	"github.com/gin-gonic/gin"
	"light-controller/internal/domain/user"
	"light-controller/internal/types"
	"net/http"
	"time"
)

func instantiatePublicRoutes(s *Server, router *gin.RouterGroup) {
	loginHandler := handleLogin(s)
	router.POST("/login", loginHandler)
}

func handleLogin(s *Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form user.Login
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		token, err := s.bus.AuthorizeLogin()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		if s.config.GetGinMode() == types.GinRelease {
			c.SetSameSite(http.SameSiteStrictMode)
			expireCookie := time.Now().Add(time.Hour)
			maxAge := int(expireCookie.Unix() - time.Now().Unix())
			c.SetCookie(s.config.GetGinJwtCookie(), token, maxAge, "/", s.config.GetGinDomain(), true, true)
		} else {
			c.SetSameSite(http.SameSiteNoneMode)
			expireCookie := time.Now().Add(time.Hour * 8760)
			maxAge := int(expireCookie.Unix() - time.Now().Unix())
			c.SetCookie(s.config.GetGinJwtCookie(), token, maxAge, "/", "*", false, false)
		}
	}
}
