package builder

import (
	"github.com/Sandhya-Pratama/Ngonser-App/internal/config"
	"github.com/Sandhya-Pratama/Ngonser-App/internal/http/handler"
	"github.com/Sandhya-Pratama/Ngonser-App/internal/http/router"
	"github.com/Sandhya-Pratama/Ngonser-App/internal/repository"
	"github.com/Sandhya-Pratama/Ngonser-App/internal/service"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func BuildPublicRoutes(cfg *config.Config, db *gorm.DB, redisClient *redis.Client) []*router.Route {
	registrationRepository := repository.NewRegistrationRepository(db)
	registrationService := service.NewRegistrationService(registrationRepository)

	userRepository := repository.NewUserRepository(db, redisClient)
	loginService := service.NewLoginService(userRepository)
	tokenService := service.NewTokenService(cfg)

	ticketRepository := repository.NewTicketRepository(db, redisClient)
	ticketService := service.NewTicketService(ticketRepository)
	ticketHandler := handler.NewTicketHandler(ticketService)

	authHandler := handler.NewAuthHandler(registrationService, loginService, tokenService)
	return router.PublicRoutes(authHandler, ticketHandler)
}

func BuildPrivateRoutes(cfg *config.Config, db *gorm.DB, redisClient *redis.Client) []*router.Route {
	// Create a user handler
	userRepository := repository.NewUserRepository(db, redisClient)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	// Menggunakan PrivateRoutes dengan kedua handler
	return router.PrivateRoutes(userHandler)
}
