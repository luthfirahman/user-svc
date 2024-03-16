package main

import (
	"log"
	"os"

	"github.com/luthfirahman/user-svc/config"
	"github.com/luthfirahman/user-svc/internal"
	"github.com/luthfirahman/user-svc/internal/handler"
	"github.com/luthfirahman/user-svc/internal/repository"
	"github.com/luthfirahman/user-svc/internal/service"
	"github.com/luthfirahman/user-svc/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	var (
		db             *gorm.DB                   = config.SetupDatabaseConnection()
		jwt            config.JWTService          = config.NewJWTService()
		userRepository repository.IUserRepository = repository.NewUserRepository(db)
		userService    service.IUserService       = service.NewUserService(userRepository)
		handler        handler.Handler            = handler.NewHandler(userService)
	)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())
	internal.Router(server, handler, jwt)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := server.Run(":" + port); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
