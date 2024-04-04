package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/luthfirahman/user-svc/config"
	"github.com/luthfirahman/user-svc/docs"
	"github.com/luthfirahman/user-svc/internal/handler"
	"github.com/luthfirahman/user-svc/internal/middleware"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

func Router(route *gin.Engine, handler handler.Handler, jwtService config.JWTService) {
	// programatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example APIss"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"

	routes := route.Group("/api/v1")
	{
		routes.POST("/register", handler.Register)
		routes.POST("/login", handler.Login)

		routes.GET("/me", middleware.Authenticate(jwtService), handler.GetMyProfile)
		routes.PATCH("", middleware.Authenticate(jwtService), handler.UpdateMyProfile)
	}
}
