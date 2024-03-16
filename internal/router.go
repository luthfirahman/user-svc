package internal

import (
	"github.com/luthfirahman/user-svc/config"
	"github.com/luthfirahman/user-svc/internal/handler"
	"github.com/luthfirahman/user-svc/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Router(route *gin.Engine, handler handler.Handler, jwtService config.JWTService) {
	routes := route.Group("/api/v1/user")
	{
		routes.POST("/register", handler.Register)
		routes.POST("/login", handler.Login)

		routes.GET("/me", middleware.Authenticate(jwtService), handler.GetMyProfile)
		routes.PATCH("", middleware.Authenticate(jwtService), handler.UpdateMyProfile)
	}
}
