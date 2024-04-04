package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/luthfirahman/user-svc/config"
	"github.com/luthfirahman/user-svc/internal"
	"github.com/luthfirahman/user-svc/internal/handler"
	"github.com/luthfirahman/user-svc/internal/middleware"
	"github.com/luthfirahman/user-svc/internal/repository"
	"github.com/luthfirahman/user-svc/internal/service"
	"github.com/luthfirahman/user-svc/internal/utils"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.Use(middleware.CORSMiddleware())
	internal.Router(server, handler, jwt)
	utils.LoadEnv()
	port := os.Getenv("PORT")

	if err := server.Run(":" + port); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
