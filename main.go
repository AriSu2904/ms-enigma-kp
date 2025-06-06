package main

import (
	"awesomeProject/config"
	"awesomeProject/connector"
	"awesomeProject/controllers"
	"awesomeProject/database"
	"awesomeProject/development"
	"awesomeProject/repositories"
	"awesomeProject/services"
	"awesomeProject/utils"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	if err := database.AutoMigrate(db); err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}

	// inject admin
	//seeder := development.NewSeeder(db)
	//if err := seeder.SeedAdminUser(); err != nil {
	//	log.Printf("Failed to seed admin user: %v", err)
	//}

	// inject candidate
	seeder := development.NewSeeder(db)
	if err := seeder.SeedCandidate(); err != nil {
		log.Printf("Failed to seed candidate: %v", err)
	}

	jwtConfig := utils.NewJWTConfig(cfg.JWTSecret, 24*time.Hour)
	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo, jwtConfig)
	authController := controllers.NewAuthController(authService)
	machineLearningConnector := connector.NewMachineLearningConnector()
	candidateRepository := repositories.NewCandidateRepository(db)
	predictorService := services.NewPredictorService(machineLearningConnector)
	candidateService := services.NewCandidateService(candidateRepository, predictorService)
	candidateController := controllers.NewCandidateController(candidateService)

	router := gin.Default()

	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", authController.Login)
		}

		candidates := api.Group("/candidates")
		{
			candidates.GET("", candidateController.List)
			candidates.POST("/predict/:id", candidateController.PredictById)
		}
	}

	log.Println("Server running on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
